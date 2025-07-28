# E-commerce API in Go

A simple e-commerce REST API built with Go's standard library.

## Features

- RESTful API endpoints
- Product listing and creation functionality
- Simple HTTP server implementation
- CORS support with standardized headers
- JSON request/response handling
- Code reusability with helper functions

## API Endpoints

- `GET /products` - Get all products
- `POST /create-product` - Create a new product
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

### API Usage Examples

#### Creating a new product

Send a POST request to `/create-product` with the following JSON body:

```json
{
  "title": "Product Name",
  "description": "Product Description",
  "price": 19.99,
  "imageUrl": "https://example.com/image.jpg"
}
```

The ID will be automatically assigned.

## Project Structure

```
├── main.go     # Main application file
├── go.mod      # Go module file
└── .gitignore  # Git ignore file
```

## License

This project is licensed under the MIT License - see the LICENSE file for details.