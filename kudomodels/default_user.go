package kudomodels

import (
	"main/kudostore"
	"main/kudotypes"
	"sync"
)

var (
	UserID = uint64(3)
	mu     sync.Mutex
)

func DefaultCreateUser(name, email string) kudotypes.User {
	mu.Lock()
	defer mu.Unlock()

	UserID++

	newUser := kudotypes.User{
		ID:    UserID,
		Name:  name,
		Email: email,
	}

	kudostore.Store = append(kudostore.Store, newUser)

	return newUser
}
