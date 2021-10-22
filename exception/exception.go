package exception

import (
	"fmt"

	"github.com/gofrs/uuid"
	// "github.com/labstack/gommon/log"
)

type ErrorHandle struct {
	TracingId    string      `json:"tracing_id"`
	ErrorStatus  int         `json:"error_status"`
	ErrorMessage string      `json:"error_message"`
	ErrorInput   interface{} `json:"error_input"`
	ErrorFrom    string      `json:"error_from"`
}

func ExceptionError(errorStatus int, err error) (int, ErrorHandle) {
	var err_input interface{}
	var tracing_id, _ = uuid.NewV4()

	return errorStatus, ErrorHandle{
		TracingId:    tracing_id.String(),
		ErrorStatus:  errorStatus,
		ErrorMessage: fmt.Sprintf("%s", err.Error()),
		ErrorInput:   err_input,
	}
}

func ExceptionErrorValidation(errorStatus int, err error, err_input []map[string]interface{}) (int, ErrorHandle) {
	var tracing_id, _ = uuid.NewV4()

	return errorStatus, ErrorHandle{
		TracingId:    tracing_id.String(),
		ErrorStatus:  errorStatus,
		ErrorMessage: fmt.Sprintf("%s", err.Error()),
		ErrorInput:   err_input,
	}
}
