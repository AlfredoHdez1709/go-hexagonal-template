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

Click the **"Use this template"** button on GitHub to create a new repository based on this template, or follow these steps:


1. Copy the example environment file and configure as needed: **.env.example** in **.env**
2. Change the module name in **go.mod** to match your project name. <br>

example:
```sh
    go mod edit -module github.com/yourusername/yourprojectname
```

3. Install dependencies:

```sh
    go mod tidy
```


4. Running the Application

```sh
    go run cmd/main.go
```
