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

// *TODO: db Find & Create Relationship Call
func (nr *Neo4JRepo) Follow(usernameOfFollower string, usernameToFollow string, isPrivate bool) (string, error) {
	nr.log.Println("RepoNeo4j - Follow User")

	// TODO status = property of "follows" relationship ( pending, follows )
	var status string = ""

	//	TODO db call to find if relationship exists - return relationship status
	//	...

	//	TODO if relationship doesn't exist - create "follows" relationship between 2 User Nodes,
	//	 with property status = follows (if not "isPrivate") / pending (if "isPrivate")
	//	 + return relationship status
	//	...

	//

	// TODO on db functions error:
	// return status, err

	// Final return after no errors:
	return status, nil
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

	result := []*social.PendingFollower{}
	//result, err := query()

	//	TODO get all "follows" relationships w/ status property == pending
	//if err != nil {
	//	nr.log.Println("RepoNeo4j - Error Getting Pending Followers: ", err)
	//	return nil, err
	//}
	return result, nil
}
