# Go Server

## What this project does

This project is a RESTful API server written in Go, designed to manage users and "chirps" (short messages), with authentication, metrics, and webhook support. It uses PostgreSQL for data storage and provides endpoints for user management, authentication, and message handling.

## Why someone should care

- **Modern Go backend**: Demonstrates idiomatic Go server patterns, middleware, and modular design.
- **Database integration**: Shows how to connect and interact with PostgreSQL using Go.
- **Authentication**: Includes user authentication and token management.
- **Metrics and health checks**: Useful for monitoring and administration.
- **Extensible**: Good starting point for building more complex Go web services.

## How to install and run the project

1. **Clone the repository:**
   ```sh
   git clone https://github.com/nunseik/go-server.git
   cd go-server
   ```

2. **Set up environment variables:**
   Create a `.env` file in the project root with the following variables:
   ```
   DB_URL=postgres://user:password@localhost:5432/dbname?sslmode=disable
   PLATFORM=your_platform
   SECRET_KEY=your_secret_key
   POLKA_KEY=your_polka_key
   ```

3. **Install dependencies:**
   ```sh
   go mod tidy
   ```

4. **Run the server:**
   ```sh
   go run main.go
   ```

5. **Access the API:**
   The server will be running on `http://localhost:8080`. You can use tools like `curl` or Postman to interact with the endpoints.

---
For more details on API endpoints and usage, see the code and comments in `main.go`.
