package data

import (
	"TweeterMicro/TweetService/proto/tweet"
	"github.com/gocql/gocql"
	"log"
)

type TweetRepoCassandra struct {
	log     *log.Logger
	session *gocql.Session
}

//var Session *gocql.Session

func CassandraConnection(log *log.Logger) (TweetRepoCassandra, error) {
	log.Println("Connecting to Cassandra database...")
	cluster := gocql.NewCluster("127.0.0.1")
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
			Id:      int32(m["id"].(int)),
			Text:    m["text"].(string),
			Picture: m["picture"].(string),
		})
		m = map[string]interface{}{}
	}
	return tweets
}

func (t *TweetRepoCassandra) CreateTweet(tw *Tweet) error {
	t.log.Println("{TweetRepoCassandra} - create tweet")

	query := "INSERT INTO tweets(id,text, picture) VALUES(?,?, ?)"

	err := t.session.Query(query).Bind(2, tw.Text, tw.Picture).Exec()
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
