everate/
├── cmd/
│   ├── server/
│   │   └── main.go               # Entry point for the web server
│   └── wallet-worker/
│       └── main.go               # Entry point for the background ISK wallet worker
│
├── internal/
│   ├── auth/                     # EVE SSO logic, session management
│   │   └── sso.go
│   │
│   ├── config/                   # Configuration loading from environment variables
│   │   └── config.go
│   │
│   ├── datastore/                # Database interaction layer (interfaces + implementation)
│   │   ├── datastore.go          # Defines interfaces (e.g., UserStore, RatingStore)
│   │   └── mongodb/              # MongoDB implementation of the datastore interfaces
│   │       ├── user.go
│   │       ├── rating.go
│   │       ├── report.go
│   │       └── db.go             # Connection logic
│   │
│   ├── handler/                  # HTTP handlers (the "Controllers" in MVC)
│   │   ├── handler.go            # Base handler struct with dependencies (logger, services)
│   │   ├── auth.go               # SSO login/callback handlers
│   │   ├── entity.go             # Character/Corp/Alliance page handlers
│   │   ├── rating.go             # Comment posting, voting, reporting handlers
│   │   ├── draft.go              # Draft saving/loading handlers
│   │   ├── search.go             # Autocomplete search handler
│   │   └── admin/                # Sub-package for admin panel handlers
│   │       └── admin_handler.go
│   │
│   ├── logger/                   # Zap logger initialization
│   │   └── logger.go             # Defines New() constructor for zap.Logger
│   │
│   ├── middleware/               # Custom Echo middleware
│   │   ├── auth.go               # Check for login, admin roles, ban status
│   │   ├── request_logger.go     # Custom Zap request logger
│   │   └── ratelimit.go          # Rate limiting middleware
│   │
│   ├── model/                    # Core application data structures (User, Rating, etc.)
│   │   ├── user.go
│   │   ├── entity.go
│   │   └── ...
│   │
│   ├── server/                   # Echo server setup and routing
│   │   └── server.go             # NewServer(), setupRoutes(), Start()
│   │
│   ├── service/                  # Core business logic (decoupled from web layer)
│   │   ├── rating_service.go     # Logic for calculating costs, updating scores
│   │   └── wallet_service.go     # Logic for processing ESI journal transactions
│   │
│   └── view/                     # HTML template management
│       ├── view.go               # Template parsing and rendering engine
│       ├── layout/               # Base layouts (*.html)
│       ├── page/                 # Full page templates (*.html)
│       └── component/            # Reusable htmx partials (*.html)
│
├── pkg/
│   └── esi/                      # Reusable EVE Swagger Interface client wrapper
│       ├── client.go             # The ESI client with caching
│       └── models.go             # Structs for ESI API responses
│
├── web/
│   └── static/                   # Static assets (CSS, JS, images)
│       ├── css/
│       │   └── style.css
│       └── js/
│           ├── main.js
│           └── editor.js
│
├── .env.example                  # Example environment variables for local development
├── .gitignore
├── go.mod
├── go.sum
└── Makefile                      # For common dev tasks (build, run, test)
