# Togo - Todo and User Management Application

A full-stack application with Go backend API and Vue.js frontend for managing todos and users.

## Prerequisites

Before running this project, make sure you have the following installed:

- **Go** (version 1.22 or higher)
- **Node.js** (version 18 or higher)
- **Yarn** package manager
- **PostgreSQL** database server

## Project Structure

```
togo/
├── togo-api/          # Go backend API
├── togo-ui/           # Vue.js frontend
├── Procfile           # Development process configuration
└── README.md          # This file
```

## Quick Start

### 1. Database Setup

First, ensure PostgreSQL is running and create the database:

```bash
# Create the database
cd togo-api
go run cmd/db/create/main.go
```

### 2. Database Schema

Create the required tables:

```bash
# From the togo-api directory
psql -U postgres -d togo_dev -f cmd/db/structure/structure.sql
```

### 3. Seed Data (Optional)

Load sample data into the database:

```bash
# From the togo-api directory
# Seed todos
go run cmd/db/seed/main.go

# Seed users
go run cmd/db/seed_users/main.go
```

### 4. Install Dependencies

Install backend dependencies:
```bash
cd togo-api
go mod download
```

Install frontend dependencies:
```bash
cd togo-ui
yarn install
```

### 5. Run the Application

#### Option 1: Using Procfile (Recommended for Development)

From the root directory:
```bash
# Install foreman or similar tool first
npm install -g foreman
# or use your preferred process manager

# Run both frontend and backend
foreman start
```

#### Option 2: Manual Start

Run backend (from `togo-api` directory):
```bash
go run main.go
```

Run frontend (from `togo-ui` directory):
```bash
yarn dev
```

The application will be available at:
- **Frontend**: http://localhost:5173
- **Backend API**: http://localhost:8080

## Database Configuration

The application connects to PostgreSQL using the following default connection string:
```
postgres://postgres:@localhost/togo_dev?sslmode=disable
```

To modify the database connection, update the connection string in:
`togo-api/internal/dbclient/dbclient.go`

## API Endpoints

### Todos
- `GET /todos` - List all todos

### Users
- `GET /users` - List all users
- `GET /users/:id` - Get single user by ID
- `POST /users` - Create new user
- `PUT /users/:id` - Update existing user

## Frontend Features

- **Todo Management**: View existing todos
- **User Management**: 
  - View list of all users
  - View individual user details
  - Create new users
  - Edit existing users
- **Responsive UI**: Works on desktop and mobile devices
- **Form Validation**: Input validation with error messages

## Development

### Backend Development

The backend is built with:
- **Go 1.22**
- **Gin** web framework
- **PostgreSQL** database with `lib/pq` driver

To add new API endpoints, modify `togo-api/main.go` and create corresponding handlers in the `internal/` packages.

### Frontend Development

The frontend is built with:
- **Vue.js 3**
- **TypeScript**
- **Vite** build tool
- **Pinia** for state management

Frontend development commands:
```bash
cd togo-ui

# Development server with hot-reload
yarn dev

# Type checking
yarn type-check

# Linting
yarn lint

# Build for production
yarn build
```

## Database Management

### Manual Database Operations

```bash
# Connect to database
psql -U postgres -d togo_dev

# Reset database (drops and recreates tables)
psql -U postgres -d togo_dev -f cmd/db/structure/structure.sql
```

### Adding New Database Tables

1. Update `cmd/db/structure/structure.sql` with new table definitions
2. Create seed data CSV files if needed
3. Create corresponding Go structs and handlers in `internal/` packages

## Troubleshooting

### Database Connection Issues

1. Ensure PostgreSQL is running
2. Verify the database `togo_dev` exists
3. Check connection string in `dbclient.go`
4. Ensure PostgreSQL user has proper permissions

### Frontend Build Issues

1. Delete `node_modules` and `yarn.lock`, then run `yarn install`
2. Check Node.js version compatibility
3. Run `yarn type-check` to identify TypeScript issues

### Backend Build Issues

1. Run `go mod download` to ensure all dependencies are installed
2. Check Go version compatibility (requires Go 1.22+)
3. Verify PostgreSQL driver is properly installed

## Production Deployment

For production deployment:

1. Build the frontend:
   ```bash
   cd togo-ui
   yarn build
   ```

2. Set production environment variables for database connection
3. Build the Go backend for your target platform
4. Serve the built frontend files through your web server
5. Ensure PostgreSQL is configured for production use

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test locally
5. Submit a pull request

## License

This project is open source and available under the MIT License.