// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/takatomatsumura/echo_todo/ent/todo"
	"github.com/takatomatsumura/echo_todo/ent/user"
)

// TodoCreate is the builder for creating a Todo entity.
type TodoCreate struct {
	config
	mutation *TodoMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetTitle sets the "title" field.
func (tc *TodoCreate) SetTitle(s string) *TodoCreate {
	tc.mutation.SetTitle(s)
	return tc
}

// SetContent sets the "content" field.
func (tc *TodoCreate) SetContent(s string) *TodoCreate {
	tc.mutation.SetContent(s)
	return tc
}

// SetTodoComplete sets the "todoComplete" field.
func (tc *TodoCreate) SetTodoComplete(b bool) *TodoCreate {
	tc.mutation.SetTodoComplete(b)
	return tc
}

// SetNillableTodoComplete sets the "todoComplete" field if the given value is not nil.
func (tc *TodoCreate) SetNillableTodoComplete(b *bool) *TodoCreate {
	if b != nil {
		tc.SetTodoComplete(*b)
	}
	return tc
}

// SetDeadline sets the "deadline" field.
func (tc *TodoCreate) SetDeadline(t time.Time) *TodoCreate {
	tc.mutation.SetDeadline(t)
	return tc
}

// SetNillableDeadline sets the "deadline" field if the given value is not nil.
func (tc *TodoCreate) SetNillableDeadline(t *time.Time) *TodoCreate {
	if t != nil {
		tc.SetDeadline(*t)
	}
	return tc
}

// SetCreatedAt sets the "created_at" field.
func (tc *TodoCreate) SetCreatedAt(t time.Time) *TodoCreate {
	tc.mutation.SetCreatedAt(t)
	return tc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tc *TodoCreate) SetNillableCreatedAt(t *time.Time) *TodoCreate {
	if t != nil {
		tc.SetCreatedAt(*t)
	}
	return tc
}

// SetUpdatedAt sets the "updated_at" field.
func (tc *TodoCreate) SetUpdatedAt(t time.Time) *TodoCreate {
	tc.mutation.SetUpdatedAt(t)
	return tc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tc *TodoCreate) SetNillableUpdatedAt(t *time.Time) *TodoCreate {
	if t != nil {
		tc.SetUpdatedAt(*t)
	}
	return tc
}

// SetImagePath sets the "image_path" field.
func (tc *TodoCreate) SetImagePath(s string) *TodoCreate {
	tc.mutation.SetImagePath(s)
	return tc
}

// SetNillableImagePath sets the "image_path" field if the given value is not nil.
func (tc *TodoCreate) SetNillableImagePath(s *string) *TodoCreate {
	if s != nil {
		tc.SetImagePath(*s)
	}
	return tc
}

// AddOwnerIDs adds the "owner" edge to the User entity by IDs.
func (tc *TodoCreate) AddOwnerIDs(ids ...int) *TodoCreate {
	tc.mutation.AddOwnerIDs(ids...)
	return tc
}

// AddOwner adds the "owner" edges to the User entity.
func (tc *TodoCreate) AddOwner(u ...*User) *TodoCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tc.AddOwnerIDs(ids...)
}

// Mutation returns the TodoMutation object of the builder.
func (tc *TodoCreate) Mutation() *TodoMutation {
	return tc.mutation
}

