package mongoutils

import (
	"os"

	"github.com/betonr/go-utils/base"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectMongo - open the connection with MONGO database
func ConnectMongo() (*mongo.Database, error) {
	dbHost := base.GetBetween([]string{os.Getenv("DBHOST"), "localhost"})
	dbPort := base.GetBetween([]string{os.Getenv("DBPORT"), "27017"})
	dbName := base.GetBetween([]string{os.Getenv("DBNAME"), "db"})
	dbUser := base.GetBetween([]string{os.Getenv("DBUSER"), "root"})
	dbPass := base.GetBetween([]string{os.Getenv("DBPASS"), "mongo"})

	dbInfos := &options.ClientOptions{
		Hosts: []string{dbHost + ":" + dbPort},
		Auth: &options.Credential{
			AuthSource: "admin",
			Username:   dbUser,
			Password:   dbPass,
		},
	}

	client, err := mongo.NewClient(dbInfos)
	if err != nil {
		return nil, err
	}
	return client.Database(dbName), nil
}
