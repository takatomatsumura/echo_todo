// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/takatomatsumura/echo_todo/ent/predicate"
	"github.com/takatomatsumura/echo_todo/ent/todo"
	"github.com/takatomatsumura/echo_todo/ent/user"

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
	TypeTodo = "Todo"
	TypeUser = "User"
)

// TodoMutation represents an operation that mutates the Todo nodes in the graph.
type TodoMutation struct {
	config
	op            Op
	typ           string
	id            *int
	title         *string
	content       *string
	todoComplete  *bool
	deadline      *time.Time
	created_at    *time.Time
	updated_at    *time.Time
	image_path    *string
	clearedFields map[string]struct{}
	owner         map[int]struct{}
	removedowner  map[int]struct{}
	clearedowner  bool
	done          bool
	oldValue      func(context.Context) (*Todo, error)
	predicates    []predicate.Todo
}

var _ ent.Mutation = (*TodoMutation)(nil)

// todoOption allows management of the mutation configuration using functional options.
type todoOption func(*TodoMutation)

// newTodoMutation creates new mutation for the Todo entity.
func newTodoMutation(c config, op Op, opts ...todoOption) *TodoMutation {
	m := &TodoMutation{
		config:        c,
		op:            op,
		typ:           TypeTodo,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withTodoID sets the ID field of the mutation.
func withTodoID(id int) todoOption {
	return func(m *TodoMutation) {
		var (
			err   error
			once  sync.Once
			value *Todo
		)
		m.oldValue = func(ctx context.Context) (*Todo, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Todo.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withTodo sets the old Todo of the mutation.
func withTodo(node *Todo) todoOption {
	return func(m *TodoMutation) {
		m.oldValue = func(context.Context) (*Todo, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m TodoMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m TodoMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *TodoMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *TodoMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Todo.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetTitle sets the "title" field.
func (m *TodoMutation) SetTitle(s string) {
	m.title = &s
}

// Title returns the value of the "title" field in the mutation.
func (m *TodoMutation) Title() (r string, exists bool) {
	v := m.title
	if v == nil {
		return
	}
	return *v, true
}

// OldTitle returns the old "title" field's value of the Todo entity.
// If the Todo object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TodoMutation) OldTitle(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldTitle is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldTitle requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTitle: %w", err)
	}
	return oldValue.Title, nil
}

// ResetTitle resets all changes to the "title" field.
func (m *TodoMutation) ResetTitle() {
	m.title = nil
}

// SetContent sets the "content" field.
func (m *TodoMutation) SetContent(s string) {
	m.content = &s
}

// Content returns the value of the "content" field in the mutation.
func (m *TodoMutation) Content() (r string, exists bool) {
	v := m.content
	if v == nil {
		return
	}
	return *v, true
}

// OldContent returns the old "content" field's value of the Todo entity.
// If the Todo object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TodoMutation) OldContent(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldContent is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldContent requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldContent: %w", err)
	}
	return oldValue.Content, nil
}

// ResetContent resets all changes to the "content" field.
func (m *TodoMutation) ResetContent() {
	m.content = nil
}

// SetTodoComplete sets the "todoComplete" field.
func (m *TodoMutation) SetTodoComplete(b bool) {
	m.todoComplete = &b
}

// TodoComplete returns the value of the "todoComplete" field in the mutation.
func (m *TodoMutation) TodoComplete() (r bool, exists bool) {
	v := m.todoComplete
	if v == nil {
		return
	}
	return *v, true
}

// OldTodoComplete returns the old "todoComplete" field's value of the Todo entity.
// If the Todo object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TodoMutation) OldTodoComplete(ctx context.Context) (v bool, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldTodoComplete is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldTodoComplete requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTodoComplete: %w", err)
	}
	return oldValue.TodoComplete, nil
}

// ResetTodoComplete resets all changes to the "todoComplete" field.
func (m *TodoMutation) ResetTodoComplete() {
	m.todoComplete = nil
}

// SetDeadline sets the "deadline" field.
func (m *TodoMutation) SetDeadline(t time.Time) {
	m.deadline = &t
}

// Deadline returns the value of the "deadline" field in the mutation.
func (m *TodoMutation) Deadline() (r time.Time, exists bool) {
	v := m.deadline
	if v == nil {
		return
	}
	return *v, true
}

// OldDeadline returns the old "deadline" field's value of the Todo entity.
// If the Todo object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TodoMutation) OldDeadline(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDeadline is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDeadline requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDeadline: %w", err)
	}
	return oldValue.Deadline, nil
}

// ClearDeadline clears the value of the "deadline" field.
func (m *TodoMutation) ClearDeadline() {
	m.deadline = nil
	m.clearedFields[todo.FieldDeadline] = struct{}{}
}

