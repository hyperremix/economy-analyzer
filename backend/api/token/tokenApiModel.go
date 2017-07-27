package token

import "github.com/hyperremix/economy-analyzer/backend/model"

type tokenApiModel struct {
	AccessToken string  `json:"access_token"`
	ExpiresIn   float64 `json:"expires_in"`
	TokenType   string  `json:"token_type"`
}

func NewTokenApiModel(token model.Token) tokenApiModel {
	return tokenApiModel{AccessToken: token.AccessToken, ExpiresIn: token.ExpiresIn, TokenType: "bearer"}
}
