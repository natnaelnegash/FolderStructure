package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"folderstructure/internal/constant/errors"

	validation "github.com/go-playground/validator/v10"
	"github.com/joomcode/errorx"
	"github.com/spf13/viper"
)

func SendSuccessResponse(w http.ResponseWriter, statusCode int, code Code, data any, message any) {

	response := SuccessResponse{
		Status:  "success",
		Code:    string(code),
		Message: message,
		Data:    data,
		Metadata: Metadata{
			ServerTimestamp: time.Now().UTC().Format("2006-01-02T15:04:05Z"),
		},
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(response)

}

func SendErrorResponse(w http.ResponseWriter, err *ErrorResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.Code)
	_ = json.NewEncoder(w).Encode(err)
}

func SendErrorResponseFormated(w http.ResponseWriter, err error) {
	statusCode := http.StatusInternalServerError
	code := "internal error"
	message := "internal server error"

	if err != nil {
		for _, registered := range errors.Error {
			if registered.Type == nil {
				continue
			}

			if !errorx.IsOfType(err, registered.Type) {
				continue
			}

			statusCode = registered.StatusCode

			code = strings.ToLower(
				strings.ReplaceAll(
					registered.Type.Namespace().String(),
					"_",
					" ",
				),
			)

			if typedErr := errorx.Cast(err); typedErr != nil {
				if customMessage := typedErr.Message(); customMessage != "" {
					message = customMessage
				} else {
					message = code
				}
			}

			break
		}
	}

	response := ErrorResponseFormat{
		Status:  "error",
		Code:    code,
		Message: message,
		Data:    nil,
		Metadata: Metadata{
			ServerTimestamp: time.Now().UTC().Format(time.RFC3339),
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	_ = json.NewEncoder(w).Encode(response)
}

func GetErrorFrom(err error) *ErrorResponse {
	debugMode := viper.GetBool("debug")

	for _, e := range errors.Error {
		if errorx.IsOfType(err, e.Type) {
			er := errorx.Cast(err)
			res := ErrorResponse{
				Code:       e.StatusCode,
				Message:    er.Message(),
				FieldError: ErrorFields(er.Cause()),
			}

			if debugMode {
				res.Description = fmt.Sprintf("Error: %v", er)
				res.StackTrace = fmt.Sprintf("%+v", errorx.EnsureStackTrace(err))
			}

			return &res
		}
	}

	return &ErrorResponse{
		Code:    http.StatusInternalServerError,
		Message: "Unknown server error",
	}
}

func ErrorFields(err error) []FieldError {
	var errs []FieldError

	if data, ok := err.(validation.ValidationErrors); ok {
		for i, v := range data {
			errs = append(errs, FieldError{
				Name:        fmt.Sprintf("%d", i+1),
				Description: v.Error(),
			},
			)
		}

		return errs
	}

	return nil
}

