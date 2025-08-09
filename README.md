# goauthic
Simple user auth in go ( in a lil scary way )

## Local Development Setup

### Prerequisites
- Go 1.19 or higher
- PostgreSQL
- Git
- Docker

### Setup Steps

#### Option 1: Using Docker (Recommended)

1. **Clone the repository**
   ```bash
   git clone https://github.com/sahilsobhani/goauthic.git
   cd goauthic
   ```

2. **Set up the environment**
   
   Create a `.env` file in the root directory:
   ```env
   DATABASE_URL=postgresql://authuser:authpass@localhost:5432/authdb
   ```

3. **Start PostgreSQL using Docker**
   ```bash
   docker-compose up -d
   ```
   This will start a PostgreSQL instance with the following credentials:
   - User: authuser
   - Password: authpass
   - Database: authdb
   - Port: 5432

4. **Install dependencies and run the server**
   ```bash
   go mod download
   go run main.go
   ```
   The server will start on `http://localhost:8080`

#### Option 2: Local Setup (without Docker)

1. **Clone the repository**
   ```bash
   git clone https://github.com/sahilsobhani/goauthic.git
   cd goauthic
   ```

2. **Set up the environment**
   
   Create a `.env` file in the root directory:
   ```env
   DATABASE_URL=postgresql://username:password@localhost:5432/goauthic
   ```
   Replace username, password with your PostgreSQL credentials.

3. **Set up the database**
   ```bash
   # Create the database
   createdb goauthic
   ```

4. **Install dependencies**
   ```bash
   go mod download
   ```

5. **Run the server**
   ```bash
   go run main.go
   ```
   The server will start on `http://localhost:8080`

### API Endpoints

- `POST /register` - Register a new user
  ```json
  {
    "email": "user@example.com",
    "password": "yourpassword"
  }
  ```

- `POST /login` - Login user
  ```json
  {
    "email": "user@example.com",
    "password": "yourpassword"
  }
  ```

- `GET /users/{id}` - Get user details (Protected route, requires JWT token)
  ```bash
  # Include the JWT token in the Authorization header
  Authorization: Bearer your.jwt.token
  ```

## 🗺️ Roadmap

### Phase 1: Enhanced Authentication
- [ ] Email verification flow
- [ ] Password reset functionality
- [ ] OAuth2 integration (Google, GitHub)
- [ ] Role-based access control (RBAC)
- [ ] Session management

### Phase 2: Security Improvements
- [ ] Rate limiting for login attempts
- [ ] IP-based blocking
- [ ] Password strength requirements
- [ ] Two-factor authentication (2FA)
- [ ] CORS configuration
- [ ] Security headers implementation

### Phase 3: API Enhancements
- [ ] User profile management
- [ ] API documentation with Swagger/OpenAPI
- [ ] Refresh token implementation
- [ ] Account deletion
- [ ] Logging and monitoring
- [ ] API versioning

### Phase 4: Infrastructure
- [ ] Dockerize the Go application
- [ ] CI/CD pipeline setup
- [ ] Automated testing
- [ ] Production deployment guide
- [ ] Performance monitoring
- [ ] Backup and restore procedures