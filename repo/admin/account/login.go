package account

import (
	"errors"
	"time"

	"github.com/go-redis/redis"
	"github.com/google/uuid"

	"github.com/setis-project/api/pkg"
	"github.com/setis-project/api/pkg/database"
	predis "github.com/setis-project/api/pkg/redis"
)

type loginCredentials struct {
	Id       uuid.UUID `json:"id"`
	Password string    `json:"password"`
}

const queryLogin = `
SELECT
	id,
	password
FROM admin_account
WHERE email=$1
`

func Login(db *database.Db, client *redis.Client, email, password string) (predis.Session, error) {
	var session predis.Session
	var row loginCredentials
	err := db.Conn.QueryRow(*db.Ctx, queryLogin, email).Scan(&row.Id, &row.Password)
	if err != nil {
		return session, err
	}
	if !pkg.CheckPasswordHash(row.Password, password) {
		return session, errors.New("invalid credentials")
	}
	now := time.Now()
	session = predis.NewSession(
		predis.NewRefreshToken(now.Add(time.Hour)),
		row.Id,
		now.Add(time.Minute*5),
	)
	err = session.Insert(client)
	if err == predis.ErrHasOpenSession {
		return predis.GetUserSession(client, row.Id)
	}
	return session, err
}
