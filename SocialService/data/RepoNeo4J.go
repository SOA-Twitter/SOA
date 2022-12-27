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
					"optional match (u1:User), (u2:User) where u1.username = $usernameOfFollower and u2.username = $usernameToFollow Merge (u1)-[r:FOLLOWS{status:\"ACCEPTED\"}]->(u2) RETURN r.status",
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
}

func (nr *Neo4JRepo) Unfollow(usernameOfRequester string, usernameToUnfollow string) error {
	nr.log.Println("RepoNeo4j - Unfollow User")

	ctx := context.Background()
	session := nr.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)

	_, err1 := session.ExecuteWrite(ctx,
		func(tx neo4j.ManagedTransaction) (any, error) {
			var result, err = tx.Run(ctx,
				"MATCH (u1:User)-[r:FOLLOWS]->(u2:User) where u1.username = $usernameOfRequester and u2.username = $usernameToUnfollow DELETE r",
				map[string]any{"usernameOfRequester": usernameOfRequester, "usernameToUnfollow": usernameToUnfollow})

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

	// Final return after no errors:
	return nil
}

// *db get username of Nodes whose relationship Status towards "usernameOfRequester" == pending
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
	nr.log.Println("RepoNeo4j - Logged: ", requesterUsername, "Target user: ", targetUsername)
	isFollowed := false

	ctx := context.Background()
	session := nr.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)
	_, err := session.ExecuteRead(ctx,
		func(transaction neo4j.ManagedTransaction) (any, error) {
			result, err := transaction.Run(ctx,
				"MATCH (u:User)-[r:FOLLOWS]->(u1:User) where u.username = $requesterUsername and u1.username = $targetUsername and r.status = \"ACCEPTED\" RETURN count(r) as relationshipCount",
				map[string]any{"requesterUsername": requesterUsername, "targetUsername": targetUsername})
			if err != nil {
				nr.log.Println("RepoNeo4j - Error getting Is Followed info")
				return isFollowed, err
			}
			for result.Next(ctx) {
				record := result.Record()
				relationshipCount, ok := record.Get("relationshipCount")
				if !ok {
					nr.log.Println("Could not get relationship Count from query result")
				}
				nr.log.Println("Casting relationshipCount found to int...")
				nr.log.Println(relationshipCount)
				if relationshipCount.(int64) == 1 {
					isFollowed = true
					nr.log.Println("isFollowed should be true here:", isFollowed)
				}

			}
			return isFollowed, nil
		})
	if err != nil {
		nr.log.Println("Error querying db:", err)
		return isFollowed, err
	}
	nr.log.Println("isFollowed bool end of func:", isFollowed)
	return isFollowed, nil
}

func (nr *Neo4JRepo) DeclineFollowRequest(usernameOfFollowed string, usernameOfFollower string) error {
	nr.log.Println("RepoNeo4j - Decline Follow Request")

	ctx := context.Background()
	session := nr.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)

	_, err1 := session.ExecuteWrite(ctx,
		func(tx neo4j.ManagedTransaction) (any, error) {
			var result, err = tx.Run(ctx,
				"MATCH (u1:User)-[r:FOLLOWS]->(u2:User) where u1.username = $usernameOfFollower and u2.username = $usernameOfFollowed DELETE r",
				map[string]any{"usernameOfFollowed": usernameOfFollowed, "usernameOfFollower": usernameOfFollower})

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

	// Final return after no errors:
	return nil
}

func (nr *Neo4JRepo) AcceptFollowRequest(usernameOfFollowed string, usernameOfFollower string) error {
	nr.log.Println("RepoNeo4j - Accept Follow Request")

	ctx := context.Background()
	session := nr.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)

	_, err1 := session.ExecuteWrite(ctx,
		func(tx neo4j.ManagedTransaction) (any, error) {
			var result, err = tx.Run(ctx,
				"MATCH (u1:User)-[r:FOLLOWS]->(u2:User) where u1.username = $usernameOfFollower and u2.username = $usernameOfFollowed and r.status = \"PENDING\" SET r.status =\"ACCEPTED\" RETURN r.status as status",
				map[string]any{"usernameOfFollowed": usernameOfFollowed, "usernameOfFollower": usernameOfFollower})

			if err != nil {
				nr.log.Println("ERROR running the query: ", err)
				return nil, err
			}

			if result.Next(ctx) {
				return result.Record().Values[0], nil
			}

			return nil, result.Err()
		})
	if err1 != nil {
		nr.log.Println("ERROR executing the querying func: ", err1)
		return err1
	}

	// Final return after no errors:
	return nil
}
