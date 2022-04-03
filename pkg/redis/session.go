package redis

import (
	"encoding/json"
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
	if client.Get("user:"+s.UserId.String()).Err() != redis.Nil {
		return ErrHasOpenSession
	}
	exp := time.Until(s.Expiry)
	jsonSession, err := json.Marshal(s)
	if err != nil {
		return err
	}
	client.Set("session:"+s.Token.String(), jsonSession, exp)
	client.Set("user:"+s.UserId.String(), s.Token.String(), exp)
	return nil
}

func (s Session) Refresh(client *redis.Client, refreshToken uuid.UUID) (Session, error) {
	var newSession Session
	if s.RefreshToken.Token != refreshToken || s.RefreshToken.IsExpired() {
		return newSession, ErrNotAuthorized
	}
	if client.Del("session:"+s.Token.String()).Val() == 0 {
		return newSession, ErrSessionNotExists
	}
	newSession = NewSession(
		s.RefreshToken,
		s.UserId,
		s.Expiry,
	)
	exp := time.Until(s.Expiry)
	jsonSession, err := json.Marshal(newSession)
	if err != nil {
		return newSession, err
	}
	client.Set("session:"+s.Token.String(), jsonSession, exp)
	client.Set("user:"+s.UserId.String(), s.Token.String(), exp)
	return newSession, err
}

func GetSession(client *redis.Client, token uuid.UUID) (Session, error) {
	var session Session
	redisSession := client.Get("session:" + token.String())
	if redisSession.Err() == redis.Nil {
		return session, ErrSessionNotExists
	}
	err := json.Unmarshal([]byte(redisSession.Val()), &session)
	return session, err
}

func GetUserSession(client *redis.Client, userId uuid.UUID) (Session, error) {
	var session Session
	redisUser := client.Get("user:" + userId.String())
	if redisUser.Err() == redis.Nil {
		return session, ErrUserNotExists
	}
	token, err := uuid.Parse(redisUser.Val())
	if err != nil {
		return session, ErrInvalidSessionToken
	}
	return GetSession(client, token)
}
