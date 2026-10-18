# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Repository Overview

This is a Go monorepo template for building microservices with shared packages. It follows [Standard Go Project Layout](https://github.com/golang-standards/project-layout) with pre-configured development tools.

## Project Structure

```text
.
├── services/           # Individual microservices
│   └── sample/        # Sample service
│       ├── cmd/sample/ # Application entry point (main.go)
│       └── internal/  # Private application code
├── pkg/               # Packages shared across services
│   └── logger/        # Shared logging package using slog
├── .devcontainer/     # VS Code DevContainer configuration
└── .github/           # GitHub Actions workflows
```

## Architecture Decisions

1. **Monorepo structure**: Services and packages are managed in a single repository, making shared code management and consistent tooling easier.

2. **Internal packages**: Each service uses the `internal/` directory to prevent other services from importing private implementation details.

3. **Module boundaries**: Each service has its own `go.mod` file and uses `replace` directives for local packages during development.

4. **Configuration management**: Uses godotenv to load environment-specific configuration from `.env/.env.{ENV}` files.

5. **Structured logging**: All services use a shared logger package with slog, producing consistent JSON logs in production.

## Development Environment

DevContainer provides:

- Go 1.27.2
- golangci-lint
- dprint
- lefthook

## Source Editing Notes

- Do not delete comments without also removing or updating their corresponding source code
- When making fixes via GitHub issues and running `git commit`, use JST timezone

## Verification

- Run the following to format:
  - `gofumpt -w .`
- Run the following for spell check:
  - `cspell .`
- Confirm successful build
- Run `go test`
- Run `golangci-lint run ./...` in the directory containing the go.mod of the target process and confirm no warnings
- Public methods should have unit tests implemented, except for very simple ones
