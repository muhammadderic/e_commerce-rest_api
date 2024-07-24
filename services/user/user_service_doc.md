# User Service API Documentation

## Overview

This API provides endpoints for user authentication, including login and registration.

## Table of Contents

1. [Authentication](#authentication)
2. [Endpoints](#endpoints)
    - [Login](#login)
    - [Register](#register)
3. [Error Codes](#error-codes)
4. [Changelog](#changelog)

## Authentication

The `register` and `login` endpoints do not require authentication. Upon successful login, a JWT token will be provided, which should be included in the `Authorization` header for authenticated requests.

## Endpoints

### Login

`POST /login`

Authenticate a user and return a JWT token.

#### Request Headers

| Header       | Type   | Description        |
| ------------ | ------ | ------------------ |
| Content-Type | String | `application/json` |

#### Request Body

```json
{
  "email": "user@example.com",
  "password": "userpassword"
}
```

| Field    | Type   | Description         |
| -------- | ------ | ------------------- |
| email    | String | The user's email.   |
| password | String | The user's password.|

#### Response

```json
{
  "token": "your-jwt-token"
}
```

| Field | Type   | Description         |
| ----- | ------ | ------------------- |
| token | String | The JWT token.      |

#### Example cURL

```bash
curl -X POST https://api.example.com/login \
     -H "Content-Type: application/json" \
     -d '{
           "email": "user@example.com",
           "password": "userpassword"
         }'
```

### Register

`POST /register`

Register a new user.

#### Request Headers

| Header       | Type   | Description        |
| ------------ | ------ | ------------------ |
| Content-Type | String | `application/json` |

#### Request Body

```json
{
  "first_name": "John",
  "last_name": "Doe",
  "email": "john.doe@example.com",
  "password": "securepassword"
}
```

| Field      | Type   | Description               |
| ---------- | ------ | ------------------------- |
| first_name | String | The user's first name.    |
| last_name  | String | The user's last name.     |
| email      | String | The user's email.         |
| password   | String | The user's password.      |

#### Response

- **Status**: `201 Created`
- **Body**: `null`

#### Example cURL

```bash
curl -X POST https://api.example.com/register \
     -H "Content-Type: application/json" \
     -d '{
           "first_name": "John",
           "last_name": "Doe",
           "email": "john.doe@example.com",
           "password": "securepassword"
         }'
```

## Error Codes

| Code | Description                         |
| ---- | ----------------------------------- |
| 400  | Bad Request - Invalid input         |
| 401  | Unauthorized - Invalid credentials  |
| 404  | Not Found - User does not exist     |
| 500  | Internal Server Error               |

### Example Error Response

#### Status: `400 Bad Request`

```json
{
  "error": "invalid request: some error details"
}
```

## Changelog

### 2024-07-24
- Initial release of the user service API.