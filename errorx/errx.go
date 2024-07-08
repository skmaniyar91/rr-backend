package errorx

import (
	"fmt"
	"net/http"
	"rr-backend/ent/entgen"

	"github.com/go-playground/validator/v10"

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
		be.Err = err
		be.StatusCode = typedError.Code

		return be

	default:
		var be BaseError
		be.Message = "Something went wrong"
		be.Discription = err.Error()
		be.Err = err
		be.StatusCode = http.StatusInternalServerError

		return be
	}
}

func HandleBaseError(be BaseError) BaseError {
	fmt.Println("Information", be)
	switch be.ErrorType {
	case ErrorTypeEnt:
		return HandlerENTError(be)

	case ErrorTypeUnproccessableEntity:
		return be
	case ErrorTypeUnauthorized:
		return be
	case ErrorTypeBindingError:
		return HandleBindingError(be)

	default:
		be.Message = "Something went wrong"
		be.StatusCode = http.StatusInternalServerError

		return be
	}
}

func HandleBindingError(be BaseError) BaseError {
	switch typedErr := be.Err.(type) {
	case validator.ValidationErrors:
		contexualError := make(map[string]string)
		for _, e := range typedErr {
			contexualError[e.StructField()] = e.Error()
		}
		be.ContextualError = contexualError
		be.Message = "Validation failed on the fields"
		be.StatusCode = http.StatusBadRequest

		return be

	default:
		be.StatusCode = http.StatusBadRequest
		return be
	}
}

func HandlerENTError(be BaseError) BaseError {
	switch typedErr := be.Err.(type) {
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
