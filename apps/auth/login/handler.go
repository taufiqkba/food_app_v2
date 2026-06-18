package login

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/taufiqkba/food_app_v2/internal/config"
	"github.com/taufiqkba/food_app_v2/internal/constant"
	"github.com/taufiqkba/food_app_v2/internal/infra/response"
)

type handlerLogin struct {
	service serviceLogin //exact object
}

func newHandlerLogin(service serviceLogin) handlerLogin {
	return handlerLogin{service: service}
}

func (h handlerLogin) login(rw http.ResponseWriter, r *http.Request) {
	var req = LoginRequest{}

	// error while get request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		resp := response.NewErrorBadRequest("bad request", err, response.WithStatusCode(
			fmt.Sprintf("%v%v%v%v", http.StatusBadRequest, constant.ServiceAuthCodeLogin, constant.ModuleAuthCode, "00"),
		))
		resp.JSON(rw)
	}

	// validate
	if err := req.Validate(); err != nil {
		resp := getResponse(err)
		resp.JSON(rw)
		return
	}

	// call service
	token, role, err := h.service.login(r.Context(), req)
	if err != nil {
		resp := getResponse(err)
		resp.JSON(rw)
		return
	}

	cfg := config.GetConfig()
	resp := response.NewSuccessOK("login success", response.WithPayload(
		map[string]interface{}{
			"token":      token,
			"role":       role,
			"token_type": cfg.App.TokenType,
		},
	),
		response.WithStatusCode(generateStatusCode("00")),
	)

	resp.JSON(rw)
}
