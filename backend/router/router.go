package router

import (
	"github.com/redha28/bookingtogo/internal/handlers"
	"github.com/redha28/bookingtogo/internal/services"
	"gorm.io/gorm"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func InitRouter(db *gorm.DB) *mux.Router {
	r := mux.NewRouter()
	
	// Initialize services
	customerService := services.NewCustomerService(db)
	nationalityService := services.NewNationalityService(db)
	familyService := services.NewFamilyService(db)
	
	// Initialize handlers
	customerHandler := handlers.NewCustomerHandler(customerService)
	nationalityHandler := handlers.NewNationalityHandler(nationalityService)
	familyHandler := handlers.NewFamilyHandler(familyService)
	
	// Swagger endpoint
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	
	// API routes
	api := r.PathPrefix("/api").Subrouter()
	
	// Customer routes
	api.HandleFunc("/customers", customerHandler.CreateCustomer).Methods("POST")
	api.HandleFunc("/customers", customerHandler.GetCustomers).Methods("GET")
	api.HandleFunc("/customers/{id}", customerHandler.GetCustomerByID).Methods("GET")
	api.HandleFunc("/customers/{id}", customerHandler.UpdateCustomer).Methods("PUT")
	api.HandleFunc("/customers/{id}", customerHandler.DeleteCustomer).Methods("DELETE")
	
	// Nationality routes
	api.HandleFunc("/nationalities", nationalityHandler.CreateNationality).Methods("POST")
	api.HandleFunc("/nationalities", nationalityHandler.GetNationalities).Methods("GET")
	api.HandleFunc("/nationalities/{id}", nationalityHandler.GetNationalityByID).Methods("GET")
	api.HandleFunc("/nationalities/{id}", nationalityHandler.UpdateNationality).Methods("PUT")
	api.HandleFunc("/nationalities/{id}", nationalityHandler.DeleteNationality).Methods("DELETE")
	
	// Family routes
	api.HandleFunc("/families", familyHandler.CreateFamily).Methods("POST")
	api.HandleFunc("/customers/{customer_id}/families", familyHandler.GetFamiliesByCustomerID).Methods("GET")
	api.HandleFunc("/families/{id}", familyHandler.GetFamilyByID).Methods("GET")
	api.HandleFunc("/families/{id}", familyHandler.UpdateFamily).Methods("PUT")
	api.HandleFunc("/families/{id}", familyHandler.DeleteFamily).Methods("DELETE")
	
	return r
}
