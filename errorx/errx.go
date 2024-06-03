package errorx

import (
	"net/http"
	"rr-backend/ent/entgen"

	"github.com/labstack/echo/v4"
)

func HandleError(err error) BaseError {
	switch typedError := err.(type) {
	case BaseError:
		return HandleBaseError(typedError)
	case *echo.HTTPError:
		var be BaseError
		be.Message, _ = typedError.Message.(string)
		be.Discription = typedError.Error()
		be.err = err
		be.StatusCode = typedError.Code

		return be

	default:
		var be BaseError
		be.Message = "Something went wrong"
		be.Discription = err.Error()
		be.err = err
		be.StatusCode = http.StatusInternalServerError

		return be
	}
}

func HandleBaseError(be BaseError) BaseError {
	switch be.ErrorType {
	case ErrorTypeEnt:
		return HandlerENTError(be)

	case ErrorTypeUnproccessableEntity:
		return be
	case ErrorTypeUnauthorized:
		return be

	default:
		be.Message = "Something went wrong"
		be.StatusCode = http.StatusInternalServerError

		return be
	}
}

func HandlerENTError(be BaseError) BaseError {
	switch typedErr := be.err.(type) {
	case *entgen.NotFoundError:
		be.Message = "Some thing not found"
		be.StatusCode = http.StatusNotFound

		return be
	case *entgen.ConstraintError:
		be.Message = typedErr.Error()
		be.StatusCode = http.StatusBadRequest

		return be

	case *entgen.NotSingularError:
		be.Message = typedErr.Error()
		be.StatusCode = http.StatusConflict

		return be

	case *entgen.ValidationError:
		be.Message = typedErr.Error()
		be.StatusCode = http.StatusBadRequest

		return be

	default:
		be.Message = "Something went wrong"
		be.StatusCode = http.StatusInternalServerError

		return be
	}
}
