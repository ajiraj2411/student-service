# 🧑‍🎓 Student Service (Go + Gin + PostgreSQL)

A clean, production-style REST API built with Go, Gin, and PostgreSQL, following layered architecture with proper unit testing and performance testing.

This project is designed to be interview-ready and demonstrates real-world backend engineering practices.

## 🚀 Features

- CRUD APIs for Student management
- Clean layered architecture (Handler → Service → Repository)
- Dependency Injection using interfaces
- API Key–based authentication middleware
- PostgreSQL integration
- Unit tests (mock-based)
- Performance & load testing
- Production-style project structure

## 🏗️ Project Structure

```
student-service/
├── main.go
├── go.mod
├── go.sum
├── config/
│   └── db.go
├── model/
│   └── student.go
├── repository/
│   └── student_repository.go
├── service/
│   ├── student_service.go
│   └── student_service_test.go
├── handler/
│   ├── student_handler.go
│   └── student_handler_test.go
├── middleware/
│   └── api_key.go
├── routes/
│   └── routes.go
├── docs/
│   └── swagger.yaml
└── README.md
```

## 📦 Tech Stack

- Go
- Gin (HTTP framework)
- PostgreSQL
- database/sql
- testing / httptest
- hey (load testing)

## 🔐 API Authentication

All endpoints require an API Key header:

```
X-API-Key: my-secret-key
```

## 🧩 Database Setup

Set environment variable:
```bash
export DATABASE_URL="postgres://username:password@localhost:5432/studentdb?sslmode=disable"
```

Create table:
```sql
CREATE TABLE students (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    age INT NOT NULL
);
```

## ▶️ Run the Application

```bash
go run main.go
```

## 📖 API Documentation

The API is fully documented using Swagger/OpenAPI. Once the server is running, you can access the interactive API documentation at:

```
http://localhost:8080/swagger/index.html
```

Alternatively, you can use the Postman collection for testing the API endpoints:

[📮 Postman Collection](https://documenter.getpostman.com/view/4451782/2sB3dTtTc4)

### API Endpoints Overview

- **GET /students** - Retrieve all students
- **POST /students** - Create a new student
- **PUT /students/{id}** - Update an existing student by ID
- **DELETE /students/{id}** - Delete a student by ID

### Student Model
```json
{
  "id": 1,
  "name": "John Doe",
  "age": 25
}
```

## 📌 API Endpoints

**Create Student**
```bash
curl -X POST http://localhost:8080/students \
  -H "Content-Type: application/json" \
  -H "X-API-Key: my-secret-key" \
  -d '{"name":"Ajith","age":29}'
```

**Get All Students**
```bash
curl -X GET http://localhost:8080/students \
  -H "X-API-Key: my-secret-key"
```

**Update Student**
```bash
curl -X PUT http://localhost:8080/students/1 \
  -H "Content-Type: application/json" \
  -H "X-API-Key: my-secret-key" \
  -d '{"name":"Ajith Updated","age":30}'
```

**Delete Student**
```bash
curl -X DELETE http://localhost:8080/students/1 \
  -H "X-API-Key: my-secret-key"
```

## 🧪 Unit Testing

```bash
go test ./...
```

Tests cover:
- Service layer with mocked repository
- Handler layer with httptest and mocked service
- No real database required

## ⚡ Performance Testing

Install hey:
```bash
brew install hey
```

Run load test:
```bash
hey -n 1000 -c 50 \
  -H "X-API-Key: my-secret-key" \
  http://localhost:8080/students
```

## 🧠 Architecture Overview

```
Client
  ↓
Handler (HTTP)
  ↓
Service (Business Logic)
  ↓
Repository (Database)
  ↓
PostgreSQL
```

**Benefits:**
- Easy unit testing
- Loose coupling
- Clear responsibility per layer

## 💬 Interview Talking Points

- Interface-driven design for testability
- Dependency Injection via constructors
- Proper error handling & status codes
- Separation of concerns
- Performance benchmarking using hey

## 🏁 Future Improvements

- Pagination & filtering
- Docker & docker-compose
- JWT authentication
- Integration tests using Testcontainers

## 👨‍💻 Author

Ajith Kumar  
[GitHub](https://github.com/ajiraj2411)

---

**Built as part of Go backend interview preparation and hands-on learning.**