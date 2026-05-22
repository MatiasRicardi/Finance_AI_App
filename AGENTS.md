# AGENTS.md

## Purpose

This file defines the working standards for code agents contributing to **Finance AI App**.

Agents must follow these instructions when modifying the project. The goal is to keep the codebase consistent, secure, maintainable, and aligned with the product scope.

---

## Project Summary

Finance AI App is a personal finance application that uses AI to extract and classify transactions from financial statement files.

Users upload files in **PDF, CSV, XLS, or XLSX** format. The backend sends the uploaded file directly to a user-configured AI provider. The AI provider returns structured transaction data. The user reviews and corrects the extracted transactions before confirming the import.

The application is built as a monorepo.

```text
finance_ai_app/
├── backend/
├── frontend/
├── docs/
├── scripts/
├── docker-compose.yml
├── README.md
└── AGENTS.md
```

---

## Product Rules

Agents must respect these product rules.

### AI Provider Requirement

The configured AI provider must support direct file input for:

- PDF.
- CSV.
- XLS.
- XLSX.

Text-only providers are not supported in the MVP.

Do not implement backend-side text extraction fallback unless explicitly requested.

### Mandatory User Review

AI-extracted transactions must never be saved automatically.

The user must review and confirm extracted transactions before they become final records.

### User-Owned AI Provider

Each user configures their own AI provider credentials.

API keys are user-provided and must be treated as secrets.

### One Active AI Provider

The MVP supports one active AI provider per user.

Do not implement multi-provider routing unless explicitly requested.

---

## Technology Stack

### Backend

- Go.
- Chi.
- MariaDB.
- `database/sql`.
- SQL migrations.
- No ORM.
- No `sqlc`.

### Frontend

- Vue.js.
- Ionic Vue.
- Capacitor.
- Vite.
- Pinia.
- Vue Router.
- pnpm.

### Mobile

The mobile app is generated from the frontend project using Capacitor.

Do not create a separate mobile repository or separate mobile application unless explicitly requested.

---

## Repository Structure Standards

Expected structure:

```text
finance_ai_app/
├── backend/
│   ├── cmd/
│   ├── internal/
│   ├── migrations/
│   ├── README.md
│   └── AGENTS.md
│
├── frontend/
│   ├── src/
│   ├── public/
│   ├── android/
│   ├── ios/
│   ├── README.md
│   └── AGENTS.md
│
├── docs/
│   ├── project-brief.md
│   ├── project-summary.md
│   ├── mvp-scope.md
│   ├── main-user-flow.md
│   ├── architecture.md
│   └── api-contracts.md
│
├── scripts/
├── docker-compose.yml
├── README.md
└── AGENTS.md
```

Do not move major folders without updating documentation.

---

## Documentation Rules

Agents should update documentation when changing:

- Project structure.
- API contracts.
- Environment variables.
- Setup commands.
- Main user flow.
- Database schema.
- Import flow.
- AI provider behavior.
- Security behavior.

Documentation must be written in English.

Relevant files:

```text
README.md
backend/README.md
frontend/README.md
docs/project-brief.md
docs/project-summary.md
docs/mvp-scope.md
docs/main-user-flow.md
docs/architecture.md
docs/api-contracts.md
```

If a task changes behavior, update or create the appropriate document.

---

## Security Rules

Agents must treat this application as handling sensitive financial data.

Never:

- Log full API keys.
- Log full uploaded financial files.
- Return full API keys to the frontend.
- Store API keys in plain text.
- Expose another user's data.
- Trust client-side validation only.
- Save AI-extracted transactions without review.

Always:

- Validate authenticated user ownership.
- Validate file type and size on the backend.
- Use safe error messages.
- Encrypt user-provided API keys at rest.
- Mask API keys in frontend responses.
- Keep uploaded files private.
- Clean up temporary files when applicable.

---

## AI Integration Rules

The AI provider integration must be isolated behind a provider interface or service layer.

Do not call provider APIs directly from unrelated modules.

The provider integration must support:

- Configurable base URL.
- Configurable API key.
- Configurable model.
- File input.
- Typed provider errors.
- Structured JSON response validation.

