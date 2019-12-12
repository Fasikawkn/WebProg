package repository

import (
	"database/sql"
	"errors"

	"WebProg/entity"
)

type UsersRepositoryImpl struct {
	conn *sql.DB
}

func (cri *UsersRepositoryImpl) User(id int) (entity.User, error) {

	row := cri.conn.QueryRow("SELECT * FROM users WHERE id = $1", id)

	c := entity.User{}

	err := row.Scan(&c.Id, &c.Uuid, &c.Name, &c.Email, &c.Phone, &c.Password)
	if err != nil {
		return c, err
	}

	return c, nil
}

func NewUsersRepositoryImpl(Conn *sql.DB) *UsersRepositoryImpl {
	return &UsersRepositoryImpl{conn: Conn}
}

func (cri *UsersRepositoryImpl) Users() ([]entity.User, error) {

	rows, err := cri.conn.Query("SELECT * FROM users;")
	if err != nil {
		return nil, errors.New("Could not query the database")
	}
	defer rows.Close()

	users := []entity.User{}

	for rows.Next() {
		user := entity.User{}
		err = rows.Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Phone, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (cri *UsersRepositoryImpl) UpdateUser(c entity.User) error {

	_, err := cri.conn.Exec("UPDATE users SET uuid=$1,full_name=$2, email=$3,phone = $4,password=$5 WHERE id=$6", c.Uuid, c.Name, c.Email, c.Phone, c.Password, c.Id)
	if err != nil {
		return errors.New("Update has failed")
	}

	return nil
}
func (cri *UsersRepositoryImpl) DeleteUser(id int) error {

	_, err := cri.conn.Exec("DELETE FROM users WHERE id=$1", id)
	if err != nil {
		return errors.New("Delete has failed")
	}

	return nil
}

func (cri *UsersRepositoryImpl) AddUser(c entity.User) error {

	_, err := cri.conn.Exec("INSERT INTO users (uuid,full_name,email,phone,password) values($1, $2, $3,$4,$5)", c.Uuid, c.Name, c.Email, c.Phone, c.Password)
	if err != nil {
		return errors.New("Insertion has failed")
	}

	return nil
}
