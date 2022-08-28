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

func (d *DB) FindUsersByIDs(userIDs []string) ([]*User, error) {
	log.Printf("FindUsersByIDs is called (ids: %+v)\n", userIDs)

	users := make([]*User, len(userIDs))

	for i, userID := range userIDs {
		user, ok := dummyUsers[userID]
		if !ok {
			return nil, fmt.Errorf("user not found (id: %s)", userID)
		}
		users[i] = user
	}

	return users, nil
}
