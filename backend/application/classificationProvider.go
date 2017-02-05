package application

import (
	"strings"

	"github.com/hyperremix/economy-analyzer/backend/model"
)

type ClassificationProvider struct{}

func (classificationProvider *ClassificationProvider) Get(transaction model.Transaction, classifications []model.Classification) model.ClassificationType {
	for _, classification := range classifications {
		if strings.Contains(transaction.Client, classification.Client) && strings.Contains(transaction.Purpose, classification.Purpose) {
			return classification.Type
		}
	}

	return model.Unclassified
}
