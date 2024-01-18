package mongo_db

import (
	"context"

	"github.com/rattapon001/porter-management/internal/porter/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDbRepository struct {
	Coll *mongo.Collection
}

// NewMongoDbRepository creates a new instance of the MongoDbRepository.
// It takes a *mongo.Collection as a parameter and returns a pointer to the MongoDbRepository.
func NewMongoDbRepository(coll *mongo.Collection) *MongoDbRepository {
	return &MongoDbRepository{
		Coll: coll,
	}
}

// Save saves the given porter to the MongoDB repository.
// If the porter already exists, it updates the existing record.
func (r *MongoDbRepository) Save(porter *domain.Porter) error {
	var existingPorter domain.Porter
	err := r.Coll.FindOne(context.Background(), bson.M{"_id": porter.ID}).Decode(&existingPorter)
	if err == mongo.ErrNoDocuments {
		_, err = r.Coll.InsertOne(context.Background(), porter)
		if err != nil {
			return err
		}
	} else if err != nil {
		// Some error occurred during the find operation
		return err
	} else {
		_, err := r.Coll.UpdateOne(context.Background(), bson.M{"_id": porter.ID}, porter)
		if err != nil {
			return err
		}
	}
	return nil
}

// FindAvailablePorter retrieves an available porter from the MongoDB repository.
// It queries the database for a porter with the status set to domain.PorterStatusAvailable.
// If a matching porter is found, it returns a pointer to the porter.
func (r *MongoDbRepository) FindAvailablePorter() (*domain.Porter, error) {
	var porter domain.Porter
	err := r.Coll.FindOne(context.Background(), bson.M{"status": domain.PorterStatusAvailable}).Decode(&porter)
	if err != nil {
		return nil, err
	}
	return &porter, nil
}

// FindByID retrieves a porter from the database based on the given ID.
// It returns a pointer to the porter and an error if any occurred.
func (r *MongoDbRepository) FindByID(id domain.PorterId) (*domain.Porter, error) {
	var porter domain.Porter
	err := r.Coll.FindOne(context.Background(), bson.M{"_id": id}).Decode(&porter)
	if err != nil {
		return nil, err
	}
	return &porter, nil
}

// FindByCode retrieves a porter from the database based on the given code.
// It returns a pointer to the porter and an error if any occurred.
func (r *MongoDbRepository) FindByCode(code domain.PorterCode) (*domain.Porter, error) {
	var porter domain.Porter
	err := r.Coll.FindOne(context.Background(), bson.M{"code": code}).Decode(&porter)
	if err != nil {
		return nil, err
	}
	return &porter, nil
}
