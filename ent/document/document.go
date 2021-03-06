// Code generated by entc, DO NOT EDIT.

package document

import (
	"time"
)

const (
	// Label holds the string label denoting the document type in the database.
	Label = "document"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldFilename holds the string denoting the filename field in the database.
	FieldFilename = "filename"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgePosts holds the string denoting the posts edge name in mutations.
	EdgePosts = "posts"
	// Table holds the table name of the document in the database.
	Table = "documents"
	// PostsTable is the table that holds the posts relation/edge. The primary key declared below.
	PostsTable = "document_posts"
	// PostsInverseTable is the table name for the Post entity.
	// It exists in this package in order to avoid circular dependency with the "post" package.
	PostsInverseTable = "posts"
)

// Columns holds all SQL columns for document fields.
var Columns = []string{
	FieldID,
	FieldFilename,
	FieldCreatedAt,
	FieldUpdatedAt,
}

var (
	// PostsPrimaryKey and PostsColumn2 are the table columns denoting the
	// primary key for the posts relation (M2M).
	PostsPrimaryKey = []string{"document_id", "post_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
)
