// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tenkoh/exsql/ent/document"
	"github.com/tenkoh/exsql/ent/post"
	"github.com/tenkoh/exsql/ent/predicate"
)

// DocumentUpdate is the builder for updating Document entities.
type DocumentUpdate struct {
	config
	hooks    []Hook
	mutation *DocumentMutation
}

// Where appends a list predicates to the DocumentUpdate builder.
func (du *DocumentUpdate) Where(ps ...predicate.Document) *DocumentUpdate {
	du.mutation.Where(ps...)
	return du
}

// SetFilename sets the "filename" field.
func (du *DocumentUpdate) SetFilename(s string) *DocumentUpdate {
	du.mutation.SetFilename(s)
	return du
}

// SetCreatedAt sets the "created_at" field.
func (du *DocumentUpdate) SetCreatedAt(t time.Time) *DocumentUpdate {
	du.mutation.SetCreatedAt(t)
	return du
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (du *DocumentUpdate) SetNillableCreatedAt(t *time.Time) *DocumentUpdate {
	if t != nil {
		du.SetCreatedAt(*t)
	}
	return du
}

// SetUpdatedAt sets the "updated_at" field.
func (du *DocumentUpdate) SetUpdatedAt(t time.Time) *DocumentUpdate {
	du.mutation.SetUpdatedAt(t)
	return du
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (du *DocumentUpdate) SetNillableUpdatedAt(t *time.Time) *DocumentUpdate {
	if t != nil {
		du.SetUpdatedAt(*t)
	}
	return du
}

// AddPostIDs adds the "posts" edge to the Post entity by IDs.
func (du *DocumentUpdate) AddPostIDs(ids ...int) *DocumentUpdate {
	du.mutation.AddPostIDs(ids...)
	return du
}

// AddPosts adds the "posts" edges to the Post entity.
func (du *DocumentUpdate) AddPosts(p ...*Post) *DocumentUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return du.AddPostIDs(ids...)
}

// Mutation returns the DocumentMutation object of the builder.
func (du *DocumentUpdate) Mutation() *DocumentMutation {
	return du.mutation
}

// ClearPosts clears all "posts" edges to the Post entity.
func (du *DocumentUpdate) ClearPosts() *DocumentUpdate {
	du.mutation.ClearPosts()
	return du
}

// RemovePostIDs removes the "posts" edge to Post entities by IDs.
func (du *DocumentUpdate) RemovePostIDs(ids ...int) *DocumentUpdate {
	du.mutation.RemovePostIDs(ids...)
	return du
}

// RemovePosts removes "posts" edges to Post entities.
func (du *DocumentUpdate) RemovePosts(p ...*Post) *DocumentUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return du.RemovePostIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (du *DocumentUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(du.hooks) == 0 {
		affected, err = du.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DocumentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			du.mutation = mutation
			affected, err = du.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(du.hooks) - 1; i >= 0; i-- {
			if du.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = du.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, du.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (du *DocumentUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DocumentUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DocumentUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

func (du *DocumentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   document.Table,
			Columns: document.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: document.FieldID,
			},
		},
	}
	if ps := du.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.Filename(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: document.FieldFilename,
		})
	}
	if value, ok := du.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: document.FieldCreatedAt,
		})
	}
	if value, ok := du.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: document.FieldUpdatedAt,
		})
	}
	if du.mutation.PostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   document.PostsTable,
			Columns: document.PostsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: post.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.RemovedPostsIDs(); len(nodes) > 0 && !du.mutation.PostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   document.PostsTable,
			Columns: document.PostsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: post.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.PostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   document.PostsTable,
			Columns: document.PostsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: post.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{document.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// DocumentUpdateOne is the builder for updating a single Document entity.
type DocumentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DocumentMutation
}

// SetFilename sets the "filename" field.
func (duo *DocumentUpdateOne) SetFilename(s string) *DocumentUpdateOne {
	duo.mutation.SetFilename(s)
	return duo
}

// SetCreatedAt sets the "created_at" field.
func (duo *DocumentUpdateOne) SetCreatedAt(t time.Time) *DocumentUpdateOne {
	duo.mutation.SetCreatedAt(t)
	return duo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (duo *DocumentUpdateOne) SetNillableCreatedAt(t *time.Time) *DocumentUpdateOne {
	if t != nil {
		duo.SetCreatedAt(*t)
	}
	return duo
}

// SetUpdatedAt sets the "updated_at" field.
func (duo *DocumentUpdateOne) SetUpdatedAt(t time.Time) *DocumentUpdateOne {
	duo.mutation.SetUpdatedAt(t)
	return duo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (duo *DocumentUpdateOne) SetNillableUpdatedAt(t *time.Time) *DocumentUpdateOne {
	if t != nil {
		duo.SetUpdatedAt(*t)
	}
	return duo
}

// AddPostIDs adds the "posts" edge to the Post entity by IDs.
func (duo *DocumentUpdateOne) AddPostIDs(ids ...int) *DocumentUpdateOne {
	duo.mutation.AddPostIDs(ids...)
	return duo
}

// AddPosts adds the "posts" edges to the Post entity.
func (duo *DocumentUpdateOne) AddPosts(p ...*Post) *DocumentUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return duo.AddPostIDs(ids...)
}

// Mutation returns the DocumentMutation object of the builder.
func (duo *DocumentUpdateOne) Mutation() *DocumentMutation {
	return duo.mutation
}

// ClearPosts clears all "posts" edges to the Post entity.
func (duo *DocumentUpdateOne) ClearPosts() *DocumentUpdateOne {
	duo.mutation.ClearPosts()
	return duo
}

// RemovePostIDs removes the "posts" edge to Post entities by IDs.
func (duo *DocumentUpdateOne) RemovePostIDs(ids ...int) *DocumentUpdateOne {
	duo.mutation.RemovePostIDs(ids...)
	return duo
}

// RemovePosts removes "posts" edges to Post entities.
func (duo *DocumentUpdateOne) RemovePosts(p ...*Post) *DocumentUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return duo.RemovePostIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (duo *DocumentUpdateOne) Select(field string, fields ...string) *DocumentUpdateOne {
	duo.fields = append([]string{field}, fields...)
	return duo
}

// Save executes the query and returns the updated Document entity.
func (duo *DocumentUpdateOne) Save(ctx context.Context) (*Document, error) {
	var (
		err  error
		node *Document
	)
	if len(duo.hooks) == 0 {
		node, err = duo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DocumentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			duo.mutation = mutation
			node, err = duo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(duo.hooks) - 1; i >= 0; i-- {
			if duo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = duo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, duo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DocumentUpdateOne) SaveX(ctx context.Context) *Document {
	node, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (duo *DocumentUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DocumentUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (duo *DocumentUpdateOne) sqlSave(ctx context.Context) (_node *Document, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   document.Table,
			Columns: document.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: document.FieldID,
			},
		},
	}
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Document.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := duo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, document.FieldID)
		for _, f := range fields {
			if !document.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != document.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := duo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := duo.mutation.Filename(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: document.FieldFilename,
		})
	}
	if value, ok := duo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: document.FieldCreatedAt,
		})
	}
	if value, ok := duo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: document.FieldUpdatedAt,
		})
	}
	if duo.mutation.PostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   document.PostsTable,
			Columns: document.PostsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: post.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.RemovedPostsIDs(); len(nodes) > 0 && !duo.mutation.PostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   document.PostsTable,
			Columns: document.PostsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: post.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.PostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   document.PostsTable,
			Columns: document.PostsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: post.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Document{config: duo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{document.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
