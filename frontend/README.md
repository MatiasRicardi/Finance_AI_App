# Finance AI App Frontend

This is the frontend application for Finance AI App.

The frontend is built with Vue.js and Ionic. It is designed to work as a web application and later be packaged for Android and iOS using Capacitor.

## Technology Stack

| Area | Technology |
|---|---|
| Framework | Vue.js |
| UI Framework | Ionic Vue |
| Mobile Runtime | Capacitor |
| Build Tool | Vite |
| State Management | Pinia |
| Routing | Vue Router |
| Package Manager | pnpm |

## Responsibilities

The frontend should provide the user interface for:

- User registration.
- User login.
- Dashboard.
- AI provider settings.
- AI provider validation results.
- Financial statement upload.
- Import processing state.
- Import review.
- Transaction management.
- Categories.
- Reports.
- Profile and settings.

## Suggested Project Structure

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
└── README.md
```

## Main Screens

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

### Optional but Useful

- Import History.
- Import Detail.
- Data Export.

## Navigation

Recommended mobile navigation:

```text
Dashboard
Transactions
Import
Reports
Profile
```

On desktop, the app can use a sidebar layout while keeping the same routes.

## Environment Variables

Create a `.env` file in the `frontend/` directory.

```env
VITE_API_BASE_URL=http://localhost:8080
VITE_APP_NAME=Finance AI App
```

## Local Development

### Install Dependencies

```bash
pnpm install
```

### Start Development Server

```bash
pnpm dev
```

### Build for Production

```bash
pnpm build
```

### Preview Production Build

```bash
pnpm preview
```

## Ionic and Capacitor

The same frontend codebase is used for web and mobile.

Capacitor allows the Vue/Ionic app to be packaged for Android and iOS.

### Add Android Platform

```bash
pnpm ionic capacitor add android
```

### Add iOS Platform

```bash
pnpm ionic capacitor add ios
```

### Sync Web Build with Native Projects

```bash
pnpm ionic capacitor sync
```

### Open Android Project

```bash
pnpm ionic capacitor open android
```

### Open iOS Project

```bash
pnpm ionic capacitor open ios
```

## API Client

The frontend should use a centralized API client.

Responsibilities:

- Configure backend base URL.
- Attach authentication token/session if needed.
- Handle JSON requests.
- Handle multipart file uploads.
- Normalize API errors.
- Return user-friendly errors to pages and components.

Suggested location:

```text
src/api/
```

## State Management

Pinia should be used for shared application state.

Suggested stores:

```text
src/stores/auth.store.ts
src/stores/ai-provider.store.ts
src/stores/imports.store.ts
src/stores/transactions.store.ts
```

## Import Flow

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

## AI Provider Settings Flow

The AI provider settings screen should allow the user to configure:

- Configuration name.
- Provider type.
- Base URL.
- API key.
- Model name.

The user should be able to run a provider test.

The UI should display validation results for:

- Connection.
- Credentials.
- Model availability.
- PDF support.
- CSV support.
- Excel support.
- JSON output.

The import flow must be disabled unless the provider is ready.

## Error and Empty States

The frontend should include reusable components for:

- Loading state.
- Error message.
- Empty state.
- Retry action.
- Form validation errors.
- Unsupported file format message.
- AI provider not ready message.

## File Upload Rules

The frontend should allow the user to select:

- PDF.
- CSV.
- XLS.
- XLSX.

Frontend validation is useful for UX, but backend validation remains mandatory.

The frontend should show:

- Selected file name.
- Selected file size.
- Upload/processing state.
- Clear error message if the file is unsupported.

## Styling

The app should use Ionic components and responsive layouts.

Mobile should prioritize:

- Tabs.
- Cards.
- Simple lists.
- Bottom-friendly actions.

Desktop should prioritize:

- Wider tables.
- Sidebar navigation.
- More visible dashboard/report content.

## License

MIT
