package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/knwoop/lets-gnomock/src/models"
	"github.com/knwoop/lets-gnomock/src/service"
)

var (
	_ http.Handler = (*createHandler)(nil)
	_ http.Handler = (*getHandler)(nil)
)

type createHandler struct {
	userService *service.UserService
}

func (h *createHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var u models.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		fmt.Printf("error bind: %s", err.Error())

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	if _, err := h.userService.Create(ctx, u.Username); err != nil {
		fmt.Printf("error bad request: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := json.NewEncoder(w).Encode(&u); err != nil {
		fmt.Printf("error internal: %s", err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Printf("good")

	return
}

type getHandler struct {
	userService *service.UserService
}

func (h *getHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}
