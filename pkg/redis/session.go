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

func (s Session) Insert(client *redis.Client) error {
	exp := time.Until(s.Expiry)
	if client.Get("user:"+s.UserId.String()).Err() != redis.Nil {
		return errors.New("user has an open session")
	}
	client.Set("session:"+s.Token.String(), s, exp)
	client.Set("user:"+s.UserId.String(), s.Token, exp)
	return nil
}

func (s Session) Refresh(client *redis.Client, refreshToken uuid.UUID) (Session, error) {
	var newSession Session
	if s.RefreshToken.Token != refreshToken || s.RefreshToken.IsExpired() {
		return newSession, errors.New("not authorized")
	}
	if client.Del("session:"+s.Token.String()).Val() == 0 {
		return newSession, errors.New("session doesnt exists")
	}
	newSession = NewSession(
		s.RefreshToken,
		s.UserId,
		s.Expiry,
	)
	newSession.Insert(client)
	return newSession, nil
}

func GetSession(client *redis.Client, token uuid.UUID) (Session, error) {
	var session Session
	err := json.Unmarshal([]byte(client.Get("session:"+token.String()).Val()), &session)
	return session, err
}

func GetUserSession(client *redis.Client, userId uuid.UUID) (Session, error) {
	var session Session
	token, err := uuid.Parse(client.Get("user:" + userId.String()).Val())
	if err != nil {
		return session, errors.New("invalid value for session token")
	}
	return GetSession(client, token)
}
