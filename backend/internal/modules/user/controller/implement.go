package controller

import (
	"backend/internal/lib/rr"
	"backend/internal/modules/user/service"
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

func (u *UserControl) GetByUsername(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (u *UserControl) Update(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (u *UserControl) Delete(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (u *UserControl) Login(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (u *UserControl) Logout(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (u *UserControl) Create(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (u *UserControl) CreateWithArray(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}
