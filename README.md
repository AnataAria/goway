# Goway

![Go Version](https://img.shields.io/badge/go-1.25-blue)
![License](https://img.shields.io/badge/license-MIT-green)

Welcome to **Goway**, a robust and modern Go backend template designed to get you up and running with production-ready standards in no time. 

This isn't just another "Hello World" example. It's a fully structured foundation built on **Hexagonal Architecture (Ports and Adapters)**, ensuring that your business logic remains pure, testable, and independent of external frameworks or databases. We've hand-picked a stack that balances performance with developer experience.

## Why this template?

Starting a new Go project often involves a lot of boilerplate decision-making. "Where do I put my models?", "How should I structure my handlers?", "What router should I use?". 

Goway answers these questions with an opinionated yet flexible structure. It demonstrates:
- **Clean Architecture principles**: Clear separation of concerns between Domain, Application, and adapter layers.
- **Modern Web Framework**: Powered by [Fuego](https://github.com/go-fuego/fuego) for a developer-friendly API experience with built-in OpenAPI generation.
- **Production Basics**: Includes database migrations, configuration management, structured logging, and authentication out of the box.

## Features

- **Standard Library compatible**: Uses `net/http` compatible handlers via Fuego.
- **Structured Project Layout**: Follows standard Go project layout conventions.
- **Database**: PostgreSQL with `sqlx` for powerful yet simple DB interactions.
- **Authentication**: JWT-based authentication (Access & Refresh tokens).
- **Validation**: Request validation using `go-playground/validator`.
- **Environment Config**: easy 12-factor app configuration via `.env`.
- **Rate Limiting**: Built-in middleware to protect your API.

## Tech Stack

- **Language**: Go 1.25+
- **Web Framework**: [Fuego](https://github.com/go-fuego/fuego)
- **Database**: PostgreSQL (Driver: `lib/pq`, SQL Builder: `sqlx`)
- **Caching/KV**: Redis
- **Auth**: JWT (`golang-jwt/jwt/v5`)
- **Crypto**: bcrypt
- **Migrations**: `golang-migrate/migrate`

## Project Structure

This project follows the **Hexagonal Architecture**. Here is a high-level overview:

- **`internal/domain`**: The heart of the application. Contains business entities and logic. Pure Go code, no external dependencies.
- **`internal/application`**: The "use cases" of the application. Orchestrates the flow of data to and from the domain entities. Use cases orchestrate the dance of the domain objects.
- **`internal/port`**: Interfaces (contracts) that define how the application interacts with the outside world (repositories, services).
- **`internal/adapter`**: Implementation of the ports.
    - **`in`**: Driving adapters (e.g., HTTP handlers) that trigger the application.
    - **`out`**: Driven adapters (e.g., Postgres repositories, Redis cache) called by the application.

For a deep dive into the architecture, check out [docs/ARCHITECTURE.md](docs/ARCHITECTURE.md).

## Getting Started

### Prerequisites

- [Go](https://go.dev/dl/) (version 1.25 or higher)
- [PostgreSQL](https://www.postgresql.org/)
- [Make](https://www.gnu.org/software/make/) (optional, but recommended)

### 1. Clone the repository

```bash
git clone https://github.com/AnataAria/goway.git
cd goway
```

### 2. Configure Environment

Copy the example environment file:

```bash
cp .env.example .env
```

Open `.env` and configure your database credentials and secret keys.

### 3. Setup Database

You can use the provided Makefile to create the database and run migrations.

```bash
# Create database
make db-create

# Run migrations
make migrate-up
```

### 4. Run the application

```bash
make run
```

The server will start at `http://localhost:8080` (or whatever port you configured).

### 5. Explore the API

Fuego automatically generates an OpenAPI specification and Swagger UI.
Once the server is running, visit:

`http://localhost:8080/swagger`

## Contributing

We welcome contributions! Please check out [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines on how to proceed.
