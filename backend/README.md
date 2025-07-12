# Customer Family API

A RESTful API for managing customers and their family members built with Go, Gorilla Mux, GORM, and
PostgreSQL.

## Features

- Complete CRUD operations for Customers, Nationalities, and Family members
- Swagger documentation
- Input validation
- Database relationships with foreign keys
- Clean architecture with separated services and handlers
- PostgreSQL database with connection pooling

## Quick Start

### Prerequisites

- Go 1.19 or later
- PostgreSQL database
- Swagger CLI (optional, for regenerating docs)
- Fresh (for live reload development, optional)

### Installation

1. Clone the repository
2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Set up PostgreSQL database:

   ```sql
   CREATE DATABASE customerdb;
   ```

4. Update database connection in `config/db.go` if needed

5. Generate Swagger documentation:

   ```bash
   swag init -g cmd/main.go
   ```

6. Run the application:

   ```bash
   go run cmd/main.go
   ```

   Or for development with live reload:

   ```bash
   fresh
   ```

The server will start on `http://localhost:8080`

## API Documentation

### Swagger UI

Access the interactive API documentation at: `http://localhost:8080/swagger/index.html`

### API Endpoints

#### Customers

- `POST /api/customers` - Create a new customer with family members
- `GET /api/customers` - Get all customers with their families
- `GET /api/customers/{id}` - Get customer by ID with family members
- `PUT /api/customers/{id}` - Update customer and family members
- `DELETE /api/customers/{id}` - Delete customer and all family members

#### Nationalities

- `POST /api/nationalities` - Create a new nationality
- `GET /api/nationalities` - Get all nationalities
- `GET /api/nationalities/{id}` - Get nationality by ID
- `PUT /api/nationalities/{id}` - Update nationality
- `DELETE /api/nationalities/{id}` - Delete nationality

#### Family Members

- `POST /api/families` - Create a new family member
- `GET /api/customers/{customer_id}/families` - Get family members by customer ID
- `GET /api/families/{id}` - Get family member by ID
- `PUT /api/families/{id}` - Update family member
- `DELETE /api/families/{id}` - Delete family member

## Example Usage

### Create a Nationality

```bash
curl -X POST http://localhost:8080/api/nationalities \
  -H "Content-Type: application/json" \
  -d '{
    "nationality_name": "Indonesian",
    "nationality_code": "ID"
  }'
```

### Create a Customer with Family

```bash
curl -X POST http://localhost:8080/api/customers \
  -H "Content-Type: application/json" \
  -d '{
    "cst_name": "John Doe",
    "cst_email": "john@example.com",
    "cst_dob": "1990-01-01",
    "cst_phoneNum": "123456789",
    "nationality_id": 1,
    "family_list": [
      {
        "fl_name": "Jane Doe",
        "fl_dob": "1992-02-02",
        "fl_relation": "Wife"
      }
    ]
  }'
```

### Get All Customers

```bash
curl http://localhost:8080/api/customers
```

## Database Schema

### Tables

#### nationality

- `nationality_id` (SERIAL PRIMARY KEY)
- `nationality_name` (VARCHAR(50) NOT NULL)
- `nationality_code` (CHAR(2) NOT NULL)

#### customer

- `cst_id` (SERIAL PRIMARY KEY)
- `nationality_id` (INT NOT NULL, FOREIGN KEY)
- `cst_name` (VARCHAR(50) NOT NULL)
- `cst_dob` (DATE NOT NULL)
- `cst_phone_num` (VARCHAR(20) NOT NULL)
- `cst_email` (VARCHAR(50) NOT NULL)

#### family

- `fl_id` (SERIAL PRIMARY KEY)
- `cst_id` (INT NOT NULL, FOREIGN KEY)
- `fl_relation` (VARCHAR(50) NOT NULL)
- `fl_name` (VARCHAR(50) NOT NULL)
- `fl_dob` (DATE NOT NULL)

## Project Structure

```
.
├── cmd/
│   └── main.go                 # Application entry point
├── config/
│   └── db.go                   # Database configuration
├── internal/
│   ├── dto/
│   │   └── customer/
│   │       ├── request.go      # Request DTOs
│   │       └── response.go     # Response DTOs
│   ├── entities/
│   │   ├── customer.go         # Customer entity
│   │   ├── family.go           # Family entity
│   │   └── nationality.go      # Nationality entity
│   ├── handlers/
│   │   ├── customer.handler.go # Customer HTTP handlers
│   │   ├── family.handler.go   # Family HTTP handlers
│   │   └── nationality.handler.go # Nationality HTTP handlers
│   └── services/
│       ├── customer.service.go # Customer business logic
│       ├── family.service.go   # Family business logic
│       └── nationality.service.go # Nationality business logic
├── router/
│   └── router.go               # HTTP routes configuration
├── docs/                       # Generated Swagger documentation
├── Makefile                    # Build automation
└── README.md
```

## Development

### Live Reload Development

For development with automatic reload on file changes:

```bash
make dev
```

Or directly:

```bash
fresh
```

Fresh will automatically reload the application when you make changes to Go files.

### Makefile Commands

- `make docs` - Generate Swagger documentation
- `make run` - Run the application
- `make build` - Build the application
- `make clean` - Clean generated files
- `make deps` - Install/update dependencies
- `make dev` - Run with live reload using Fresh
- `make docs-run` - Generate docs and run
- `make help` - Show help message

### Regenerating Swagger Documentation

After making changes to the API handlers, regenerate the documentation:

```bash
make docs
```

Or manually:

```bash
swag init -g cmd/main.go
```

## Technologies Used

- **Go** - Programming language
- **Gorilla Mux** - HTTP router
- **GORM** - ORM library
- **PostgreSQL** - Database
- **Swagger** - API documentation
- **pgxpool** - PostgreSQL connection pooling
- **Validator** - Input validation
- **Fresh** - Live reload for development

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Update documentation
6. Submit a pull request

## License

MIT License
