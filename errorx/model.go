package errorx

import (
	"net/http"

	"github.com/oklog/ulid/v2"
)

type BaseError struct {
	Id string `json:"id"`

	Domain      string `json:"domain"`
	Discription string `json:"discription"`
	Message     string `json:"message"`
	err         error

	ErrorType  string `json:"errorType"`
	StatusCode int    `json:"statusCode"`
}

func (b BaseError) Error() string {
	return b.Discription
}

func WrapENTError(domain string, err error) error {
	return &BaseError{
		Id: ulid.Make().String(),

		Domain:      domain,
		Discription: err.Error(),

		ErrorType: ErrorTypeEnt,
	}
}

func WrapError(domain string, err error) error {
	return BaseError{
		Id: ulid.Make().String(),

		Domain:      domain,
		Discription: err.Error(),
		err:         err,
	}
}

func NewUnProccessableEntity(domain, message string) error {
	return BaseError{
		Id: ulid.Make().String(),

		Domain:  domain,
		Message: message,

		ErrorType: ErrorTypeUnproccessableEntity,

		StatusCode: http.StatusUnprocessableEntity,
	}
}

func NewUnauthorizedError(domain, message string) error {
	return BaseError{
		Id: ulid.Make().String(),

		Domain:  domain,
		Message: message,

		ErrorType: ErrorTypeUnauthorized,

		StatusCode: http.StatusUnauthorized,
	}
}
