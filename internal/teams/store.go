package teams

import "go.mongodb.org/mongo-driver/mongo"

type TeamStore struct {
	mongoDriver    *mongo.Database
	collection     *mongo.Collection
	collectionName string
}

func NewStore(db *mongo.Database) *TeamStore {
	return &TeamStore{
		mongoDriver:    db,
		collectionName: "team",
		collection:     db.Collection("team"),
	}
}
