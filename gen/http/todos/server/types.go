// Code generated by goa v3.13.2, DO NOT EDIT.
//
// todos HTTP server types
//
// Command:
// $ goa gen github.com/akm/goa_todo_example/design

package server

import (
	todos "github.com/akm/goa_todo_example/gen/todos"
	todosviews "github.com/akm/goa_todo_example/gen/todos/views"
	goa "goa.design/goa/v3/pkg"
)

// CreateRequestBody is the type of the "todos" service "create" endpoint HTTP
// request body.
type CreateRequestBody struct {
	// Title
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// State
	State *string `form:"State,omitempty" json:"State,omitempty" xml:"State,omitempty"`
}

// UpdateRequestBody is the type of the "todos" service "update" endpoint HTTP
// request body.
type UpdateRequestBody struct {
	// Title
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// State
	State *string `form:"State,omitempty" json:"State,omitempty" xml:"State,omitempty"`
}

// ListResponseBody is the type of the "todos" service "list" endpoint HTTP
// response body.
type ListResponseBody struct {
	// Items
	Items TodoListItemResponseBodyCollection `form:"items" json:"items" xml:"items"`
	// Total number of items
	Total uint64 `form:"total" json:"total" xml:"total"`
	// Offset
	Offset uint64 `form:"offset" json:"offset" xml:"offset"`
}

// ShowResponseBody is the type of the "todos" service "show" endpoint HTTP
// response body.
type ShowResponseBody struct {
	// ID
	ID uint64 `form:"id" json:"id" xml:"id"`
	// Title
	Title string `form:"title" json:"title" xml:"title"`
	// State
	State string `form:"State" json:"State" xml:"State"`
	// CreatedAt
	CreatedAt string `form:"created_at" json:"created_at" xml:"created_at"`
	// UpdatedAt
	UpdatedAt string `form:"updated_at" json:"updated_at" xml:"updated_at"`
}

// CreateResponseBody is the type of the "todos" service "create" endpoint HTTP
// response body.
type CreateResponseBody struct {
	// ID
	ID uint64 `form:"id" json:"id" xml:"id"`
	// Title
	Title string `form:"title" json:"title" xml:"title"`
	// State
	State string `form:"State" json:"State" xml:"State"`
	// CreatedAt
	CreatedAt string `form:"created_at" json:"created_at" xml:"created_at"`
	// UpdatedAt
	UpdatedAt string `form:"updated_at" json:"updated_at" xml:"updated_at"`
}

// UpdateResponseBody is the type of the "todos" service "update" endpoint HTTP
// response body.
type UpdateResponseBody struct {
	// ID
	ID uint64 `form:"id" json:"id" xml:"id"`
	// Title
	Title string `form:"title" json:"title" xml:"title"`
	// State
	State string `form:"State" json:"State" xml:"State"`
	// CreatedAt
	CreatedAt string `form:"created_at" json:"created_at" xml:"created_at"`
	// UpdatedAt
	UpdatedAt string `form:"updated_at" json:"updated_at" xml:"updated_at"`
}

// DeleteResponseBody is the type of the "todos" service "delete" endpoint HTTP
// response body.
type DeleteResponseBody struct {
	// ID
	ID uint64 `form:"id" json:"id" xml:"id"`
	// Title
	Title string `form:"title" json:"title" xml:"title"`
	// State
	State string `form:"State" json:"State" xml:"State"`
	// CreatedAt
	CreatedAt string `form:"created_at" json:"created_at" xml:"created_at"`
	// UpdatedAt
	UpdatedAt string `form:"updated_at" json:"updated_at" xml:"updated_at"`
}

// TodoListItemResponseBodyCollection is used to define fields on response body
// types.
type TodoListItemResponseBodyCollection []*TodoListItemResponseBody

// TodoListItemResponseBody is used to define fields on response body types.
type TodoListItemResponseBody struct {
	// ID
	ID uint64 `form:"id" json:"id" xml:"id"`
	// Title
	Title string `form:"title" json:"title" xml:"title"`
	// State
	State string `form:"State" json:"State" xml:"State"`
	// CreatedAt
	CreatedAt string `form:"created_at" json:"created_at" xml:"created_at"`
	// UpdatedAt
	UpdatedAt string `form:"updated_at" json:"updated_at" xml:"updated_at"`
}

