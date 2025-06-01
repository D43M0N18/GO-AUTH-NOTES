# Go Auth Notes API

A secure REST API built with Golang (Fiber) for user authentication and personal notes management. This API allows users to register, login, and manage their personal notes with JWT authentication.

## Features

- User authentication (register/login) with JWT
- Secure password hashing using bcrypt
- CRUD operations for personal notes
- MySQL database with GORM
- Docker and Docker Compose setup
- Environment configuration using .env file

## Tech Stack

- **Backend**: Golang with Fiber framework
- **Database**: MySQL 8.0
- **ORM**: GORM
- **Authentication**: JWT
- **Containerization**: Docker & Docker Compose
- **Password Hashing**: bcrypt

## Prerequisites

- Docker and Docker Compose
- Go 1.21 or later (for local development)

## Project Structure

```
.
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
├── main.go
├── handlers/
│   ├── auth.go
│   └── note.go
├── middleware/
│   └── auth.go
├── models/
│   ├── user.go
│   └── note.go
├── routes/
│   └── routes.go
└── utils/
    ├── database.go
    └── jwt.go
```

## Setup and Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/D43M0N18/GO-AUTH-NOTES/tree/main
   cd go-auth-notes
   ```

2. Create a `.env` file in the root directory:
   ```
   DB_HOST=mysql
   DB_USER=notes_user
   DB_PASSWORD=notes_password
   DB_NAME=notes_db
   DB_ROOT_PASSWORD=root_password
   JWT_SECRET=your-secret-key-here
   PORT=8000
   ```

3. Build and run with Docker Compose:
   ```bash
   docker-compose up --build
   ```

The API will be available at `http://localhost:8000`

## API Endpoints

### Authentication

#### Register User
- **POST** `/register`
- **Body**:
  ```json
  {
    "name": "John Doe",
    "email": "john@example.com",
    "password": "yourpassword"
  }
  ```

#### Login
- **POST** `/login`
- **Body**:
  ```json
  {
    "email": "john@example.com",
    "password": "yourpassword"
  }
  ```
- **Response**: JWT token

### Notes (Protected Routes - Requires JWT)

#### Create Note
- **POST** `/notes`
- **Headers**: `Authorization: Bearer <jwt_token>`
- **Body**:
  ```json
  {
    "title": "My Note",
    "content": "Note content here"
  }
  ```

#### Get All Notes
- **GET** `/notes`
- **Headers**: `Authorization: Bearer <jwt_token>`

#### Get Single Note
- **GET** `/notes/:id`
- **Headers**: `Authorization: Bearer <jwt_token>`

#### Update Note
- **PUT** `/notes/:id`
- **Headers**: `Authorization: Bearer <jwt_token>`
- **Body**:
  ```json
  {
    "title": "Updated Title",
    "content": "Updated content"
  }
  ```

#### Delete Note
- **DELETE** `/notes/:id`
- **Headers**: `Authorization: Bearer <jwt_token>`

## Error Handling

The API returns appropriate HTTP status codes and error messages:

- `400 Bad Request`: Invalid input data
- `401 Unauthorized`: Invalid or missing JWT token
- `403 Forbidden`: Insufficient permissions
- `404 Not Found`: Resource not found
- `409 Conflict`: Resource already exists
- `500 Internal Server Error`: Server-side error

## Security Features

- Password hashing using bcrypt
- JWT-based authentication
- Protected routes with middleware
- Input validation
- Secure database connections
- Environment variable configuration

## Development

For local development without Docker:

1. Install Go dependencies:
   ```bash
   go mod download
   ```

2. Run the application:
   ```bash
   go run main.go
   ```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details. 
