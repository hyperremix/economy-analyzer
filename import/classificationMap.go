package main

import "github.com/hyperremix/economy-analyzer/import/model"

type classificationMap struct{}

func (classificationMap *classificationMap) TransformMany(records [][]string) []interface{} {

	var classifications = make([]interface{}, len(records))

	for i, record := range records {
		classifications[i] = classificationMap.Transform(record)
	}

	return classifications
}

func (classificationMap *classificationMap) Transform(record []string) interface{} {
	return model.Classification{Client: record[0], Purpose: record[1], Type: record[2]}
}
