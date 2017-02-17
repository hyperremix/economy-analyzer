package token

import (
	"encoding/csv"
	"os"

	"log"

	"github.com/hyperremix/economy-analyzer/backend/model"
)

const repositoryPath string = "C:\\Tokens\\fredr_000\\goplayground\\src\\github.com\\hyperremix\\economy-analyzer\\token.csv"

type TokenRepository struct{}

func (tokenRepository *TokenRepository) Find() []model.Token {
	fileReader, err := os.Open(repositoryPath)

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

	return new(TokenMap).TransformMany(records)
}
