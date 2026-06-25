package create

import (
	"github.com/taufiqkba/food_app_v2/internal/constant"
	"github.com/taufiqkba/food_app_v2/internal/utils/encryption"
	"github.com/taufiqkba/food_app_v2/internal/utils/generator"
	"github.com/taufiqkba/food_app_v2/internal/utils/validation"
)

type CreateEmployeeRequest struct {
	Name     string `json:"name"`
	Role     string `json:"role"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Profile  string `json:"profile"`
}

type CreateEmployeeResponse struct {
}

func (r CreateEmployeeRequest) Validate() error {
	if r.Name == "" || r.Email == "" || r.Password == "" || r.Role == "" || r.Profile == "" {
		return errNameEmailPasswordRoleProfileIsEmpty
	}

	if !constant.IsRoleCanBeCreated(r.Role) {
		return errRoleIsNotSupported
	}

	if len(r.Password) < 6 {
		return errPasswordInvalidLength
	}

	if !validation.IsValidEmail(r.Email) {
		return errEmailIsNotValid
	}

	return nil
}

func (r CreateEmployeeRequest) ToAuthModel() (auth Auth) {
	hashed, err := encryption.GenerateFromPassword(r.Password)
	if err != nil {
		return
	}

	auth = Auth{
		PublicId: generator.GeneratePublicID(),
		Email:    r.Email,
		Password: hashed,
		Role:     r.Role,
		IsActive: true,
	}
	return
}

func (r CreateEmployeeRequest) ToEmployeeModel(authId string) (employee Employee) {
	employee = Employee{
		PublicId: generator.GeneratePublicID(),
		Name:     r.Name,
		Profile:  r.Profile,
		AuthId:   authId,
	}

	return
}
