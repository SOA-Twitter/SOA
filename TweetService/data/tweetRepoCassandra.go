package data

import (
	"TweetService/proto/tweet"
	"github.com/gocql/gocql"
	"github.com/google/uuid"
	"log"
	"os"
)

type TweetRepoCassandra struct {
	log     *log.Logger
	session *gocql.Session
}

//var Session *gocql.Session

func CassandraConnection(log *log.Logger) (TweetRepoCassandra, error) {
	log.Println("Connecting to Cassandra database...")
	cassUri := os.Getenv("CASS_URI")
	cluster := gocql.NewCluster(cassUri)
	cluster.Keyspace = "tweets"
	//	cluster.ProtoVersion = 5
	var err error
	session, err := cluster.CreateSession()
	if err != nil {
		log.Println("Error establishing a database connection")
		return TweetRepoCassandra{}, err
	}
	log.Println("Connected to database")
	return TweetRepoCassandra{log, session}, nil
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
	id := uuid.New().String()
	t.log.Println(tw.UserId)

	query := "INSERT INTO tweets(id,text, picture,user_id) VALUES(?,?, ?,?)"
	err := t.session.Query(query).Bind(id, tw.Text, tw.Picture, tw.UserId).Exec()
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
