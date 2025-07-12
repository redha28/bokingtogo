package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	familydto "github.com/redha28/bookingtogo/backend/internal/dto/family"
	"github.com/redha28/bookingtogo/backend/internal/services"
)

type FamilyHandler struct {
	FamilyService *services.FamilyService
	Validator     *validator.Validate
}

func NewFamilyHandler(familyService *services.FamilyService) *FamilyHandler {
	return &FamilyHandler{
		FamilyService: familyService,
		Validator:     validator.New(),
	}
}

// CreateFamily creates a new family member
// @Summary Create a new family member
// @Description Create a new family member
// @Tags families
// @Accept json
// @Produce json
// @Param family body familydto.CreateFamilyRequest true "Family data"
// @Success 201 {object} familydto.FamilyResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /families [post]
func (h *FamilyHandler) CreateFamily(w http.ResponseWriter, r *http.Request) {
	var req familydto.CreateFamilyRequest
	
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	if err := h.Validator.Struct(req); err != nil {
		http.Error(w, "Validation failed: "+err.Error(), http.StatusBadRequest)
		return
	}

	familyData, err := h.FamilyService.CreateFamily(r.Context(), req)
	if err != nil {
		http.Error(w, "Failed to create family: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := familydto.FamilyResponse{
		Data: *familyData,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// GetFamiliesByCustomerID retrieves all family members for a customer
// @Summary Get family members by customer ID
// @Description Get all family members for a specific customer
// @Tags families
// @Accept json
// @Produce json
// @Param customer_id path int true "Customer ID"
// @Success 200 {object} familydto.FamilyListResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /customers/{customer_id}/families [get]
func (h *FamilyHandler) GetFamiliesByCustomerID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerID, err := strconv.Atoi(vars["customer_id"])
	if err != nil {
		http.Error(w, "Invalid customer ID format", http.StatusBadRequest)
		return
	}

	families, err := h.FamilyService.GetFamiliesByCustomerID(r.Context(), customerID)
	if err != nil {
		http.Error(w, "Failed to get families: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := familydto.FamilyListResponse{
		Data: families,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetFamilyByID retrieves a family member by ID
// @Summary Get family member by ID
// @Description Get a family member by ID
// @Tags families
// @Accept json
// @Produce json
// @Param id path int true "Family ID"
// @Success 200 {object} familydto.FamilyResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /families/{id} [get]
func (h *FamilyHandler) GetFamilyByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	family, err := h.FamilyService.GetFamilyByID(r.Context(), id)
	if err != nil {
		http.Error(w, "Family not found", http.StatusNotFound)
		return
	}

	response := familydto.FamilyResponse{
		Data: *family,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// UpdateFamily updates a family member by ID
// @Summary Update family member by ID
// @Description Update a family member by ID
// @Tags families
// @Accept json
// @Produce json
// @Param id path int true "Family ID"
// @Param family body familydto.UpdateFamilyRequest true "Family data"
// @Success 200 {object} familydto.FamilyResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /families/{id} [put]
func (h *FamilyHandler) UpdateFamily(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var req familydto.UpdateFamilyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	if err := h.Validator.Struct(req); err != nil {
		http.Error(w, "Validation failed: "+err.Error(), http.StatusBadRequest)
		return
	}

	family, err := h.FamilyService.UpdateFamily(r.Context(), id, req)
	if err != nil {
		http.Error(w, "Failed to update family: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := familydto.FamilyResponse{
		Data: *family,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// DeleteFamily deletes a family member by ID
// @Summary Delete family member by ID
// @Description Delete a family member by ID
// @Tags families
// @Accept json
// @Produce json
// @Param id path int true "Family ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /families/{id} [delete]
func (h *FamilyHandler) DeleteFamily(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	if err := h.FamilyService.DeleteFamily(r.Context(), id); err != nil {
		http.Error(w, "Failed to delete family: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Family member deleted successfully",
	})
}
