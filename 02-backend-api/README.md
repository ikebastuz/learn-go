# Calculator API

A simple REST API calculator service written in Go. This project demonstrates Go best practices, including:

- Clean architecture and project structure
- Proper error handling
- Middleware for logging and recovery
- Configuration management
- Graceful shutdown
- Input validation
- Structured logging

## Features

- Basic arithmetic operations (add, subtract, multiply, divide)
- Sum of an array of numbers
- Input validation
- Structured logging
- Graceful shutdown
- Configuration management

## Prerequisites

- Go 1.22 or later
- Make (optional, for using Makefile commands)

## Project Structure

```
02-backend-api/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── api/
│   │   ├── handlers.go
│   │   └── middleware.go
│   ├── calculator/
│   │   ├── operations.go
│   │   └── types.go
│   └── config/
│       └── config.go
├── config/
│   └── config.yaml
├── api-spec.yaml
├── go.mod
└── README.md
```

## Installation

1. Clone the repository:

   ```bash
   git clone <repository-url>
   cd 02-backend-api
   ```

2. Install dependencies:

   ```bash
   go mod download
   ```

3. Build the application:
   ```bash
   go build -o calculator ./cmd/server
   ```

## Running the Application

1. Start the server:

   ```bash
   ./calculator
   ```

   Or run directly:

   ```bash
   go run ./cmd/server
   ```

2. The server will start on port 1337 by default. You can change this in `config/config.yaml`.

## API Endpoints

### Add

```http
POST /add
Content-Type: application/json

{
    "number1": 5,
    "number2": 3
}
```

### Subtract

```http
POST /subtract
Content-Type: application/json

{
    "number1": 5,
    "number2": 3
}
```

### Multiply

```http
POST /multiply
Content-Type: application/json

{
    "number1": 5,
    "number2": 3
}
```

### Divide

```http
POST /divide
Content-Type: application/json

{
    "number1": 6,
    "number2": 2
}
```

### Sum

```http
POST /sum
Content-Type: application/json

[1, 2, 3, 4, 5]
```

## Configuration

The application can be configured through `config/config.yaml`:

```yaml
server:
  port: ":1337"
logLevel: "info"
```

## Development

### Running Tests

```bash
go test ./...
```

### Linting

```bash
golangci-lint run
```

## Best Practices Implemented

1. **Project Structure**

   - Separation of concerns
   - Clean architecture principles
   - Internal packages for private code

2. **Error Handling**

   - Custom error types
   - Proper error wrapping
   - Graceful error responses

3. **Logging**

   - Structured logging with slog
   - Request/response logging
   - Error logging

4. **Configuration**

   - External configuration file
   - Environment variable support
   - Default values

5. **HTTP Server**

   - Graceful shutdown
   - Middleware support
   - Proper routing

6. **Input Validation**
   - Request validation
   - Error messages
   - Type safety

## Contributing

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a new Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.
