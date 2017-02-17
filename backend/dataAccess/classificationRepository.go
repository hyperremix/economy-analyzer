package dataAccess

import "github.com/hyperremix/economy-analyzer/backend/model"

type ClassificationRepository struct{}

func NewClassificationRepository() *ClassificationRepository {
	return new(ClassificationRepository)
}

func (classificationRepository *ClassificationRepository) FindMany() []model.Classification {
	var results []model.Classification

	findMany("classifications", &results)
	return results
}
