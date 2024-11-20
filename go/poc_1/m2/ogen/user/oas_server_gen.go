// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// DeleteUser implements deleteUser operation.
	//
	// DELETE /user/{id}
	DeleteUser(ctx context.Context, params DeleteUserParams) (DeleteUserRes, error)
	// GetUser implements getUser operation.
	//
	// GET /user/{id}
	GetUser(ctx context.Context, params GetUserParams) (GetUserRes, error)
	// PostUser implements postUser operation.
	//
	// POST /user
	PostUser(ctx context.Context, req OptUser) (PostUserRes, error)
	// PutUser implements putUser operation.
	//
	// PUT /user/{id}
	PutUser(ctx context.Context, params PutUserParams) (PutUserRes, error)
	// NewError creates *UserStatusCode from error returned by handler.
	//
	// Used for common default response.
	NewError(ctx context.Context, err error) *UserStatusCode
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h Handler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		baseServer: s,
	}, nil
}
