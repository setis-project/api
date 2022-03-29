package account

import (
	"github.com/setis-project/api/pkg"
	"github.com/setis-project/api/pkg/database"
)

const queryRegister = `
INSERT INTO admin_account (
	first_name,
	last_name,
	email,
	password
) VALUES ($1, $2, $3, $4);
`

func Register(db *database.Db, firstName, lastName, email, password string) error {
	passHash, err := pkg.HashPassword(password)
	if err != nil {
		return err
	}
	_, err = db.Conn.Exec(*db.Ctx, queryRegister, firstName, lastName, email, passHash)
	return err
}
