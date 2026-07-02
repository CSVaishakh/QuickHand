# QuickHand

QuickHand is a modern, high-performance platform designed to connect **Clients** with local **Handymen** (plumbers, electricians, carpenters, etc.) to get jobs done efficiently. 

This repository is organized as a monorepo containing both the backend and frontend components, along with shared libraries and database configurations.

---

## 🛠️ Project Structure

The project code is located in the `src/` directory, divided into applications (`apps/`) and shared packages (`packages/`):

```text
QuickHand/
├── src/
│   ├── apps/
│   │   ├── server/       # Go-based Fiber (v3) backend API server
│   │   └── web/          # React frontend built with TanStack Start, TypeScript, and Tailwind CSS
│   └── packages/
│       ├── db/           # Database models, GORM repositories, and database migrations
│       └── auth/         # JWT-based authentication, session management, and auth services
├── go.work               # Go workspace configuration
├── package.json          # Root scripts for database and orchestration
└── README.md             # Project documentation (this file)
```

---

## 🚀 Prerequisites

To run this project locally, ensure you have the following installed:

- **Go**: `v1.26.3` or later
- **Bun**: Fast JavaScript package manager & runtime (used for the frontend)
- **Docker**: For running the PostgreSQL database container
- **golang-migrate CLI**: If you wish to run migrations manually outside of Go code

---

## ⚡ Getting Started

Follow these steps to set up and run the application locally:

### 1. Set Up the Database

Start the PostgreSQL database container via Docker:

```bash
bun db:start
# or using npm
npm run db:start
```

This starts a PostgreSQL instance on port `5432` with the database `quickhand`.

### 2. Run Database Migrations

Apply the database migrations to set up tables (users, sessions, handymen, addresses, and jobs):

```bash
bun migrate
# or using npm
npm run migrate
```

### 3. Start the Go Backend Server

Run the development server script which starts the database container and runs the Fiber API server:

```bash
bun dev:server
# or using npm
npm run dev:server
```

Alternatively, navigate to `src/apps/server/` and run:

```bash
go run main.go
```

The server runs on `http://localhost:3000` (or the port specified in `.env.local`).

### 4. Start the Web Frontend

In a separate terminal, navigate to the frontend directory, install dependencies, and run the development server:

```bash
cd src/apps/web
bun install
bun --bun run dev
```

The frontend will be available at `http://localhost:3000` (or the port specified in `src/apps/web/.env.local`).

---

## 📁 Package Breakdown

### 1. `src/apps/server` (Backend APIs)
Built using the **Fiber v3** framework. It registers routes, sets up middlewares (such as JWT/session authentication checks), and runs handlers via controllers:
- `controllers/`: Handles incoming HTTP request binding, route mapping, and calls appropriate services.
- `services/`: Encapsulates business logic (e.g., address creation and updates).

### 2. `src/packages/auth` (Authentication Package)
Implements JWT-based authentication and session verification:
- Sign-up and sign-in handlers for both clients and handymen.
- JWT token generation, signature validation, and secure OTP generation for password resets.
- Database session verification and revocation.

### 3. `src/packages/db` (Database & GORM Package)
Initializes database connections and houses the schema models and repository operations:
- `models/`: GORM struct definitions mapping to PostgreSQL tables (`User`, `Handyman`, `Client`, `Session`, `Address`, `Job`).
- `repositories/`: Database query abstraction layer using GORM to query and modify records.
- `migrations/`: Raw SQL migration files (`.up.sql` and `.down.sql`) for creating database schemas.

---

## 🎛️ Root Scripts Reference

The following commands are available from the root directory:

| Command | Action |
|:---|:---|
| `bun db:start` | Spins up the local PostgreSQL Docker container. |
| `bun db:stop` | Stops the PostgreSQL Docker container. |
| `bun db:reset` | Wipes the database volume and restarts the container from scratch. |
| `bun migrate` | Runs the database migration script (`src/packages/db/cmd/main.go`). |
| `bun dev:server` | Verifies DB is running and launches the Go backend server. |
