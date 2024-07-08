package errorx

import (
	"net/http"
	"rr-backend/config/logx"

	"github.com/oklog/ulid/v2"
)

type BaseError struct {
	Id string `json:"id"`

	Domain      string `json:"domain"`
	Discription string `json:"discription"`
	Message     string `json:"message"`
	Err         error  `json:"error"`

	ErrorType       string            `json:"errorType"`
	ContextualError map[string]string `json:contextualError`
	StatusCode      int               `json:"statusCode"`
}

func (b BaseError) Error() string {
	return b.Discription
}

func WrapENTError(domain string, err error) error {
	be := BaseError{
		Id: ulid.Make().String(),

		Domain:      domain,
		Discription: err.Error(),

		ErrorType: ErrorTypeEnt,
		Err:       err,
	}
	logx.Logger().Error(be.Message, be)

	return be
}

func WrapBindingError(domain string, err error) error {
	be := BaseError{
		Id: ulid.Make().String(),

		Domain:      domain,
		ErrorType:   ErrorTypeBindingError,
		Discription: err.Error(),
		Err:         err,
	}
	logx.Logger().Error(err.Error(), be)

	return be
}

func WrapError(domain string, err error) error {
	be := BaseError{
		Id: ulid.Make().String(),

		Domain:      domain,
		Discription: err.Error(),
		Message:     err.Error(),
		Err:         err,
	}
	logx.Logger().Error(be.Message, be)

	return be
}

func NewUnProccessableEntity(domain, message string) error {
	be := BaseError{
		Id: ulid.Make().String(),

		Domain:  domain,
		Message: message,

		ErrorType: ErrorTypeUnproccessableEntity,

		StatusCode: http.StatusUnprocessableEntity,
	}

	logx.Logger().Error(be.Message, be)

	return be
}

func NewUnauthorizedError(domain, message string) error {
	be := BaseError{
		Id: ulid.Make().String(),

		Domain:  domain,
		Message: message,

		ErrorType: ErrorTypeUnauthorized,

		StatusCode: http.StatusUnauthorized,
	}
	logx.Logger().Error(be.Message, be)

	return be
}
