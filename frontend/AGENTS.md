# AGENTS.md

## Purpose

This file defines frontend-specific standards for code agents working inside the `frontend/` directory of Finance AI App.

Agents must follow these instructions when modifying frontend code.

---

## Frontend Summary

The frontend is a Vue.js and Ionic application.

It provides the user interface for:

- Authentication.
- Dashboard.
- AI provider settings.
- Financial statement upload.
- Import processing.
- Import review.
- Transactions.
- Categories.
- Reports.
- Profile and settings.

The same frontend codebase is used for web and mobile through Capacitor.

---

## Frontend Stack

Use:

- Vue.js.
- Ionic Vue.
- Capacitor.
- Vite.
- Pinia.
- Vue Router.
- pnpm.

Do not create a separate mobile app.

Android and iOS projects, when generated, belong inside the frontend project.

---

## Expected Frontend Structure

Recommended structure:

```text
frontend/
├── src/
│   ├── api/
│   ├── components/
│   ├── pages/
│   ├── router/
│   ├── stores/
│   ├── theme/
│   ├── utils/
│   ├── App.vue
│   └── main.ts
├── public/
├── android/
├── ios/
├── capacitor.config.ts
├── index.html
├── package.json
├── pnpm-lock.yaml
├── vite.config.ts
├── .env.example
├── README.md
└── AGENTS.md
```

---

## Main Screens

The MVP frontend should include:

### Authentication

- Login.
- Register.

### Main App

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
- General Settings if needed.

### Optional MVP Screens

- Import History.
- Import Detail.
- Data Export.

---

## Routing Rules

Use Vue Router.

Routes should be organized clearly and protected when needed.

Authenticated routes should require the user to be logged in.

Unauthenticated users should be redirected to login.

Suggested route groups:

```text
/auth/login
/auth/register
/app/dashboard
/app/transactions
/app/import
/app/import/:id/review
/app/import/:id/result
/app/reports
/app/profile
/app/settings/ai-provider
/app/settings/categories
```

The exact paths may change, but the route purpose should remain clear.

---

## Layout Rules

Use Ionic layouts.

Mobile navigation should prioritize tabs:

```text
Dashboard
Transactions
Import
Reports
Profile
```

Desktop may use a sidebar layout while keeping the same routes.

Do not design screens only for desktop.

Do not design screens only for mobile.

All primary screens should be responsive.

---

## API Client Rules

All backend communication should go through a centralized API client.

Suggested location:

```text
src/api/
```

The API client should handle:

- Base API URL from environment variables.
- Auth token/session handling.
- JSON requests.
- Multipart file uploads.
- Standard API error parsing.
- Request cancellation when useful.
- User-friendly error mapping.

Do not call `fetch` or `axios` directly from many unrelated components.

Components should call stores or composables that use the API client.

---

## Environment Variables

Use Vite environment variables.

Required example:

```text
VITE_API_BASE_URL=http://localhost:8080
VITE_APP_NAME=Finance AI App
```

Do not hardcode backend URLs in components.

Update `.env.example` when adding new variables.

---

## State Management Rules

Use Pinia for shared state.

Suggested stores:

```text
src/stores/auth.store.ts
src/stores/ai-provider.store.ts
src/stores/imports.store.ts
src/stores/transactions.store.ts
src/stores/categories.store.ts
```

Rules:

- Keep stores focused.
- Do not store secrets.
- Do not store full API keys.
- Keep server state reloadable.
- Avoid duplicating large data unnecessarily.
- Clear auth-related state on logout.

---

## Security Rules

The frontend must not:

- Store provider API keys after submission.
- Display full API keys.
- Call AI providers directly.
- Save transactions without backend confirmation.
- Trust AI output as final.
- Expose sensitive provider error details.
- Bypass backend validation.

The frontend may display masked API keys only if returned by the backend.

Example:

```text
sk-...ab91
```

---

## AI Provider Settings UI Rules

The AI provider settings screen should allow the user to configure:

- Configuration name.
- Provider type.
- Base URL.
- API key.
- Model name.

The screen should support:

- Save settings.
- Test provider.
- Show validation results.
- Show provider status.
- Display actionable error messages.

Provider readiness must be clearly visible.

Import must be blocked unless provider status is `ready`.

Validation result categories:

- Connection.
- Credentials.
- Model availability.
- PDF support.
- CSV support.
- Excel support.
- JSON output.

---

## Import Flow UI Rules

The import flow is central to the app.

Required flow:

```text
Import Statement Screen
→ Select PDF/CSV/XLS/XLSX
→ Upload file
→ Show processing state
→ Navigate to Import Review Screen
→ User edits extracted transactions
→ User confirms import
→ Show Import Result Screen
→ User views Transactions or Dashboard
```

