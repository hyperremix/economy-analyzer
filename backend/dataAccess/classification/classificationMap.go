package classification

import "github.com/hyperremix/economy-analyzer/backend/model"

type ClassificationMap struct{}

func (classificationMap *ClassificationMap) TransformMany(records [][]string) []model.Classification {

	var classifications = make([]model.Classification, len(records))

	for i, record := range records {
		classifications[i] = classificationMap.Transform(record)
	}

	return classifications
}

func (classificationMap *ClassificationMap) Transform(record []string) model.Classification {
	return model.Classification{Client: record[0], Purpose: record[1], Type: model.ClassificationType(record[2])}
}
