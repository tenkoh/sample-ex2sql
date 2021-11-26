// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/tenkoh/exsql/ent/document"
	"github.com/tenkoh/exsql/ent/post"
	"github.com/tenkoh/exsql/ent/predicate"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeDocument = "Document"
	TypePost     = "Post"
)

// DocumentMutation represents an operation that mutates the Document nodes in the graph.
type DocumentMutation struct {
	config
	op            Op
	typ           string
	id            *int
	filename      *string
	created_at    *time.Time
	updated_at    *time.Time
	clearedFields map[string]struct{}
	posts         map[int]struct{}
	removedposts  map[int]struct{}
	clearedposts  bool
	done          bool
	oldValue      func(context.Context) (*Document, error)
	predicates    []predicate.Document
}

var _ ent.Mutation = (*DocumentMutation)(nil)

// documentOption allows management of the mutation configuration using functional options.
type documentOption func(*DocumentMutation)

// newDocumentMutation creates new mutation for the Document entity.
func newDocumentMutation(c config, op Op, opts ...documentOption) *DocumentMutation {
	m := &DocumentMutation{
		config:        c,
		op:            op,
		typ:           TypeDocument,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withDocumentID sets the ID field of the mutation.
func withDocumentID(id int) documentOption {
	return func(m *DocumentMutation) {
		var (
			err   error
			once  sync.Once
			value *Document
		)
		m.oldValue = func(ctx context.Context) (*Document, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Document.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withDocument sets the old Document of the mutation.
func withDocument(node *Document) documentOption {
	return func(m *DocumentMutation) {
		m.oldValue = func(context.Context) (*Document, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m DocumentMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m DocumentMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *DocumentMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetFilename sets the "filename" field.
func (m *DocumentMutation) SetFilename(s string) {
	m.filename = &s
}

// Filename returns the value of the "filename" field in the mutation.
func (m *DocumentMutation) Filename() (r string, exists bool) {
	v := m.filename
	if v == nil {
		return
	}
	return *v, true
}

// OldFilename returns the old "filename" field's value of the Document entity.
// If the Document object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *DocumentMutation) OldFilename(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldFilename is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldFilename requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldFilename: %w", err)
	}
	return oldValue.Filename, nil
}

// ResetFilename resets all changes to the "filename" field.
func (m *DocumentMutation) ResetFilename() {
	m.filename = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *DocumentMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *DocumentMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Document entity.
// If the Document object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *DocumentMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *DocumentMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *DocumentMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *DocumentMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Document entity.
// If the Document object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *DocumentMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *DocumentMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// AddPostIDs adds the "posts" edge to the Post entity by ids.
func (m *DocumentMutation) AddPostIDs(ids ...int) {
	if m.posts == nil {
		m.posts = make(map[int]struct{})
	}
	for i := range ids {
		m.posts[ids[i]] = struct{}{}
	}
}

// ClearPosts clears the "posts" edge to the Post entity.
func (m *DocumentMutation) ClearPosts() {
	m.clearedposts = true
}

// PostsCleared reports if the "posts" edge to the Post entity was cleared.
func (m *DocumentMutation) PostsCleared() bool {
	return m.clearedposts
}

// RemovePostIDs removes the "posts" edge to the Post entity by IDs.
func (m *DocumentMutation) RemovePostIDs(ids ...int) {
	if m.removedposts == nil {
		m.removedposts = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.posts, ids[i])
		m.removedposts[ids[i]] = struct{}{}
	}
}

// RemovedPosts returns the removed IDs of the "posts" edge to the Post entity.
func (m *DocumentMutation) RemovedPostsIDs() (ids []int) {
	for id := range m.removedposts {
		ids = append(ids, id)
	}
	return
}

// PostsIDs returns the "posts" edge IDs in the mutation.
func (m *DocumentMutation) PostsIDs() (ids []int) {
	for id := range m.posts {
		ids = append(ids, id)
	}
	return
}

// ResetPosts resets all changes to the "posts" edge.
func (m *DocumentMutation) ResetPosts() {
	m.posts = nil
	m.clearedposts = false
	m.removedposts = nil
}

// Where appends a list predicates to the DocumentMutation builder.
func (m *DocumentMutation) Where(ps ...predicate.Document) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *DocumentMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Document).
func (m *DocumentMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *DocumentMutation) Fields() []string {
	fields := make([]string, 0, 3)
	if m.filename != nil {
		fields = append(fields, document.FieldFilename)
	}
	if m.created_at != nil {
		fields = append(fields, document.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, document.FieldUpdatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *DocumentMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case document.FieldFilename:
		return m.Filename()
	case document.FieldCreatedAt:
		return m.CreatedAt()
	case document.FieldUpdatedAt:
		return m.UpdatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *DocumentMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case document.FieldFilename:
		return m.OldFilename(ctx)
	case document.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case document.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Document field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *DocumentMutation) SetField(name string, value ent.Value) error {
	switch name {
	case document.FieldFilename:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetFilename(v)
		return nil
	case document.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case document.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Document field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *DocumentMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *DocumentMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *DocumentMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Document numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *DocumentMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *DocumentMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *DocumentMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Document nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *DocumentMutation) ResetField(name string) error {
	switch name {
	case document.FieldFilename:
		m.ResetFilename()
		return nil
	case document.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case document.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	}
	return fmt.Errorf("unknown Document field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *DocumentMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.posts != nil {
		edges = append(edges, document.EdgePosts)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *DocumentMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case document.EdgePosts:
		ids := make([]ent.Value, 0, len(m.posts))
		for id := range m.posts {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *DocumentMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedposts != nil {
		edges = append(edges, document.EdgePosts)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *DocumentMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case document.EdgePosts:
		ids := make([]ent.Value, 0, len(m.removedposts))
		for id := range m.removedposts {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *DocumentMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedposts {
		edges = append(edges, document.EdgePosts)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *DocumentMutation) EdgeCleared(name string) bool {
	switch name {
	case document.EdgePosts:
		return m.clearedposts
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *DocumentMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown Document unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *DocumentMutation) ResetEdge(name string) error {
	switch name {
	case document.EdgePosts:
		m.ResetPosts()
		return nil
	}
	return fmt.Errorf("unknown Document edge %s", name)
}

// PostMutation represents an operation that mutates the Post nodes in the graph.
type PostMutation struct {
	config
	op              Op
	typ             string
	id              *int
	title           *string
	img_path        *string
	created_at      *time.Time
	updated_at      *time.Time
	clearedFields   map[string]struct{}
	document        map[int]struct{}
	removeddocument map[int]struct{}
	cleareddocument bool
	done            bool
	oldValue        func(context.Context) (*Post, error)
	predicates      []predicate.Post
}

var _ ent.Mutation = (*PostMutation)(nil)

// postOption allows management of the mutation configuration using functional options.
type postOption func(*PostMutation)

// newPostMutation creates new mutation for the Post entity.
func newPostMutation(c config, op Op, opts ...postOption) *PostMutation {
	m := &PostMutation{
		config:        c,
		op:            op,
		typ:           TypePost,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withPostID sets the ID field of the mutation.
func withPostID(id int) postOption {
	return func(m *PostMutation) {
		var (
			err   error
			once  sync.Once
			value *Post
		)
		m.oldValue = func(ctx context.Context) (*Post, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Post.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withPost sets the old Post of the mutation.
func withPost(node *Post) postOption {
	return func(m *PostMutation) {
		m.oldValue = func(context.Context) (*Post, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m PostMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m PostMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *PostMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetTitle sets the "title" field.
func (m *PostMutation) SetTitle(s string) {
	m.title = &s
}

// Title returns the value of the "title" field in the mutation.
func (m *PostMutation) Title() (r string, exists bool) {
	v := m.title
	if v == nil {
		return
	}
	return *v, true
}

// OldTitle returns the old "title" field's value of the Post entity.
// If the Post object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PostMutation) OldTitle(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldTitle is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldTitle requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTitle: %w", err)
	}
	return oldValue.Title, nil
}

// ResetTitle resets all changes to the "title" field.
func (m *PostMutation) ResetTitle() {
	m.title = nil
}

// SetImgPath sets the "img_path" field.
func (m *PostMutation) SetImgPath(s string) {
	m.img_path = &s
}

// ImgPath returns the value of the "img_path" field in the mutation.
func (m *PostMutation) ImgPath() (r string, exists bool) {
	v := m.img_path
	if v == nil {
		return
	}
	return *v, true
}

// OldImgPath returns the old "img_path" field's value of the Post entity.
// If the Post object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PostMutation) OldImgPath(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldImgPath is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldImgPath requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldImgPath: %w", err)
	}
	return oldValue.ImgPath, nil
}

// ResetImgPath resets all changes to the "img_path" field.
func (m *PostMutation) ResetImgPath() {
	m.img_path = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *PostMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *PostMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Post entity.
// If the Post object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PostMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *PostMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *PostMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *PostMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Post entity.
// If the Post object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *PostMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *PostMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// AddDocumentIDs adds the "document" edge to the Document entity by ids.
func (m *PostMutation) AddDocumentIDs(ids ...int) {
	if m.document == nil {
		m.document = make(map[int]struct{})
	}
	for i := range ids {
		m.document[ids[i]] = struct{}{}
	}
}

// ClearDocument clears the "document" edge to the Document entity.
func (m *PostMutation) ClearDocument() {
	m.cleareddocument = true
}

// DocumentCleared reports if the "document" edge to the Document entity was cleared.
func (m *PostMutation) DocumentCleared() bool {
	return m.cleareddocument
}

// RemoveDocumentIDs removes the "document" edge to the Document entity by IDs.
func (m *PostMutation) RemoveDocumentIDs(ids ...int) {
	if m.removeddocument == nil {
		m.removeddocument = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.document, ids[i])
		m.removeddocument[ids[i]] = struct{}{}
	}
}

// RemovedDocument returns the removed IDs of the "document" edge to the Document entity.
func (m *PostMutation) RemovedDocumentIDs() (ids []int) {
	for id := range m.removeddocument {
		ids = append(ids, id)
	}
	return
}

// DocumentIDs returns the "document" edge IDs in the mutation.
func (m *PostMutation) DocumentIDs() (ids []int) {
	for id := range m.document {
		ids = append(ids, id)
	}
	return
}

// ResetDocument resets all changes to the "document" edge.
func (m *PostMutation) ResetDocument() {
	m.document = nil
	m.cleareddocument = false
	m.removeddocument = nil
}

// Where appends a list predicates to the PostMutation builder.
func (m *PostMutation) Where(ps ...predicate.Post) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *PostMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Post).
func (m *PostMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *PostMutation) Fields() []string {
	fields := make([]string, 0, 4)
	if m.title != nil {
		fields = append(fields, post.FieldTitle)
	}
	if m.img_path != nil {
		fields = append(fields, post.FieldImgPath)
	}
	if m.created_at != nil {
		fields = append(fields, post.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, post.FieldUpdatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *PostMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case post.FieldTitle:
		return m.Title()
	case post.FieldImgPath:
		return m.ImgPath()
	case post.FieldCreatedAt:
		return m.CreatedAt()
	case post.FieldUpdatedAt:
		return m.UpdatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *PostMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case post.FieldTitle:
		return m.OldTitle(ctx)
	case post.FieldImgPath:
		return m.OldImgPath(ctx)
	case post.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case post.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Post field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *PostMutation) SetField(name string, value ent.Value) error {
	switch name {
	case post.FieldTitle:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTitle(v)
		return nil
	case post.FieldImgPath:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetImgPath(v)
		return nil
	case post.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case post.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Post field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *PostMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *PostMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *PostMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Post numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *PostMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *PostMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *PostMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Post nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *PostMutation) ResetField(name string) error {
	switch name {
	case post.FieldTitle:
		m.ResetTitle()
		return nil
	case post.FieldImgPath:
		m.ResetImgPath()
		return nil
	case post.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case post.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	}
	return fmt.Errorf("unknown Post field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *PostMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.document != nil {
		edges = append(edges, post.EdgeDocument)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *PostMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case post.EdgeDocument:
		ids := make([]ent.Value, 0, len(m.document))
		for id := range m.document {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *PostMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removeddocument != nil {
		edges = append(edges, post.EdgeDocument)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *PostMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case post.EdgeDocument:
		ids := make([]ent.Value, 0, len(m.removeddocument))
		for id := range m.removeddocument {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *PostMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.cleareddocument {
		edges = append(edges, post.EdgeDocument)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *PostMutation) EdgeCleared(name string) bool {
	switch name {
	case post.EdgeDocument:
		return m.cleareddocument
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *PostMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown Post unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *PostMutation) ResetEdge(name string) error {
	switch name {
	case post.EdgeDocument:
		m.ResetDocument()
		return nil
	}
	return fmt.Errorf("unknown Post edge %s", name)
}
