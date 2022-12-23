package data

import (
	"SocialService/proto/social"
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"log"
	"os"
)

type Neo4JRepo struct {
	log    *log.Logger
	driver neo4j.DriverWithContext
}

func Neo4JConnection(log *log.Logger) (*Neo4JRepo, error) {
	uri := os.Getenv("NEO4J_DB")
	user := os.Getenv("NEO4J_USERNAME")
	pass := os.Getenv("NEO4J_PASS")
	auth := neo4j.BasicAuth(user, pass, "")

	driver, err := neo4j.NewDriverWithContext(uri, auth)
	if err != nil {
		log.Panic(err)
		return nil, err
	}
	return &Neo4JRepo{
		driver: driver,
		log:    log,
	}, nil
}

func (nr *Neo4JRepo) CloseDriverConnection(ctx context.Context) {
	nr.driver.Close(ctx)
}

func (nr *Neo4JRepo) CheckConnection() {
	ctx := context.Background()
	err := nr.driver.VerifyConnectivity(ctx)
	if err != nil {
		nr.log.Println("Cannot establish database connection")
		nr.log.Panic(err)
		return
	}
	nr.log.Printf(`Neo4J server address: %s`, nr.driver.Target().Host)
}

func (nr *Neo4JRepo) RegUser(username string) error {
	nr.log.Println("RepoNeo4j - Register User")

	ctx := context.Background()
	session := nr.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)

	savedPerson, err := session.ExecuteWrite(ctx,
		func(transaction neo4j.ManagedTransaction) (any, error) {
			result, err := transaction.Run(ctx,
				"CREATE (u:User) SET u.username = $username RETURN u.username + ', from node ' + id(u)",
				map[string]any{"username": username})
			if err != nil {
				return nil, err
			}

			if result.Next(ctx) {
				return result.Record().Values[0], nil
			}

			return nil, result.Err()
		})
	if err != nil {
		nr.log.Println("Error inserting Person:", err)
		return err
	}
	nr.log.Println(savedPerson.(string))
	return nil
}

func (nr *Neo4JRepo) Follow(usernameOfFollower string, usernameToFollow string, isPrivate bool) error {
	nr.log.Println("RepoNeo4j - Follow User")

	ctx := context.Background()
	session := nr.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)

	if isPrivate {
		createdRelationship, err1 := session.ExecuteWrite(ctx,
			func(tx neo4j.ManagedTransaction) (any, error) {
				var result, err = tx.Run(ctx,
					"optional match (u1:User), (u2:User) where u1.username = $usernameOfFollower and u2.username = $usernameToFollow Merge (u1)-[r:FOLLOWS{status:\"PENDING\"}]->(u2) RETURN r.status",
					map[string]any{"usernameOfFollower": usernameOfFollower, "usernameToFollow": usernameToFollow})

				if err != nil {
					nr.log.Println(err)
					return nil, err
				}

				if result.Next(ctx) {
					return result.Record().Values[0], nil
				}

				return nil, result.Err()
			})
		if err1 != nil {
			nr.log.Println(err1)
			return err1
		}
		nr.log.Println(createdRelationship)
		return nil
	} else {
		createdRelationship, err1 := session.ExecuteWrite(ctx,
			func(tx neo4j.ManagedTransaction) (any, error) {
				var result, err = tx.Run(ctx,
					"optional match (u1:User), (u2:User) where u1.username = $usernameOfFollower and u2.username = $usernameToFollow Merge (u1)-[r:FOLLOWS{status:\"FOLLOWS\"}]->(u2) RETURN r.status",
					map[string]any{"usernameOfFollower": usernameOfFollower, "usernameToFollow": usernameToFollow})

				if err != nil {
					nr.log.Println(err)
					return nil, err
				}
				if result.Next(ctx) {
					return result.Record().Values[0], nil
				}

				return nil, result.Err()
			})
		if err1 != nil {
			nr.log.Println(err1)
			return err1
		}
		nr.log.Println(createdRelationship)
		return nil
	}
	return nil
}

// *TODO: db delete relationship query
func (nr *Neo4JRepo) Unfollow(usernameOfRequester string, usernameToUnfollow string) error {
	nr.log.Println("RepoNeo4j - Unfollow User")

	// TODO db delete relationship query:
	// ...
	// return err

	// Final return after no errors:
	return nil
}

// *TODO: db get username of Nodes whose relationship Status towards "usernameOfRequester" == pending
func (nr *Neo4JRepo) GetPendingFollowers(usernameOfRequester string) ([]*social.PendingFollower, error) {
	nr.log.Println("RepoNeo4j - Get Pending Followers")
	ctx := context.Background()
	session := nr.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)
	pendingRequests, err := session.ExecuteRead(ctx,
		func(transaction neo4j.ManagedTransaction) (any, error) {
			result, err := transaction.Run(ctx,
				"MATCH (u:User)-[r:FOLLOWS]->(u1:User) where u1.username = $username and r.status = \"PENDING\" RETURN u.username as username",
				map[string]any{"username": usernameOfRequester})
			if err != nil {
				return nil, err
			}

			var pendingList []*social.PendingFollower
			for result.Next(ctx) {
				record := result.Record()
				username, ok := record.Get("username")
				if !ok {
					nr.log.Println("Nije dobro ne valja")
					//username = "toki"
				}
				pendingList = append(pendingList, &social.PendingFollower{
					Username: username.(string),
				})
			}

			return pendingList, nil
		})
	if err != nil {
		nr.log.Println("Error querying search:", err)
		return nil, err
	}
	return pendingRequests.([]*social.PendingFollower), nil
}

func (nr *Neo4JRepo) IsFollowed(requesterUsername string, targetUsername string) (bool, error) {
	nr.log.Println("RepoNeo4j - Is Followed")

	follows := false
	//	TODO db query requesterUsername-->follows (accepted)-->targetUsername
	//   follows, errDB := query()
	//	 if errDB != nil {
	//		 nr.log.Println("RepoNeo4j - Error getting Is Followed info from db")
	//		 return follows, errDB

	return follows, nil
}
