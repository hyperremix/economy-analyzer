package dataAccess

import (
	"github.com/hyperremix/economy-analyzer/backend/model"
	"gopkg.in/mgo.v2/bson"
)

type TokenRepository struct{}

const collectionName = "tokens"

func NewTokenRepository() *TokenRepository {
	return new(TokenRepository)
}

func (tr *TokenRepository) Upsert(token model.Token) (err error) {
	err = upsert(collectionName, bson.M{"userid": token.UserID}, token)
	return
}

func (tr *TokenRepository) FindSingleByAccessToken(accessToken string) (token model.Token, err error) {
	err = findSingle(collectionName, bson.M{"accesstoken": accessToken}, &token)
	return
}
