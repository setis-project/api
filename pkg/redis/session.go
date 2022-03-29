package redis

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/go-redis/redis"
	"github.com/google/uuid"
)

type Session struct {
	Token        uuid.UUID
	RefreshToken RefreshToken
	UserId       uuid.UUID
	Expiry       time.Time
}

func NewSession(rt RefreshToken, userId uuid.UUID, exp time.Time) Session {
	return Session{
		Token:        uuid.New(),
		RefreshToken: rt,
		UserId:       userId,
		Expiry:       exp,
	}
}

func (s Session) IsExpired() bool {
	return s.Expiry.Before(time.Now())
}

func (s Session) Insert(client *redis.Client) {
	exp := time.Until(s.Expiry)
	client.Set("session:"+s.Token.String(), s, exp)
	client.Set("user:"+s.UserId.String(), s.Token, exp)
}

func (s Session) Refresh(client *redis.Client, oldSessionId, refreshToken uuid.UUID) error {
	oldSession := Session{}
	err := json.Unmarshal([]byte(client.Get(oldSessionId.String()).Val()), &oldSession)
	if err != nil {
		return err
	}
	if oldSession.RefreshToken.Token != refreshToken || oldSession.RefreshToken.IsExpired() {
		return errors.New("not authorized")
	}
	if client.Del("session:"+oldSessionId.String()).Val() == 0 {
		return errors.New("session doesnt exists")
	}
	s.Insert(client)
	return nil
}
