package errors

import (
	"net/http"

	"github.com/joomcode/errorx"
)

type ErrorType struct {
	StatusCode int
	Type       *errorx.Type
}

var Error = []ErrorType{

	{
		StatusCode: http.StatusBadRequest,
		Type:       ErrInvalidUserInput,
	},
	
	{
		StatusCode: http.StatusBadRequest,
		Type:       ErrPoolNotStarted,
	},
}

var (
	invalidInput = errorx.NewNamespace("validation error").ApplyModifiers(errorx.TypeModifierOmitStackTrace)
	poolNotStarted=errorx.NewNamespace("pool not started")
)

var (
	ErrInvalidUserInput = errorx.NewType(invalidInput, "invalid user input")
	ErrPoolNotStarted=errorx.NewType(poolNotStarted,"pool not started")
)
