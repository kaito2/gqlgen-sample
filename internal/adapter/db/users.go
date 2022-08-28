package db

import (
	"fmt"
	"log"
)

type User struct {
	ID   string
	Name string
}

var dummyUsers = map[string]*User{
	"a": {
		ID:   "a",
		Name: "Name of a",
	},
	"b": {
		ID:   "b",
		Name: "Name of b",
	},
}

func (d *DB) GetUser(userID string) (*User, error) {
	log.Printf("GetUser is called (id: %s)\n", userID)

	user, ok := dummyUsers[userID]
	if !ok {
		return nil, fmt.Errorf("user not found (id: %s)", userID)
	}
	return user, nil
}