// Save creates the Todo in the database.
func (tc *TodoCreate) Save(ctx context.Context) (*Todo, error) {
	var (
		err  error
		node *Todo
	)
	tc.defaults()
	if len(tc.hooks) == 0 {
		if err = tc.check(); err != nil {
			return nil, err
		}
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TodoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tc.check(); err != nil {
				return nil, err
			}
			tc.mutation = mutation
			if node, err = tc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			if tc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TodoCreate) SaveX(ctx context.Context) *Todo {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TodoCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TodoCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TodoCreate) defaults() {
	if _, ok := tc.mutation.TodoComplete(); !ok {
		v := todo.DefaultTodoComplete
		tc.mutation.SetTodoComplete(v)
	}
	if _, ok := tc.mutation.CreatedAt(); !ok {
		v := todo.DefaultCreatedAt()
		tc.mutation.SetCreatedAt(v)
	}
	if _, ok := tc.mutation.UpdatedAt(); !ok {
		v := todo.DefaultUpdatedAt()
		tc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TodoCreate) check() error {
	if _, ok := tc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Todo.title"`)}
	}
	if _, ok := tc.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New(`ent: missing required field "Todo.content"`)}
	}
	if _, ok := tc.mutation.TodoComplete(); !ok {
		return &ValidationError{Name: "todoComplete", err: errors.New(`ent: missing required field "Todo.todoComplete"`)}
	}
	return nil
}

func (tc *TodoCreate) sqlSave(ctx context.Context) (*Todo, error) {
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (tc *TodoCreate) createSpec() (*Todo, *sqlgraph.CreateSpec) {
	var (
		_node = &Todo{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: todo.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: todo.FieldID,
			},
		}
	)
	_spec.OnConflict = tc.conflict
	if value, ok := tc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: todo.FieldTitle,
		})
		_node.Title = value
	}
	if value, ok := tc.mutation.Content(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: todo.FieldContent,
		})
		_node.Content = value
	}
	if value, ok := tc.mutation.TodoComplete(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: todo.FieldTodoComplete,
		})
		_node.TodoComplete = value
	}
	if value, ok := tc.mutation.Deadline(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: todo.FieldDeadline,
		})
		_node.Deadline = value
	}
	if value, ok := tc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: todo.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := tc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: todo.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := tc.mutation.ImagePath(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: todo.FieldImagePath,
		})
		_node.ImagePath = value
	}
	if nodes := tc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   todo.OwnerTable,
			Columns: todo.OwnerPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Todo.Create().
