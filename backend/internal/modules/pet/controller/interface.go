package controller

import (
	"net/http"
)

type PetController interface {
	GetById(w http.ResponseWriter, r *http.Request)
	UpdateWithForm(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	UploadImage(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	GetByStatus(w http.ResponseWriter, r *http.Request)
}
