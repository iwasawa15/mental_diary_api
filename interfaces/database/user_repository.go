package database

import (
	"security_sample/domain"
)

type UserRepository struct {
	SqlHandler
}

func (repo *UserRepository) Store(u domain.User) (id int, err error) {
	result, err := repo.Execute(
		"INSERT INTO users (name, email) VALUES (?,?)", u.Name, u.Email,
	)
	if err != nil {
		return
	}
	id64, err := result.LastInsertId()
	if err != nil {
		return
	}
	id = int(id64)
	return
}

func (repo *UserRepository) FindById(identifier int) (user domain.User, err error) {
	row, err := repo.Query("SELECT id, name, email FROM users WHERE id = ?", identifier)
	defer row.Close()
	if err != nil {
		return
	}
	var id int
	var name string
	var email string
	row.Next()
	if err = row.Scan(&id, &name, &email); err != nil {
		return
	}
	user.ID = id
	user.Name = name
	user.Email = email
	return
}

func (repo *UserRepository) FindAll() (users domain.Users, err error) {
	rows, err := repo.Query("SELECT id, name, email FROM users")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var id int
		var name string
		var email string
		if err := rows.Scan(&id, &name, &email); err != nil {
			continue
		}
		user := domain.User{
			ID:    id,
			Name:  name,
			Email: email,
		}
		users = append(users, user)
	}
	return
}
