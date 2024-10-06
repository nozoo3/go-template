
# To start the project
```sh
# Start the Containers
dip provision

# Generate the GraphQL Code
go run github.com/99designs/gqlgen generate
```
# Docker Commands
```sh
# Create an Image and Start the Containers
docker compose up -d --build

# Stop the Containers
docker compose down
```

```sh
This project sets up a GraphQL server using the Gin framework. It leverages Docker for containerization and `dip` for easy management of the development environment.
```