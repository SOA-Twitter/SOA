package data

import (
	"TweetService/proto/tweet"
	"fmt"
	"github.com/gocql/gocql"
	"log"
	"os"
)

type TweetRepoCassandra struct {
	log     *log.Logger
	session *gocql.Session
}

//var Session *gocql.Session

func CassandraConnection(log *log.Logger) (*TweetRepoCassandra, error) {
	log.Println("Connecting to Cassandra database...")
	cassUri := os.Getenv("CASS_URI")
	cluster := gocql.NewCluster(cassUri)
	cluster.Keyspace = "system"
	//	cluster.ProtoVersion = 5
	var err error
	session, err := cluster.CreateSession()
	if err != nil {
		log.Println("Error establishing a database connection")
		return nil, err
	}
	err = session.Query(
		fmt.Sprintf(`CREATE KEYSPACE IF NOT EXISTS %s
					WITH replication = {
						'class' : 'SimpleStrategy',
						'replication_factor' : %d
					}`, "tweets", 1)).Exec()
	if err != nil {
		log.Println(err)
	}
	err = session.Query(
		fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s 
					(id UUID PRIMARY KEY, text text, picture text, user_email text)`,
			"tweets")).Exec()
	if err != nil {
		log.Println(err)
	}
	session.Close()
	cluster.Keyspace = "tweets"
	cluster.Consistency = gocql.One
	session, err = cluster.CreateSession()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println("Connected to database")
	return &TweetRepoCassandra{log: log, session: session}, nil
}

func (t *TweetRepoCassandra) GetAll() []*tweet.Tweet {
	t.log.Println("{TweetRepoCassandra} - Getting all tweets")
	tweets := []*tweet.Tweet{}
	m := map[string]interface{}{}
	iter := t.session.Query("SELECT * FROM tweets").Iter()
	for iter.MapScan(m) {
		tweets = append(tweets, &tweet.Tweet{
			Text:    m["text"].(string),
			Picture: m["picture"].(string),
		})
		m = map[string]interface{}{}
	}
	return tweets
}

func (t *TweetRepoCassandra) CreateTweet(tw *Tweet) error {
	t.log.Println("{TweetRepoCassandra} - create tweet")
	id, _ := gocql.RandomUUID()

	err := t.session.Query(`INSERT INTO tweets(id, text, picture, user_email) VALUES(?,?, ?,?)`, id, tw.Text, tw.Picture, tw.UserEmail).Exec()
	if err != nil {
		t.log.Println("Error happened during Querying")
		return err
	}
	return nil

}

func (t TweetRepoCassandra) PutTweet(tw *Tweet, id int) error {
	//TODO implement me
	panic("implement me")
}

func (t TweetRepoCassandra) DeleteTweet(id int) error {
	//TODO implement me
	panic("implement me")
}
