# AGENTS.md

## Purpose

This file defines backend-specific standards for code agents working inside the `backend/` directory of Finance AI App.

Agents must follow these instructions when modifying backend code.

---

## Backend Summary

The backend is a Go service responsible for:

- Authentication.
- User data.
- AI provider configuration.
- API key encryption.
- File uploads.
- AI provider validation.
- AI extraction.
- Import processing.
- Transaction storage.
- Categories.
- Dashboard data.
- Reports.
- Import history.

---

## Backend Stack

Use:

- Go.
- Chi.
- MariaDB.
- `database/sql`.
- SQL migrations.
- MySQL-compatible MariaDB driver.

Do not use:

- ORM libraries.
- `sqlc`.
- Global mutable database state outside controlled initialization.
- Direct AI provider calls from handlers.

---

## Expected Backend Structure

Recommended structure:

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
│   ├── logging/
│   └── http/
├── migrations/
├── tmp/
├── go.mod
├── go.sum
├── .env.example
├── README.md
└── AGENTS.md
```

Use `internal/` packages for application code.

Keep handlers, services, and repositories separated.

---

## Layering Rules

Follow this flow:

```text
HTTP handler
→ request validation
→ service
→ repository / external provider
→ response mapper
```

### Handlers

Handlers should:

- Decode requests.
- Validate request shape.
- Read authenticated user from context.
- Call services.
- Return HTTP responses.

Handlers should not:

- Contain business logic.
- Build SQL queries.
- Call AI providers directly.
- Encrypt or decrypt secrets directly unless delegated to a service.

### Services

Services should:

- Contain business logic.
- Enforce product rules.
- Manage import state transitions.
- Coordinate repositories and provider clients.
- Validate ownership when applicable.

### Repositories

Repositories should:

- Use `database/sql`.
- Execute explicit SQL.
- Return domain models.
- Handle `sql.ErrNoRows`.
- Accept `context.Context`.
- Avoid business logic.

---

## Database Rules

Use MariaDB.

Use Go's standard `database/sql`.

Use explicit SQL queries.

Do not use ORM.

Do not use `sqlc`.

Recommended driver:

```text
github.com/go-sql-driver/mysql
```

Repository methods should receive `context.Context`.

Example:

```go
func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*User, error) {
    const query = `
        SELECT id, name, email, password_hash, created_at, updated_at
        FROM users
        WHERE email = ?
        LIMIT 1
    `

    var user User
    err := r.db.QueryRowContext(ctx, query, email).Scan(
        &user.ID,
        &user.Name,
        &user.Email,
        &user.PasswordHash,
        &user.CreatedAt,
        &user.UpdatedAt,
    )
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, ErrUserNotFound
        }
        return nil, err
    }

    return &user, nil
}
```

---

## Migration Rules

Use SQL migration files.

Migration files should live in:

```text
backend/migrations/
```

Recommended format:

```text
000001_create_users_table.up.sql
000001_create_users_table.down.sql
000002_create_ai_provider_settings_table.up.sql
000002_create_ai_provider_settings_table.down.sql
```

Rules:

- Every schema change needs a migration.
- Every `.up.sql` should have a matching `.down.sql` when practical.
- Do not modify old migrations after they are shared or applied.
- Add new migrations for schema changes.
- Use explicit indexes for common lookup fields.
- Use foreign keys where useful and not harmful.

---

## Transaction Rules

Use database transactions for multi-step writes.

Examples:

- Confirming an import and saving multiple transactions.
- Updating import status and inserting related rows.
- Creating user-related default data.

Use:

```go
tx, err := db.BeginTx(ctx, nil)
```

Always rollback on error.

Always commit explicitly.

---

## Configuration Rules

Load configuration from environment variables.

Do not hardcode environment-specific values.

Required backend variables should be documented in `backend/.env.example` and `backend/README.md`.

Typical variables:

```text
APP_ENV
HTTP_PORT
DATABASE_HOST
DATABASE_PORT
DATABASE_NAME
DATABASE_USER
DATABASE_PASSWORD
AUTH_SECRET
AUTH_TOKEN_TTL_MINUTES
ENCRYPTION_KEY
CORS_ALLOWED_ORIGINS
MAX_UPLOAD_SIZE_MB
UPLOAD_TEMP_DIR
AI_PROVIDER_TEST_TIMEOUT_SECONDS
AI_EXTRACTION_TIMEOUT_SECONDS
LOG_LEVEL
```

---

## Authentication Rules

Protected routes must require authentication.

Authenticated user identity should be stored in request context.

Do not trust user IDs from request bodies when the authenticated user is available from context.

Users must only access their own:

- AI provider settings.
- Imports.
- Transactions.
- Categories.
- Reports.
- Dashboard data.

---

## AI Provider Rules

AI provider code should live under a focused package such as:

```text
internal/ai/
```

The AI package should expose interfaces such as:

```go
type Provider interface {
    TestConnection(ctx context.Context, cfg ProviderConfig) error
    TestCapabilities(ctx context.Context, cfg ProviderConfig) (*CapabilityResult, error)
    ExtractTransactions(ctx context.Context, input ExtractionInput) (*ExtractionResult, error)
}
```

Rules:

- Provider base URL must be configurable.
- API key must be configurable.
- Model must be configurable.
- API key must be decrypted only when needed.
- API key must never be logged.
- Provider errors must be mapped to internal error codes.
- Raw provider responses should not be exposed to frontend unless sanitized.

---

## AI Provider Validation Rules

Provider validation must check:

- Base URL is reachable.
- API key is valid.
- Model can be used.
- PDF file input is supported.
- CSV file input is supported.
- XLS/XLSX file input is supported.
- Structured JSON output is valid.

Possible statuses:

```text
draft
testing
ready
invalid_credentials
model_not_found
file_input_not_supported
pdf_not_supported
csv_not_supported
excel_not_supported
json_output_invalid
```

The import flow must be blocked unless the provider status is `ready`.

---

## File Upload Rules

Supported formats:

```text
.pdf
.csv
.xls
.xlsx
```

Backend must validate:

- Authentication.
- Provider readiness.
- Extension.
- MIME type.
- File size.
- Empty file.
- Temporary storage location.

Do not trust frontend validation.

Temporary files must not be publicly accessible.

Temporary files should be cleaned up after processing or after a retention period.

---

## AI Extraction Response Rules

The AI extraction result must be validated before it reaches the review screen.

Expected fields:

- Statement metadata.
- Transactions.
- Warnings.

Each transaction should include:

- Date.
- Description.
- Amount.
- Currency.
- Type.
- Category.
- Confidence.

Validation rules:

- JSON must be valid.
- Required fields must exist.
- Date must use ISO format.
- Amount must be numeric.
- Type must be valid.
- Confidence must be between 0 and 1.
- Unknown categories should map to `other`.

---

## Import Lifecycle Rules

Use these statuses:

```text
uploaded
processing
awaiting_review
completed
failed
cancelled
```

Rules:

- Do not save final transactions during AI processing.
- Store AI result as review/draft data.
- Save final transactions only after user confirmation.
- A failed import should store a safe error code/message.
- A cancelled import should not create final transactions.

---

## API Response Rules

Use consistent JSON response formats.

Recommended error shape:

```json
{
  "error": {
    "code": "invalid_request",
    "message": "The request is invalid.",
    "details": {}
  }
}
```

Rules:

- Use stable error codes.
- Use user-safe messages.
- Do not expose stack traces.
- Do not expose secrets.
- Do not expose full provider responses.

---

## Logging Rules

Logs should help debug without exposing sensitive data.

Allowed:

- Import ID.
- User ID.
- Provider name.
- Model name.
- File type.
- Import status.
- Error code.
- Duration.

Not allowed:

- Full API key.
- Full uploaded file contents.
- Full financial statement contents.
- Full provider raw response if it contains financial data.
- Passwords.
- Auth tokens.

---

## Testing Rules

Add tests for:

- Auth service.
- User repository.
- Provider settings service.
- API key encryption.
- Provider validation.
- File upload validation.
- AI response validation.
- Import state transitions.
- Confirm import transaction.
- Transaction repository.
- Dashboard calculations.
- Report calculations.

Run:

```bash
go test ./...
```

If applicable:

```bash
go test -race ./...
```

---

## Code Style

- Use `gofmt`.
- Use `go vet` when available.
- Keep functions small and explicit.
- Return errors instead of panicking.
- Use `context.Context` for request-scoped operations.
- Avoid package-level mutable state.
- Keep SQL readable.
- Prefer constants for statuses and error codes.
- Keep domain models separate from API DTOs when useful.

---

## Do Not Implement Without Explicit Request

- ORM.
- `sqlc`.
- Bank integrations.
- OCR fallback.
- Text extraction fallback.
- Ollama integration.
- llama.cpp integration.
- Payment system.
- Multi-provider routing.
- App store release pipeline.
