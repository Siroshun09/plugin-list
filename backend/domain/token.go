package domain

import "time"

type Token struct {
	Value   string
	Created time.Time
}

func NewToken(token string, created time.Time) *Token {
	return &Token{token, created}
}
