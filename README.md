# 🎟️ GoTicket — Ticket Booking API

GoTicket is a backend REST API for a ticket booking system built with **Go**. The project follows a clean, modular architecture designed to keep business logic, database access, and HTTP handling separated and maintainable.

## 🚀 Tech Stack

- **Go**
- **Echo v5** — HTTP web framework
- **GORM** — ORM
- **PostgreSQL** — Database
- **Neon PostgreSQL** — Cloud database
- **Go Playground Validator** — Request validation
- **JWT** — Authentication
- **REST API**
- **Layered Architecture**

## 📁 Project Structure

```text
goticket/
│
├── cmd/
│   └── main.go
│
├── internal/
│   ├── config/
│   │   ├── config.go
│   │   └── database.go
│   │
│   ├── server/
│   │   └── server.go
│   │
│   ├── httpresponse/
│   │   └── error.go
│   │
│   └── user/
│       ├── dto/
│       │   └── user.go
│       ├── errors.go
│       ├── handler.go
│       ├── model.go
│       ├── repository.go
│       └── service.go
│
├── .env
├── .gitignore
├── go.mod
├── go.sum
└── README.md
```

## 🏗️ Architecture

GoTicket follows a layered architecture:

```text
                Client
                  │
                  ▼
             HTTP Request
                  │
                  ▼
              Handler
                  │
                  ▼
               Service
                  │
                  ▼
             Repository
                  │
                  ▼
              Database
```

### Handler

Responsible for:

- Receiving HTTP requests
- Binding request data
- Request validation
- Returning HTTP responses

### Service

Responsible for:

- Business logic
- Processing application rules
- Communicating with repositories

### Repository

Responsible for:

- Database operations
- Creating records
- Querying records
- Handling database-specific errors

## ⚙️ Requirements

Before running the project, make sure you have installed:

- Go 1.27+
- PostgreSQL or a Neon PostgreSQL database
- Git

Check your Go version:

```bash
go version
```

## 📦 Installation

Clone the repository:

```bash
git clone https://github.com/your-username/goticket.git
```

Move into the project:

```bash
cd goticket
```

Install dependencies:

```bash
go mod tidy
```

## 🔐 Environment Variables

Create a `.env` file in the project root:

```env
DATABASE_URL=your_database_connection_string
PORT=8080
```

Example:

```env
DATABASE_URL=postgresql://username:password@host/database?sslmode=require
PORT=8080
```

> Never commit your `.env` file or database credentials to GitHub.

Add this to `.gitignore`:

```gitignore
.env
```

## ▶️ Run the Project

Start the development server:

```bash
go run ./cmd
```

Or, if your `main.go` is in the root:

```bash
go run .
```

The API will run at:

```text
http://localhost:8080
```

## 🗄️ Database

GoTicket uses **PostgreSQL** with **GORM**.

Database connection is handled through the configuration package:

```go
db := config.ConnectDatabase(cfg)
```

Models can be automatically migrated using GORM:

```go
db.AutoMigrate(&User{})
```

## 👤 User API

### Create User

**POST**

```text
/users
```

### Request

```json
{
  "name": "Zamirul Kabir",
  "email": "zamirul@example.com",
  "password": "Admin1234@"
}
```

### Successful Response

```json
{
  "id": 1,
  "name": "Zamirul Kabir",
  "email": "zamirul@example.com",
  "createdAt": "2026-10-08T10:30:00Z"
}
```

### Duplicate Email Response

If the email already exists:

```json
{
  "code": 409,
  "message": "Email already exists",
  "details": "email already exists"
}
```

HTTP status:

```text
409 Conflict
```

## 📡 API Endpoints

| Method | Endpoint              | Description       |
| ------ | --------------------- | ----------------- |
| GET    | `/`                   | API health check  |
| POST   | `/users`              | Create a new user |
| GET    | `/users/:id`          | Get user by ID    |
| GET    | `/users/email/:email` | Get user by email |
| POST   | `/auth/login`         | User login        |

> Additional ticket, booking, authentication, and event endpoints will be added as development continues.

## 🧩 User Model

The current user model contains:

```go
type User struct {
    gorm.Model

    Name     string `json:"name" gorm:"type:varchar(100);not null"`
    Email    string `json:"email" gorm:"type:varchar(100);unique;not null"`
    Password string `json:"password" gorm:"type:varchar(100);not null"`
}
```

GORM provides:

```text
ID
CreatedAt
UpdatedAt
DeletedAt
```

through:

```go
gorm.Model
```

## 🔒 Error Handling

GoTicket uses application-level errors for common business cases.

Example:

```go
var ErrorEmailAlreadyExists = errors.New("email already exists")
```

Repository errors are converted into domain-level errors:

```go
if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
    return ErrorEmailAlreadyExists
}
```

The handler can then return the appropriate HTTP status:

```text
409 Conflict
```

## 🧪 Testing

Run all tests:

```bash
go test ./...
```

Run tests with verbose output:

```bash
go test -v ./...
```

## 🛠️ Development

Format Go files:

```bash
gofmt -w .
```

Check the project:

```bash
go vet ./...
```

Download/update dependencies:

```bash
go mod tidy
```

## 🔑 Planned Features

### Authentication

- [ ] User registration
- [ ] Password hashing
- [ ] Login
- [ ] JWT authentication
- [ ] Refresh tokens
- [ ] Logout

### Users

- [x] Create user
- [x] Unique email
- [ ] Get user
- [ ] Update user
- [ ] Delete user
- [ ] User profile

### Events

- [ ] Create event
- [ ] Update event
- [ ] Delete event
- [ ] List events
- [ ] Event details
- [ ] Event categories

### Tickets

- [ ] Create ticket type
- [ ] Ticket availability
- [ ] Ticket pricing
- [ ] Ticket inventory
- [ ] Ticket purchase

### Booking

- [ ] Create booking
- [ ] Booking confirmation
- [ ] Booking cancellation
- [ ] Booking history
- [ ] Booking status

### Payment

- [ ] Payment integration
- [ ] Payment verification
- [ ] Transaction history
- [ ] Refund handling

## 🔮 Future Improvements

- Redis caching
- Background jobs
- Email notifications
- QR code ticket generation
- Payment gateway integration
- Rate limiting
- API documentation with Swagger/OpenAPI
- Docker support
- CI/CD pipeline
- Automated tests
- Structured logging
- Production monitoring

## 📌 Development Status

🚧 **GoTicket is currently under active development.**

The project is being developed incrementally, starting with the user management and authentication system before expanding into events, tickets, bookings, and payments.

## 👨‍💻 Author

**Zamirul Kabir**

Full-Stack Developer

- GitHub: `github.com/alive1258`
- LinkedIn: `linkedin.com/in/zamirul-kabir-575a41279/`

---

⭐ If you find this project useful, consider giving it a star!
