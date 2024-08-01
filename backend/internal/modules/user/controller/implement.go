package controller

import (
	"backend/internal/lib/e"
	"backend/internal/lib/rr"
	"backend/internal/lib/u"
	"backend/internal/modules/user/entities"
	"backend/internal/modules/user/service"
	"errors"
	"fmt"
	"net/http"
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
	username := u.ParamFromPath(r.URL.Path)

	if username == "" {
		_ = c.rr.WriteJSONError(w, errors.New("invalid username"))
		return
	}

	user, err := c.service.GetByName(r.Context(), username)
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
	username := u.ParamFromPath(r.URL.Path)

	if username == "" {
		_ = c.rr.WriteJSONError(w, errors.New("invalid username"))
		return
	}

	if _, err := c.service.GetByName(r.Context(), username); err != nil {
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

	if err := c.service.Update(r.Context(), userUpdate); err != nil {
		_ = c.rr.WriteJSONError(w, e.Wrap("couldn't update user", err))
		return
	}

	resp := rr.JSONResponse{Error: false, Message: "user updated"}
	_ = c.rr.WriteJSON(w, 200, resp)

}

// Create godoc
// @Summary create user
// @Description Create user
// @Tags user
// @Accept json
// @Produce json
// @Param body body entities.User true "User object"
// @Success 201 {object} rr.JSONResponse
// @Failure 400 {object} rr.JSONResponse
// @Router /user [post]
func (c *UserControl) Create(w http.ResponseWriter, r *http.Request) {
	var user entities.User
	_ = c.rr.ReadJSON(w, r, &user)

	if err := c.service.Create(r.Context(), user); err != nil {
		_ = c.rr.WriteJSONError(w, e.Wrap("couldn't create user", err))
		return
	}

	resp := rr.JSONResponse{Error: false, Message: "user created"}
	_ = c.rr.WriteJSON(w, 201, resp)
}

// Delete godoc
// @Summary delete user
// @Description Delete user
// @Tags user
// @Produce json
// @Param username path string true "The name that needs to be deleted"
// @Success 200 {object} rr.JSONResponse
// @Failure 400,404 {object} rr.JSONResponse
// @Router /user/{username} [delete]
func (c *UserControl) Delete(w http.ResponseWriter, r *http.Request) {
	username := u.ParamFromPath(r.URL.Path)
	ctx := r.Context()

	if username == "" {
		_ = c.rr.WriteJSONError(w, errors.New("invalid username"))
		return
	}

	if _, err := c.service.GetByName(ctx, username); err != nil {
		_ = c.rr.WriteJSONError(w, e.Wrap("couldn't get user", err), 404)
		return
	}

	if err := c.service.Delete(ctx, username); err != nil {
		_ = c.rr.WriteJSONError(w, e.Wrap("couldn't delete user", err))
		return
	}

	resp := rr.JSONResponse{Error: false, Message: "user deleted"}
	_ = c.rr.WriteJSON(w, 200, resp)
}

// CreateWithArray godoc
// @Summary create with array
// @Description Create list of users with given input array
// @Tags user
// @Accept json
// @Produce json
// @Param body body entities.Users true "List of user objects"
// @Success 201 {object} rr.JSONResponse
// @Failure 400 {object} rr.JSONResponse
// @Router /user/createWithArray [post]
func (c *UserControl) CreateWithArray(w http.ResponseWriter, r *http.Request) {
	var users entities.Users
	_ = c.rr.ReadJSON(w, r, &users)

	for _, user := range users {
		if err := c.service.Create(r.Context(), user); err != nil {
			_ = c.rr.WriteJSONError(w, e.Wrap("couldn't create user "+user.Username, err))
			return
		}
	}

	resp := rr.JSONResponse{Error: false, Message: fmt.Sprintf("%d user created", len(users))}
	_ = c.rr.WriteJSON(w, 201, resp)
}

// Login godoc
// @Summary login
// @Description Log user into the system
// @Tags user
// @Produce json
// @Param username query string true "The username for login"
// @Param password query string true "The password for login in clear text"
// @Success 200 {object} rr.JSONResponse
// @Failure 401 {object} rr.JSONResponse
// @Router /user/login [get]
func (c *UserControl) Login(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	username := query.Get("username")
	password := query.Get("password")

	tokens, cookie, err := c.service.Authorize(r.Context(), username, password)
	if err != nil {
		_ = c.rr.WriteJSONError(w, e.Wrap("couldn't authorize user", err), 401)
		return
	}

	resp := rr.JSONResponse{Error: false, Message: "user authorized", Data: tokens}

	http.SetCookie(w, cookie)
	_ = c.rr.WriteJSON(w, 200, resp)

}

// Logout godoc
// @Summary logout
// @Description Logs out currently logged user
// @Tags user
// @Produce json
// @Success 200 {object} rr.JSONResponse
// @Router /user/logout [get]
func (c *UserControl) Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, c.service.ResetCookie())

	resp := rr.JSONResponse{Error: false, Message: "user logged out"}
	_ = c.rr.WriteJSON(w, 200, resp)
}
