package controller

import (
	"backend/internal/lib/e"
	"backend/internal/lib/rr"
	"backend/internal/modules/pet/entities"
	"backend/internal/modules/pet/service"
	"errors"
	"fmt"
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
// @Security ApiKeyAuth
// @Description Find pet by ID
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

// UpdateWithForm godoc
// @Summary update pet
// @Security ApiKeyAuth
// @Description Updates a pet in the store with form data
// @Tags pet
// @Produce json
// @Param petId path int true "Pet ID"
// @Param name formData string false "Pet name"
// @Param status formData string false "Pet status"
// @Success 200 {object} rr.JSONResponse
// @Failure 400 {object} rr.JSONResponse
// @Router /pet/{petId} [post]
func (c *PetControl) UpdateWithForm(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	id := parts[len(parts)-1]

	petId, err := strconv.Atoi(id)
	if err != nil {
		_ = c.rr.WriteJSONError(w, errors.New("invalid id supplied"))
		return
	}

	_ = r.ParseForm()
	name := r.FormValue("name")
	status := r.FormValue("status")

	if name == "" && status == "" {
		_ = c.rr.WriteJSONError(w, errors.New("at least one field must be not empty"))
		return
	}

	if err = c.service.UpdateWithForm(r.Context(), petId, name, status); err != nil {
		_ = c.rr.WriteJSONError(w, err)
		return
	}

	resp := rr.JSONResponse{Error: false, Message: "pet updated"}
	_ = c.rr.WriteJSON(w, 200, resp)

}

// DeleteById godoc
// @Summary delete pet
// @Security ApiKeyAuth
// @Description Deletes a pet
// @Tags pet
// @Produce json
// @Param petId path int true "Pet ID"
// @Success 200 {object} rr.JSONResponse
// @Failure 400,404 {object} rr.JSONResponse
// @Router /pet/{petId} [delete]
func (c *PetControl) DeleteById(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	id := parts[len(parts)-1]

	petId, err := strconv.Atoi(id)
	if err != nil {
		_ = c.rr.WriteJSONError(w, errors.New("invalid id supplied"))
		return
	}

	if err = c.service.DeleteById(r.Context(), petId); err != nil {
		_ = c.rr.WriteJSONError(w, err, 404)
		return
	}

	resp := rr.JSONResponse{Error: false, Message: "pet deleted"}
	_ = c.rr.WriteJSON(w, 200, resp)
}

// UploadImage godoc
// @Summary upload image
// @Security ApiKeyAuth
// @Description Uploads an image
// @Tags pet
// @Produce json
// @Param petId path int true "Pet ID"
// @Param additionalMetadata formData string false "Additional data to pass to server"
// @Param file formData file true "File to upload"
// @Success 200 {object} rr.JSONResponse
// @Failure 400,404 {object} rr.JSONResponse
// @Router /pet/{petId}/uploadImage [post]
func (c *PetControl) UploadImage(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	id := parts[len(parts)-2]

	petId, err := strconv.Atoi(id)
	if err != nil {
		_ = c.rr.WriteJSONError(w, errors.New("invalid id supplied"))
		return
	}

	if _, err = c.service.GetById(r.Context(), petId); err != nil {
		_ = c.rr.WriteJSONError(w, err, 404)
		return
	}

	_ = r.ParseForm()
	file, _, err := r.FormFile("file")
	if err != nil || file == nil {
		_ = c.rr.WriteJSONError(w, errors.New("invalid file supplied"))
		return
	}

	resp := rr.JSONResponse{Error: false, Message: "pet image uploaded"}
	_ = c.rr.WriteJSON(w, 200, resp)
}

// CreatePet godoc
// @Summary create pet
// @Security ApiKeyAuth
// @Description Add a new pet to the store
// @Tags pet
// @Accept json
// @Produce json
// @Param body body entities.Pet true "Pet object that needs to be added to the store"
// @Success 200 {object} rr.JSONResponse
// @Failure 400 {object} rr.JSONResponse
// @Router /pet [post]
func (c *PetControl) CreatePet(w http.ResponseWriter, r *http.Request) {
	var pet entities.Pet
	_ = c.rr.ReadJSON(w, r, &pet)

	petId, err := c.service.Create(r.Context(), pet)
	if err != nil {
		_ = c.rr.WriteJSONError(w, err)
		return
	}

	resp := rr.JSONResponse{Error: false, Message: fmt.Sprintf("pet created, id: %d", petId)}
	_ = c.rr.WriteJSON(w, 200, resp)

}

// UpdatePet godoc
// @Summary update pet
// @Security ApiKeyAuth
// @Description Update an existing pet
// @Tags pet
// @Accept json
// @Produce json
// @Param body body entities.Pet true "Pet object that needs to be added to the store"
// @Success 200 {object} rr.JSONResponse
// @Failure 400,404 {object} rr.JSONResponse
// @Router /pet [put]
func (c *PetControl) UpdatePet(w http.ResponseWriter, r *http.Request) {
	var pet entities.Pet
	_ = c.rr.ReadJSON(w, r, &pet)

	if _, err := c.service.GetById(r.Context(), pet.Id); err != nil {
		_ = c.rr.WriteJSONError(w, err, 404)
		return
	}

	if err := c.service.Update(r.Context(), pet); err != nil {
		_ = c.rr.WriteJSONError(w, err)
		return
	}

	resp := rr.JSONResponse{Error: false, Message: "pet updated"}
	_ = c.rr.WriteJSON(w, 200, resp)
}

// GetByStatus godoc
// @Summary get pets by status
// @Security ApiKeyAuth
// @Description Finds pets by status
// @Tags pet
// @Produce json
// @Param status query []string true "Status values that need to be considered for filter<br>Available values : <i>available, pending, sold</i>"
// @Success 200 {object} entities.Pets
// @Failure 400 {object} rr.JSONResponse
// @Router /pet/findByStatus [get]
func (c *PetControl) GetByStatus(w http.ResponseWriter, r *http.Request) {
	statuses := strings.Split(r.URL.Query()["status"][0], ",")

	if len(statuses) == 0 {
		_ = c.rr.WriteJSONError(w, errors.New("at least one field must be filled"))
		return
	}

	result := make(entities.Pets, 0)
	ctx := r.Context()

	for _, status := range statuses {
		pets, err := c.service.GetByStatus(ctx, status)
		if err != nil {
			_ = c.rr.WriteJSONError(w, e.Wrap("couldn't get pets by status "+status, err))
			return
		}
		result = append(result, pets...)
	}

	_ = c.rr.WriteJSON(w, 200, result)
}
