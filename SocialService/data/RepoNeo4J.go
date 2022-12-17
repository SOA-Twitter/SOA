package data

import (
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

func (n Neo4JRepo) AddUser(username string) error {
	//TODO implement me
	panic("implement me")
}
