package db

import "log"

type Todo struct {
	ID     string
	Text   string
	Done   bool
	UserID string
}

var dummyTodos = []*Todo{
	{
		ID:     "1",
		Text:   "Text1",
		Done:   false,
		UserID: "a",
	},
	{
		ID:     "2",
		Text:   "Text2",
		Done:   true,
		UserID: "a",
	},
	{
		ID:     "3",
		Text:   "Text3",
		Done:   false,
		UserID: "b",
	},
}

func (d *DB) GetTodos() []*Todo {
	log.Println("GetTodos is called.")

	return dummyTodos
}