// DeadlineCleared returns if the "deadline" field was cleared in this mutation.
func (m *TodoMutation) DeadlineCleared() bool {
	_, ok := m.clearedFields[todo.FieldDeadline]
	return ok
}

// ResetDeadline resets all changes to the "deadline" field.
func (m *TodoMutation) ResetDeadline() {
	m.deadline = nil
	delete(m.clearedFields, todo.FieldDeadline)
}

// SetCreatedAt sets the "created_at" field.
func (m *TodoMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *TodoMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Todo entity.
// If the Todo object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TodoMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ClearCreatedAt clears the value of the "created_at" field.
func (m *TodoMutation) ClearCreatedAt() {
	m.created_at = nil
	m.clearedFields[todo.FieldCreatedAt] = struct{}{}
}

// CreatedAtCleared returns if the "created_at" field was cleared in this mutation.
func (m *TodoMutation) CreatedAtCleared() bool {
	_, ok := m.clearedFields[todo.FieldCreatedAt]
	return ok
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *TodoMutation) ResetCreatedAt() {
	m.created_at = nil
	delete(m.clearedFields, todo.FieldCreatedAt)
}

// SetUpdatedAt sets the "updated_at" field.
func (m *TodoMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *TodoMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Todo entity.
// If the Todo object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TodoMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (m *TodoMutation) ClearUpdatedAt() {
	m.updated_at = nil
	m.clearedFields[todo.FieldUpdatedAt] = struct{}{}
}

// UpdatedAtCleared returns if the "updated_at" field was cleared in this mutation.
func (m *TodoMutation) UpdatedAtCleared() bool {
	_, ok := m.clearedFields[todo.FieldUpdatedAt]
	return ok
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *TodoMutation) ResetUpdatedAt() {
	m.updated_at = nil
	delete(m.clearedFields, todo.FieldUpdatedAt)
}

// SetImagePath sets the "image_path" field.
func (m *TodoMutation) SetImagePath(s string) {
	m.image_path = &s
}

// ImagePath returns the value of the "image_path" field in the mutation.
func (m *TodoMutation) ImagePath() (r string, exists bool) {
	v := m.image_path
	if v == nil {
		return
	}
	return *v, true
}

// OldImagePath returns the old "image_path" field's value of the Todo entity.
// If the Todo object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TodoMutation) OldImagePath(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldImagePath is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldImagePath requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldImagePath: %w", err)
	}
	return oldValue.ImagePath, nil
}

// ClearImagePath clears the value of the "image_path" field.
func (m *TodoMutation) ClearImagePath() {
	m.image_path = nil
	m.clearedFields[todo.FieldImagePath] = struct{}{}
}

// ImagePathCleared returns if the "image_path" field was cleared in this mutation.
func (m *TodoMutation) ImagePathCleared() bool {
	_, ok := m.clearedFields[todo.FieldImagePath]
	return ok
}

// ResetImagePath resets all changes to the "image_path" field.
func (m *TodoMutation) ResetImagePath() {
	m.image_path = nil
	delete(m.clearedFields, todo.FieldImagePath)
}

// AddOwnerIDs adds the "owner" edge to the User entity by ids.
func (m *TodoMutation) AddOwnerIDs(ids ...int) {
	if m.owner == nil {
		m.owner = make(map[int]struct{})
	}
	for i := range ids {
		m.owner[ids[i]] = struct{}{}
	}
}

// ClearOwner clears the "owner" edge to the User entity.
func (m *TodoMutation) ClearOwner() {
	m.clearedowner = true
}

// OwnerCleared reports if the "owner" edge to the User entity was cleared.
func (m *TodoMutation) OwnerCleared() bool {
	return m.clearedowner
}

// RemoveOwnerIDs removes the "owner" edge to the User entity by IDs.
func (m *TodoMutation) RemoveOwnerIDs(ids ...int) {
	if m.removedowner == nil {
		m.removedowner = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.owner, ids[i])
		m.removedowner[ids[i]] = struct{}{}
	}
}

// RemovedOwner returns the removed IDs of the "owner" edge to the User entity.
func (m *TodoMutation) RemovedOwnerIDs() (ids []int) {
	for id := range m.removedowner {
		ids = append(ids, id)
	}
	return
}

// OwnerIDs returns the "owner" edge IDs in the mutation.
func (m *TodoMutation) OwnerIDs() (ids []int) {
	for id := range m.owner {
		ids = append(ids, id)
	}
	return
}

// ResetOwner resets all changes to the "owner" edge.
func (m *TodoMutation) ResetOwner() {
	m.owner = nil
	m.clearedowner = false
	m.removedowner = nil
}

