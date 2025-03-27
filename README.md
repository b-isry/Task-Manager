# Task Manager API

A robust task management API built with Go, featuring JWT authentication and role-based access control.

## Features

- 🔐 JWT Authentication
- 👥 User Management with Role-Based Access Control
- ✅ Task CRUD Operations
- 🗄️ MongoDB Integration
- 🛡️ Input Validation
- ⚡ RESTful API Design

## Prerequisites

- Go 1.16 or higher
- MongoDB 4.4 or higher
- Git

## Installation

1. Clone the repository:
```bash
git clone https://github.com/Bisrat/task-manager.git
cd task-manager
```

2. Install dependencies:
```bash
go mod download
```

3. Set up environment variables:
```bash
# Windows PowerShell
$env:MONGODB_URI="mongodb://localhost:27017"
$env:JWT_SECRET="your-secure-secret-key"

# Windows Command Prompt
set MONGODB_URI=mongodb://localhost:27017
set JWT_SECRET=your-secure-secret-key

# Linux/Mac
export MONGODB_URI="mongodb://localhost:27017"
export JWT_SECRET="your-secure-secret-key"
```

4. Start MongoDB service

5. Run the application:
```bash
go run main.go
```

The server will start on `http://localhost:8080`

## Project Structure

```
task-manager/
├── controllers/     # HTTP request handlers
├── data/           # Database operations and services
├── docs/           # Documentation
├── errors/         # Custom error types
├── middleware/     # Authentication and authorization
├── models/         # Data models
├── router/         # Route definitions
├── main.go         # Application entry point
└── go.mod          # Go module file
```

## API Documentation

Detailed API documentation is available in the [docs/api_documentation.md](docs/api_documentation.md) file.

### Quick Start Guide

1. Register a new user:
```bash
curl -X POST http://localhost:8080/register \
-H "Content-Type: application/json" \
-d '{
    "name": "Test User",
    "email": "test@example.com",
    "password": "password123",
    "role": "user"
}'
```

2. Login to get a JWT token:
```bash
curl -X POST http://localhost:8080/login \
-H "Content-Type: application/json" \
-d '{
    "email": "test@example.com",
    "password": "password123"
}'
```

3. Use the token for protected endpoints:
```bash
curl -X GET http://localhost:8080/tasks \
-H "Authorization: Bearer <your_token>"
```

## Security Features

- Password hashing using bcrypt
- JWT-based authentication
- Role-based access control
- Input validation
- Secure headers
- MongoDB injection prevention

## Error Handling

The API implements comprehensive error handling for:
- Validation errors
- Authentication errors
- Authorization errors
- Database errors
- Not found errors

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [Gin](https://github.com/gin-gonic/gin) - Web framework
- [MongoDB Go Driver](https://github.com/mongodb/mongo-go-driver) - MongoDB driver
- [JWT-Go](https://github.com/golang-jwt/jwt) - JWT implementation 