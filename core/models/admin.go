package models

import "github.com/setis-project/api/core"

type Admin struct {
	Id        string `json:"id,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
}

func NewAdmin(firstName, lastName, email, password string) *Admin {
	return &Admin{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  password,
	}
}

func (admin *Admin) Insert(db *core.Db) error {
	const query = `
	INSERT INTO admin_account (
		first_name,
		last_name,
		email,
		password
	) VALUES ($1, $2, $3, $4);
	`
	_, err := db.Conn.Exec(
		*db.Ctx,
		query,
		admin.FirstName,
		admin.LastName,
		admin.Email,
		admin.Password,
	)
	return err
}

func (admin *Admin) GetByEmail(db *core.Db, email string) error {
	const query = `
	SELECT
		id,
		first_name,
		last_name,
		password
	FROM admin_account
	WHERE email=$1
	`
	row := db.Conn.QueryRow(*db.Ctx, query, email)
	return row.Scan(
		&admin.Id,
		&admin.FirstName,
		&admin.LastName,
		&admin.Password,
	)
}
