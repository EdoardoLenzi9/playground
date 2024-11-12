// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

	ht "github.com/ogen-go/ogen/http"
)

func encodeDeleteUserResponse(response DeleteUserRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *DeleteUserBadRequest:
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))

		return nil

	case *DeleteUserInternalServerError:
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeGetUserResponse(response GetUserRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *GetUserBadRequest:
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))

		return nil

	case *GetUserInternalServerError:
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodePostUserResponse(response PostUserRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *PostUserBadRequest:
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))

		return nil

	case *PostUserInternalServerError:
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodePutUserResponse(response PutUserRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *PutUserBadRequest:
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))

		return nil

	case *PutUserInternalServerError:
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeErrorResponse(response *UserStatusCode, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	code := response.StatusCode
	if code == 0 {
		// Set default status code.
		code = http.StatusOK
	}
	w.WriteHeader(code)
	if st := http.StatusText(code); code >= http.StatusBadRequest {
		span.SetStatus(codes.Error, st)
	} else {
		span.SetStatus(codes.Ok, st)
	}

	e := new(jx.Encoder)
	response.Response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	if code >= http.StatusInternalServerError {
		return errors.Wrapf(ht.ErrInternalServerErrorResponse, "code: %d, message: %s", code, http.StatusText(code))
	}
	return nil

}
