package databases

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Client       *mongo.Client
	DatabaseName string
}

var (
	Database MongoDB
)

func (db *MongoDB) Connect() error {
	var client_options = options.Client()

	client_options.ApplyURI("mongodb://ikanpakus:ikanpakus@localhost:27017")

	client, err := mongo.Connect(context.TODO(), client_options)

	if err != nil {
		return err
	}

	db.DatabaseName = "belajar-golang"
	db.Client = client

	return err
}

func (db *MongoDB) Disconnect() {
	if err := db.Client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}
