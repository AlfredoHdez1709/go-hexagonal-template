# Go Hexagonal Template

A robust template for building scalable Go applications using Hexagonal (Ports and Adapters) Architecture.

## Overview

This project follows the hexagonal architecture, ensuring a clear separation between domain logic, application services, and infrastructure. It is designed for maintainability, testability, and scalability.

## Project Structure

## Features

- Hexagonal architecture (ports and adapters)
- MongoDB integration
- Environment-based configuration
- Structured logging with Zap
- Fiber HTTP server
- Dependency injection for testability

## Getting Started

### Prerequisites

- Go 1.24+
- MongoDB

### Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/your-org/go-hexagonal-template.git
   cd go-hexagonal-template

Copy the example environment file and configure as needed:
cp .env.example .env
Install dependencies:

```sh
    go mod tidy
```


Running the Application

```sh
    go run cmd/main.go
```