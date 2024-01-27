package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/eduardogomesf/go-api/internal/dto"
	"github.com/eduardogomesf/go-api/internal/entity"
	"github.com/eduardogomesf/go-api/internal/infra/database"
	"github.com/go-chi/jwtauth"
)

type UserHandler struct {
	UserDB       database.UserInterface
	Jwt          *jwtauth.JWTAuth
	JwtExpiresIn int
}

func NewUserHandler(db database.UserInterface, Jwt *jwtauth.JWTAuth, JwtExpiresIn int) *UserHandler {
	return &UserHandler{UserDB: db, Jwt: Jwt, JwtExpiresIn: JwtExpiresIn}
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

func (h *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	var jwtDto dto.GetJWTInput

	err := json.NewDecoder(r.Body).Decode(&jwtDto)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := h.UserDB.FindByEmail(jwtDto.Email)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if !user.ValidatePassword(jwtDto.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	_, tokenString, _ := h.Jwt.Encode(map[string]interface{}{
		"sub": user.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(h.JwtExpiresIn)).Unix(),
	})

	accessToken := struct {
		AccessToken string `json:"access_token"`
	}{
		AccessToken: tokenString,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)

}
