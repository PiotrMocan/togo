package todo

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type Todos []Todo

func TodoList() Todos {
	db, err := sql.Open("postgres", "postgres://postgres:@localhost/togo_dev?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT id, title, completed FROM todos")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var todos Todos
	for rows.Next() {
		var todo Todo
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed)
		if err != nil {
			log.Fatal(err)
		}
		todos = append(todos, todo)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return todos
}
