package mongo_db

import (
	"context"

	"github.com/rattapon001/porter-management/internal/job/domain"
	infra_errors "github.com/rattapon001/porter-management/internal/job/infra/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDbRepository struct {
	Coll *mongo.Collection
}

// Save saves the given job to the MongoDB repository.
// If the job already exists, it checks the version for optimistic locking.
// If the version is outdated, it returns ErrVersionMismatch.
// Otherwise, it increments the version and updates the job in the repository.
// If the job doesn't exist, it inserts it into the repository.
// Returns an error if any error occurs during the operation.
func (r *MongoDbRepository) Save(job *domain.Job) error {
	currentVersion := job.Version
	var existingJob domain.Job
	err := r.Coll.FindOne(context.Background(), bson.M{"id": job.ID}).Decode(&existingJob)
	if err == mongo.ErrNoDocuments {
		// Job doesn't exist, insert it
		_, err = r.Coll.InsertOne(context.Background(), job)
		if err != nil {
			return err
		}
	} else if err != nil {
		// Some error occurred during the find operation
		return err
	} else {
		// Job already exists, check version for optimistic locking
		if existingJob.Version != currentVersion {
			return infra_errors.ErrVersionMismatch
		}
		// Increment version and update
		job.Version++
		updateResult, err := r.Coll.ReplaceOne(
			context.Background(),
			bson.M{"id": job.ID, "version": existingJob.Version},
			job,
		)

		if err != nil {
			return err
		}

		// If no documents were updated, it means the version was outdated
		if updateResult.MatchedCount == 0 {
			return infra_errors.ErrVersionMismatch
		}

	}
	return nil
}

// FindById finds the job with the given ID in the MongoDB repository.
// Returns the job if found, nil otherwise.
// Returns an error if any error occurs during the operation.
func (r *MongoDbRepository) FindById(id domain.JobId) (*domain.Job, error) {
	var job domain.Job
	err := r.Coll.FindOne(context.Background(), bson.M{"id": id}).Decode(&job)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &job, nil
}
