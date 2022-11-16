package data

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

type ProfileRepo struct {
	l   *log.Logger
	cli *mongo.Client
}

func MongoConnection(ctx context.Context, l *log.Logger) (*ProfileRepo, error) {
	dbUri := os.Getenv("MONGO_DB_URI")
	client, err := mongo.NewClient(options.Client().ApplyURI(dbUri))
	if err != nil {
		return nil, err
	}
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	return &ProfileRepo{
		cli: client,
		l:   l,
	}, nil

}
func (pr *ProfileRepo) getCollection() *mongo.Collection {
	userDatabase := pr.cli.Database("userDatabase")
	userCollection := userDatabase.Collection("users")
	return userCollection
}
func (pr *ProfileRepo) Disconnect(ctx context.Context) error {
	err := pr.cli.Disconnect(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (pr *ProfileRepo) Register(user *User) (string, error) {
	pr.l.Println("RepoMongo - Register")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	userCollection := pr.getCollection()
	result, err := userCollection.InsertOne(ctx, &user)
	if err != nil {
		pr.l.Println("Error inserting user into database")
	}
	pr.l.Printf("User ID: %v\n", result.InsertedID)

	useId := result.InsertedID.(primitive.ObjectID)

	return useId.Hex(), nil

}
