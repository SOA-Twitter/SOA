package data

import (
	"TweetService/proto/tweet"
	"fmt"
	"github.com/gocql/gocql"
	"log"
	"os"
	"time"
)

type TweetRepoCassandra struct {
	log     *log.Logger
	session *gocql.Session
}

func CassandraConnection(log *log.Logger) (*TweetRepoCassandra, error) {
	log.Println("Connecting to Cassandra database...")
	cassUri := os.Getenv("CASS_URI")
	cluster := gocql.NewCluster(cassUri)
	cluster.Keyspace = "system"
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

func (t *TweetRepoCassandra) CreateTable() {
	err := t.session.Query(
		fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s 
					(id UUID, text text, username text, creationdate timestamp,
					PRIMARY KEY ((username), id))`,
			"tweets_by_username")).Exec()
	if err != nil {
		t.log.Println(err)
	}

	err1 := t.session.Query(
		fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s
				(tweet_id UUID, username text, liked boolean,
				PRIMARY KEY ((tweet_id), username))`,
			"likes_by_tweet")).Exec()
	if err1 != nil {
		t.log.Println(err)
	}
}

//type TweetsByUser []*Tweet

func (t *TweetRepoCassandra) GetTweetsByUsername(username string) ([]*tweet.Tweet, error) {
	t.log.Println("TweetRepoCassandra - Get tweets by username")

	scanner := t.session.Query(`SELECT id, text, username FROM tweets_by_username WHERE username = ?`, username).Iter().Scanner()
	var tweetsByUser []*tweet.Tweet
	for scanner.Next() {
		var tweet tweet.Tweet
		err := scanner.Scan(&tweet.Id, &tweet.Text, &tweet.Username)
		if err != nil {
			t.log.Println(err)
			return nil, err
		}
		tweetsByUser = append(tweetsByUser, &tweet)
	}
	if err := scanner.Err(); err != nil {
		t.log.Println(err)
		return nil, err
	}
	return tweetsByUser, nil
}

func (t *TweetRepoCassandra) CreateTweet(tw *Tweet) error {
	t.log.Println("TweetRepoCassandra - Create tweet")
	id, _ := gocql.RandomUUID()

	err := t.session.Query(`INSERT INTO tweets_by_username(id, text, username, creationdate) VALUES(?, ?, ?, ?)`, id, tw.Text, tw.Username, time.Now()).Exec()
	if err != nil {
		t.log.Println("Error happened during Querying")
		return err
	}
	return nil
}

func (t *TweetRepoCassandra) LikeTweet(id string, username string, like bool) error {
	t.log.Println("TweetRepoCassandra - Like tweet")

	err := t.session.Query(`INSERT INTO likes_by_tweet(tweet_id, username, liked) VALUES(?,?,?)`, id, username, like).Exec()
	if err != nil {
		t.log.Println("Error happened during Querying")
		return err
	}
	return nil
}

func (t *TweetRepoCassandra) GetLikesByTweetId(id string) ([]*tweet.Like, error) {
	t.log.Println("TweetRepoCassandra - Get Likes By Tweet Id")

	scanner := t.session.Query(`SELECT username, liked FROM likes_by_tweet WHERE tweet_id = ?`, id).Iter().Scanner()

	var likes []*tweet.Like
	for scanner.Next() {
		var like tweet.Like
		err := scanner.Scan(&like.Username, &like.Liked)
		if err != nil {
			t.log.Println(err)
			return likes, err
		}
		likes = append(likes, &like)
	}
	if err := scanner.Err(); err != nil {
		t.log.Println(err)
		return likes, err
	}
	return likes, nil
}
