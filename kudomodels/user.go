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

func nextIdGenerator() {
	mu.Lock()
	defer mu.Unlock()

	UserID++
}

func CreateUser(name, email string) kudotypes.User {
	nextIdGenerator()

	mu.Lock()
	defer mu.Unlock()

	newUser := kudotypes.User{
		ID:    UserID,
		Name:  name,
		Email: email,
	}

	kudostore.Store = append(kudostore.Store, newUser)

	return newUser
}
