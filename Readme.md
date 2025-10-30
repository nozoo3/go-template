# Go GraphQL Template

A template project for building web servers using Go and GraphQL. Built with the Gin framework and gqlgen library for GraphQL API development, with Docker for development environment management.

## 🚀 Features

- **GraphQL API**: Type-safe GraphQL implementation using gqlgen
- **Gin Framework**: High-performance HTTP router
- **Docker Support**: Easy development environment setup
- **Dip Integration**: Streamlined development workflow
- **Hot Reload**: Automatic reload during development

## 📋 Prerequisites

- Go 1.25 or higher
- Docker & Docker Compose
- [Dip](https://github.com/bibendi/dip) (recommended)

## 🛠️ Setup

### Using Dip (Recommended)

```bash
# Start development environment
dip provision

# Generate GraphQL code
go run github.com/99designs/gqlgen generate
```

### Using Docker Compose directly

```bash
# Build image and start containers
docker compose up -d --build

# Stop containers
docker compose down
```

## 🌐 Access

After starting the server, you can access:

- **GraphQL Playground**: http://localhost:8080/
- **GraphQL API**: http://localhost:8080/query
- **Health Check**: http://localhost:8080/ping

## 📁 Project Structure

```
.
├── cmd/                    # Command-line tools
│   └── hello/             # Sample command
├── graph/                 # GraphQL related files
│   ├── model/            # Generated models
│   ├── generated.go      # Generated GraphQL code
│   ├── resolver.go       # Base resolver structure
│   ├── schema.graphqls   # GraphQL schema definition
│   └── schema.resolvers.go # Resolver implementations
├── build/                # Build related files
│   └── Dockerfile       # Docker image definition
├── server.go            # Main server file
├── go.mod              # Go module definition
├── gqlgen.yml          # gqlgen configuration file
├── docker-compose.yml  # Docker Compose configuration
└── dip.yml            # Dip configuration file
```

## 🔧 Development

### Updating GraphQL Schema

1. Edit `graph/schema.graphqls`
2. Run code generation:
   ```bash
   go run github.com/99designs/gqlgen generate
   ```
3. Implement new resolvers in `graph/schema.resolvers.go`

### Adding New Endpoints

You can add new endpoints to the Gin router in `server.go`:

```go
r.GET("/new-endpoint", func(c *gin.Context) {
    c.JSON(200, gin.H{"message": "New endpoint"})
})
```

## 📝 GraphQL API

### Sample Queries

```graphql
# Get all todos
query {
  todos {
    id
    text
    done
    user {
      id
      name
    }
  }
}
```

### Sample Mutations

```graphql
# Create a new todo
mutation {
  createTodo(input: { text: "New task", userId: "1" }) {
    id
    text
    done
  }
}
```

## 📚 Tech Stack

- **[Go](https://golang.org/)** - Programming language
- **[Gin](https://gin-gonic.com/)** - HTTP web framework
- **[gqlgen](https://gqlgen.com/)** - GraphQL server library
- **[Docker](https://www.docker.com/)** - Containerization platform
- **[Dip](https://github.com/bibendi/dip)** - Development environment management tool

## 🤝 Contributing

1. Fork this repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Create a Pull Request

## 📄 License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
