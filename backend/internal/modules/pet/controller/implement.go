package controller

import (
	"backend/internal/lib/rr"
	"backend/internal/modules/pet/service"
	"errors"
	"net/http"
	"strconv"
	"strings"
)

type PetControl struct {
	service service.PetServicer
	rr      rr.ReadResponder
}

func NewPetControl(service service.PetServicer, readResponder rr.ReadResponder) *PetControl {
	return &PetControl{
		service: service,
		rr:      readResponder,
	}
}

// GetById godoc
// @Summary get pet by id
// @Description Return pet object provided pet id
// @Tags pet
// @Produce json
// @Param petId path int true "Pet ID"
// @Success 200 {object} entities.Pet
// @Failure 400,404 {object} rr.JSONResponse
// @Router /pet/{petId} [get]
func (c *PetControl) GetById(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	id := parts[len(parts)-1]

	petId, err := strconv.Atoi(id)
	if err != nil {
		_ = c.rr.WriteJSONError(w, errors.New("invalid id supplied"))
		return
	}

	pet, err := c.service.GetById(r.Context(), petId)
	if err != nil {
		_ = c.rr.WriteJSONError(w, err, 404)
		return
	}

	_ = c.rr.WriteJSON(w, 200, pet)
}

func (c *PetControl) UpdateWithForm(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (c *PetControl) DeleteById(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (c *PetControl) UploadImage(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (c *PetControl) AddPet(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (c *PetControl) UpdatePet(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (c *PetControl) GetByStatus(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}
