package controllers

import (
	"errors"
	"time"

	"github.com/go-redis/redis"

	"github.com/setis-project/api/core"
	"github.com/setis-project/api/core/models"
)

func Login(db *core.Db, client *redis.Client, email, password string) (core.Session, error) {
	var session core.Session
	account := new(models.Admin)
	if err := account.GetByEmail(db, email); err != nil {
		return session, err
	}
	if !core.CheckPasswordHash(account.Password, password) {
		return session, errors.New("invalid credentials")
	}

	now := time.Now()
	session = core.NewSession(
		core.NewRefreshToken(now.Add(time.Hour)),
		account.Id,
		now.Add(time.Minute*5),
	)
	err := session.Insert(client)
	if err == core.ErrHasOpenSession {
		return core.GetUserSession(client, account.Id)
	}

	return session, err
}

func Register(db *core.Db, firstName, lastName, email, password string) error {
	passHash, err := core.HashPassword(password)
	if err != nil {
		return err
	}
	return models.NewAdmin(
		firstName,
		lastName,
		email,
		passHash,
	).Insert(db)
}
