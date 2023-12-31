// Code generated by goa v3.13.2, DO NOT EDIT.
//
// todos views
//
// Command:
// $ goa gen github.com/akm/goa_todo_example/design

package views

import (
	goa "goa.design/goa/v3/pkg"
)

// TodoList is the viewed result type that is projected based on a view.
type TodoList struct {
	// Type to project
	Projected *TodoListView
	// View to render
	View string
}

// Todo is the viewed result type that is projected based on a view.
type Todo struct {
	// Type to project
	Projected *TodoView
	// View to render
	View string
}

// TodoListView is a type that runs validations on a projected type.
type TodoListView struct {
	// Items
	Items TodoListItemCollectionView
	// Total number of items
	Total *uint64
	// Offset
	Offset *uint64
}

// TodoListItemCollectionView is a type that runs validations on a projected
// type.
type TodoListItemCollectionView []*TodoListItemView

// TodoListItemView is a type that runs validations on a projected type.
type TodoListItemView struct {
	// ID
	ID *uint64
	// Title
	Title *string
	// State
	State *string
	// CreatedAt
	CreatedAt *string
	// UpdatedAt
	UpdatedAt *string
}

// TodoView is a type that runs validations on a projected type.
type TodoView struct {
	// ID
	ID *uint64
	// Title
	Title *string
	// State
	State *string
	// CreatedAt
	CreatedAt *string
	// UpdatedAt
	UpdatedAt *string
}

var (
	// TodoListMap is a map indexing the attribute names of TodoList by view name.
	TodoListMap = map[string][]string{
		"default": {
			"items",
			"total",
			"offset",
		},
	}
	// TodoMap is a map indexing the attribute names of Todo by view name.
	TodoMap = map[string][]string{
		"default": {
			"id",
			"title",
			"State",
			"created_at",
			"updated_at",
		},
	}
	// TodoListItemCollectionMap is a map indexing the attribute names of
	// TodoListItemCollection by view name.
	TodoListItemCollectionMap = map[string][]string{
		"default": {
			"id",
			"title",
			"State",
			"created_at",
			"updated_at",
		},
	}
	// TodoListItemMap is a map indexing the attribute names of TodoListItem by
	// view name.
	TodoListItemMap = map[string][]string{
		"default": {
			"id",
			"title",
			"State",
			"created_at",
			"updated_at",
		},
	}
)

// ValidateTodoList runs the validations defined on the viewed result type
// TodoList.
func ValidateTodoList(result *TodoList) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateTodoListView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []any{"default"})
	}
	return
}

// ValidateTodo runs the validations defined on the viewed result type Todo.
func ValidateTodo(result *Todo) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateTodoView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []any{"default"})
	}
	return
}

// ValidateTodoListView runs the validations defined on TodoListView using the
// "default" view.
func ValidateTodoListView(result *TodoListView) (err error) {
	if result.Total == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("total", "result"))
	}
	if result.Offset == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("offset", "result"))
	}
	if result.Items != nil {
		if err2 := ValidateTodoListItemCollectionView(result.Items); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateTodoListItemCollectionView runs the validations defined on
// TodoListItemCollectionView using the "default" view.
func ValidateTodoListItemCollectionView(result TodoListItemCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateTodoListItemView(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateTodoListItemView runs the validations defined on TodoListItemView
// using the "default" view.
func ValidateTodoListItemView(result *TodoListItemView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Title == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("title", "result"))
	}
	if result.State == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("State", "result"))
	}
	if result.CreatedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("created_at", "result"))
	}
	if result.UpdatedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("updated_at", "result"))
	}
	if result.State != nil {
		if !(*result.State == "open" || *result.State == "closed") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.State", *result.State, []any{"open", "closed"}))
		}
	}
	if result.CreatedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.created_at", *result.CreatedAt, goa.FormatDateTime))
	}
	if result.UpdatedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.updated_at", *result.UpdatedAt, goa.FormatDateTime))
	}
	return
}

// ValidateTodoView runs the validations defined on TodoView using the
// "default" view.
func ValidateTodoView(result *TodoView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Title == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("title", "result"))
	}
	if result.State == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("State", "result"))
	}
	if result.CreatedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("created_at", "result"))
	}
	if result.UpdatedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("updated_at", "result"))
	}
	if result.State != nil {
		if !(*result.State == "open" || *result.State == "closed") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.State", *result.State, []any{"open", "closed"}))
		}
	}
	if result.CreatedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.created_at", *result.CreatedAt, goa.FormatDateTime))
	}
	if result.UpdatedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.updated_at", *result.UpdatedAt, goa.FormatDateTime))
	}
	return
}
