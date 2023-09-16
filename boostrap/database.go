package bootstrap

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/FxIvan/go-backend/mongo"
)

func NewMongoDatabase(env *Env) mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dbHost := env.DBHost
	dbPort := env.DBPort
	dbUser := env.DBUser
	dbPass := env.DBPass

	mongodbURI := fmt.Sprintf("mongodb+srv://%s:%s@%s", dbUser, dbPass, dbHost)

	if dbUser == "" || dbPass == "" {
		mongodbURI = fmt.Sprintf("mongodb+srv://%s:%s", dbHost, dbPort)
	}

	client, _ := mongo.NewClient(mongodbURI)

	/*if err != nil {
		log.Fatal(err)
	}*/

	client.Connect(ctx)
	/*
		if err != nil {
			log.Fatal(err)
		}

		err = client.Ping(ctx)
		if err != nil {
			log.Fatal(err)
		}
	*/
	return client
}

func CloseMongoDBConnection(client mongo.Client) {
	if client == nil {
		return
	}

	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection to MongoDB closed.")
}
