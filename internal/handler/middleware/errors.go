package middleware

import (
	"fmt"
	"net/http"

	"folderstructure/internal/constant/errors"
	response2 "folderstructure/internal/constant/response"

	"github.com/joomcode/errorx"

	"github.com/spf13/viper"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type contextKey string

var ErrorKey contextKey = "error"

func ErrorHandler(next http.Handler) http.Handler {
	debugMode := viper.GetBool("debug")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)

		err, ok := r.Context().Value(ErrorKey).(error)
		if !ok || err == nil {
			return
		}

		response := CastErrorResponse(err)
		if response != nil {
			if debugMode {
				er := errorx.Cast(err)
				response.Description = fmt.Sprintf("Error: %v", er)
				response.StackTrace = fmt.Sprintf("%+v", errorx.EnsureStackTrace(err))
			}
			response2.SendErrorResponse(w, response)
			return
		}

		response2.SendErrorResponse(w, &response2.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Unknown server error",
		})
	})
}

func ErrorFields(err error) []response2.FieldError {
	var errs []response2.FieldError

	if data, ok := err.(validation.Errors); ok {
		for i, v := range data {
			nestedErrors := ErrorFields(v)
			if len(nestedErrors) > 0 {
				errs = append(errs, nestedErrors...)
			} else {
				errs = append(errs, response2.FieldError{
					Name:        i,
					Description: v.Error(),
				})
			}
		}

		return errs
	}

	return nil
}

func CastErrorResponse(err error) *response2.ErrorResponse {
	for _, e := range errors.Error {
		if errorx.IsOfType(err, e.Type) {
			er := errorx.Cast(err)
			response := response2.ErrorResponse{
				Code:       e.StatusCode,
				Message:    er.Message(),
				FieldError: ErrorFields(er.Cause()),
			}
			return &response
		}
	}
	return nil
}
