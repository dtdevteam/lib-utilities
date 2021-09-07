package utils

import "fmt"

type ErrorHandle struct {
	ErrorStatus  int         `json:"error_status"`
	ErrorMessage string      `json:"error_message"`
	ErrorInput   interface{} `json:"error_input"`
	ErrorFrom    string      `json:"error_from"`
}

func ExceptionError(errorStatus int, err error) (int, ErrorHandle) {
	var err_input interface{}

	return errorStatus, ErrorHandle{
		ErrorStatus:  errorStatus,
		ErrorMessage: fmt.Sprintf("%s", err.Error()),
		ErrorInput:   err_input,
	}
}

func ExceptionErrorValidation(errorStatus int, err error, err_input []map[string]interface{}) (int, ErrorHandle) {
	return errorStatus, ErrorHandle{
		ErrorStatus:  errorStatus,
		ErrorMessage: fmt.Sprintf("%s", err.Error()),
		ErrorInput:   err_input,
	}
}
