package create

import (
	"errors"

	"github.com/taufiqkba/food_app_v2/internal/constant"
	"github.com/taufiqkba/food_app_v2/internal/infra/response"
)

var (
	errNameEmailPasswordRoleProfileIsEmpty = errors.New("name, email, password, role and profile can't be empty")
	errRoleIsNotSupported                  = errors.New("role is not supported, should be one of cashier or warehouse")
	errEmailAlreadyRegistered              = errors.New("email already registered")
	errPasswordInvalidLength               = errors.New("password length minimum is 6")
	errEmailNotFound                       = errors.New("email is not found")
	errEmailIsNotValid                     = errors.New("email is not valild")
	errInternalServerrError                = errors.New("internal server error")
)

var (
	errCreateMap = map[string]response.Response{
		errNameEmailPasswordRoleProfileIsEmpty.Error(): response.NewErrorBadRequest("bad request", errNameEmailPasswordRoleProfileIsEmpty, response.WithStatusCode(
			generateStatusCode(constant.ServiceEmployeeCreate, "01"),
		)),
		errRoleIsNotSupported.Error(): response.NewErrorBadRequest("bad request", errRoleIsNotSupported, response.WithStatusCode(
			generateStatusCode(constant.ServiceEmployeeCreate, "02"),
		)),
		errEmailAlreadyRegistered.Error(): response.NewErrorConflict("data already exists", errEmailAlreadyRegistered, response.WithStatusCode(
			generateStatusCode(constant.ServiceEmployeeCreate, "01"),
		)),
		errPasswordInvalidLength.Error(): response.NewErrorBadRequest("bad request", errPasswordInvalidLength, response.WithStatusCode(
			generateStatusCode(constant.ServiceEmployeeCreate, "03"),
		)),
		errEmailIsNotValid.Error(): response.NewErrorBadRequest("bad request", errEmailIsNotValid, response.WithStatusCode(
			generateStatusCode(constant.ServiceEmployeeCreate, "04"),
		)),
		// errEmailNotFound.Error(): response.NewErrorBadRequest("bad request", errEmailNotFound, response.WithStatusCode(
		// 	generateStatusCode(constant.ServiceEmployeeCreate, "01"),
		// )),
		// errInternalServerrError.Error(): response.NewErrorBadRequest("bad request", errInternalServerrError, response.WithStatusCode(
		// 	generateStatusCode(constant.ServiceEmployeeCreate, "01"),
		// )),
	}

	errListMapp = map[string]response.Response{}
)

func generateStatusCode(serviceCode string, statusCode string) string {
	return constant.ModuleEmployeeCode + serviceCode + statusCode
}

func getResponse(err error, errorMap map[string]response.Response) response.Response {
	if r, ok := errorMap[err.Error()]; ok {
		return r
	}

	return response.NewErrorGeneral("unknown error", errInternalServerrError, response.WithStatusCode(
		generateStatusCode("00", "99"),
	))
}
