# eCommerce REST API

This project is an eCommerce REST API built with Golang by muhammadderic. It includes services for user authentication and product management.

## Table of Contents

1. [Features](#features)
2. [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Installation](#installation)
    - [Database Migration](#database-migration)
3. [API Documentation](#api-documentation)
    - [User Service](#user-service)
    - [Product Service](#product-service)
4. [Built With](#built-with)
5. [Contributing](#contributing)
6. [License](#license)

## Features

- **User Service**:
  - User Registration
  - User Login with JWT tokenization
  - Form Validation

- **Product Service**:
  - Get All Products
  - Create a Product

## Getting Started

### Prerequisites

- Go 1.22+
- MySQL
- `golang-migrate` for database migration

### Installation

1. Clone the repository
   ```bash
   git clone https://github.com/muhammadderic/e_commerce-rest_api.git
   cd ecommerce-api
   ```

2. Install dependencies
   ```bash
   go mod tidy
   ```

3. Set up environment variables. Create a `.env` file in the root directory and add the following variables:
   ```env
   PUBLIC_HOST=http://localhost
   PORT=8080
   DB_USER=your_db_user
   DB_PASSWORD=your_db_password
   DB_HOST=localhost
   DB_PORT=3306
   DB_NAME=your_database_name
   JWT_SECRET=your_jwt_secret
   ```

### Database Migration

1. Install `golang-migrate`:
   ```bash
   go get github.com/golang-migrate/migrate/v4
   ```

2. Run the migrations:
   ```bash
   migrate -database "mysql://user:password@tcp(localhost:3306)/ecommerce" -path ./your_folder/migrations up
   ```

## API Documentation

### User Service

#### Register

`POST /register`

Register a new user.

**Request Body**:
```json
{
  "first_name": "John",
  "last_name": "Doe",
  "email": "john.doe@example.com",
  "password": "securepassword"
}
```

**Response**:
- `201 Created`

#### Login

`POST /login`

Authenticate a user and return a JWT token.

**Request Body**:
```json
{
  "email": "john.doe@example.com",
  "password": "securepassword"
}
```

**Response**:
```json
{
  "token": "your-jwt-token"
}
```

### Product Service

#### Get All Products

`GET /products`

Retrieve a list of all products.

**Response**:
```json
[
  {
    "id": 1,
    "name": "Product 1",
    "price": 100
  },
  {
    "id": 2,
    "name": "Product 2",
    "price": 150
  }
]
```

#### Create a Product

`POST /products`

Create a new product.

**Request Body**:
```json
{
  "name": "Product 1",
  "price": 100
}
```

**Response**:
- `201 Created`

## Built With

- [Golang](https://golang.org/) - The programming language used
- [mux](https://github.com/gorilla/mux) - HTTP router for Go
- [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql) - MySQL driver for Go
- [golang-migrate](https://github.com/golang-migrate/migrate) - Database migrations

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any changes.

## License

This project is licensed under the MIT License.