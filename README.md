# Go-Booking-Management

Go-Booking-Management is a Go-based booking management system designed for handling reservations, scheduling, and customer management efficiently.

## Features
- User authentication and authorization
- Booking creation, update, and cancellation
- Availability checking and scheduling
- Customer management
- Payment integration (optional)
- API for third-party integration

## Requirements
- Go 1.18 or later
- PostgreSQL/MySQL database
- Redis (optional, for caching)
- Docker (optional, for containerization)

## Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/yourusername/go-booking-management.git
   cd go-booking-management
   ```

2. Install dependencies:
   ```sh
   go mod tidy
   ```

3. Set up the environment variables:
   ```sh
   cp .env.example .env
   ```
   Update `.env` with your database and other configuration details.

4. Run database migrations:
   ```sh
   go run migrate/migrate.go
   ```

5. Start the application:
   ```sh
   go run main.go
   ```

## Packages Used
This project makes use of the following Go packages:
- [github.com/alexedwards/scs/v2](https://github.com/alexedwards/scs) v2.8.0 - Session management for Go web applications.
- [github.com/go-chi/chi/v5](https://github.com/go-chi/chi) v5.2.0 - A lightweight, idiomatic and composable router for building Go HTTP services.
- [github.com/justinas/nosurf](https://github.com/justinas/nosurf) v1.1.1 - CSRF protection middleware for Go.

## Docker Support (Optional)
To run the project using Docker, execute:
```sh
docker-compose up --build
```

## Contribution
1. Fork the repository.
2. Create a new branch: `git checkout -b feature-branch`
3. Commit your changes: `git commit -m 'Add new feature'`
4. Push to the branch: `git push origin feature-branch`
5. Submit a pull request.

## License
This project is licensed under the MIT License.

## Author
Heinrich Smith