// NewListResponseBody builds the HTTP response body from the result of the
// "list" endpoint of the "todos" service.
func NewListResponseBody(res *todosviews.TodoListView) *ListResponseBody {
	body := &ListResponseBody{
		Total:  *res.Total,
		Offset: *res.Offset,
	}
	if res.Items != nil {
		body.Items = make([]*TodoListItemResponseBody, len(res.Items))
		for i, val := range res.Items {
			body.Items[i] = marshalTodosviewsTodoListItemViewToTodoListItemResponseBody(val)
		}
	} else {
		body.Items = []*TodoListItemResponseBody{}
	}
	return body
}

// NewShowResponseBody builds the HTTP response body from the result of the
// "show" endpoint of the "todos" service.
func NewShowResponseBody(res *todosviews.TodoView) *ShowResponseBody {
	body := &ShowResponseBody{
		ID:        *res.ID,
		Title:     *res.Title,
		State:     *res.State,
		CreatedAt: *res.CreatedAt,
		UpdatedAt: *res.UpdatedAt,
	}
	return body
}

// NewCreateResponseBody builds the HTTP response body from the result of the
// "create" endpoint of the "todos" service.
func NewCreateResponseBody(res *todosviews.TodoView) *CreateResponseBody {
	body := &CreateResponseBody{
		ID:        *res.ID,
		Title:     *res.Title,
		State:     *res.State,
		CreatedAt: *res.CreatedAt,
		UpdatedAt: *res.UpdatedAt,
	}
	return body
}

// NewUpdateResponseBody builds the HTTP response body from the result of the
// "update" endpoint of the "todos" service.
func NewUpdateResponseBody(res *todosviews.TodoView) *UpdateResponseBody {
	body := &UpdateResponseBody{
		ID:        *res.ID,
		Title:     *res.Title,
		State:     *res.State,
		CreatedAt: *res.CreatedAt,
		UpdatedAt: *res.UpdatedAt,
	}
	return body
}

// NewDeleteResponseBody builds the HTTP response body from the result of the
// "delete" endpoint of the "todos" service.
func NewDeleteResponseBody(res *todosviews.TodoView) *DeleteResponseBody {
	body := &DeleteResponseBody{
		ID:        *res.ID,
		Title:     *res.Title,
		State:     *res.State,
		CreatedAt: *res.CreatedAt,
		UpdatedAt: *res.UpdatedAt,
	}
	return body
}

// NewShowPayload builds a todos service show endpoint payload.
func NewShowPayload(id uint64) *todos.ShowPayload {
	v := &todos.ShowPayload{}
	v.ID = id

	return v
}

// NewCreateTodoCreatePayload builds a todos service create endpoint payload.
func NewCreateTodoCreatePayload(body *CreateRequestBody) *todos.TodoCreatePayload {
	v := &todos.TodoCreatePayload{
		Title: *body.Title,
		State: *body.State,
	}

	return v
}

// NewUpdateTodoUpdatePayload builds a todos service update endpoint payload.
func NewUpdateTodoUpdatePayload(body *UpdateRequestBody, id uint64) *todos.TodoUpdatePayload {
	v := &todos.TodoUpdatePayload{
		Title: *body.Title,
		State: *body.State,
	}
	v.ID = id

	return v
}

// NewDeletePayload builds a todos service delete endpoint payload.
func NewDeletePayload(id uint64) *todos.DeletePayload {
	v := &todos.DeletePayload{}
	v.ID = id

	return v
}

// ValidateCreateRequestBody runs the validations defined on CreateRequestBody
func ValidateCreateRequestBody(body *CreateRequestBody) (err error) {
	if body.Title == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("title", "body"))
	}
	if body.State == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("State", "body"))
	}
	if body.State != nil {
		if !(*body.State == "open" || *body.State == "closed") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.State", *body.State, []any{"open", "closed"}))
		}
	}
	return
}

// ValidateUpdateRequestBody runs the validations defined on UpdateRequestBody
func ValidateUpdateRequestBody(body *UpdateRequestBody) (err error) {
	if body.Title == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("title", "body"))
	}
	if body.State == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("State", "body"))
	}
	if body.State != nil {
		if !(*body.State == "open" || *body.State == "closed") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.State", *body.State, []any{"open", "closed"}))
		}
	}
	return
}
