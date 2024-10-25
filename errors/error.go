package errors

import (
	"fmt"
	"log"
	"net/http"
)

type HttpError struct {
	StatusCd int
	Message  string
}

func (e *HttpError) Error() string {
	return fmt.Sprintf("Status: %d - %s", e.StatusCd, e.Message)
}

func NewHttpError(statusCode int, message string) *HttpError {
	return &HttpError{
		StatusCd: statusCode,
		Message:  message,
	}
}

func New500Error(err error) *HttpError {
	log.Println("An internal server error happened while searching coupons", err.Error())
	return NewHttpError(http.StatusInternalServerError, "an unknown error has occurred")
}
