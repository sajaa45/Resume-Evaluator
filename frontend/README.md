# Candidate Screener — Frontend

A recruiter-facing web app for screening candidates by uploading a resume PDF and pasting a job description. The UI sends both to the Go backend, which runs the analysis via Amazon Bedrock and returns a structured result.

## Stack

- **React 19** + **TypeScript**
- **Vite 8** — dev server and bundler
- Plain **CSS** — no Tailwind, no component library

## Project Structure

```
frontend/
├── public/
│   └── logo.png              # Your brand logo (drop it here)
├── src/
│   ├── App.tsx               # Root component, state, API call
│   ├── App.css               # All component styles
│   └── index.css             # Global reset and CSS variables
└── components/
    ├── Header.tsx            # Top bar with logo and title
    ├── ResumeUploader.tsx    # Drag-and-drop PDF upload
    ├── JobDescriptionInput.tsx  # Job description textarea
    └── ResultsDisplay.tsx    # Score ring, conclusion, points cards
```

## Setup

### Prerequisites

- Node.js 18+
- Backend running on `http://localhost:8080` (see `../backend-go/README.md`)

### Install & Run

```bash
npm install
npm run dev
```

App runs at `http://localhost:5173`.

### Build

```bash
npm run build
```

## Logo

Drop your icon as `public/logo.png`. It will appear in the header automatically. PNG, SVG, or JPEG all work.

## How It Works

1. Recruiter uploads a candidate's resume (PDF) and pastes the job description
2. On clicking **Screen Candidate**, the app sends a `multipart/form-data` POST to `/analyze`
3. Results are displayed as:
   - An animated ring showing the **match percentage** (green ≥ 70%, amber 45–69%, red below 45%)
   - An amber **Conclusion** card (Yes / No / Maybe)
   - Side-by-side **Strong Points** (green) and **Weak Points** (red) cards