Provider validation must check:

- Base URL reachability.
- API key validity.
- Model availability.
- PDF support.
- CSV support.
- Excel support.
- Structured JSON output.

---

## Import Flow Rules

The import lifecycle should use explicit statuses.

Recommended statuses:

```text
uploaded
processing
awaiting_review
completed
failed
cancelled
```

Rules:

- A file upload creates an import record.
- AI processing moves the import to `processing`.
- A valid AI response moves the import to `awaiting_review`.
- User confirmation moves the import to `completed`.
- User cancellation moves the import to `cancelled`.
- Provider or validation failure moves the import to `failed`.

Do not skip the `awaiting_review` step.

---

## Backend/Frontend Boundary

The backend is responsible for:

- Auth.
- Authorization.
- File validation.
- AI provider calls.
- AI response validation.
- Data persistence.
- Security-sensitive logic.

The frontend is responsible for:

- UI.
- Forms.
- User interaction.
- Client-side validation for UX.
- Displaying import review data.
- Submitting user-confirmed data.

The frontend must not:

- Call AI providers directly.
- Store provider API keys.
- Trust AI output as final.
- Bypass backend validation.

---

## Environment Variable Rules

Do not hardcode secrets.

Use environment variables for:

- Database credentials.
- Auth secrets.
- Encryption keys.
- CORS settings.
- Upload limits.
- API URLs.
- Timeouts.

Provide `.env.example` files when adding new variables.

---

## Testing Expectations

Agents should add or update tests when modifying:

- Auth.
- AI provider validation.
- File upload validation.
- AI response validation.
- Import state transitions.
- Transaction persistence.
- Dashboard/report calculations.
- Frontend forms.
- Import review behavior.

Run relevant checks before considering work complete.

Backend:

```bash
go test ./...
```

Frontend:

```bash
pnpm lint
pnpm test
pnpm build
```

If a command is not yet available, do not invent passing results. Add the script only if the project setup supports it.

---

## Coding Standards

### General

- Prefer simple, explicit code.
- Keep modules focused.
- Avoid hidden side effects.
- Use clear names.
- Avoid premature abstractions.
- Prefer small functions.
- Handle errors explicitly.
- Do not leave dead code.

### Language

All code comments, documentation, error codes, and user-facing English strings should be in English unless the task explicitly requests otherwise.

### Commits and Changes

When working on a task:

- Keep changes focused.
- Avoid unrelated refactors.
- Update documentation when behavior changes.
- Do not modify generated/native platform files unless required.
- Do not introduce new major dependencies without a clear reason.

---

## Dependency Rules

Before adding a dependency, verify that:

- It solves a real problem.
- It is maintained.
- It does not duplicate existing functionality.
- It does not introduce unnecessary complexity.
- It fits the project stack.

Avoid adding large frameworks where a small library or standard library solution is enough.

---

## File Upload Rules

The backend must validate:

- Extension.
- MIME type.
- File size.
- Empty files.
- User authentication.
- AI provider readiness.

Supported formats:

```text
.pdf
.csv
.xls
.xlsx
```

Frontend validation is useful for UX but must never be the only validation.

---

## Error Handling Rules

Errors should be consistent and safe.

Backend errors should include:

- Stable error code.
- User-safe message.
- Optional details when safe.

Do not expose:

- API keys.
- Full provider raw responses.
- Internal stack traces.
- Other users' data.
- Full financial file contents.

---

## When Unsure

If a task is ambiguous:

1. Check `docs/`.
2. Check the relevant `README.md`.
3. Follow the MVP scope.
4. Prefer the simplest implementation aligned with the documented product rules.
5. Ask for clarification only when the decision would significantly change architecture or product behavior.

---

## Non-Goals for Agents

Do not implement the following unless explicitly requested:

- Bank API synchronization.
- Plaid or Belvo integration.
- OCR fallback.
- Local model support.
- Ollama support.
- llama.cpp support.
- Multi-provider routing.
- Budgeting.
- Payment system.
- App Store / Google Play release pipeline.
- Offline-first mode.
