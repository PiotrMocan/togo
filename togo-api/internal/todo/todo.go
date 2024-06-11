package todo

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type Todos []Todo

func TodoList() Todos {
	return []Todo{
		{ID: 1, Title: "Go to the store", Completed: false},
		{ID: 2, Title: "Buy milk", Completed: false},
		{ID: 3, Title: "Feed the cat", Completed: false},
	}
}
