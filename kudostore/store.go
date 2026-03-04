package kudostore

import (
	"context"
	"fmt"
	"log"
	"os"

	"main/ent"

	_ "github.com/lib/pq"
)

// NewClient opens a Postgres connection via Ent,
// runs schema migration, and returns a ready-to-use *ent.Client.
// The caller is responsible for calling client.Close().
func NewClient() (*ent.Client, error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)

	client, err := ent.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		client.Close()
		return nil, fmt.Errorf("failed to run schema migration: %w", err)
	}

	log.Println("Database connected and schema migrated successfully")
	return client, nil
}

// WithClient opens a DB connection, passes it to fn, and closes it
// automatically when fn returns — whether it errors or not.
func WithClient(fn func(client *ent.Client) error) error {
	client, err := NewClient()
	if err != nil {
		return err
	}
	defer client.Close()

	return fn(client)
}

// WithTx opens a DB connection, begins a transaction, and passes a
// transaction-scoped *ent.Client to fn.
// - If fn returns nil  → the transaction is committed.
// - If fn returns an error or panics → the transaction is rolled back.
// The connection is always closed when done.
func WithTx(ctx context.Context, fn func(client *ent.Client) error) error {
	dbClient, err := NewClient()
	if err != nil {
		return err
	}
	defer dbClient.Close()

	tx, err := dbClient.Tx(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	// Ensure rollback on panic so the DB is never left in a dirty state.
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v) // re-panic after rollback
		}
	}()

	if err := fn(tx.Client()); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			return fmt.Errorf("rolling back transaction: %w (original error: %v)", rerr, err)
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}

	return nil
}
