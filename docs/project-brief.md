# Project Brief

## Project Name

Finance AI App

## Repository Name

`finance_ai_app`

---

## Project Summary

Finance AI App is a personal finance application that uses AI to extract and classify transactions from financial statement files.

Users can upload financial statements in **PDF, CSV, XLS, or XLSX** format. The backend sends the uploaded file directly to a user-configured AI provider. The AI provider extracts transactions and returns structured data. The user then reviews, corrects, and confirms the extracted transactions before they are saved.

The application is designed as a practical product and learning project focused on building a real AI-powered workflow for financial document processing.

---

## Problem Statement

Many users receive financial data in different file formats, such as bank statements, credit card summaries, CSV exports, or Excel files.

Manually reviewing those files, identifying transactions, classifying expenses, and creating reports is repetitive and time-consuming.

This application aims to simplify that process by using AI to extract and classify financial transactions while still keeping the user in control through a mandatory review step.

---

## Product Goal

The main goal is to build a personal finance application that allows users to:

- Upload financial statement files.
- Process those files with a configurable AI provider.
- Extract structured transaction data.
- Classify expenses and income.
- Review and correct AI-generated results.
- Save confirmed transactions.
- View basic financial summaries and reports.

The MVP should validate the complete flow from file upload to transaction storage and reporting.

---

## Target Users

The initial target user is an individual who wants to track personal finances by importing financial statement files manually.

Typical user needs:

- Understand monthly expenses.
- Identify spending by category.
- Track income and expenses.
- Import data from bank or credit card statements.
- Avoid manually entering every transaction.
- Keep control over AI-generated results.

---

## Core Use Case

The core use case is:

```text
A user uploads a financial statement file.
The app sends the file to the configured AI provider.
The AI extracts transactions.
The user reviews and corrects the transactions.
The user confirms the import.
The app saves the transactions and updates dashboard/reports.
```

---

## MVP Scope

The MVP includes the following core capabilities.

### User Management

- User registration.
- User login.
- Authenticated app access.
- Basic user profile.

### AI Provider Configuration

- The user can configure an AI provider.
- The provider must be OpenAI-compatible or compatible with the expected API integration.
- The provider must support direct file input.
- The provider must support PDF, CSV, and Excel files.
- The system validates the provider before allowing imports.

### File Import

Supported file types:

- PDF.
- CSV.
- XLS.
- XLSX.

The backend validates:

- File extension.
- MIME type.
- File size.
- Empty file.
- AI provider readiness.

### AI Extraction

The backend sends the uploaded file directly to the configured AI provider.

The AI provider should return structured JSON containing:

- Statement metadata.
- Transactions.
- Categories.
- Confidence values.
- Warnings.

### Import Review

Before saving transactions, the user can:

- Review extracted transactions.
- Edit dates.
- Edit descriptions.
- Edit amounts.
- Edit currencies.
- Change transaction types.
- Change categories.
- Remove incorrect transactions.
- Confirm or cancel the import.

### Transaction Management

The app allows users to:

- View saved transactions.
- Search transactions.
- Filter transactions.
- Edit transactions.
- Delete transactions.
- Create manual transactions.

### Dashboard and Reports

The MVP includes basic reporting:

- Monthly income.
- Monthly expenses.
- Monthly balance.
- Recent transactions.
- Spending by category.
- Income vs expenses.
- Monthly spending trend.

---

## Out of Scope for MVP

The following features are intentionally excluded from the first version:

- Bank account synchronization.
- Plaid, Belvo, or Open Finance integrations.
- Automatic recurring imports.
- OCR fallback for scanned PDFs.
- Manual backend text extraction fallback.
- Text-only local model support.
- Ollama support.
- llama.cpp support.
- Multi-provider routing.
- Budgeting features.
- Shared accounts.
- Family accounts.
- Business accounting features.
- Investment tracking.
- App Store release.
- Google Play release.
- Offline-first mode.
- Push notifications.
- Payment or subscription system.

---

## Key Product Decisions

### 1. AI Provider Must Support Files

The MVP requires the configured AI provider to support direct file input.

The required formats are:

- PDF.
- CSV.
- XLS.
- XLSX.

A provider that only supports text chat completions is not enough.

### 2. User Review Is Mandatory

AI-generated data is never saved automatically.

The user must review and confirm extracted transactions before they become final.

### 3. One Active AI Provider Per User

The MVP supports one active AI provider configuration per user.

Future versions may support multiple providers.

### 4. No ORM

