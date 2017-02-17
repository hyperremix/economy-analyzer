package token

import (
	"fmt"
	"time"

	"github.com/hyperremix/economy-analyzer/backend/model"
)

type TokenMap struct{}

func (tokenMap *TokenMap) TransformMany(records [][]string) []model.Token {

	var tokens = make([]model.Token, len(records))

	for i, record := range records {
		tokens[i] = tokenMap.Transform(record)
	}

	return tokens
}

func (tokenMap *TokenMap) Transform(record []string) model.Token {

	t, err := time.Parse("02.01.2006", record[1])

	if err != nil {
		panic(fmt.Sprintf("Unable to transform date of record %v for a token", record))
	}

	return model.Token{AccessToken: record[0], CreatedAt: t, TokenType: record[2]}
}
