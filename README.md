# Fiber Backend Template

This is a simple backend template built with [Go Fiber v2](https://github.com/gofiber/fiber) to provide a clean foundation for new projects. It includes environment variable loading, basic CORS support, and request logging middleware.

## Features

- ⚡ Fast and lightweight Fiber framework
- 🌍 CORS enabled by default
- 🔐 `.env` configuration support
- 📜 Request logging with middleware
- 📘 Swagger (OpenAPI) documentation
- 🧪 Easily extendable for routes, services, and database layers

## Getting Started

### Prerequisites

- Go 1.18+
- Git

### Initialize the project

```bash
git clone https://github.com/patrikduch/fiber-be-template.git
cd fiber-be-template
go mod init github.com/patrikduch/fiber-be-template



## 📘 API Documentation (Swagger)

This project uses [Swaggo](https://github.com/swaggo/swag) to generate Swagger 2.0 API documentation and serve it via Fiber.

### 📦 Installation

Make sure you have the Swagger CLI tool installed:

```
go install github.com/swaggo/swag/cmd/swag@latest
```

Ensure `$GOPATH/bin` is in your `PATH` so `swag` can be run globally.

Then in your project root:

```
swag init
```

This generates the `docs/` folder containing Swagger definitions.

### ✍️ Documenting Handlers

Use special comments above your handler functions. Example:

```
```go
// GetUsers godoc
// @Summary Get all users
// @Description Returns list of users
// @Tags users
// @Produce json
// @Success 200 {array} models.User
// @Router /users [get]
func GetUsers(c *fiber.Ctx) error {
    ...
}
```
```

Each route should have:
- `@Summary`, `@Description`
- `@Tags` to group endpoints
- `@Accept` / `@Produce` if applicable
- `@Param` and `@Success` / `@Failure`
- `@Router` with method and path

### 🖥️ Accessing Swagger UI

In your `main.go`, Swagger is registered at:

```
```go
app.Get("/swagger/*", swagger.HandlerDefault)
```
```

Once the app is running, open:

```
http://localhost:3000/swagger/index.html
```

### 🔁 Regenerate Docs

Whenever you update Swagger annotations, run:

```
swag init
```

### 🧪 Example

Sample `CreateUser` annotation:

```
```go
// CreateUser godoc
// @Summary Create a new user
// @Description Accepts name and email to create a new user
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "User input"
// @Success 201 {object} models.User
// @Failure 400 {object} map[string]string
// @Router /users [post]
```
```
