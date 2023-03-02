package database

import (
	"context"
	"errors"
	"log"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DatabaseInt interface {
	Connect()
	InsertOne(collectionName string, data interface{}) error
	FindOne(collectionName string, filter bson.D, data interface{}) error
	UpdateOne(collectionName string, filter bson.M, data interface{}) error
	DeleteOne(collectionName string, filter bson.D) error
}

type Database struct {
	ctx          context.Context
	client       *mongo.Client
	databaseName string
}

func NewDatabase(ctx context.Context, databaseName string) *Database {
	return &Database{
		ctx:          ctx,
		databaseName: databaseName,
	}
}

func (db *Database) Connect() {
	go func() {
		<-db.ctx.Done()
		db.disconnect()
	}()

	ctx, cancel := context.WithTimeout(db.ctx, 5*time.Second)
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
	db.client = client

	err = db.client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Error while pinging to db: %s", err.Error())
	}
}

func (db *Database) disconnect() {
	if db.client == nil {
		return
	}

	ctx, cancel := context.WithTimeout(db.ctx, 1*time.Second)
	defer cancel()

	err := db.client.Disconnect(ctx)
	if err != nil {
		logrus.Fatalf("Error while disconnecting from db %s", err.Error())
	}
}

func (db *Database) InsertOne(collectionName string, data interface{}) error {
	coll := db.client.Database(db.databaseName).Collection(collectionName)

	ctx, cancel := context.WithTimeout(db.ctx, 1*time.Second)
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

func (db *Database) FindOne(collectionName string, filter bson.D, data interface{}) error {
	coll := db.client.Database(db.databaseName).Collection(collectionName)

	ctx, cancel := context.WithTimeout(db.ctx, 1*time.Second)
	defer cancel()

	err := coll.FindOne(ctx, filter).Decode(data)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return &NoDocumentsFoundError{errorMessage: err.Error()}
		}

		logrus.Errorf("Error while retrieving from db: %s", err.Error())

		return &DatabaseError{errorMessage: "Database error"}
	}

	return nil
}

func (db *Database) UpdateOne(collectionName string, filter bson.M, data interface{}) error {
	coll := db.client.Database(db.databaseName).Collection(collectionName)

	ctx, cancel := context.WithTimeout(db.ctx, 1*time.Second)
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

func (db *Database) DeleteOne(collectionName string, filter bson.D) error {
	coll := db.client.Database(db.databaseName).Collection(collectionName)

	ctx, cancel := context.WithTimeout(db.ctx, 1*time.Second)
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

func (dbe *DatabaseError) Error() string {
	return dbe.errorMessage
}

type NoDocumentsFoundError struct {
	errorMessage string
}

func (nde *NoDocumentsFoundError) Error() string {
	return nde.errorMessage
}

type DuplicateKeyError struct {
	errorMessage string
}

func (dke *DuplicateKeyError) Error() string {
	return dke.errorMessage
}