// Where appends a list predicates to the TodoMutation builder.
func (m *TodoMutation) Where(ps ...predicate.Todo) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *TodoMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Todo).
func (m *TodoMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *TodoMutation) Fields() []string {
	fields := make([]string, 0, 7)
	if m.title != nil {
		fields = append(fields, todo.FieldTitle)
	}
	if m.content != nil {
		fields = append(fields, todo.FieldContent)
	}
	if m.todoComplete != nil {
		fields = append(fields, todo.FieldTodoComplete)
	}
	if m.deadline != nil {
		fields = append(fields, todo.FieldDeadline)
	}
	if m.created_at != nil {
		fields = append(fields, todo.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, todo.FieldUpdatedAt)
	}
	if m.image_path != nil {
		fields = append(fields, todo.FieldImagePath)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *TodoMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case todo.FieldTitle:
		return m.Title()
	case todo.FieldContent:
		return m.Content()
	case todo.FieldTodoComplete:
		return m.TodoComplete()
	case todo.FieldDeadline:
		return m.Deadline()
	case todo.FieldCreatedAt:
		return m.CreatedAt()
	case todo.FieldUpdatedAt:
		return m.UpdatedAt()
	case todo.FieldImagePath:
		return m.ImagePath()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *TodoMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case todo.FieldTitle:
		return m.OldTitle(ctx)
	case todo.FieldContent:
		return m.OldContent(ctx)
	case todo.FieldTodoComplete:
		return m.OldTodoComplete(ctx)
	case todo.FieldDeadline:
		return m.OldDeadline(ctx)
	case todo.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case todo.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	case todo.FieldImagePath:
		return m.OldImagePath(ctx)
	}
	return nil, fmt.Errorf("unknown Todo field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *TodoMutation) SetField(name string, value ent.Value) error {
	switch name {
	case todo.FieldTitle:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTitle(v)
		return nil
	case todo.FieldContent:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetContent(v)
		return nil
	case todo.FieldTodoComplete:
		v, ok := value.(bool)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTodoComplete(v)
		return nil
	case todo.FieldDeadline:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDeadline(v)
		return nil
	case todo.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case todo.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	case todo.FieldImagePath:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetImagePath(v)
		return nil
	}
	return fmt.Errorf("unknown Todo field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *TodoMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *TodoMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *TodoMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Todo numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *TodoMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(todo.FieldDeadline) {
		fields = append(fields, todo.FieldDeadline)
	}
	if m.FieldCleared(todo.FieldCreatedAt) {
		fields = append(fields, todo.FieldCreatedAt)
	}
	if m.FieldCleared(todo.FieldUpdatedAt) {
		fields = append(fields, todo.FieldUpdatedAt)
	}
	if m.FieldCleared(todo.FieldImagePath) {
		fields = append(fields, todo.FieldImagePath)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *TodoMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *TodoMutation) ClearField(name string) error {
	switch name {
	case todo.FieldDeadline:
		m.ClearDeadline()
		return nil
	case todo.FieldCreatedAt:
		m.ClearCreatedAt()
		return nil
	case todo.FieldUpdatedAt:
		m.ClearUpdatedAt()
		return nil
	case todo.FieldImagePath:
		m.ClearImagePath()
		return nil
	}
	return fmt.Errorf("unknown Todo nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *TodoMutation) ResetField(name string) error {
	switch name {
	case todo.FieldTitle:
		m.ResetTitle()
		return nil
	case todo.FieldContent:
		m.ResetContent()
		return nil
	case todo.FieldTodoComplete:
		m.ResetTodoComplete()
		return nil
	case todo.FieldDeadline:
		m.ResetDeadline()
		return nil
	case todo.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case todo.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	case todo.FieldImagePath:
		m.ResetImagePath()
		return nil
	}
	return fmt.Errorf("unknown Todo field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *TodoMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.owner != nil {
		edges = append(edges, todo.EdgeOwner)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *TodoMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case todo.EdgeOwner:
		ids := make([]ent.Value, 0, len(m.owner))
		for id := range m.owner {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *TodoMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedowner != nil {
		edges = append(edges, todo.EdgeOwner)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *TodoMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case todo.EdgeOwner:
		ids := make([]ent.Value, 0, len(m.removedowner))
		for id := range m.removedowner {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *TodoMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedowner {
		edges = append(edges, todo.EdgeOwner)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *TodoMutation) EdgeCleared(name string) bool {
	switch name {
	case todo.EdgeOwner:
		return m.clearedowner
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *TodoMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown Todo unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *TodoMutation) ResetEdge(name string) error {
	switch name {
	case todo.EdgeOwner:
		m.ResetOwner()
		return nil
	}
	return fmt.Errorf("unknown Todo edge %s", name)
}

// UserMutation represents an operation that mutates the User nodes in the graph.
type UserMutation struct {
	config
	op            Op
	typ           string
	id            *int
	myuuid        *string
	name          *string
	clearedFields map[string]struct{}
	todos         map[int]struct{}
	removedtodos  map[int]struct{}
	clearedtodos  bool
	done          bool
	oldValue      func(context.Context) (*User, error)
	predicates    []predicate.User
}

var _ ent.Mutation = (*UserMutation)(nil)

// userOption allows management of the mutation configuration using functional options.
type userOption func(*UserMutation)

// newUserMutation creates new mutation for the User entity.
func newUserMutation(c config, op Op, opts ...userOption) *UserMutation {
	m := &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserID sets the ID field of the mutation.
func withUserID(id int) userOption {
	return func(m *UserMutation) {
		var (
			err   error
			once  sync.Once
			value *User
		)
		m.oldValue = func(ctx context.Context) (*User, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().User.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUser sets the old User of the mutation.
func withUser(node *User) userOption {
	return func(m *UserMutation) {
		m.oldValue = func(context.Context) (*User, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *UserMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *UserMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().User.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetMyuuid sets the "myuuid" field.
func (m *UserMutation) SetMyuuid(s string) {
	m.myuuid = &s
}

// Myuuid returns the value of the "myuuid" field in the mutation.
func (m *UserMutation) Myuuid() (r string, exists bool) {
	v := m.myuuid
	if v == nil {
		return
	}
	return *v, true
}

// OldMyuuid returns the old "myuuid" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldMyuuid(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldMyuuid is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldMyuuid requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldMyuuid: %w", err)
	}
	return oldValue.Myuuid, nil
}

// ResetMyuuid resets all changes to the "myuuid" field.
func (m *UserMutation) ResetMyuuid() {
	m.myuuid = nil
}

// SetName sets the "name" field.
func (m *UserMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *UserMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *UserMutation) ResetName() {
	m.name = nil
}

// AddTodoIDs adds the "todos" edge to the Todo entity by ids.
func (m *UserMutation) AddTodoIDs(ids ...int) {
	if m.todos == nil {
		m.todos = make(map[int]struct{})
	}
	for i := range ids {
		m.todos[ids[i]] = struct{}{}
	}
}

// ClearTodos clears the "todos" edge to the Todo entity.
func (m *UserMutation) ClearTodos() {
	m.clearedtodos = true
}

// TodosCleared reports if the "todos" edge to the Todo entity was cleared.
func (m *UserMutation) TodosCleared() bool {
	return m.clearedtodos
}

// RemoveTodoIDs removes the "todos" edge to the Todo entity by IDs.
func (m *UserMutation) RemoveTodoIDs(ids ...int) {
	if m.removedtodos == nil {
		m.removedtodos = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.todos, ids[i])
		m.removedtodos[ids[i]] = struct{}{}
	}
}

// RemovedTodos returns the removed IDs of the "todos" edge to the Todo entity.
func (m *UserMutation) RemovedTodosIDs() (ids []int) {
	for id := range m.removedtodos {
		ids = append(ids, id)
	}
	return
}

// TodosIDs returns the "todos" edge IDs in the mutation.
func (m *UserMutation) TodosIDs() (ids []int) {
	for id := range m.todos {
		ids = append(ids, id)
	}
	return
}

// ResetTodos resets all changes to the "todos" edge.
func (m *UserMutation) ResetTodos() {
	m.todos = nil
	m.clearedtodos = false
	m.removedtodos = nil
}

// Where appends a list predicates to the UserMutation builder.
func (m *UserMutation) Where(ps ...predicate.User) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.myuuid != nil {
		fields = append(fields, user.FieldMyuuid)
	}
	if m.name != nil {
		fields = append(fields, user.FieldName)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldMyuuid:
		return m.Myuuid()
	case user.FieldName:
		return m.Name()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case user.FieldMyuuid:
		return m.OldMyuuid(ctx)
	case user.FieldName:
		return m.OldName(ctx)
	}
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldMyuuid:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetMyuuid(v)
		return nil
	case user.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UserMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UserMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldMyuuid:
		m.ResetMyuuid()
		return nil
	case user.FieldName:
		m.ResetName()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.todos != nil {
		edges = append(edges, user.EdgeTodos)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeTodos:
		ids := make([]ent.Value, 0, len(m.todos))
		for id := range m.todos {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedtodos != nil {
		edges = append(edges, user.EdgeTodos)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeTodos:
		ids := make([]ent.Value, 0, len(m.removedtodos))
		for id := range m.removedtodos {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedtodos {
		edges = append(edges, user.EdgeTodos)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	switch name {
	case user.EdgeTodos:
		return m.clearedtodos
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	switch name {
	case user.EdgeTodos:
		m.ResetTodos()
		return nil
	}
	return fmt.Errorf("unknown User edge %s", name)
}