# go-api
This repository is meant to serve as a template for future projects.  
It uses a robust and flexible project structure and contains examples of api-features and mock response data.

Overview of the project structure
```
/project-root
├── /cmd
│   └── main.go                      # Entry point for the application
│
├── /routes                          # API routing logic
│   ├── api_router.go                # Main API router
│   └── /v1                          # Version 1 routes
│       ├── v1_router.go             # Combines v1-specific routers
│       ├── /users                   # Routes related to users
│       │   └── user_router.go
│       └── /tasks                   # Routes related to tasks
│           └── task_router.go
│
├── /middlewares                     # Middleware (grouped by domain and purpose)
│   ├── /auth                        # Authentication-related middleware
│   │   └── auth_middleware.go
│   └── /logger                      # Logging middleware
│       └── logger_middleware.go
│
├── /controllers                     # Controllers (organized by domain)
│   ├── /users                       # User-specific controllers
│   │   └── user_controller.go
│   └── /tasks                       # Task-specific controllers
│       └── task_controller.go
│
├── /models                          # Models (organized by domain)
│   ├── /users                       # User models
│   │   ├── user.go                  # Main user model
│   │   └── user_dto.go              # DTO for user responses
│   └── /tasks                       # Task models
│       ├── task.go                  # Main task model
│       └── task_dto.go              # DTO for task responses
│
├── /services                        # Business logic layer (grouped by domain)
│   ├── /users                       # User-specific services
│   │   └── user_service.go
│   └── /tasks                       # Task-specific services
│       └── task_service.go
│
├── /validators                      # Reusable validation utilities
│   ├── email_validator.go           # Email validation logic
│   └── /users                       # User-specific validation logic
│       └── name_validator.go        
│
└── go.mod                           # Go module definition
```