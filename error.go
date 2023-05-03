package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	g "wistefan/satellite-wrapper/generated"

	"github.com/valyala/fasthttp"
)

type ErrorHandler func(ctx *fasthttp.RequestCtx, err error, result *g.ImplResponse)

var (
	// ErrTypeAssertionError is thrown when type an interface does not match the asserted type
	ErrTypeAssertionError = errors.New("unable to assert type")
)

// ParsingError indicates that an error has occurred when parsing request parameters
type ParsingError struct {
	Err error
}

func (e *ParsingError) Unwrap() error {
	return e.Err
}

func (e *ParsingError) Error() string {
	return e.Err.Error()
}

// RequiredError indicates that an error has occurred when parsing request parameters
type RequiredError struct {
	Field string
}

func (e *RequiredError) Error() string {
	return fmt.Sprintf("required field '%s' is zero value.", e.Field)
}

func DefaultErrorHandler(ctx *fasthttp.RequestCtx, err error, result *g.ImplResponse) {
	encoder := json.NewEncoder(ctx.Response.BodyWriter())
	var problemDetails g.ProblemDetails
	if _, ok := err.(*ParsingError); ok {
		// Handle parsing errors
		problemDetails = g.ProblemDetails{Status: http.StatusBadRequest, Title: "Bad Request", Detail: err.Error()}
		ctx.Response.SetStatusCode(http.StatusBadRequest)
	} else if _, ok := err.(*RequiredError); ok {
		// Handle missing required errors
		problemDetails = g.ProblemDetails{Status: http.StatusUnprocessableEntity, Title: "Unprocessable entity", Detail: err.Error()}
		ctx.Response.SetStatusCode(http.StatusUnprocessableEntity)
	} else {
		// Handle all other errors
		problemDetails = g.ProblemDetails{Status: http.StatusInternalServerError, Title: err.Error()}
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
	}
	err = encoder.Encode(problemDetails)
	if err != nil {
		ctx.Error(err.Error(), http.StatusInternalServerError)
	}

}
