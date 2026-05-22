# Finance AI App — Project Summary

## Overview

**Finance AI App** is a personal finance application designed to help users import, review, classify, and analyze their financial transactions using artificial intelligence.

The main idea is that users can upload financial statement files, such as **PDF, CSV, XLS, or XLSX**, and the application will send those files to a configured AI provider. The AI provider will analyze the file, extract the transactions, classify them into categories, and return structured data that the app can process.

The user always keeps final control. Before saving anything, the app shows a review screen where the user can check the extracted transactions, correct mistakes, change categories, remove incorrect rows, and then confirm the import.

---

## General Idea

The project is based on this flow:

```text
User uploads a financial statement
   ↓
Backend validates the file
   ↓
Backend sends the file to the configured AI provider
   ↓
AI extracts and classifies transactions
   ↓
Backend validates the AI response
   ↓
User reviews the extracted data
   ↓
User confirms the import
   ↓
Transactions are saved
   ↓
Dashboard and reports are updated
```

The purpose of the application is to reduce the manual work involved in tracking personal finances.

Instead of manually entering every transaction, the user can import a statement and let AI perform the first extraction and classification step.

---

## Main Goal

The goal of the project is to build a personal finance app that can:

- Import financial statements.
- Extract transactions using AI.
- Classify income and expenses.
- Let users review and correct AI results.
- Store confirmed transactions.
- Show dashboards and reports.

The project is also useful as a learning experience because it involves real-world software development concepts such as:

- Go backend development.
- Vue and Ionic frontend development.
- AI provider integration.
- File uploads.
- Secure API key storage.
- Database design.
- Transaction processing.
- Mobile-ready web applications.

---

## AI Provider Concept

One important part of the project is that the user configures their own AI provider.

The app includes a settings screen where the user can enter:

- Provider name.
- Base URL.
- API key.
- Model name.

The first version focuses on **OpenAI-compatible providers**, but with a strict requirement:

> The configured provider must support direct file input.

The provider must support:

- PDF.
- CSV.
- XLS.
- XLSX.

If the provider does not support files, the app will not allow imports and will show a clear error.

---

## Why User Review Is Required

AI can make mistakes, especially when reading financial statements.

Because of that, the app does not automatically save extracted transactions.

The user must first review the extracted data.

During review, the user can:

- Fix dates.
- Fix descriptions.
- Fix amounts.
- Change categories.
- Change transaction type.
- Remove incorrect transactions.
- Confirm or cancel the import.

Only after confirmation are the transactions saved permanently.

---

## Main Parts of the App

The app has three main areas:

1. Backend.
2. Frontend.
3. Documentation.

---

## Backend

The backend will be built with:

- Go.
- Chi.
- MariaDB.
- `database/sql`.

The backend is responsible for:

- Authentication.
- User data.
- AI provider settings.
- API key encryption.
- File uploads.
- AI provider validation.
- AI extraction.
- Transaction storage.
- Dashboard APIs.
- Report APIs.

The backend will not use an ORM. Database access will use explicit SQL through Go's standard `database/sql` package.

---

## Frontend

The frontend will be built with:

- Vue.js.
- Ionic.
- Capacitor.
- Pinia.
- Vue Router.
- pnpm.

The same frontend project will support:

- Web app.
- Android app through Capacitor.
- iOS app through Capacitor.

There is no need for a separate mobile repository because Ionic and Capacitor allow the same frontend codebase to be packaged for mobile platforms.

---

## Documentation

The project includes documentation inside the `docs/` folder.

Suggested documents:

- `project-brief.md`
- `project-summary.md`
- `mvp-scope.md`
- `main-user-flow.md`
- `architecture.md`
- `api-contracts.md`

---

## Main Screens

The app will include screens such as:

- Login.
- Register.
- Dashboard.
- Import Statement.
- Import Processing.
- Import Review.
- Transactions.
- Reports.
- Profile.
- AI Provider Settings.
- Categories.
- Import History.

---

## MVP Focus

The first version of the app should focus on completing the full import flow:

```text
Register
→ Configure AI Provider
→ Test AI Provider
→ Upload Statement
→ Process with AI
→ Review Transactions
→ Confirm Import
→ View Transactions
→ View Dashboard
```

Advanced features are intentionally left for future versions.

Examples of future features:

- Bank synchronization.
- OCR support.
- Local model support.
- Budgeting.
- App Store release.
- Google Play release.
- Advanced duplicate detection.
- Multi-provider support.
- Merchant-based categorization rules.

---

## In Simple Terms

Finance AI App is an app where the user uploads a bank or credit card statement, the AI reads it and extracts the transactions, the user reviews the result, and then the app saves the data to show expenses, income, categories, and reports.

The project combines personal finance tracking with AI-powered file processing, using a Go backend and a Vue/Ionic frontend.

---

## One-Sentence Summary

Finance AI App is a personal finance application that uses a user-configured AI provider to extract, classify, review, and store transactions from PDF, CSV, and Excel financial statements.
