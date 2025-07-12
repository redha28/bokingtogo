package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	nationalitydto "github.com/redha28/bokingtogo/backend/internal/dto/nationality"
	"github.com/redha28/bokingtogo/backend/internal/services"
)

type NationalityHandler struct {
	NationalityService *services.NationalityService
	Validator          *validator.Validate
}

func NewNationalityHandler(nationalityService *services.NationalityService) *NationalityHandler {
	return &NationalityHandler{
		NationalityService: nationalityService,
		Validator:          validator.New(),
	}
}

// CreateNationality creates a new nationality
// @Summary Create a new nationality
// @Description Create a new nationality
// @Tags nationalities
// @Accept json
// @Produce json
// @Param nationality body nationalitydto.CreateNationalityRequest true "Nationality data"
// @Success 201 {object} nationalitydto.NationalityResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /nationalities [post]
func (h *NationalityHandler) CreateNationality(w http.ResponseWriter, r *http.Request) {
	var req nationalitydto.CreateNationalityRequest
	
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	if err := h.Validator.Struct(req); err != nil {
		http.Error(w, "Validation failed: "+err.Error(), http.StatusBadRequest)
		return
	}

	nationalityData, err := h.NationalityService.CreateNationality(r.Context(), req)
	if err != nil {
		http.Error(w, "Failed to create nationality: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := nationalitydto.NationalityResponse{
		Data: *nationalityData,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// GetNationalities retrieves all nationalities
// @Summary Get all nationalities
// @Description Get all nationalities
// @Tags nationalities
// @Accept json
// @Produce json
// @Success 200 {object} nationalitydto.NationalityListResponse
// @Failure 500 {object} map[string]string
// @Router /nationalities [get]
func (h *NationalityHandler) GetNationalities(w http.ResponseWriter, r *http.Request) {
	nationalities, err := h.NationalityService.GetNationalities(r.Context())
	if err != nil {
		http.Error(w, "Failed to get nationalities: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := nationalitydto.NationalityListResponse{
		Data: nationalities,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetNationalityByID retrieves a nationality by ID
// @Summary Get nationality by ID
// @Description Get a nationality by ID
// @Tags nationalities
// @Accept json
// @Produce json
// @Param id path int true "Nationality ID"
// @Success 200 {object} nationalitydto.NationalityResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /nationalities/{id} [get]
func (h *NationalityHandler) GetNationalityByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	nationality, err := h.NationalityService.GetNationalityByID(r.Context(), id)
	if err != nil {
		http.Error(w, "Nationality not found", http.StatusNotFound)
		return
	}

	response := nationalitydto.NationalityResponse{
		Data: *nationality,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// UpdateNationality updates a nationality by ID
// @Summary Update nationality by ID
// @Description Update a nationality by ID
// @Tags nationalities
// @Accept json
// @Produce json
// @Param id path int true "Nationality ID"
// @Param nationality body nationalitydto.UpdateNationalityRequest true "Nationality data"
// @Success 200 {object} nationalitydto.NationalityResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /nationalities/{id} [put]
func (h *NationalityHandler) UpdateNationality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var req nationalitydto.UpdateNationalityRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	if err := h.Validator.Struct(req); err != nil {
		http.Error(w, "Validation failed: "+err.Error(), http.StatusBadRequest)
		return
	}

	nationality, err := h.NationalityService.UpdateNationality(r.Context(), id, req)
	if err != nil {
		http.Error(w, "Failed to update nationality: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := nationalitydto.NationalityResponse{
		Data: *nationality,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// DeleteNationality deletes a nationality by ID
// @Summary Delete nationality by ID
// @Description Delete a nationality by ID
// @Tags nationalities
// @Accept json
// @Produce json
// @Param id path int true "Nationality ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /nationalities/{id} [delete]
func (h *NationalityHandler) DeleteNationality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	if err := h.NationalityService.DeleteNationality(r.Context(), id); err != nil {
		http.Error(w, "Failed to delete nationality: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Nationality deleted successfully",
	})
}
