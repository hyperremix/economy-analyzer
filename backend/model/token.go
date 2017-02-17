package model

import "time"

type Token struct {
	AccessToken string
	CreatedAt   time.Time
	TokenType   string
}
