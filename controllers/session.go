package controllers

import (
	"github.com/go-redis/redis"
	"github.com/google/uuid"

	"github.com/setis-project/api/core"
)

func RefreshSession(client *redis.Client, sessionToken, refreshToken uuid.UUID) (core.Session, error) {
	session, err := core.GetSession(client, sessionToken)
	if err != nil {
		return session, err
	}
	return session.Refresh(client, refreshToken)
}
