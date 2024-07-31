package controller

import (
	"net/http"
)

type PetController interface {
	GetById(w http.ResponseWriter, r *http.Request)
	UpdateWithForm(w http.ResponseWriter, r *http.Request)
	DeleteById(w http.ResponseWriter, r *http.Request)
	UploadImage(w http.ResponseWriter, r *http.Request)
	CreatePet(w http.ResponseWriter, r *http.Request)
	UpdatePet(w http.ResponseWriter, r *http.Request)
	GetByStatus(w http.ResponseWriter, r *http.Request)
}
