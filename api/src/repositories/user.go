package repositories

import (
	"api/src/models"
	"database/sql"
)

type User struct {
	db *sql.DB
}

func NewRepositoryUser(db *sql.DB) *User {

	return &User{db}

}

func (repo User) Created(user models.User) (uint64, error) {

	statement, erro := repo.db.Prepare("insert into users (name, nick, email, password)value (?, ?, ?, ?)")
	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	result, erro := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if erro != nil {
		return 0, erro
	}
	lastID, erro := result.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(lastID), nil
}
