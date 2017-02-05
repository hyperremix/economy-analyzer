package classification

import (
	"encoding/csv"
	"os"

	"github.com/hyperremix/economy-analyzer/backend/model"
)

const repositoryPath string = "C:\\Users\\fredr_000\\goplayground\\src\\github.com\\hyperremix\\economy-analyzer\\classification.csv"

type ClassificationRepository struct{}

func (classificationRepository *ClassificationRepository) Find() []model.Classification {
	fileReader, err := os.Open(repositoryPath)

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

	return new(ClassificationMap).TransformMany(records)
}
