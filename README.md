# KotoNilo - Go E-Commerce Project

## Quick Start with Docker

### Prerequisites

- Docker
- Docker Compose

### Running the Application

1. **Clone the repository** (if not already done)

2. **Start all services**:

   ```bash
   docker-compose up -d
   ```

3. **Check service status**:

   ```bash
   docker-compose ps
   ```

4. **View logs**:

   ```bash
   # All services
   docker-compose logs -f

   # Specific service
   docker-compose logs -f api
   docker-compose logs -f postgres
   ```

5. **Stop services**:

   ```bash
   docker-compose down
   ```

6. **Stop and remove volumes** (removes database data):
   ```bash
   docker-compose down -v
   ```

### Service Endpoints

- **API Health Check**: http://localhost:8080/health
- **Products Endpoint**: http://localhost:8080/products
- **PostgreSQL Database**: localhost:5432

### Database Access

Connect to PostgreSQL using these credentials:

- **Host**: localhost
- **Port**: 5432
- **Database**: kotonilo_db
- **User**: kotonilo
- **Password**: kotonilo_secret

Using `psql`:

```bash
docker-compose exec postgres psql -U kotonilo -d kotonilo_db
```

### Environment Variables

The application uses the following environment variables (see `.env.example`):

- `POSTGRES_USER`: Database user
- `POSTGRES_PASSWORD`: Database password
- `POSTGRES_DB`: Database name
- `DB_DSN`: Go application database connection string
- `SERVER_ADDR`: Server address and port

**Note**: For production, always change default passwords and use proper secrets management.

### Development

#### Rebuild the application after code changes:

```bash
docker-compose up -d --build api
```

#### Run database migrations manually:

```bash
# Access the postgres container
docker-compose exec postgres bash

# Run goose migrations (if you have goose installed in the container)
goose -dir /docker-entrypoint-initdb.d postgres "postgres://kotonilo:kotonilo_secret@localhost:5432/kotonilo_db?sslmode=disable" up
```

### Troubleshooting

1. **Port already in use**:

   - Change the port mapping in `docker-compose.yaml`
   - Example: Change `"8080:8080"` to `"8081:8080"`

2. **Database connection failed**:

   - Check if PostgreSQL is healthy: `docker-compose ps`
   - View logs: `docker-compose logs postgres`
   - Ensure the API waits for the database health check

3. **Clean start**:
   ```bash
   docker-compose down -v
   docker-compose up -d --build
   ```

### Project Structure

```
.
├── cmd/                    # Application entry points
│   ├── main.go            # Main application
│   └── api.go             # API routes and handlers
├── internal/              # Internal packages
│   ├── adapters/          # Database adapters
│   │   └── postgresql/
│   │       └── migrations/  # Database migrations
│   ├── products/          # Products domain
│   └── ...
├── docker-compose.yaml    # Docker Compose configuration
├── Dockerfile             # Application Docker image
├── .dockerignore          # Files to exclude from Docker build
├── .env.example           # Environment variables template
└── README.md              # This file
```

## Local Development (without Docker)

### Prerequisites

- Go 1.23+
- PostgreSQL 16+
- SQLC

### Setup

1. **Install dependencies**:

   ```bash
   go mod download
   ```

2. **Set up PostgreSQL database**:

   ```bash
   createdb kotonilo_db
   ```

3. **Run migrations**:

   ```bash
   goose -dir ./internal/adapters/postgresql/migrations postgres "postgres://kotonilo:kotonilo_secret@localhost:5432/kotonilo_db?sslmode=disable" up
   ```

4. **Generate SQLC code**:

   ```bash
   sqlc generate
   ```

5. **Run the application**:
   ```bash
   go run ./cmd/main.go ./cmd/api.go
   ```

## API Documentation

### Endpoints

#### Health Check

- **GET** `/health`
- **Response**: `doing good.`

#### List Products

- **GET** `/products` [List of all products]
- **POST** `/products` [Adding products]
- **Response**: JSON array of products
