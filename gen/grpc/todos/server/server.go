// Code generated by goa v3.13.2, DO NOT EDIT.
//
// todos gRPC server
//
// Command:
// $ goa gen github.com/akm/goa_todo_example/design

package server

import (
	"context"

	todospb "github.com/akm/goa_todo_example/gen/grpc/todos/pb"
	todos "github.com/akm/goa_todo_example/gen/todos"
	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
)

// Server implements the todospb.TodosServer interface.
type Server struct {
	ListH   goagrpc.UnaryHandler
	ShowH   goagrpc.UnaryHandler
	CreateH goagrpc.UnaryHandler
	UpdateH goagrpc.UnaryHandler
	DeleteH goagrpc.UnaryHandler
	todospb.UnimplementedTodosServer
}

// New instantiates the server struct with the todos service endpoints.
func New(e *todos.Endpoints, uh goagrpc.UnaryHandler) *Server {
	return &Server{
		ListH:   NewListHandler(e.List, uh),
		ShowH:   NewShowHandler(e.Show, uh),
		CreateH: NewCreateHandler(e.Create, uh),
		UpdateH: NewUpdateHandler(e.Update, uh),
		DeleteH: NewDeleteHandler(e.Delete, uh),
	}
}

// NewListHandler creates a gRPC handler which serves the "todos" service
// "list" endpoint.
func NewListHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, nil, EncodeListResponse)
	}
	return h
}

// List implements the "List" method in todospb.TodosServer interface.
func (s *Server) List(ctx context.Context, message *todospb.ListRequest) (*todospb.ListResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "list")
	ctx = context.WithValue(ctx, goa.ServiceKey, "todos")
	resp, err := s.ListH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*todospb.ListResponse), nil
}

// NewShowHandler creates a gRPC handler which serves the "todos" service
// "show" endpoint.
func NewShowHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeShowRequest, EncodeShowResponse)
	}
	return h
}

// Show implements the "Show" method in todospb.TodosServer interface.
func (s *Server) Show(ctx context.Context, message *todospb.ShowRequest) (*todospb.ShowResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "show")
	ctx = context.WithValue(ctx, goa.ServiceKey, "todos")
	resp, err := s.ShowH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*todospb.ShowResponse), nil
}

// NewCreateHandler creates a gRPC handler which serves the "todos" service
// "create" endpoint.
func NewCreateHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeCreateRequest, EncodeCreateResponse)
	}
	return h
}

// Create implements the "Create" method in todospb.TodosServer interface.
func (s *Server) Create(ctx context.Context, message *todospb.CreateRequest) (*todospb.CreateResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "create")
	ctx = context.WithValue(ctx, goa.ServiceKey, "todos")
	resp, err := s.CreateH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*todospb.CreateResponse), nil
}

// NewUpdateHandler creates a gRPC handler which serves the "todos" service
// "update" endpoint.
func NewUpdateHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeUpdateRequest, EncodeUpdateResponse)
	}
	return h
}

// Update implements the "Update" method in todospb.TodosServer interface.
func (s *Server) Update(ctx context.Context, message *todospb.UpdateRequest) (*todospb.UpdateResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "update")
	ctx = context.WithValue(ctx, goa.ServiceKey, "todos")
	resp, err := s.UpdateH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*todospb.UpdateResponse), nil
}

// NewDeleteHandler creates a gRPC handler which serves the "todos" service
// "delete" endpoint.
func NewDeleteHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeDeleteRequest, EncodeDeleteResponse)
	}
	return h
}

// Delete implements the "Delete" method in todospb.TodosServer interface.
func (s *Server) Delete(ctx context.Context, message *todospb.DeleteRequest) (*todospb.DeleteResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "delete")
	ctx = context.WithValue(ctx, goa.ServiceKey, "todos")
	resp, err := s.DeleteH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*todospb.DeleteResponse), nil
}
