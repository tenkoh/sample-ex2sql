package db

import (
	"context"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/tenkoh/exsql/ent"
)

const dsn = "file:ent?mode=memory&cache=shared&_fk=1"

// Initialize opens db and creates schema, then returns client.
func Initialize() (*ent.Client, error) {
	client, err := Connect()
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		return nil, fmt.Errorf("fail to create schema; %w", err)
	}
	return client, nil
}

func Connect() (*ent.Client, error) {
	client, err := ent.Open("sqlite3", dsn)
	if err != nil {
		return nil, fmt.Errorf("fail to connect with db; %w", err)
	}
	return client, nil
}
