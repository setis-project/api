package session

import (
	"github.com/go-redis/redis"
	"github.com/google/uuid"
	predis "github.com/setis-project/api/pkg/redis"
)

func RefreshSession(client *redis.Client, sessionToken, refreshToken uuid.UUID) (predis.Session, error) {
	session, err := predis.GetSession(client, sessionToken)
	if err != nil {
		return session, err
	}
	return session.Refresh(client, refreshToken)
}
