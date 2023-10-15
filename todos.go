package todoapi

import (
	"context"
	"log"

	todos "github.com/akm/goa_todo_example/gen/todos"
)

// todos service example implementation.
// The example methods log the requests and return zero values.
type todossrvc struct {
	logger *log.Logger
}

// NewTodos returns the todos service implementation.
func NewTodos(logger *log.Logger) todos.Service {
	return &todossrvc{logger}
}

// List implements list.
func (s *todossrvc) List(ctx context.Context) (res *todos.TodoList, err error) {
	res = &todos.TodoList{}
	s.logger.Print("todos.list")
	return
}

// Show implements show.
func (s *todossrvc) Show(ctx context.Context, p *todos.ShowPayload) (res *todos.Todo, err error) {
	res = &todos.Todo{}
	s.logger.Print("todos.show")
	return
}

// Create implements create.
func (s *todossrvc) Create(ctx context.Context, p *todos.TodoCreatePayload) (res *todos.Todo, err error) {
	res = &todos.Todo{}
	s.logger.Print("todos.create")
	return
}

// Update implements update.
func (s *todossrvc) Update(ctx context.Context, p *todos.UpdatePayload) (res *todos.Todo, err error) {
	res = &todos.Todo{}
	s.logger.Print("todos.update")
	return
}

// Delete implements delete.
func (s *todossrvc) Delete(ctx context.Context, p *todos.DeletePayload) (res *todos.Todo, err error) {
	res = &todos.Todo{}
	s.logger.Print("todos.delete")
	return
}
