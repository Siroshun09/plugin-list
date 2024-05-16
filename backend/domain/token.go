package domain

import "time"

type Token struct {
	Value   string
	Created time.Time
}
