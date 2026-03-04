## Module 1: System Design & Architecture
### RESTful API Design

**1. Base Setup (In-Memory REST API)**
- **Concept:** Built a foundational HTTP server using Go's standard `net/http` library.
- **Key Takeaway:** Utilized Go 1.22+ method-based routing (`GET /users`) to handle requests without external frameworks, serving JSON from an in-memory struct slice.

## Running the Project

The easiest way to run both the Go application and the PostgreSQL database is using Docker Compose.

### 1. Start the services
```bash
docker-compose up --build -d
```
This command will:
- Build the Go application into a lightweight container
- Pull the official PostgreSQL 17 image
- Start both services and connect them automatically on port 8080.

### 2. Verify the Server
Once the containers are running, you can interact with the API:

```bash
# Get all users
curl http://localhost:8080/users/

# Create a new user
curl -X POST -H "Content-Type: application/json" \
  -d '{"name":"Alice", "email":"alice@example.com"}' \
  http://localhost:8080/users/
```

### 3. Stop the services
```bash
docker-compose down
```

---

## Database Management (Ent ORM)
If you are developing locally without Docker and need to generate new database schemas:

```bash
go generate ./...
go run -mod=mod entgo.io/ent/cmd/ent generate ./ent/schema
```