//		SetTitle(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TodoUpsert) {
//			SetTitle(v+v).
//		}).
//		Exec(ctx)
//
func (tc *TodoCreate) OnConflict(opts ...sql.ConflictOption) *TodoUpsertOne {
	tc.conflict = opts
	return &TodoUpsertOne{
		create: tc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Todo.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (tc *TodoCreate) OnConflictColumns(columns ...string) *TodoUpsertOne {
	tc.conflict = append(tc.conflict, sql.ConflictColumns(columns...))
	return &TodoUpsertOne{
		create: tc,
	}
}

type (
	// TodoUpsertOne is the builder for "upsert"-ing
	//  one Todo node.
	TodoUpsertOne struct {
		create *TodoCreate
	}

	// TodoUpsert is the "OnConflict" setter.
	TodoUpsert struct {
		*sql.UpdateSet
	}
)

// SetTitle sets the "title" field.
func (u *TodoUpsert) SetTitle(v string) *TodoUpsert {
	u.Set(todo.FieldTitle, v)
	return u
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *TodoUpsert) UpdateTitle() *TodoUpsert {
	u.SetExcluded(todo.FieldTitle)
	return u
}

// SetContent sets the "content" field.
func (u *TodoUpsert) SetContent(v string) *TodoUpsert {
	u.Set(todo.FieldContent, v)
	return u
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *TodoUpsert) UpdateContent() *TodoUpsert {
	u.SetExcluded(todo.FieldContent)
	return u
}

// SetTodoComplete sets the "todoComplete" field.
func (u *TodoUpsert) SetTodoComplete(v bool) *TodoUpsert {
	u.Set(todo.FieldTodoComplete, v)
	return u
}

// UpdateTodoComplete sets the "todoComplete" field to the value that was provided on create.
func (u *TodoUpsert) UpdateTodoComplete() *TodoUpsert {
	u.SetExcluded(todo.FieldTodoComplete)
	return u
}

// SetDeadline sets the "deadline" field.
func (u *TodoUpsert) SetDeadline(v time.Time) *TodoUpsert {
	u.Set(todo.FieldDeadline, v)
	return u
}

// UpdateDeadline sets the "deadline" field to the value that was provided on create.
func (u *TodoUpsert) UpdateDeadline() *TodoUpsert {
	u.SetExcluded(todo.FieldDeadline)
	return u
}

// ClearDeadline clears the value of the "deadline" field.
func (u *TodoUpsert) ClearDeadline() *TodoUpsert {
	u.SetNull(todo.FieldDeadline)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *TodoUpsert) SetCreatedAt(v time.Time) *TodoUpsert {
	u.Set(todo.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TodoUpsert) UpdateCreatedAt() *TodoUpsert {
	u.SetExcluded(todo.FieldCreatedAt)
	return u
}

// ClearCreatedAt clears the value of the "created_at" field.
func (u *TodoUpsert) ClearCreatedAt() *TodoUpsert {
	u.SetNull(todo.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TodoUpsert) SetUpdatedAt(v time.Time) *TodoUpsert {
	u.Set(todo.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TodoUpsert) UpdateUpdatedAt() *TodoUpsert {
	u.SetExcluded(todo.FieldUpdatedAt)
	return u
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *TodoUpsert) ClearUpdatedAt() *TodoUpsert {
	u.SetNull(todo.FieldUpdatedAt)
	return u
}

// SetImagePath sets the "image_path" field.
func (u *TodoUpsert) SetImagePath(v string) *TodoUpsert {
	u.Set(todo.FieldImagePath, v)
	return u
}

// UpdateImagePath sets the "image_path" field to the value that was provided on create.
func (u *TodoUpsert) UpdateImagePath() *TodoUpsert {
	u.SetExcluded(todo.FieldImagePath)
	return u
}

// ClearImagePath clears the value of the "image_path" field.
func (u *TodoUpsert) ClearImagePath() *TodoUpsert {
	u.SetNull(todo.FieldImagePath)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Todo.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *TodoUpsertOne) UpdateNewValues() *TodoUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Todo.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *TodoUpsertOne) Ignore() *TodoUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TodoUpsertOne) DoNothing() *TodoUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TodoCreate.OnConflict
// documentation for more info.
func (u *TodoUpsertOne) Update(set func(*TodoUpsert)) *TodoUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TodoUpsert{UpdateSet: update})
	}))
	return u
}

