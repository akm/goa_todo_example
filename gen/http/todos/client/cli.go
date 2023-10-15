// Code generated by goa v3.13.2, DO NOT EDIT.
//
// todos HTTP client CLI support package
//
// Command:
// $ goa gen github.com/akm/goa_todo_example/design

package client

import (
	"encoding/json"
	"fmt"
	"strconv"

	todos "github.com/akm/goa_todo_example/gen/todos"
	goa "goa.design/goa/v3/pkg"
)

// BuildShowPayload builds the payload for the todos show endpoint from CLI
// flags.
func BuildShowPayload(todosShowID string) (*todos.ShowPayload, error) {
	var err error
	var id uint64
	{
		id, err = strconv.ParseUint(todosShowID, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be UINT64")
		}
	}
	v := &todos.ShowPayload{}
	v.ID = id

	return v, nil
}

// BuildCreatePayload builds the payload for the todos create endpoint from CLI
// flags.
func BuildCreatePayload(todosCreateBody string) (*todos.TodoCreatePayload, error) {
	var err error
	var body CreateRequestBody
	{
		err = json.Unmarshal([]byte(todosCreateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"State\": \"open\",\n      \"title\": \"Nesciunt itaque ratione quis id et blanditiis.\"\n   }'")
		}
		if !(body.State == "open" || body.State == "closed") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.State", body.State, []any{"open", "closed"}))
		}
		if err != nil {
			return nil, err
		}
	}
	v := &todos.TodoCreatePayload{
		Title: body.Title,
		State: body.State,
	}

	return v, nil
}

// BuildUpdatePayload builds the payload for the todos update endpoint from CLI
// flags.
func BuildUpdatePayload(todosUpdateBody string, todosUpdateID string) (*todos.UpdatePayload, error) {
	var err error
	var body UpdateRequestBody
	{
		err = json.Unmarshal([]byte(todosUpdateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"body\": {\n         \"State\": \"closed\",\n         \"title\": \"Accusamus tempora dolores facere nulla.\"\n      }\n   }'")
		}
		if body.Body == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("body", "body"))
		}
		if body.Body != nil {
			if err2 := ValidateTodoUpdatePayloadRequestBody(body.Body); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
		if err != nil {
			return nil, err
		}
	}
	var id uint64
	{
		id, err = strconv.ParseUint(todosUpdateID, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be UINT64")
		}
	}
	v := &todos.UpdatePayload{}
	if body.Body != nil {
		v.Body = marshalTodoUpdatePayloadRequestBodyToTodosTodoUpdatePayload(body.Body)
	}
	v.ID = id

	return v, nil
}

// BuildDeletePayload builds the payload for the todos delete endpoint from CLI
// flags.
func BuildDeletePayload(todosDeleteID string) (*todos.DeletePayload, error) {
	var err error
	var id uint64
	{
		id, err = strconv.ParseUint(todosDeleteID, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be UINT64")
		}
	}
	v := &todos.DeletePayload{}
	v.ID = id

	return v, nil
}
