# Resume Analyzer — Backend (Go)

A Go backend that parses a resume PDF and evaluates it against a job description using Amazon Bedrock.

## Stack

- **Go** — HTTP server
- **Amazon Bedrock** — LLM inference (`openai.gpt-oss-20b-1:0`)
- **MuPDF / go-fitz** — PDF text extraction

## Project Structure

```
backend-go/
├── main.go                 # Server setup and routes
├── .env                    # Environment variables
├── go.mod
├── handlers/
│   └── application.go      # HTTP endpoints
├── services/
│   ├── llm.go              # Bedrock LLM client
│   └── pdf.go              # PDF text extractor
├── models/
│   └── models.go           # Request/response structs
└── uploads/                # Uploaded files and saved analysis
```

## Setup

### 1. Prerequisites

- Go 1.22+
- MuPDF installed via Homebrew:
  ```bash
  brew install mupdf
  ```
- An AWS account with Bedrock access and the model enabled

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

## API Endpoints

### `GET /health`
Returns server status.

```bash
curl http://localhost:8080/health
```

---

### `POST /analyze`
Upload a resume PDF and a job description. Returns a structured analysis.

**Form fields:**
| Field | Type | Description |
|-------|------|-------------|
| `resume` | file | Resume in PDF format |
| `job` | text | Job description as plain text |

```bash
curl -X POST http://localhost:8080/analyze \
  -F "resume=@uploads/resume.pdf" \
  -F "job=We are looking for an AI Engineer with Python and LangChain experience..."
```

**Response:**
```json
{
  "strong_points": [
    "Proficient in Python and LangChain",
    "Experience with RAG pipelines"
  ],
  "weak_points": [
    "No experience with Google ADK",
    "Limited MLOps exposure"
  ],
  "conclusion": "Maybe",
  "match_percentage": "65%"
}
```

---

### `GET /analysis`
Returns the last saved analysis result.

```bash
curl http://localhost:8080/analysis
```
