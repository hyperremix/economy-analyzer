package dataAccess

import "github.com/hyperremix/economy-analyzer/backend/model"

type TokenRepository struct{}

func NewTokenRepository() *TokenRepository {
	return new(TokenRepository)
}

func (TokenRepository *TokenRepository) FindMany() []model.Token {
	var results []model.Token

	findMany("tokens", &results)
	return results
}
