package main

import (
	"log"
	"net/http"

	_ "github.com/joho/godotenv/autoload"
	"github.com/redha28/bookingtogo/config"
	_ "github.com/redha28/bookingtogo/docs"
	"github.com/redha28/bookingtogo/internal/entities"
	"github.com/redha28/bookingtogo/router"
)

// @title Customer Family API
// @version 1.0
// @description This is a sample API for managing customers and their family members.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /api
func main() {
	config.InitDB()
	defer config.CloseDB()

	config.DB.AutoMigrate(&entities.Nationality{}, &entities.Customer{}, &entities.Family{})

	r := router.InitRouter(config.DB)

	log.Println("🚀 Server running at http://localhost:8080")
	log.Println("📚 Swagger docs available at http://localhost:8080/swagger/index.html")
	http.ListenAndServe(":8080", r)
}
