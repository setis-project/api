package redis

import (
	"time"

	"github.com/google/uuid"
)

type RefreshToken struct {
	Token  uuid.UUID
	Expiry time.Time
}

func NewRefreshToken(exp time.Time) RefreshToken {
	return RefreshToken{
		Token:  uuid.New(),
		Expiry: exp,
	}
}

func (rt RefreshToken) IsExpired() bool {
	return rt.Expiry.Before(time.Now())
}
