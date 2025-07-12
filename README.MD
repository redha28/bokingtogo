# BookingTogo - Customer Family Management System

A full-stack application for managing customers and their family members, built with Go backend API
and Laravel frontend.

## 🏗️ Architecture

This project consists of two main components:

- **Backend API**: Go-based REST API with PostgreSQL database
- **Frontend**: Laravel web application with Tailwind CSS

## 🛠️ Technology Stack

### Backend (Go API)

- **Language**: Go 1.23.5
- **Framework**: Gorilla Mux (HTTP router)
- **Database**: PostgreSQL
- **ORM**: GORM
- **Documentation**: Swagger/OpenAPI
- **Validation**: Go Playground Validator
- **Environment**: GoDotEnv

### Frontend (Laravel)

- **Language**: PHP 8.2+
- **Framework**: Laravel 12.0 (Full-stack framework - no separate backend needed)
- **CSS Framework**: Tailwind CSS 4.0
- **Build Tool**: Vite 6.2.4
- **Package Manager**: Composer & NPM

### Database

- **Database**: PostgreSQL
- **Tables**:
  - `nationality` - Store nationality information
  - `customer` - Store customer details
  - `family` - Store family member information

## 📋 Prerequisites

Before running this application, make sure you have the following installed:

- Go 1.23.5 or higher
- PHP 8.2 or higher
- PostgreSQL
- Composer (PHP package manager)
- Node.js and NPM
- Git

## 🚀 Installation & Setup

### 1. Clone the Repository

```bash
git clone https://github.com/redha28/bookingtogo.git
cd bookingtogo
```

### 2. Database Setup

1. Create a PostgreSQL database:

```sql
CREATE DATABASE bookingtogo;
```

2. Import the database schema:

```bash
psql -U your_username -d bookingtogo -f bokingtogo.sql
```

### 3. Backend Setup (Go API)

1. Navigate to the backend directory:

```bash
cd backend
```

2. Install Go dependencies:

```bash
go mod download
```

3. Create a `.env` file in the backend directory:

```env
DB_URL=postgres://postgres:postgres@localhost:5432/bokingtogo?sslmode=disable
```

4. Generate Swagger documentation:

```bash
make docs
```

5. Run the backend server:

```bash
make run
```

The API will be available at `http://localhost:8080`

### 4. Frontend Setup (Laravel)

1. Navigate to the Laravel directory:

```bash
cd ../laravel
```

2. Install PHP dependencies:

```bash
composer install
```

3. Install Node.js dependencies:

```bash
npm install
```

4. Create a `.env` file:

```bash
cp .env.example .env
```

5. Generate application key:

```bash
php artisan key:generate
```

6. Configure the `.env` file with your database settings:

```env
DB_CONNECTION=pgsql
DB_HOST=127.0.0.1
DB_PORT=5432
DB_DATABASE=bookingtogo
DB_USERNAME=your_username
DB_PASSWORD=your_password
```

7. Build frontend assets:

```bash
npm run build
```

8. Start the Laravel development server:

```bash
php artisan serve --port=3000
```

The web application will be available at `http://localhost:3000`

## 🔧 Development

### Backend Development

1. **Hot Reload**: Use `make run` to start the server with automatic restart on file changes.

2. **API Documentation**: Access Swagger UI at `http://localhost:8080/swagger/index.html`

3. **Available Make Commands**:
   - `make docs` - Generate Swagger documentation
   - `make run` - Run the application
   - `make build` - Build the application
   - `make clean` - Clean build artifacts

### Frontend Development

1. **Hot Reload**: Use `npm run dev` for development with hot reload:

```bash
npm run dev
```

2. **Build for Production**:

```bash
npm run build
```

3. **Laravel Commands**:
   - `php artisan serve --port=3000` - Start development server on port 3000
   - `php artisan migrate` - Run database migrations
   - `php artisan tinker` - Access Laravel REPL

## 📊 Database Schema

### Nationality Table

- `nationality_id` (SERIAL PRIMARY KEY)
- `nationality_name` (VARCHAR(50))
- `nationality_code` (CHAR(2))

### Customer Table

- `cst_id` (SERIAL PRIMARY KEY)
- `nationality_id` (INT, Foreign Key)
- `cst_name` (VARCHAR(50))
- `cst_dob` (DATE)
- `cst_phone_num` (VARCHAR(20))
- `cst_email` (VARCHAR(50))

### Family Table

- `fl_id` (SERIAL PRIMARY KEY)
- `cst_id` (INT, Foreign Key)
- `fl_relation` (VARCHAR(50))
- `fl_name` (VARCHAR(50))
- `fl_dob` (DATE)

## 🔍 API Endpoints

The Go backend provides RESTful API endpoints for:

- **Nationality Management**
  - GET, POST, PUT, DELETE operations
- **Customer Management**
  - GET, POST, PUT, DELETE operations
- **Family Management**
  - GET, POST, PUT, DELETE operations

Access the complete API documentation at `http://localhost:8080/swagger/index.html` when the backend
is running.

## 📝 Project Structure

```
bookingtogo/
├── backend/                 # Go API backend
│   ├── cmd/                # Application entry point
│   ├── config/             # Configuration files
│   ├── internal/           # Internal packages
│   │   ├── dto/           # Data Transfer Objects
│   │   ├── entities/      # Database entities
│   │   ├── handlers/      # HTTP handlers
│   │   └── services/      # Business logic
│   ├── router/            # HTTP routing
│   └── docs/              # Swagger documentation
├── laravel/               # Laravel frontend
│   ├── app/               # Application code
│   ├── config/            # Configuration files
│   ├── database/          # Database migrations/seeds
│   ├── public/            # Public assets
│   ├── resources/         # Views, CSS, JS
│   └── routes/            # Route definitions
├── bokingtogo.sql         # Database schema
└── README.MD              # This file
```

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## 📄 License

This project is licensed under the MIT License.

## 👥 Contact

- **Developer**: Redha
- **Repository**: https://github.com/redha28/bookingtogo

---

For more detailed information about specific components, please refer to the README files in the
`backend/` and `laravel/` directories.
