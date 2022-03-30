package session

import (
	"encoding/json"

	"github.com/go-redis/redis"
	"github.com/google/uuid"
	predis "github.com/setis-project/api/pkg/redis"
)

func RefreshSession(client *redis.Client, sessionToken, refreshToken uuid.UUID) (predis.Session, error) {
	var session predis.Session
	err := json.Unmarshal([]byte(client.Get("session:"+sessionToken.String()).Val()), &session)
	if err != nil {
		return session, err
	}
	newSession, err := session.Refresh(client, refreshToken)
	return newSession, err
}
