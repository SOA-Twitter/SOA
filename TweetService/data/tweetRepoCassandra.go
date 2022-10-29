package data

import (
	"github.com/gocql/gocql"
	"log"
	"time"
)

type TweetRepoCassandra struct {
	log *log.Logger
}

var Session *gocql.Session

func CassandraConnection(log *log.Logger) (TweetRepoCassandra, error) {
	log.Println("Connecting to Cassandra database...")
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "tweets"
	//	cluster.ProtoVersion = 5
	var err error
	Session, err = cluster.CreateSession()
	if err != nil {
		log.Println("Error establishing a database connection")
		return TweetRepoCassandra{}, err
	}
	log.Println("Connected to database")
	return TweetRepoCassandra{log}, nil
}

func (t TweetRepoCassandra) GetAll() Tweets {
	t.log.Println("{TweetRepoCassandra} - getting all tweets")
	var tweets []*Tweet
	m := map[string]interface{}{}
	iter := Session.Query("SELECT * FROM tweets").Iter()
	for iter.MapScan(m) {
		tweets = append(tweets, &Tweet{
			ID:        m["id"].(int),
			Text:      m["text"].(string),
			Picture:   m["picture"].(string),
			CreatedOn: m["created_on"].(string),
			DeletedOn: m["deleted_on"].(string),
		})
		m = map[string]interface{}{}
	}

	return tweets
}

func (t TweetRepoCassandra) CreateTweet(tw *Tweet) {
	t.log.Println("{TweetRepoCassandra} - create tweet")

	query := "INSERT INTO tweets(id,text, picture,created_on, deleted_on) VALUES(?,?, ?, ?, ?)"

	createdOn := time.Now().UTC().String()
	deletedOn := ""
	err := ExecuteQuery(query, 2, tw.Text, tw.Picture, createdOn, deletedOn)
	if err != nil {
		t.log.Println("Error happened during Querying")
	}

}

func (t TweetRepoCassandra) PutTweet(tw *Tweet, id int) error {
	//TODO implement me
	panic("implement me")
}

func (t TweetRepoCassandra) DeleteTweet(id int) error {
	//TODO implement me
	panic("implement me")
}
