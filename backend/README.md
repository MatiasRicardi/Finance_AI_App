# Finance AI App Backend

This is the backend service for Finance AI App.

The backend is responsible for authentication, AI provider configuration, secure API key storage, file upload handling, AI extraction, transaction validation, transaction storage, dashboard data, and reports.

## Technology Stack

| Area | Technology |
|---|---|
| Language | Go |
| HTTP Router | Chi |
| Database | MariaDB |
| Database Access | Explicit SQL, no ORM |
| Migrations | SQL migration files |
| Authentication | Token-based or session-based authentication |
| File Uploads | Multipart form uploads |

## Database Approach

This project uses Go's standard `database/sql` package with a MySQL-compatible driver.

SQL queries are written explicitly. There is no ORM and no code generator.

This gives full control over queries and keeps the implementation simple and readable.

## Responsibilities

The backend should handle:

- User registration.
- User login.
- Authenticated API routes.
- User profile data.
- AI provider settings.
- API key encryption.
- AI provider validation.
- PDF, CSV, XLS, and XLSX upload validation.
- Sending files directly to the configured AI provider.
- Parsing and validating AI responses.
- Import status lifecycle.
- Transaction review data.
- Confirmed transaction storage.
- Transaction listing and filtering.
- Categories.
- Dashboard summaries.
- Basic reports.
- Import history.
- Safe logging.

## Suggested Project Structure

```text
backend/
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── auth/
│   ├── users/
│   ├── ai/
│   ├── imports/
│   ├── transactions/
│   ├── categories/
│   ├── dashboard/
│   ├── reports/
│   ├── config/
│   ├── database/
│   ├── encryption/
│   └── http/
├── migrations/
├── go.mod
├── go.sum
├── .env.example
└── README.md
```

## Environment Variables

Create a `.env` file in the `backend/` directory.

```env
# Application
APP_ENV=development

# HTTP
HTTP_PORT=8080
CORS_ALLOWED_ORIGINS=http://localhost:5173

# Database (MariaDB)
DATABASE_HOST=localhost
DATABASE_PORT=3306
DATABASE_NAME=finance_ai_app
DATABASE_USER=finance_user
DATABASE_PASSWORD=change_me

# Authentication
# Generate with: openssl rand -hex 32
AUTH_SECRET=change_me_to_a_random_32_byte_hex_string
AUTH_TOKEN_TTL_MINUTES=60

# Encryption (for AI provider API keys at rest)
# Generate with: openssl rand -hex 32
ENCRYPTION_KEY=change_me_to_a_random_32_byte_hex_string

# File Upload
MAX_UPLOAD_SIZE_MB=20
UPLOAD_TEMP_DIR=/tmp/finance-ai-app

# AI Provider
AI_PROVIDER_TEST_TIMEOUT_SECONDS=15
AI_EXTRACTION_TIMEOUT_SECONDS=120

# Logging
LOG_LEVEL=info
```

## Local Development

### Setup

```bash
cp .env.example .env
# Edit .env and fill in your database credentials and secrets
```

### Run the API

```bash
export $(cat .env | grep -v '^#' | xargs) && go run ./cmd/api
```

### Build

```bash
go build ./...
```

### Run Tests

```bash
go test ./...
```

### Run with Race Detector

```bash
go test -race ./...
```

## Database

The application uses MariaDB.

The database is expected to be hosted on an external server. Configure the connection via the environment variables in `.env`.

Required variables:

```text
DATABASE_HOST
DATABASE_PORT
DATABASE_NAME
DATABASE_USER
DATABASE_PASSWORD
```

## Migrations

Migrations are versioned SQL files used to create and update the database schema.

Example structure:

```text
backend/migrations/
├── 000001_create_users_table.up.sql
├── 000001_create_users_table.down.sql
├── 000002_create_ai_provider_settings_table.up.sql
├── 000002_create_ai_provider_settings_table.down.sql
└── ...
```

A migration tool such as `golang-migrate` can be used.

Example:

```bash
migrate -path ./migrations -database "mysql://finance_ai_app:finance_ai_app_password@tcp(localhost:3306)/finance_ai_app" up
```

## API Areas

### Authentication

- Register user.
- Login user.
- Get current user.
- Logout if sessions are used.

### AI Provider Settings

- Get provider settings.
- Save provider settings.
- Test provider configuration.
- Delete provider settings if needed.

### Imports

- Upload financial statement.
- Get import status.
- Get import review data.
- Confirm import.
- Cancel import.
- Get import history.
- Get import detail.

### Transactions

- List transactions.
- Create manual transaction.
- Update transaction.
- Delete transaction.

### Categories

- List categories.
- Create custom category.
- Update category.
- Deactivate category.

### Dashboard

- Monthly income.
- Monthly expenses.
- Monthly balance.
- Top spending category.
- Recent transactions.

### Reports

- Spending by category.
- Income vs expenses.
- Monthly spending trend.

## AI Provider Requirements

The configured provider must support direct file input for:

- PDF.
- CSV.
- XLS.
- XLSX.

Provider validation should check:

- Base URL is reachable.
- API key is valid.
- Model can be used.
- PDF file input is supported.
- CSV file input is supported.
- Excel file input is supported.
- Structured JSON output is valid.

## Security Requirements

The backend must follow these rules:

- Never store API keys in plain text.
- Encrypt API keys before saving them.
- Never return full API keys to the frontend.
- Never log API keys.
- Never log full financial file contents.
- Validate all uploaded files.
- Restrict users to their own data.
- Return safe error messages.
- Store only necessary provider error details.

## Import Status Lifecycle

Recommended import statuses:

```text
uploaded
processing
awaiting_review
completed
failed
cancelled
```

## License

MIT
