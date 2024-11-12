// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// DeleteUser implements deleteUser operation.
//
// DELETE /user/{id}
func (UnimplementedHandler) DeleteUser(ctx context.Context, params DeleteUserParams) (r DeleteUserRes, _ error) {
	return r, ht.ErrNotImplemented
}

// GetUser implements getUser operation.
//
// GET /user/{id}
func (UnimplementedHandler) GetUser(ctx context.Context, params GetUserParams) (r GetUserRes, _ error) {
	return r, ht.ErrNotImplemented
}

// PostUser implements postUser operation.
//
// POST /user
func (UnimplementedHandler) PostUser(ctx context.Context, req OptUser) (r PostUserRes, _ error) {
	return r, ht.ErrNotImplemented
}

// PutUser implements putUser operation.
//
// PUT /user/{id}
func (UnimplementedHandler) PutUser(ctx context.Context, params PutUserParams) (r PutUserRes, _ error) {
	return r, ht.ErrNotImplemented
}

// NewError creates *UserStatusCode from error returned by handler.
//
// Used for common default response.
func (UnimplementedHandler) NewError(ctx context.Context, err error) (r *UserStatusCode) {
	r = new(UserStatusCode)
	return r
}