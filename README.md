# Fiber Backend Template

This is a simple backend template built with [Go Fiber v2](https://github.com/gofiber/fiber) to provide a clean foundation for new projects. It includes environment variable loading, basic CORS support, request logging middleware, Ent ORM integration, and CQRS architecture support.

## Features

- ⚡ Fast and lightweight Fiber framework
- 🧱 Built-in Ent ORM support for schema-first database modeling
- 🧭 CQRS (Command Query Responsibility Segregation) pattern support
- 🌍 CORS enabled by default
- 🔐 `.env` configuration support
- 📜 Request logging with middleware
- 📘 Swagger (OpenAPI) documentation
- 🧪 Easily extendable for routes, services, and database layers
- 🐳 Docker-based production deployment

---

## 🚀 Getting Started

### Prerequisites

- Go 1.18+
- Git

### Initialize the project

```bash
git clone https://github.com/patrikduch/fiber-be-template.git
cd fiber-be-template
go mod tidy
```

---

## 🧱 ENT (Entity Framework for Go)

This project uses [Ent](https://entgo.io/) as the ORM.

### 📦 Install Ent Codegen

```bash
go install entgo.io/ent/cmd/ent@latest
```

### ✍️ Define Schemas

Create or modify your Ent schema files inside `./ent/schema`.

Example file:

```go
package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
)

type User struct {
    ent.Schema
}

func (User) Fields() []ent.Field {
    return []ent.Field{
        field.String("username").NotEmpty(),
        field.String("email").Unique(),
    }
}
```

### ⚙️ Generate Ent Code

After modifying your schemas, run:

```bash
go run entgo.io/ent/cmd/ent generate ./ent/schema
```

This generates the necessary Ent code in the `ent/` directory.

---

## 🧭 CQRS Pattern

This project supports the **CQRS** (Command Query Responsibility Segregation) pattern by organizing logic into `queries/` and `commands/` folders.

### 📂 Recommended Folder Structure

```
queries/
└── get_all_users/
    ├── query.go      // The Query struct (input)
    └── handler.go    // The Handler logic (output)
```

### 🛠️ Example: GetAllUsers

**`query.go`**

```go
package get_all_users

type Query struct{}
```

**`handler.go`**

```go
package get_all_users

import (
    "context"
    "fmt"

    "fiber-be-template/database"
    "fiber-be-template/dtos/users/responses"
    "fiber-be-template/mappers/users"
    "fiber-be-template/models"
)

type Handler struct{}

func NewHandler() *Handler {
    return &Handler{}
}

func (h *Handler) Handle(ctx context.Context, _ Query) ([]responses.UserResponseDto, error) {
    entUsers, err := database.EntClient.User.Query().All(ctx)
    if err != nil {
        return nil, fmt.Errorf("failed querying users: %w", err)
    }

    result := make([]responses.UserResponseDto, len(entUsers))
    for i, entUser := range entUsers {
        u := models.User{
            ID:    entUser.ID,
            Name:  entUser.Username,
            Email: entUser.Email,
        }
        result[i] = users.ToUserResponseDto(u)
    }

    return result, nil
}
```

**Using in Controller:**

```go
package controllers

import (
    "context"
    "github.com/gofiber/fiber/v2"
    "fiber-be-template/queries/get_all_users"
)

var getAllUsersHandler = get_all_users.NewHandler()

func GetUsers(c *fiber.Ctx) error {
    result, err := getAllUsersHandler.Handle(context.Background(), get_all_users.Query{})
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(result)
}
```

---

## 📘 API Documentation (Swagger)

This project uses [Swaggo](https://github.com/swaggo/swag) to generate Swagger 2.0 API documentation and serve it via Fiber.

### 📦 Installation

Install the `swag` CLI:

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

Ensure `$GOPATH/bin` is in your `PATH`.

### 🛠️ Generate Swagger Docs

```bash
swag init
```

This will generate the `docs/` folder containing Swagger JSON definitions.

### ✍️ Documenting Handlers

Use structured comments above your route handlers:

```go
// GetUsers godoc
// @Summary Get all users
// @Description Returns list of users
// @Tags users
// @Produce json
// @Success 200 {array} responses.UserResponseDto
// @Router /api/users [get]
```

Each route should have:
- `@Summary`, `@Description`
- `@Tags` to group
- `@Accept` / `@Produce` as needed
- `@Param` and `@Success` / `@Failure`
- `@Router` with method and path

### 🖥️ Accessing Swagger UI

In your `main.go`, register the route:

```go
app.Get("/swagger/*", swagger.HandlerDefault)
```

Then open:

```
http://localhost:3000/swagger/index.html
```

### 🔁 Regenerate Docs

After changes to your route annotations:

```bash
swag init
```



## 🐳 Docker Deployment (Production-Ready)

This project includes a production-ready Dockerfile using a multi-stage build for minimal image size and fast startup.

---

### 🛠️ Build the Docker Image

```bash
docker build -t fiber-be-template:latest .
```



## 🛠️ Contribution

Pull requests are welcome. For major changes, please open an issue first to discuss what you'd like to change.

---

## 📄 License

This project is licensed under the GNU General Public License v3.0.  
See the [LICENSE](LICENSE) file for details.