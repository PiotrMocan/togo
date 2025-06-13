package user

import (
	"log"
	"time"

	_ "github.com/lib/pq"
	"togoapi.com/internal/dbclient"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Users []User

func UserList() Users {
	db, err := dbclient.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.DB.Query("SELECT id, name, email, created_at, updated_at FROM users ORDER BY id")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users Users
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return users
}

func GetUser(id int) (*User, error) {
	db, err := dbclient.NewClient()
	if err != nil {
		return nil, err
	}

	var user User
	err = db.DB.QueryRow("SELECT id, name, email, created_at, updated_at FROM users WHERE id = $1", id).
		Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func CreateUser(name, email string) (*User, error) {
	db, err := dbclient.NewClient()
	if err != nil {
		return nil, err
	}

	var user User
	err = db.DB.QueryRow(
		"INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id, name, email, created_at, updated_at",
		name, email).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func UpdateUser(id int, name, email string) (*User, error) {
	db, err := dbclient.NewClient()
	if err != nil {
		return nil, err
	}

	var user User
	err = db.DB.QueryRow(
		"UPDATE users SET name = $1, email = $2, updated_at = CURRENT_TIMESTAMP WHERE id = $3 RETURNING id, name, email, created_at, updated_at",
		name, email, id).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}