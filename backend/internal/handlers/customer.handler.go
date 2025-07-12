package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	customerdto "github.com/redha28/bookingtogo/internal/dto/customer"
	"github.com/redha28/bookingtogo/internal/services"
)

type CustomerHandler struct {
	CustomerService *services.CustomerService
	Validator       *validator.Validate
}

func NewCustomerHandler(customerService *services.CustomerService) *CustomerHandler {
	return &CustomerHandler{
		CustomerService: customerService,
		Validator:       validator.New(),
	}
}

// CreateCustomer creates a new customer with family members
// @Summary Create a new customer
// @Description Create a new customer with family members
// @Tags customers
// @Accept json
// @Produce json
// @Param customer body customerdto.CreateCustomerRequest true "Customer data"
// @Success 201 {object} customerdto.CustomerResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /customers [post]
func (h *CustomerHandler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var req customerdto.CreateCustomerRequest
	
	// Decode JSON request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Validate request
	if err := h.Validator.Struct(req); err != nil {
		http.Error(w, "Validation failed: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Create customer via service
	customerData, err := h.CustomerService.CreateCustomer(r.Context(), req)
	if err != nil {
		http.Error(w, "Failed to create customer: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Create response using CustomerResponse struct
	response := customerdto.CustomerResponse{
		Data: *customerData,
	}

	// Set response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// Send JSON response
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

// GetCustomers retrieves all customers
// @Summary Get all customers
// @Description Get all customers with their family members
// @Tags customers
// @Accept json
// @Produce json
// @Success 200 {object} customerdto.CustomerListResponse
// @Failure 500 {object} map[string]string
// @Router /customers [get]
func (h *CustomerHandler) GetCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := h.CustomerService.GetCustomers(r.Context())
	if err != nil {
		http.Error(w, "Failed to get customers: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := customerdto.CustomerListResponse{
		Data: customers,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetCustomerByID retrieves a customer by ID
// @Summary Get customer by ID
// @Description Get a customer by ID with family members
// @Tags customers
// @Accept json
// @Produce json
// @Param id path int true "Customer ID"
// @Success 200 {object} customerdto.CustomerResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /customers/{id} [get]
func (h *CustomerHandler) GetCustomerByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	customer, err := h.CustomerService.GetCustomerByID(r.Context(), id)
	if err != nil {
		http.Error(w, "Customer not found", http.StatusNotFound)
		return
	}

	response := customerdto.CustomerResponse{
		Data: *customer,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// UpdateCustomer updates a customer by ID
// @Summary Update customer by ID
// @Description Update a customer by ID with family members
// @Tags customers
// @Accept json
// @Produce json
// @Param id path int true "Customer ID"
// @Param customer body customerdto.CreateCustomerRequest true "Customer data"
// @Success 200 {object} customerdto.CustomerResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /customers/{id} [put]
func (h *CustomerHandler) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var req customerdto.CreateCustomerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	if err := h.Validator.Struct(req); err != nil {
		http.Error(w, "Validation failed: "+err.Error(), http.StatusBadRequest)
		return
	}

	customer, err := h.CustomerService.UpdateCustomer(r.Context(), id, req)
	if err != nil {
		http.Error(w, "Failed to update customer: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response := customerdto.CustomerResponse{
		Data: *customer,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// DeleteCustomer deletes a customer by ID
// @Summary Delete customer by ID
// @Description Delete a customer by ID and all family members
// @Tags customers
// @Accept json
// @Produce json
// @Param id path int true "Customer ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /customers/{id} [delete]
func (h *CustomerHandler) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	if err := h.CustomerService.DeleteCustomer(r.Context(), id); err != nil {
		http.Error(w, "Failed to delete customer: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Customer deleted successfully",
	})
}