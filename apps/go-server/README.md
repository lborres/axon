# Axon API Server

# Project Structure
```
/axon/apps/server/
├── cmd/                 # Entry point(s)
├── internal/            # Core application logic
│   └── config           # App config and env
│   └── handlers         # API Handlers
│   └── http             # Bootstraps the API Server & other http utility
│   └── middlerwares     # API Middlewares
│   └── models           # Data models
│   └── routes           # API Routes definitions
│   └── services         # API Business Logic and Contracts(Interfaces)
│   └── stores           # Data Storage(DB or other) and Contracts(Interfaces)
├── pkg/                 # External systems / Reusable Packages
```
