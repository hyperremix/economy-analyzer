package dataAccess

import "github.com/hyperremix/economy-analyzer/backend/model"

type TokenRepository struct{}

func NewTokenRepository() *TokenRepository {
	return new(TokenRepository)
}

func (tr *TokenRepository) Insert(token model.Token) (err error) {
	err = insert("tokens", token)
	return
}
