# Finance AI App

Finance AI App is a personal finance application that uses AI to extract and classify transactions from PDF, CSV, and Excel financial statements.

The application allows users to upload financial statement files, send them to a configured AI provider, review the extracted transactions, and save them for tracking expenses, income, and basic financial reports.

## Overview

Core flow:

```text
User uploads PDF/CSV/XLSX
→ Backend sends the file to the configured AI provider
→ AI provider extracts structured transaction data
→ Backend validates the response
→ User reviews and corrects transactions
→ User confirms the import
→ Transactions are saved
→ Dashboard and reports are updated
```

## Monorepo Structure

```text
finance_ai_app/
├── backend/
├── frontend/
├── docs/
├── scripts/
├── README.md
└── .gitignore
```

## Applications

### Backend

The backend is responsible for authentication, AI provider configuration, secure API key storage, file uploads, AI provider validation, transaction storage, dashboard data, and reports.

Technology:

- Go
- Chi
- MariaDB
- SQL migrations
- Explicit SQL queries without an ORM

See [backend/README.md](backend/README.md).

### Frontend

The frontend is responsible for the web and mobile-ready user interface.

Technology:

- Vue.js
- Ionic Vue
- Capacitor
- Vite
- Pinia
- Vue Router
- pnpm

See [frontend/README.md](frontend/README.md).

## Technology Stack

| Area | Technology |
|---|---|
| Backend | Go |
| HTTP Router | Chi |
| Database | MariaDB |
| Database Access | Explicit SQL, no ORM |
| Frontend | Vue.js |
| UI Framework | Ionic Vue |
| Mobile Runtime | Capacitor |
| Build Tool | Vite |
| Package Manager | pnpm |
| AI Providers | OpenAI-compatible APIs with direct file input support |

## AI Provider Requirement

The configured AI provider must support direct file input for:

- PDF files.
- CSV files.
- XLS/XLSX files.
- Structured JSON output.

Text-only OpenAI-compatible APIs are not supported in the first version.

## Main Features

- User registration and login.
- User profile.
- AI provider configuration per user.
- AI provider validation.
- PDF, CSV, XLS, and XLSX upload.
- AI-based transaction extraction.
- Editable import review screen.
- Transaction confirmation.
- Transaction list.
- Manual transaction management.
- Categories.
- Dashboard.
- Basic reports.
- Import history.

## Local Development

### Requirements

- Go 1.26+
- Node.js
- pnpm
- Access to a MariaDB server (configured via environment variables)

### Start Backend

```bash
cd backend
cp .env.example .env   # first time only — fill in your credentials
export $(cat .env | grep -v '^#' | xargs) && go run ./cmd/api
```

### Start Frontend

```bash
cd frontend
cp .env.example .env   # first time only
pnpm install
pnpm dev
```

## Environment Variables

Recommended environment files:

```text
backend/.env
frontend/.env
```

See each application README for details.

## Documentation

Project documentation should live in the `docs/` folder.

## License

MIT
