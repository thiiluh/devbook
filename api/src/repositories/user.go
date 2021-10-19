package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
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

func (repo User) Search(NickOrName string) ([]models.User, error) {

	NickOrName = fmt.Sprintf("%%%s%%", NickOrName)

	lines, erro := repo.db.Query("select id, name, nick, email, createdOn from users	where name LIKE ? or nick LIKE ?", NickOrName, NickOrName)
	if erro != nil {
		return nil, erro
	}
	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User

		if erro = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedOn,
		); erro != nil {
			return nil, erro
		}
		users = append(users, user)
	}
	return users, nil
}

func (repo User) SearchId(userId uint64) (models.User, error) {

	lines, erro := repo.db.Query("select id, name, nick, email, createdOn from users where id = ?", userId)

	if erro != nil {
		return models.User{}, erro
	}
	defer repo.db.Close()

	var user models.User

	if lines.Next() {
		if erro = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedOn,
		); erro != nil {
			return models.User{}, erro
		}
	}

	return user, nil
}

func (repo User) Update(userID uint64, user models.User) error {

	statement, erro := repo.db.Prepare("update users set name = ?, nick = ?, email = ? where id = ?")

	if erro != nil {
		return erro
	}
	defer repo.db.Close()

	if _, erro = statement.Exec(user.Name, user.Nick, user.Email, userID); erro != nil {
		return erro
	}
	return nil
}

func (repo User) Delete(userId uint64) error {
	statement, erro := repo.db.Prepare("delete from users where id = ?")
	if erro != nil {
		return erro
	}

	defer repo.db.Close()

	if _, erro = statement.Exec(userId); erro != nil {
		return erro
	}

	return nil

}
