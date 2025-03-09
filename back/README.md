# Pissaze Backend API

This is a backend service built with [Gin](https://gin-gonic.com/), a high-performance HTTP web framework written in Go. The API provides product management, user authentication, and shopping cart functionality.

## Getting Started

### Prerequisites
- Go 1.21+
- PostgreSQL 15+ (with `pg_cron` extension)
- Environment variables configured as shown below

### Environment Setup
Create a `.env` file in the `pissaze/back/` directory:
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD="your_db_password"
DB_NAME=pissaze_system
JWT_SECRET_KEY="your_secret_key"
HOST="localhost"
PORT="8082"
```

### Installation

1. Clone the repository:
```bash
git clone https://github.com/SaeedMPro/pissaze.git
```
2. Database Setup 
- Create PostgreSQL database
```sql
CREATE DATABASE pissaze_system;
CREATE EXTENSION pg_cron;
```
- in `storage/` directory :
    - Run `createDB.sql query` for create tables 
    - Run `persian_dataset.sql` or `pissaze_dataset_test.aql` query for add dataset (Optional)


3. Install Dependency & run server
```bash
cd Pissaze/back
go mod tidy
go run main.go
```

## Project Structure

### Directories:
```
back/                         # Root directory
├── docs/                     # Swagger/API documentation
├── internal/                 # Internal application components
│   ├── dto/                  # Data Transfer Objects 
│   ├── middleware/           # Gin middleware 
│   ├── models/               # Database models/entities
│   ├── repositories/         # Database repository layer 
│   ├── server/               # HTTP server configuration and routes
│   ├── service/              # Business logic layer
│   ├── storage/              # Database connection and storage utilities
│   └── util/                 # Helper/utility functions
├── go.mod                    # Go module dependencies
├── go.sum                    # Go module checksums
├── main.go                   # Application entry point
├── .env                      # Environment variables
├── README.md                 # Project documentation
└── src.md                    # Source documentation            
```

### End points :

| Endpoint                 | Method   | Description         |
| :----------------------  | :------: | :-----------------  |
| /api/login               |   POST   | User authentication |
| /api/client              |   GET    | Client profile      |
| /api/client/cart	       |   GET    | Shopping carts      |
| /api/client/lockCart     |   GET    | Locked cart history |
| /api/product/list	       |   GET    | Product catalog     |
| /api/product/compatible	 |   POST   | Compatibility check | 



## Running Swagger API Documentation

Swagger is used for API documentation and testing. To access Swagger:

1. Start the server:
```bash
go run main.go
```
2. Open your browser and visit:
    
```
http://localhost:<PORT>/swagger/index.html
```

Swagger provides an interactive UI to test API endpoints.