The backend will not use an ORM.

Database access will use Go's standard `database/sql` package with explicit SQL queries.

### 5. Monorepo

The project uses a monorepo structure with separate `backend/`, `frontend/`, and `docs/` folders.

---

## Technology Stack

### Backend

- Go.
- Chi.
- MariaDB.
- `database/sql`.
- SQL migrations.
- No ORM.

### Frontend

- Vue.js.
- Ionic Vue.
- Capacitor.
- Vite.
- Pinia.
- Vue Router.
- pnpm.

### Platforms

- Web.
- Android through Capacitor.
- iOS through Capacitor.

---

## Proposed Repository Structure

```text
finance_ai_app/
├── backend/
│   ├── cmd/
│   ├── internal/
│   ├── migrations/
│   ├── go.mod
│   └── README.md
│
├── frontend/
│   ├── src/
│   ├── public/
│   ├── android/
│   ├── ios/
│   ├── capacitor.config.ts
│   ├── package.json
│   └── README.md
│
├── docs/
│   ├── project-brief.md
│   ├── mvp-scope.md
│   ├── main-user-flow.md
│   ├── architecture.md
│   └── api-contracts.md
│
├── scripts/
├── docker-compose.yml
├── README.md
└── .gitignore
```

---

## Main User Flow

```text
Register
→ Login
→ Configure AI Provider
→ Test AI Provider
→ Upload PDF/CSV/XLS/XLSX
→ Process with AI
→ Validate AI Response
→ Review Transactions
→ Confirm Import
→ Save Transactions
→ View Dashboard and Reports
```

---

## Main Screens

### Authentication

- Login.
- Register.

### Main Application

- Dashboard.
- Transactions.
- Import Statement.
- Import Processing.
- Import Review.
- Import Result.
- Reports.
- Profile.

### Settings

- AI Provider Settings.
- Categories.
- General Settings.

### Optional MVP Screens

- Import History.
- Import Detail.
- Data Export.

---

## Main Backend Modules

The backend should be organized around the following modules:

- Auth.
- Users.
- AI provider settings.
- AI provider validation.
- File uploads.
- Import processing.
- AI extraction.
- AI response validation.
- Transactions.
- Categories.
- Dashboard.
- Reports.
- Import history.
- Encryption.
- Logging.
- Database.

---

## Main Frontend Modules

The frontend should be organized around:

- API client.
- Auth store.
- AI provider store.
- Import store.
- Transaction store.
- Router.
- Pages.
- Shared components.
- Ionic layouts.
- Error components.
- Loading components.
- Empty state components.

---

## AI Provider Validation

Before a provider can be used, the backend must validate:

- Base URL is reachable.
- API key is valid.
- Model can be used.
- PDF input is supported.
- CSV input is supported.
- XLS/XLSX input is supported.
- Structured JSON output is valid.

Possible provider statuses:

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

---

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

### uploaded

The backend accepted the file and created an import record.

### processing

The backend is sending the file to the AI provider and validating the response.

### awaiting_review

The AI response was valid and extracted transactions are ready for user review.

### completed

The user confirmed the import and the transactions were saved.

### failed

The import failed because of upload, provider, validation, or processing error.

### cancelled

The user cancelled the import before saving transactions.

---

## Security Considerations

The application handles sensitive financial data.

The MVP must follow these rules:

- Store API keys encrypted at rest.
- Never return full API keys to the frontend.
- Mask API keys in the UI.
- Never log API keys.
- Never log full financial file contents.
- Validate all uploaded files.
- Restrict users to their own data.
- Use safe error messages.
- Clean up temporary files after processing or after a retention period.

---

## Success Criteria

The MVP is successful when a user can:

1. Create an account.
2. Configure an AI provider.
3. Validate that the provider supports PDF, CSV, and Excel files.
4. Upload a financial statement.
5. Receive extracted transactions from AI.
6. Review and correct the extracted data.
7. Confirm the import.
8. See transactions saved in the app.
9. View dashboard and basic reports based on imported data.

---

## Future Opportunities

After the MVP, the project may evolve with:

- Multiple AI providers per user.
- Provider selection per import.
- Support for local models through text extraction mode.
- OCR support.
- Better duplicate detection.
- Merchant-based categorization rules.
- User-trained categorization rules.
- Budgeting features.
- Recurring transactions.
- Subscription detection.
- Multi-currency improvements.
- Export to CSV, Excel, or JSON.
- Mobile app store releases.
- Bank API integrations.
