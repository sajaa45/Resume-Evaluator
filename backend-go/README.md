# Candidate Screener — Backend (Go)

Go HTTP server that accepts a candidate's resume PDF and a job description, runs them through an LLM on Amazon Bedrock, and returns a structured screening result for recruiters.

## Stack

- **Go** — HTTP server (`net/http`)
- **Amazon Bedrock** — LLM inference (`openai.gpt-oss-20b-1:0`)
- **MuPDF / go-fitz** — PDF text extraction

## Project Structure

```
backend-go/
├── main.go                 # Server setup, routes, CORS middleware
├── .env                    # Environment variables (not committed)
├── go.mod / go.sum
├── handlers/
│   └── application.go      # HTTP endpoints and response parsing
├── services/
│   ├── llm.go              # Amazon Bedrock client
│   └── pdf.go              # PDF text extractor
├── models/
│   └── models.go           # Request/response structs
└── uploads/                # Uploaded PDFs and saved analysis JSON
```

## Setup

### 1. Prerequisites

- Go 1.22+
- MuPDF installed:
  ```bash
  brew install mupdf
  ```
- AWS account with Bedrock access and the model enabled

### 2. Environment Variables

Fill in `.env`:

```env
AWS_ACCESS_KEY_ID=your_access_key
AWS_SECRET_ACCESS_KEY=your_secret_key
AWS_REGION=us-east-1
PORT=8080
```

### 3. Run

```bash
CGO_CFLAGS="-I/opt/homebrew/opt/mupdf/include" \
CGO_LDFLAGS="-L/opt/homebrew/opt/mupdf/lib -lmupdf -lmupdf-third" \
go run main.go
```

Server starts at `http://localhost:8080`.

## API

### `GET /health`
Returns server status.

```bash
curl http://localhost:8080/health
# {"status":"ok"}
```

---

### `POST /analyze`
Upload a candidate's resume PDF and your job description. Returns a structured screening result.

**Form fields:**

| Field | Type | Description |
|-------|------|-------------|
| `resume` | file | Candidate resume in PDF format |
| `job` | text | Job description as plain text |

```bash
curl -X POST http://localhost:8080/analyze \
  -F "resume=@uploads/candidate.pdf" \
  -F "job=We are looking for a Senior Go engineer with Kubernetes experience..."
```

**Response:**
```json
{
  "strong_points": [
    "5 years of Go experience",
    "Kubernetes and Docker proficiency"
  ],
  "weak_points": [
    "No mention of CI/CD pipelines",
    "Limited cloud infrastructure exposure"
  ],
  "conclusion": "Yes",
  "match_percentage": "78%"
}
```

---

### `GET /analysis`
Returns the most recently saved screening result.

```bash
curl http://localhost:8080/analysis
```

## CORS

The server allows all origins (`*`) in development. Restrict `Access-Control-Allow-Origin` to your frontend domain before deploying to production.
