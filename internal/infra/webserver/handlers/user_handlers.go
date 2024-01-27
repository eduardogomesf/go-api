package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/eduardogomesf/go-api/internal/dto"
	"github.com/eduardogomesf/go-api/internal/entity"
	"github.com/eduardogomesf/go-api/internal/infra/database"
)

type UserHandler struct {
	UserDB database.UserInterface
}

func NewUserHandler(db database.UserInterface) *UserHandler {
	return &UserHandler{UserDB: db}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userDto dto.CreateUserInput

	err := json.NewDecoder(r.Body).Decode(&userDto)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := entity.NewUser(userDto.Name, userDto.Email, userDto.Password)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.UserDB.Create(user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
