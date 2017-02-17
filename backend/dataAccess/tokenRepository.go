package dataAccess

import (
	"encoding/csv"
	"os"

	"log"

	"github.com/hyperremix/economy-analyzer/backend/model"
)

type TokenRepository struct {
	tokenMap *tokenMap
}

func NewTokenRepository() *TokenRepository {
	return &TokenRepository{tokenMap: new(tokenMap)}
}

func (TokenRepository *TokenRepository) Find() []model.Token {
	fileReader, err := os.Open("C:\\Tokens\\fredr_000\\goplayground\\src\\github.com\\hyperremix\\economy-analyzer\\token.csv")

	if err != nil {
		log.Fatal(err)
		return make([]model.Token, 0)
	}

	reader := csv.NewReader(fileReader)
	reader.Comma = ';'

	reader.Read()

	records, err := reader.ReadAll()

	if err != nil {
		log.Fatal(err)
		return make([]model.Token, 0)
	}

	return TokenRepository.tokenMap.TransformMany(records)
}
