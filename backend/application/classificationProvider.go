package application

import (
	"strings"

	"github.com/hyperremix/economy-analyzer/backend/model"
)

type classificationProvider struct{}

func (classificationProvider *classificationProvider) Get(transaction model.Transaction, classifications []model.Classification) model.ClassificationType {
	for _, classification := range classifications {
		if strings.Contains(transaction.Client, classification.Client) && strings.Contains(transaction.Purpose, classification.Purpose) {
			return classification.Type
		}
	}

	return model.Unclassified
}
