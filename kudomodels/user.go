package kudomodels

import (
	"context"
	"fmt"
	"main/ent"
	"main/ent/user"
)

// EntCreateUser inserts a new user into the database using Ent.
func CreateUser(ctx context.Context, client *ent.Client, name, email string) (*ent.User, error) {
	u, err := client.User.
		Create().
		SetName(name).
		SetEmail(email).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	return u, nil
}

// EntGetUserByID retrieves a single user by their ID.
func GetUserByID(ctx context.Context, client *ent.Client, id int) (*ent.User, error) {
	u, err := client.User.
		Query().
		Where(user.ID(id)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by id %d: %w", id, err)
	}
	return u, nil
}

// EntGetAllUsers retrieves all users from the database.
func GetAllUsers(ctx context.Context, client *ent.Client) ([]*ent.User, error) {
	users, err := client.User.
		Query().
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all users: %w", err)
	}
	return users, nil
}
