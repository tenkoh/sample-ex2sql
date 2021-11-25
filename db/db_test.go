package db_test

import (
	"context"
	"testing"

	"github.com/tenkoh/exsql/db"
)

func Test_Initialize(t *testing.T) {
	c, err := db.Initialize()
	if err != nil {
		t.Error(err)
		return
	}
	defer c.Close()
	if _, err := c.Document.Query().All(context.Background()); err != nil {
		t.Error(err)
	}
}
