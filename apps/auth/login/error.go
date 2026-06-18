package login

import (
	"errors"

	"github.com/taufiqkba/food_app_v2/internal/constant"
	"github.com/taufiqkba/food_app_v2/internal/infra/response"
)

var (
	errEmailOrPasswordEmpty  = errors.New("email or password is empty")
	errPasswordInvalidLength = errors.New("password length minimum 6")

	errEmailOrPasswordIsNotMatched = errors.New("email or password is invalid")
	errAccountIsNotActive          = errors.New("account is not active")

	errInternalSeverError = errors.New("internal server error")
)

func generateStatusCode(statusCode string) string {
	return constant.ModuleAuthCode + constant.ServiceAuthCodeLogin + statusCode
}

var (
	errMap = map[string]response.Response{
		errEmailOrPasswordEmpty.Error(): response.NewErrorBadRequest("bad request",
			errEmailOrPasswordEmpty, response.WithStatusCode(
				generateStatusCode("01"),
			)),
		errPasswordInvalidLength.Error(): response.NewErrorBadRequest("bad request", errPasswordInvalidLength, response.WithStatusCode(
			generateStatusCode("02"),
		)),
		errAccountIsNotActive.Error(): response.NewErrorStatusUnprocessableEntity("unprocessable entity", errAccountIsNotActive, response.WithStatusCode(
			generateStatusCode("01"),
		)),
		errEmailOrPasswordIsNotMatched.Error(): response.NewErrorUnauthorized("unauthorized", errEmailOrPasswordIsNotMatched, response.WithStatusCode(
			generateStatusCode("01"),
		)),
	}
)

func getResponse(err error) response.Response {
	if r, ok := errMap[err.Error()]; ok {
		return r
	}

	return response.NewErrorGeneral("unknown error", errInternalSeverError, response.WithStatusCode(
		generateStatusCode("99"),
	))
}
