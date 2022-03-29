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

func Login(db *database.Db, client *redis.Client, email, password string) (loginCredentials, error) {
	row := loginCredentials{}
	err := db.Conn.QueryRow(*db.Ctx, queryLogin, email).Scan(&row.Id, &row.Password)
	if err != nil {
		return row, err
	}
	if !pkg.CheckPasswordHash(row.Password, password) {
		return row, errors.New("invalid credentials")
	}
	now := time.Now()
	session := predis.NewSession(
		predis.NewRefreshToken(now.Add(time.Hour)),
		row.Id,
		now.Add(time.Minute*5),
	)

	return row, nil
}
