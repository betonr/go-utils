package mongoutils

import (
	"os"
	"time"

	"github.com/betonr/go-utils/base"
	"gopkg.in/mgo.v2"
)

// ConnectMongo - open the connection with MONGO database
func ConnectMongo() (*mgo.Database, error) {
	dbHost := base.GetBetween([]string{os.Getenv("DBHOST"), "localhost"})
	dbPort := base.GetBetween([]string{os.Getenv("DBPORT"), "27017"})
	dbName := base.GetBetween([]string{os.Getenv("DATABASE"), "db"})
	dbUser := base.GetBetween([]string{os.Getenv("DBUSER"), "root"})
	dbPass := base.GetBetween([]string{os.Getenv("DBPASS"), "mongo"})

	dbInfos := &mgo.DialInfo{
		Addrs:    []string{dbHost + ":" + dbPort},
		Timeout:  60 * time.Second,
		Database: "admin",
		Username: dbUser,
		Password: dbPass,
	}

	session, err := mgo.DialWithInfo(dbInfos)
	if err != nil {
		return nil, err
	}
	return session.DB(dbName), nil
}
