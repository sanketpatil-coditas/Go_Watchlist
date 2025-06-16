# Go Watchlist Microservices

This project consists of two microservices:
1. Authentication Service - Handles user authentication and management
2. Watchlist Service - Manages stock watchlists and alerts

## Project Structure
```
GO_WATCHLIST/
├── auth-service/                 # Authentication microservice
│   ├── cmd/                     # Application entry points
│   │   └── main.go             
│   ├── internal/               
│   │   ├── config/             # Service configuration
│   │   ├── handler/            # HTTP handlers
│   │   ├── middleware/         # HTTP middleware
│   │   ├── models/             # Data models
│   │   │   ├── db/            # Database models
│   │   │   ├── request/       # Request models
│   │   │   └── response/      # Response models
│   │   ├── repository/         # Data access layer
│   │   └── service/           # Business logic
│   ├── pkg/                    # Shared packages
│   └── go.mod                  # Module definition
│
├── watchlist-service/           # Watchlist microservice
│   ├── cmd/                    # Application entry points
│   │   └── main.go
│   ├── internal/
│   │   ├── config/            # Service configuration
│   │   ├── handler/           # HTTP handlers
│   │   ├── middleware/        # HTTP middleware
│   │   ├── models/            # Data models
│   │   │   ├── db/           # Database models
│   │   │   ├── request/      # Request models
│   │   │   └── response/     # Response models
│   │   ├── repository/        # Data access layer
│   │   └── service/          # Business logic
│   ├── pkg/                   # Shared packages
│   └── go.mod                 # Module definition
│
└── pkg/                        # Shared code between services
    ├── common/                # Common utilities
    └── types/                 # Shared types
```

## Setup Instructions

1. Authentication Service Setup:
```bash
cd auth-service
go mod init go-watchlist/auth-service
go mod tidy
```

2. Watchlist Service Setup:
```bash
cd watchlist-service
go mod init go-watchlist/watchlist-service
go mod tidy
```

## Running the Services

1. Start Authentication Service:
```bash
cd auth-service/cmd
go run main.go
```

2. Start Watchlist Service:
```bash
cd watchlist-service/cmd
go run main.go
```

## API Documentation

### Authentication Service (Default port: 8081)
- POST /auth/register - Register new user
- POST /auth/login - User login
- POST /auth/logout - User logout
- DELETE /auth/user - Remove user

### Watchlist Service (Default port: 8082)
- POST /watchlist/add - Add stock to watchlist
- POST /watchlist/remove - Remove stock from watchlist
- GET /watchlist - Get user's watchlist
- GET /watchlist/alerts - Get price alerts

## Environment Variables

### Authentication Service
- `AUTH_SERVICE_PORT` - Service port (default: 8081)
- `DB_CONNECTION` - Database connection string
- `JWT_SECRET` - JWT signing secret

### Watchlist Service
- `WATCHLIST_SERVICE_PORT` - Service port (default: 8082)
- `DB_CONNECTION` - Database connection string
- `AUTH_SERVICE_URL` - Authentication service URL 