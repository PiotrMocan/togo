package todo

import (
	"log"

	_ "github.com/lib/pq"
	"togoapi.com/internal/dbclient"
)

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type Todos []Todo

func TodoList() Todos {
	db, err := dbclient.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.DB.Query("SELECT id, title, completed FROM todos")
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
