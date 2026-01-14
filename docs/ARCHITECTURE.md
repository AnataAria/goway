# Architecture Guide

Goway is built upon the principles of **Hexagonal Architecture** (also known as Ports and Adapters). The core idea is to isolate the central logic of your application from the details of the outside world (web frameworks, databases, third-party services).

## The Layers

### 1. Domain (`internal/domain`)
**Dependency Level: None**

This is the innermost layer. It contains the Enterprise Business Rules.
- **Entities**: Structs representing domain concepts (`User`, `Account`).
- **Value Objects**: immutable objects defined by their attributes (e.g., `Email`, `Password`).
- **Domain Errors**: Errors specific to business rules (e.g., `ErrInsufficientFunds`).

Ideally, this layer has **zero** dependencies on other packages within the project or external libraries (except for maybe standard library or simple utility libraries).

### 2. Ports (`internal/port`)
**Dependency Level: Defer to Adapters**

This layer defines the **contracts** (interfaces) for communication.
- **Input Ports (Use Cases)**: Interfaces that define what the application *can do*. (Note: In this simplified implementation, we often skip defining interfaces for use cases and structs directly in the Application layer, but strictly speaking, they belong here conceptually).
- **Output Ports (Repositories/Services)**: Interfaces that the application needs to *work*. For example, `UserRepository` is an interface defining methods like `Save` or `FindByID`.

### 3. Application (`internal/application`)
**Dependency Level: Depends on Domain and Ports**

This layer contains the Application Business Rules. It implements the Use Cases.
- It orchestrates the domain objects.
- It calls the Output Ports (repositories) to persist state.
- It does **not** know about HTTP, JSON, or SQL.

Example `UserUseCase`:
1. Receive input (DTO).
2. load User entity from Repository (Port).
3. Call method on User entity.
4. Save User entity back to Repository.

### 4. Adapters (`internal/adapter`)
**Dependency Level: Depends on Application, Domain, and Ports**

This layer implements the interfaces defined in the Port layer. It acts as the bridge to the outside world.

#### Driving Adapters (Primary / "In")
These *drive* the application.
- **HTTP/Web**: The `in/http` package. It handles JSON decoding, validation, and calls the Application Use Cases.
- **CLI**: If we had a CLI, it would be here.

#### Driven Adapters (Secondary / "Out")
These are *driven by* the application.
- **Persistence**: `out/persistence/postgres`. Implements the Repository interfaces using SQL.
- **Cache**: `out/cache`.
- **Security**: `out/security`. BCrypt hashing, JWT generation.

## Dependency Rule

The golden rule of this architecture is: **Dependencies always point INWARDS.**

- Adapter -> depends on -> Application -> depends on -> Domain.
- Domain depends on NOTHING.
- Application depends on Domain.

This ensures that we can swap out the database (Adapter) without touching the Business Logic (Domain). We can switch from HTTP to gRPC (Adapter) without changing the core application.
