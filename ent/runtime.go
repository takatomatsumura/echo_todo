// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/takatomatsumura/echo_todo/ent/schema"
	"github.com/takatomatsumura/echo_todo/ent/todo"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	todoFields := schema.Todo{}.Fields()
	_ = todoFields
	// todoDescTodoComplete is the schema descriptor for todoComplete field.
	todoDescTodoComplete := todoFields[2].Descriptor()
	// todo.DefaultTodoComplete holds the default value on creation for the todoComplete field.
	todo.DefaultTodoComplete = todoDescTodoComplete.Default.(bool)
	// todoDescCreatedAt is the schema descriptor for created_at field.
	todoDescCreatedAt := todoFields[4].Descriptor()
	// todo.DefaultCreatedAt holds the default value on creation for the created_at field.
	todo.DefaultCreatedAt = todoDescCreatedAt.Default.(func() time.Time)
	// todoDescUpdatedAt is the schema descriptor for updated_at field.
	todoDescUpdatedAt := todoFields[5].Descriptor()
	// todo.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	todo.DefaultUpdatedAt = todoDescUpdatedAt.Default.(func() time.Time)
	// todo.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	todo.UpdateDefaultUpdatedAt = todoDescUpdatedAt.UpdateDefault.(func() time.Time)
}
