package database

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

type DatabaseInt interface {
	Connect()
	InsertOne(collectionName string, data interface{}) error
	FindOne(collectionName string, filter bson.D, data interface{}) error
	UpdateOne(collectionName string, filter bson.M, data interface{}) error
	DeleteOne(collectionName string, filter bson.D) error
}

type Database struct {
	client       *mongo.Client
	databaseName string
	ctx          context.Context
}

func NewDatabase(databaseName string, ctx context.Context) *Database {
	return &Database{
		databaseName: databaseName,
		ctx:          ctx,
	}
}

func (database *Database) Connect() {
	go func() {
		<-database.ctx.Done()
		database.disconnect()
	}()

	ctx, cancel := context.WithTimeout(database.ctx, 2*time.Second)
	defer cancel()

	clientOptions := options.Client().
		ApplyURI(
			"mongodb://" +
				os.Getenv("MONGO_USERNAME") + ":" +
				os.Getenv("MONGO_PASSWORD") + "@" +
				os.Getenv("MONGO_ADDRESS"))

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Error while connecting to db: %s", err.Error())
	}
	database.client = client

	err = database.client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Error while pinging to db: %s", err.Error())
	}
}

func (database *Database) disconnect() {
	if database.client == nil {
		return
	}

	ctx, cancel := context.WithTimeout(database.ctx, 1*time.Second)
	defer cancel()

	err := database.client.Disconnect(ctx)
	if err != nil {
		logrus.Fatalf("Error while disconnecting from db %s", err.Error())
	}
}

func (database *Database) InsertOne(collectionName string, data interface{}) error {
	coll := database.client.Database(database.databaseName).Collection(collectionName)

	ctx, cancel := context.WithTimeout(database.ctx, 1*time.Second)
	defer cancel()

	_, err := coll.InsertOne(ctx, data)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return &DuplicateKeyError{errorMessage: err.Error()}
		}

		logrus.Errorf("Error while inserting to db: %s", err.Error())
		return &DatabaseError{errorMessage: "Database error"}
	}

	return nil
}

func (database *Database) FindOne(collectionName string, filter bson.D, data interface{}) error {
	coll := database.client.Database(database.databaseName).Collection(collectionName)

	ctx, cancel := context.WithTimeout(database.ctx, 1*time.Second)
	defer cancel()

	err := coll.FindOne(ctx, filter).Decode(data)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &NoDocumentsFoundError{errorMessage: err.Error()}
		}

		logrus.Errorf("Error while retrieving from db: %s", err.Error())
		return &DatabaseError{errorMessage: "Database error"}
	}

	return nil
}

func (database *Database) UpdateOne(collectionName string, filter bson.M, data interface{}) error {
	coll := database.client.Database(database.databaseName).Collection(collectionName)

	ctx, cancel := context.WithTimeout(database.ctx, 1*time.Second)
	defer cancel()

	result, err := coll.UpdateOne(ctx, filter, data)
	if err != nil {
		logrus.Errorf("Error while updating db document: %s", err.Error())

		return &DatabaseError{errorMessage: "Database error"}
	}

	if result.MatchedCount == 0 {
		return &NoDocumentsFoundError{
			errorMessage: "company not found",
		}
	}

	return nil
}

func (database *Database) DeleteOne(collectionName string, filter bson.D) error {
	coll := database.client.Database(database.databaseName).Collection(collectionName)

	ctx, cancel := context.WithTimeout(database.ctx, 1*time.Second)
	defer cancel()

	result, err := coll.DeleteOne(ctx, filter)
	if err != nil {
		logrus.Errorf("Error while deleting from db: %s", err.Error())
		return &DatabaseError{errorMessage: "Database error"}
	}

	if result.DeletedCount == 0 {
		return &NoDocumentsFoundError{errorMessage: "No company with this id"}
	}

	return nil
}

type DatabaseError struct {
	errorMessage string
}

func (databaseError *DatabaseError) Error() string {
	return databaseError.errorMessage
}

type NoDocumentsFoundError struct {
	errorMessage string
}

func (noDocumentsFoundError *NoDocumentsFoundError) Error() string {
	return noDocumentsFoundError.errorMessage
}

type DuplicateKeyError struct {
	errorMessage string
}

func (duplicateKeyError *DuplicateKeyError) Error() string {
	return duplicateKeyError.errorMessage
}
