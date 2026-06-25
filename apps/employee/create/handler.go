package create

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/taufiqkba/food_app_v2/internal/constant"
	"github.com/taufiqkba/food_app_v2/internal/infra/response"
)

type handlerCreateEmployee struct {
	service createEmployeeService
}

func NewHandlerCreateEmployee(service createEmployeeService) handlerCreateEmployee {
	return handlerCreateEmployee{service: service}
}

func (h handlerCreateEmployee) createEmployee(rw http.ResponseWriter, r *http.Request) {
	var req = CreateEmployeeRequest{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		resp := response.NewErrorBadRequest("bad request", err, response.WithStatusCode(
			fmt.Sprintf("%v%v%v%v", http.StatusBadRequest, constant.ModuleEmployeeCode, constant.ServiceEmployeeCreate, "00"),
		))
		resp.JSON(rw)
		return
	}

	if err := req.Validate(); err != nil {
		resp := getResponse(err, errCreateMap)
		resp.JSON(rw)
		return
	}

	if err := h.service.create(r.Context(), req); err != nil {
		resp := getResponse(err, errCreateMap)
		resp.JSON(rw)
		return
	}

	resp := response.NewSuccessCreated("employee created successfully", response.WithStatusCode(
		fmt.Sprintf("%v%v%v", constant.ModuleEmployeeCode, constant.ServiceEmployeeCreate, "00"),
	))

	resp.JSON(rw)
}
