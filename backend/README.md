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

This project does not use an ORM.

The backend should use the standard `database/sql`

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
APP_ENV=development
HTTP_PORT=8080

DATABASE_HOST=localhost
DATABASE_PORT=3306
DATABASE_NAME=finance_ai_app
DATABASE_USER=finance_ai_app
DATABASE_PASSWORD=finance_ai_app_password
DATABASE_MAX_OPEN_CONNS=25
DATABASE_MAX_IDLE_CONNS=25
DATABASE_CONN_MAX_LIFETIME_MINUTES=5

AUTH_SECRET=change-me
AUTH_TOKEN_TTL_MINUTES=1440

ENCRYPTION_KEY=change-me-32-byte-key

CORS_ALLOWED_ORIGINS=http://localhost:5173,http://localhost:8100

MAX_UPLOAD_SIZE_MB=10
UPLOAD_TEMP_DIR=./tmp/uploads

AI_PROVIDER_TEST_TIMEOUT_SECONDS=60
AI_EXTRACTION_TIMEOUT_SECONDS=120

LOG_LEVEL=debug
```

## Local Development

### Install Dependencies

```bash
go mod download
```

### Run the API

```bash
go run ./cmd/api
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

Example local connection settings:

```text
Host: localhost
Port: 3306
Database: finance_ai_app
User: finance_ai_app
Password: finance_ai_app_password
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
