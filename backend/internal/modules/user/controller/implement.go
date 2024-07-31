package controller

import (
	"backend/internal/lib/e"
	"backend/internal/lib/rr"
	"backend/internal/modules/user/entities"
	"backend/internal/modules/user/service"
	"errors"
	"net/http"
	"strings"
)

type UserControl struct {
	service service.UserServicer
	rr      rr.ReadResponder
}

func NewUserController(service service.UserServicer, readResponder rr.ReadResponder) *UserControl {
	return &UserControl{
		service: service,
		rr:      readResponder,
	}
}

// GetByUsername godoc
// @Summary get user
// @Description Get user by username
// @Tags user
// @Produce json
// @Param username path string true "The name that needs to be fetched"
// @Success 200 {object} entities.User
// @Failure 400,404 {object} rr.JSONResponse
// @Router /user/{username} [get]
func (c *UserControl) GetByUsername(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	username := parts[len(parts)-1]

	if username == "" {
		_ = c.rr.WriteJSONError(w, errors.New("invalid username"))
		return
	}

	user, err := c.service.GetUserByName(r.Context(), username)
	if err != nil {
		_ = c.rr.WriteJSONError(w, e.Wrap("couldn't get user", err), 404)
		return
	}

	_ = c.rr.WriteJSON(w, 200, user)

}

// Update godoc
// @Summary update user
// @Description Updated user
// @Tags user
// @Accept json
// @Produce json
// @Param username path string true "Name that need to be updated"
// @Param body body entities.User true "Updated user object"
// @Success 200 {object} rr.JSONResponse
// @Failure 400,404 {object} rr.JSONResponse
// @Router /user/{username} [put]
func (c *UserControl) Update(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	username := parts[len(parts)-1]

	if username == "" {
		_ = c.rr.WriteJSONError(w, errors.New("invalid username"))
		return
	}

	if _, err := c.service.GetUserByName(r.Context(), username); err != nil {
		_ = c.rr.WriteJSONError(w, e.Wrap("couldn't get user", err), 404)
		return
	}

	var userUpdate entities.User
	if err := c.rr.ReadJSON(w, r, &userUpdate); err != nil {
		_ = c.rr.WriteJSONError(w, e.Wrap("couldn't read json", err))
		return
	}

	if userUpdate.Username != username {
		_ = c.rr.WriteJSONError(w, errors.New("username does not match"))
		return
	}

	if err := c.service.UpdateUser(r.Context(), userUpdate); err != nil {
		_ = c.rr.WriteJSONError(w, e.Wrap("couldn't update user", err))
		return
	}

	resp := rr.JSONResponse{Error: false, Message: "user updated"}
	_ = c.rr.WriteJSON(w, 200, resp)

}

func (c *UserControl) Delete(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (c *UserControl) Login(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (c *UserControl) Logout(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (c *UserControl) Create(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (c *UserControl) CreateWithArray(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}
