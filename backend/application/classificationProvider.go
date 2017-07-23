package application

import (
	"strings"

	"github.com/hyperremix/economy-analyzer/backend/dataAccess"
	"github.com/hyperremix/economy-analyzer/backend/model"
)

type classificationProvider struct {
	classificationRepository *dataAccess.ClassificationRepository
}

func NewClassificationProvider() *classificationProvider {
	return &classificationProvider{classificationRepository: dataAccess.NewClassificationRepository()}
}

func (classificationProvider *classificationProvider) Get(transaction model.Transaction) model.ClassificationType {
	classifications := classificationProvider.classificationRepository.FindMany()

	for _, classification := range classifications {
		if strings.Contains(transaction.Client, classification.Client) && strings.Contains(transaction.Purpose, classification.Purpose) {
			return classification.Type
		}
	}

	return model.Unclassified
}
