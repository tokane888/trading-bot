# trading-bot

trading-bot and related tools

## Directory Structure

```sh
.
├── services/          # Services
│   └── sample/       # Sample service
│       ├── cmd/
│       ├── .env/
│       └── go.mod
├── pkg/              # Shared modules
│   └── logger/       # Logging
│       └── go.mod
└── README.md
```

## Design Principles

- Directory structure follows [Standard Go Project Layout](https://github.com/golang-standards/project-layout/blob/master/README_ja.md#standard-go-project-layout)
- Multiple services managed in a Go monorepo
- Each service has its own go.mod
- Shared modules placed in `pkg/` directory
  - Referenced locally via replace directives
- Architecture follows Clean Architecture
- Commit messages follow Conventional Commits
