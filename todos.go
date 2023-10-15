package todoapi

import (
	"context"
	"errors"
	"log"
	"time"

	todos "github.com/akm/goa_todo_example/gen/todos"
	"github.com/akm/goa_todo_example/models/todo"
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
	s.logger.Print("todos.list")

	loadedItems, err := todo.Load()
	if err != nil {
		return nil, err
	}

	items := make(todos.TodoListItemCollection, len(loadedItems))
	for i, t := range loadedItems {
		items[i] = &todos.TodoListItem{
			ID:        t.ID,
			Title:     t.Title,
			State:     t.State.String(),
			CreatedAt: t.CreatedAt.Format(time.RFC3339),
			UpdatedAt: t.UpdatedAt.Format(time.RFC3339),
		}
	}

	res = &todos.TodoList{
		Items:  items,
		Total:  uint64(len(items)),
		Offset: 0,
	}
	return
}

// Show implements show.
func (s *todossrvc) Show(ctx context.Context, p *todos.ShowPayload) (res *todos.Todo, err error) {
	s.logger.Print("todos.show")
	loadedItems, err := todo.Load()
	if err != nil {
		return nil, err
	}

	m := loadedItems.Find(p.ID)
	if m == nil {
		// return nil, todos.MakeNotFound(errors.New("not found"))
		return nil, errors.New("not found")
	}

	res = s.modelToResult(m)
	return
}

// Create implements create.
func (s *todossrvc) Create(ctx context.Context, p *todos.TodoCreatePayload) (res *todos.Todo, err error) {
	s.logger.Print("todos.create")

	loadedItems, err := todo.Load()
	if err != nil {
		return nil, err
	}

	now := time.Now()
	m := &todo.Todo{
		ID:        loadedItems.LastID() + 1,
		Title:     p.Title,
		State:     todo.StateOpen,
		CreatedAt: now,
		UpdatedAt: now,
	}
	loadedItems = append(loadedItems, m)

	if err := todo.Save(loadedItems); err != nil {
		return nil, err
	}

	res = s.modelToResult(m)
	return
}

// Update implements update.
func (s *todossrvc) Update(ctx context.Context, p *todos.UpdatePayload) (res *todos.Todo, err error) {
	s.logger.Print("todos.update")

	st, err := todo.ParseState(p.Body.State)
	if err != nil {
		// return nil, todos.MakeInvalidPayload(errors.New("invalid state"))
		return nil, err
	}

	loadedItems, err := todo.Load()
	if err != nil {
		return nil, err
	}

	m := loadedItems.Find(p.ID)
	if m == nil {
		// return nil, todos.MakeNotFound(errors.New("not found"))
		return nil, errors.New("not found")
	}

	m.Title = p.Body.Title
	m.State = st
	m.UpdatedAt = time.Now()

	if err := todo.Save(loadedItems); err != nil {
		return nil, err
	}

	res = s.modelToResult(m)
	return
}

// Delete implements delete.
func (s *todossrvc) Delete(ctx context.Context, p *todos.DeletePayload) (res *todos.Todo, err error) {
	s.logger.Print("todos.delete")

	loadedItems, err := todo.Load()
	if err != nil {
		return nil, err
	}

	m := loadedItems.Find(p.ID)
	if m == nil {
		// return nil, todos.MakeNotFound(errors.New("not found"))
		return nil, errors.New("not found")
	}

	items := loadedItems.RemoveByID(p.ID)
	if err := todo.Save(items); err != nil {
		return nil, err
	}

	res = s.modelToResult(m)
	return
}

func (s *todossrvc) modelToResult(t *todo.Todo) *todos.Todo {
	return &todos.Todo{
		ID:        t.ID,
		Title:     t.Title,
		State:     t.State.String(),
		CreatedAt: t.CreatedAt.Format(time.RFC3339),
		UpdatedAt: t.UpdatedAt.Format(time.RFC3339),
	}
}
