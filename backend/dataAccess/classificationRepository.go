package dataAccess

import (
	"encoding/csv"
	"os"

	"github.com/hyperremix/economy-analyzer/backend/model"
)

type ClassificationRepository struct {
	classificationMap *classificationMap
}

func NewClassificationRepository() *ClassificationRepository {
	return &ClassificationRepository{classificationMap: new(classificationMap)}
}

func (classificationRepository *ClassificationRepository) Find() []model.Classification {
	fileReader, err := os.Open("C:\\Users\\fredr_000\\goplayground\\src\\github.com\\hyperremix\\economy-analyzer\\classification.csv")

	if err != nil {
		return make([]model.Classification, 0)
	}

	reader := csv.NewReader(fileReader)
	reader.Comma = ';'

	reader.Read()

	records, err := reader.ReadAll()

	if err != nil {
		return make([]model.Classification, 0)
	}

	return classificationRepository.classificationMap.TransformMany(records)
}