// SetTitle sets the "title" field.
func (u *TodoUpsertOne) SetTitle(v string) *TodoUpsertOne {
	return u.Update(func(s *TodoUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *TodoUpsertOne) UpdateTitle() *TodoUpsertOne {
	return u.Update(func(s *TodoUpsert) {
		s.UpdateTitle()
	})
}

// SetContent sets the "content" field.
func (u *TodoUpsertOne) SetContent(v string) *TodoUpsertOne {
	return u.Update(func(s *TodoUpsert) {
		s.SetContent(v)
	})
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *TodoUpsertOne) UpdateContent() *TodoUpsertOne {
	return u.Update(func(s *TodoUpsert) {
		s.UpdateContent()
	})
}

// SetTodoComplete sets the "todoComplete" field.
func (u *TodoUpsertOne) SetTodoComplete(v bool) *TodoUpsertOne {
	return u.Update(func(s *TodoUpsert) {
		s.SetTodoComplete(v)
	})
}

// UpdateTodoComplete sets the "todoComplete" field to the value that was provided on create.
func (u *TodoUpsertOne) UpdateTodoComplete() *TodoUpsertOne {
	return u.Update(func(s *TodoUpsert) {
		s.UpdateTodoComplete()
	})
}

// SetDeadline sets the "deadline" field.
func (u *TodoUpsertOne) SetDeadline(v time.Time) *TodoUpsertOne {
	return u.Update(func(s *TodoUpsert) {
		s.SetDeadline(v)
	})
}

// UpdateDeadline sets the "deadline" field to the value that was provided on create.
func (u *TodoUpsertOne) UpdateDeadline() *TodoUpsertOne {
	return u.Update(func(s *TodoUpsert) {
		s.UpdateDeadline()
	})
}

// ClearDeadline clears the value of the "deadline" field.
func (u *TodoUpsertOne) ClearDeadline() *TodoUpsertOne {
	return u.Update(func(s *TodoUpsert) {
		s.ClearDeadline()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *TodoUpsertOne) SetCreatedAt(v time.Time) *TodoUpsertOne {
	return u.Update(func(s *TodoUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TodoUpsertOne) UpdateCreatedAt() *TodoUpsertOne {
	return u.Update(func(s *TodoUpsert) {
		s.UpdateCreatedAt()
	})
}

// ClearCreatedAt clears the value of the "created_at" field.
func (u *TodoUpsertOne) ClearCreatedAt() *TodoUpsertOne {
	return u.Update(func(s *TodoUpsert) {
		s.ClearCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TodoUpsertOne) SetUpdatedAt(v time.Time) *TodoUpsertOne {
	return u.Update(func(s *TodoUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TodoUpsertOne) UpdateUpdatedAt() *TodoUpsertOne {
	return u.Update(func(s *TodoUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *TodoUpsertOne) ClearUpdatedAt() *TodoUpsertOne {
	return u.Update(func(s *TodoUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetImagePath sets the "image_path" field.
func (u *TodoUpsertOne) SetImagePath(v string) *TodoUpsertOne {
	return u.Update(func(s *TodoUpsert) {
		s.SetImagePath(v)
	})
}

// UpdateImagePath sets the "image_path" field to the value that was provided on create.
func (u *TodoUpsertOne) UpdateImagePath() *TodoUpsertOne {
	return u.Update(func(s *TodoUpsert) {
		s.UpdateImagePath()
	})
}

// ClearImagePath clears the value of the "image_path" field.
func (u *TodoUpsertOne) ClearImagePath() *TodoUpsertOne {
	return u.Update(func(s *TodoUpsert) {
		s.ClearImagePath()
	})
}

// Exec executes the query.
func (u *TodoUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TodoCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TodoUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *TodoUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *TodoUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// TodoCreateBulk is the builder for creating many Todo entities in bulk.
type TodoCreateBulk struct {
	config
	builders []*TodoCreate
	conflict []sql.ConflictOption
}

// Save creates the Todo entities in the database.
func (tcb *TodoCreateBulk) Save(ctx context.Context) ([]*Todo, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Todo, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TodoMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = tcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TodoCreateBulk) SaveX(ctx context.Context) []*Todo {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TodoCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TodoCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Todo.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TodoUpsert) {
//			SetTitle(v+v).
//		}).
//		Exec(ctx)
//
func (tcb *TodoCreateBulk) OnConflict(opts ...sql.ConflictOption) *TodoUpsertBulk {
	tcb.conflict = opts
	return &TodoUpsertBulk{
		create: tcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Todo.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (tcb *TodoCreateBulk) OnConflictColumns(columns ...string) *TodoUpsertBulk {
	tcb.conflict = append(tcb.conflict, sql.ConflictColumns(columns...))
	return &TodoUpsertBulk{
		create: tcb,
	}
}

// TodoUpsertBulk is the builder for "upsert"-ing
// a bulk of Todo nodes.
type TodoUpsertBulk struct {
	create *TodoCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Todo.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *TodoUpsertBulk) UpdateNewValues() *TodoUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Todo.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *TodoUpsertBulk) Ignore() *TodoUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TodoUpsertBulk) DoNothing() *TodoUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TodoCreateBulk.OnConflict
// documentation for more info.
func (u *TodoUpsertBulk) Update(set func(*TodoUpsert)) *TodoUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TodoUpsert{UpdateSet: update})
	}))
	return u
}

// SetTitle sets the "title" field.
func (u *TodoUpsertBulk) SetTitle(v string) *TodoUpsertBulk {
	return u.Update(func(s *TodoUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *TodoUpsertBulk) UpdateTitle() *TodoUpsertBulk {
	return u.Update(func(s *TodoUpsert) {
		s.UpdateTitle()
	})
}

// SetContent sets the "content" field.
func (u *TodoUpsertBulk) SetContent(v string) *TodoUpsertBulk {
	return u.Update(func(s *TodoUpsert) {
		s.SetContent(v)
	})
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *TodoUpsertBulk) UpdateContent() *TodoUpsertBulk {
	return u.Update(func(s *TodoUpsert) {
		s.UpdateContent()
	})
}

// SetTodoComplete sets the "todoComplete" field.
func (u *TodoUpsertBulk) SetTodoComplete(v bool) *TodoUpsertBulk {
	return u.Update(func(s *TodoUpsert) {
		s.SetTodoComplete(v)
	})
}

// UpdateTodoComplete sets the "todoComplete" field to the value that was provided on create.
func (u *TodoUpsertBulk) UpdateTodoComplete() *TodoUpsertBulk {
	return u.Update(func(s *TodoUpsert) {
		s.UpdateTodoComplete()
	})
}

// SetDeadline sets the "deadline" field.
func (u *TodoUpsertBulk) SetDeadline(v time.Time) *TodoUpsertBulk {
	return u.Update(func(s *TodoUpsert) {
		s.SetDeadline(v)
	})
}

// UpdateDeadline sets the "deadline" field to the value that was provided on create.
func (u *TodoUpsertBulk) UpdateDeadline() *TodoUpsertBulk {
	return u.Update(func(s *TodoUpsert) {
		s.UpdateDeadline()
	})
}

// ClearDeadline clears the value of the "deadline" field.
func (u *TodoUpsertBulk) ClearDeadline() *TodoUpsertBulk {
	return u.Update(func(s *TodoUpsert) {
		s.ClearDeadline()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *TodoUpsertBulk) SetCreatedAt(v time.Time) *TodoUpsertBulk {
	return u.Update(func(s *TodoUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TodoUpsertBulk) UpdateCreatedAt() *TodoUpsertBulk {
	return u.Update(func(s *TodoUpsert) {
		s.UpdateCreatedAt()
	})
}

// ClearCreatedAt clears the value of the "created_at" field.
func (u *TodoUpsertBulk) ClearCreatedAt() *TodoUpsertBulk {
	return u.Update(func(s *TodoUpsert) {
		s.ClearCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TodoUpsertBulk) SetUpdatedAt(v time.Time) *TodoUpsertBulk {
	return u.Update(func(s *TodoUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TodoUpsertBulk) UpdateUpdatedAt() *TodoUpsertBulk {
	return u.Update(func(s *TodoUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *TodoUpsertBulk) ClearUpdatedAt() *TodoUpsertBulk {
	return u.Update(func(s *TodoUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetImagePath sets the "image_path" field.
func (u *TodoUpsertBulk) SetImagePath(v string) *TodoUpsertBulk {
	return u.Update(func(s *TodoUpsert) {
		s.SetImagePath(v)
	})
}

// UpdateImagePath sets the "image_path" field to the value that was provided on create.
func (u *TodoUpsertBulk) UpdateImagePath() *TodoUpsertBulk {
	return u.Update(func(s *TodoUpsert) {
		s.UpdateImagePath()
	})
}

// ClearImagePath clears the value of the "image_path" field.
func (u *TodoUpsertBulk) ClearImagePath() *TodoUpsertBulk {
	return u.Update(func(s *TodoUpsert) {
		s.ClearImagePath()
	})
}

// Exec executes the query.
func (u *TodoUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the TodoCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TodoCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TodoUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