Rules:

- Show selected file name.
- Show selected file size.
- Reject obviously unsupported file types in the UI.
- Show loading state during upload.
- Prevent duplicate submissions.
- Show processing state.
- Show clear errors.
- Do not save transactions automatically.
- Always require user confirmation.

---

## Import Review UI Rules

The review screen must allow users to inspect and correct AI output.

Show:

- File name.
- Statement metadata.
- Transaction count.
- Total income.
- Total expenses.
- Net result.
- AI warnings.
- Low-confidence items.
- Editable transaction list.

Each transaction should allow editing:

- Date.
- Description.
- Amount.
- Currency.
- Type.
- Category.
- Notes if available.

The user should be able to:

- Remove incorrect transactions.
- Cancel import.
- Confirm import.

The confirmed data submitted to the backend must reflect the user's edits.

---

## Transaction UI Rules

The transaction list should support:

- List transactions.
- Search by description.
- Filter by date range.
- Filter by category.
- Filter by type.
- Filter by currency.
- Edit transaction.
- Delete transaction.
- Create manual transaction.

Mobile may use cards.

Desktop may use a table-like layout.

---

## Forms and Validation

Forms should use clear validation.

Validation should exist for:

- Required fields.
- Email format.
- Password confirmation.
- AI provider base URL.
- AI provider model.
- File type.
- Transaction date.
- Transaction amount.
- Transaction type.
- Transaction category.

Frontend validation is for user experience only.

Backend validation remains mandatory.

---

## Error Handling Rules

Use reusable error components.

Errors should be:

- Clear.
- User-friendly.
- Actionable.
- Safe.

Do not display raw stack traces or sensitive backend details.

Recommended error examples:

```text
You need to configure a compatible AI provider before importing statements.
```

```text
Unsupported file format. Please upload a PDF, CSV, XLS, or XLSX file.
```

```text
The AI provider returned an invalid response. Please try again or choose another model.
```

---

## Empty State Rules

Use empty states for:

- No transactions.
- No imports.
- No reports.
- No AI provider configured.
- No categories if custom category list is empty.

Empty states should include a next action when possible.

Example:

```text
You do not have any transactions yet.
Import your first financial statement to start tracking your expenses.
```

---

## Styling Rules

Use Ionic components first.

Prefer:

- `ion-page`.
- `ion-header`.
- `ion-content`.
- `ion-list`.
- `ion-item`.
- `ion-card`.
- `ion-modal`.
- `ion-toast`.
- `ion-alert`.
- `ion-loading`.
- `ion-button`.

Keep UI clean and responsive.

Mobile should prioritize:

- Cards.
- Tabs.
- Short actions.
- Bottom-friendly interaction.

Desktop should prioritize:

- Tables.
- Wider layouts.
- Sidebar navigation.
- Dashboard/report readability.

---

## Capacitor Rules

Capacitor is used to package the same frontend app for mobile.

Do not manually edit generated native files unless necessary.

When adding native behavior:

- Prefer Capacitor plugins.
- Document required permissions.
- Test on web and mobile when possible.
- Keep web behavior working.

Useful commands:

```bash
pnpm ionic capacitor add android
pnpm ionic capacitor add ios
pnpm ionic capacitor sync
pnpm ionic capacitor open android
pnpm ionic capacitor open ios
```

---

## Package Manager Rules

Use `pnpm`.

Do not use `npm install` or `yarn` unless explicitly requested.

Use:

```bash
pnpm install
pnpm dev
pnpm build
pnpm preview
```

---

## Testing and Quality Rules

Add or update tests when modifying:

- Auth forms.
- AI provider settings.
- Import upload.
- Import review.
- Transaction editing.
- Route guards.
- API error handling.
- Stores.

Run relevant checks:

```bash
pnpm lint
pnpm test
pnpm build
```

If scripts are not available yet, do not claim they pass.

---

## Code Style

- Use Vue single-file components.
- Prefer Composition API.
- Keep components small.
- Extract reusable components.
- Keep business logic out of templates.
- Use stores/composables for shared logic.
- Avoid duplicating API calls.
- Use TypeScript if the project is configured for it.
- Keep naming consistent.
- Use English for code, comments, and UI copy.

---

## Do Not Implement Without Explicit Request

- Separate mobile app.
- Direct frontend AI provider calls.
- Full API key display.
- Offline-first architecture.
- Push notifications.
- App Store release pipeline.
- Google Play release pipeline.
- Bank integrations.
- OCR UI flow.
- Multi-provider routing UI.
