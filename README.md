# E-commerce API in Go

A simple e-commerce REST API built with Go's standard library.

## Features

- RESTful API endpoints
- Product listing functionality
- Simple HTTP server implementation
- CORS support

## API Endpoints

- `GET /products` - Get all products
- `GET /hello` - Hello world test endpoint
- `GET /about` - About page test endpoint

## Getting Started

### Prerequisites

- Go 1.24 or higher

### Installation

1. Clone the repository
   ```bash
   git clone https://github.com/yourusername/ecommerce-by-go.git
   cd ecommerce-by-go
   ```

2. Run the server
   ```bash
   go run main.go
   ```

3. Access the API at `http://localhost:8080`

## Project Structure

```
├── main.go     # Main application file
├── go.mod      # Go module file
└── .gitignore  # Git ignore file
```

## License

This project is licensed under the MIT License - see the LICENSE file for details.