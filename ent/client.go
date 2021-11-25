// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/tenkoh/exsql/ent/migrate"

	"github.com/tenkoh/exsql/ent/document"
	"github.com/tenkoh/exsql/ent/post"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Document is the client for interacting with the Document builders.
	Document *DocumentClient
	// Post is the client for interacting with the Post builders.
	Post *PostClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Document = NewDocumentClient(c.config)
	c.Post = NewPostClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:      ctx,
		config:   cfg,
		Document: NewDocumentClient(cfg),
		Post:     NewPostClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		config:   cfg,
		Document: NewDocumentClient(cfg),
		Post:     NewPostClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Document.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Document.Use(hooks...)
	c.Post.Use(hooks...)
}

// DocumentClient is a client for the Document schema.
type DocumentClient struct {
	config
}

// NewDocumentClient returns a client for the Document from the given config.
func NewDocumentClient(c config) *DocumentClient {
	return &DocumentClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `document.Hooks(f(g(h())))`.
func (c *DocumentClient) Use(hooks ...Hook) {
	c.hooks.Document = append(c.hooks.Document, hooks...)
}

// Create returns a create builder for Document.
func (c *DocumentClient) Create() *DocumentCreate {
	mutation := newDocumentMutation(c.config, OpCreate)
	return &DocumentCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Document entities.
func (c *DocumentClient) CreateBulk(builders ...*DocumentCreate) *DocumentCreateBulk {
	return &DocumentCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Document.
func (c *DocumentClient) Update() *DocumentUpdate {
	mutation := newDocumentMutation(c.config, OpUpdate)
	return &DocumentUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DocumentClient) UpdateOne(d *Document) *DocumentUpdateOne {
	mutation := newDocumentMutation(c.config, OpUpdateOne, withDocument(d))
	return &DocumentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DocumentClient) UpdateOneID(id int) *DocumentUpdateOne {
	mutation := newDocumentMutation(c.config, OpUpdateOne, withDocumentID(id))
	return &DocumentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Document.
func (c *DocumentClient) Delete() *DocumentDelete {
	mutation := newDocumentMutation(c.config, OpDelete)
	return &DocumentDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *DocumentClient) DeleteOne(d *Document) *DocumentDeleteOne {
	return c.DeleteOneID(d.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *DocumentClient) DeleteOneID(id int) *DocumentDeleteOne {
	builder := c.Delete().Where(document.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DocumentDeleteOne{builder}
}

// Query returns a query builder for Document.
func (c *DocumentClient) Query() *DocumentQuery {
	return &DocumentQuery{
		config: c.config,
	}
}

// Get returns a Document entity by its id.
func (c *DocumentClient) Get(ctx context.Context, id int) (*Document, error) {
	return c.Query().Where(document.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DocumentClient) GetX(ctx context.Context, id int) *Document {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryPosts queries the posts edge of a Document.
func (c *DocumentClient) QueryPosts(d *Document) *PostQuery {
	query := &PostQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(document.Table, document.FieldID, id),
			sqlgraph.To(post.Table, post.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, document.PostsTable, document.PostsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *DocumentClient) Hooks() []Hook {
	return c.hooks.Document
}

// PostClient is a client for the Post schema.
type PostClient struct {
	config
}

// NewPostClient returns a client for the Post from the given config.
func NewPostClient(c config) *PostClient {
	return &PostClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `post.Hooks(f(g(h())))`.
func (c *PostClient) Use(hooks ...Hook) {
	c.hooks.Post = append(c.hooks.Post, hooks...)
}

// Create returns a create builder for Post.
func (c *PostClient) Create() *PostCreate {
	mutation := newPostMutation(c.config, OpCreate)
	return &PostCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Post entities.
func (c *PostClient) CreateBulk(builders ...*PostCreate) *PostCreateBulk {
	return &PostCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Post.
func (c *PostClient) Update() *PostUpdate {
	mutation := newPostMutation(c.config, OpUpdate)
	return &PostUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PostClient) UpdateOne(po *Post) *PostUpdateOne {
	mutation := newPostMutation(c.config, OpUpdateOne, withPost(po))
	return &PostUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PostClient) UpdateOneID(id int) *PostUpdateOne {
	mutation := newPostMutation(c.config, OpUpdateOne, withPostID(id))
	return &PostUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Post.
func (c *PostClient) Delete() *PostDelete {
	mutation := newPostMutation(c.config, OpDelete)
	return &PostDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *PostClient) DeleteOne(po *Post) *PostDeleteOne {
	return c.DeleteOneID(po.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *PostClient) DeleteOneID(id int) *PostDeleteOne {
	builder := c.Delete().Where(post.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PostDeleteOne{builder}
}

// Query returns a query builder for Post.
func (c *PostClient) Query() *PostQuery {
	return &PostQuery{
		config: c.config,
	}
}

// Get returns a Post entity by its id.
func (c *PostClient) Get(ctx context.Context, id int) (*Post, error) {
	return c.Query().Where(post.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PostClient) GetX(ctx context.Context, id int) *Post {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryDocument queries the document edge of a Post.
func (c *PostClient) QueryDocument(po *Post) *DocumentQuery {
	query := &DocumentQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := po.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(post.Table, post.FieldID, id),
			sqlgraph.To(document.Table, document.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, post.DocumentTable, post.DocumentPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(po.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *PostClient) Hooks() []Hook {
	return c.hooks.Post
}