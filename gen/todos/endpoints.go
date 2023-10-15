// Code generated by goa v3.13.2, DO NOT EDIT.
//
// todos endpoints
//
// Command:
// $ goa gen github.com/akm/goa_todo_example/design

package todos

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "todos" service endpoints.
type Endpoints struct {
	List   goa.Endpoint
	Show   goa.Endpoint
	Create goa.Endpoint
	Update goa.Endpoint
	Delete goa.Endpoint
}

// NewEndpoints wraps the methods of the "todos" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		List:   NewListEndpoint(s),
		Show:   NewShowEndpoint(s),
		Create: NewCreateEndpoint(s),
		Update: NewUpdateEndpoint(s),
		Delete: NewDeleteEndpoint(s),
	}
}

// Use applies the given middleware to all the "todos" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.List = m(e.List)
	e.Show = m(e.Show)
	e.Create = m(e.Create)
	e.Update = m(e.Update)
	e.Delete = m(e.Delete)
}

// NewListEndpoint returns an endpoint function that calls the method "list" of
// service "todos".
func NewListEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		res, err := s.List(ctx)
		if err != nil {
			return nil, err
		}
		vres := NewViewedTodoList(res, "default")
		return vres, nil
	}
}

// NewShowEndpoint returns an endpoint function that calls the method "show" of
// service "todos".
func NewShowEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*ShowPayload)
		res, err := s.Show(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedTodo(res, "default")
		return vres, nil
	}
}

// NewCreateEndpoint returns an endpoint function that calls the method
// "create" of service "todos".
func NewCreateEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*TodoCreatePayload)
		res, err := s.Create(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedTodo(res, "default")
		return vres, nil
	}
}

// NewUpdateEndpoint returns an endpoint function that calls the method
// "update" of service "todos".
func NewUpdateEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*UpdatePayload)
		res, err := s.Update(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedTodo(res, "default")
		return vres, nil
	}
}

// NewDeleteEndpoint returns an endpoint function that calls the method
// "delete" of service "todos".
func NewDeleteEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*DeletePayload)
		res, err := s.Delete(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedTodo(res, "default")
		return vres, nil
	}
}
