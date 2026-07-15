package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"backend-go/models"
	"backend-go/services"
)

type AppHandler struct {
	LLM *services.LLMService
}

const analysisPath = "uploads/analysis.json"

// POST /analyze
// Accepts resume (PDF) + job (txt), runs LLM, saves and returns JSON
func (h *AppHandler) Analyze(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)

	// Save and parse resume PDF
	resumeFile, resumeHeader, err := r.FormFile("resume")
	if err != nil {
		writeJSON(w, http.StatusBadRequest, models.ErrorResponse{Error: "missing resume file"})
		return
	}
	defer resumeFile.Close()

	resumePath := "uploads/" + resumeHeader.Filename
	out, _ := os.Create(resumePath)
	buf := make([]byte, 10<<20)
	n, _ := resumeFile.Read(buf)
	out.Write(buf[:n])
	out.Close()

	resumeText, err := services.ExtractText(resumePath)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, models.ErrorResponse{Error: "failed to parse resume: " + err.Error()})
		return
	} // Read job description as plain text field
	jobText := r.FormValue("job")
	if jobText == "" {
		writeJSON(w, http.StatusBadRequest, models.ErrorResponse{Error: "missing job field"})
		return
	}

	prompt := `You are a senior technical recruiter. No thinking, no reasoning, no explanations.
Only respond in this exact format, nothing else:

STRONG POINTS:
- point
- point
- point

WEAK POINTS:
- point
- point
- point

CONCLUSION:
Yes / No / Maybe

MATCH PERCENTAGE:
X%

Job Description:
` + jobText + `

Resume:
` + resumeText

	response, err := h.LLM.InvokeModel(prompt)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}

	// Strip <reasoning> block if present
	if start := strings.Index(response, "<reasoning>"); start != -1 {
		if end := strings.Index(response, "</reasoning>"); end != -1 {
			response = strings.TrimSpace(response[end+len("</reasoning>"):])
		}
	}

	analysis := parseAnalysis(response)

	// Save as JSON
	data, _ := json.MarshalIndent(analysis, "", "  ")
	os.WriteFile(analysisPath, data, 0644)

	writeJSON(w, http.StatusOK, analysis)
}

// GET /analysis
// Returns the last saved analysis JSON
func GetAnalysis(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile(analysisPath)
	if err != nil {
		writeJSON(w, http.StatusNotFound, models.ErrorResponse{Error: "no analysis found, run POST /analyze first"})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

// GET /health
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

// parseAnalysis splits the LLM response into structured sections
func parseAnalysis(raw string) models.AnalysisResponse {
	var result models.AnalysisResponse

	lines := strings.Split(raw, "\n")
	var current *[]string

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		upper := strings.ToUpper(line)

		switch {
		case strings.HasPrefix(upper, "STRONG POINTS:"):
			current = &result.StrongPoints
		case strings.HasPrefix(upper, "WEAK POINTS:"):
			current = &result.WeakPoints
		case strings.HasPrefix(upper, "CONCLUSION:"):
			current = nil
		case strings.HasPrefix(upper, "MATCH PERCENTAGE:"):
			current = nil
			pct := strings.TrimSpace(strings.TrimPrefix(line, "MATCH PERCENTAGE:"))
			if pct != "" {
				result.MatchPercentage = pct
			}
		case result.Conclusion == "" && (strings.HasPrefix(upper, "YES") || strings.HasPrefix(upper, "NO") || strings.HasPrefix(upper, "MAYBE")):
			result.Conclusion = line
		case strings.HasSuffix(line, "%") && result.MatchPercentage == "":
			result.MatchPercentage = line
		case current != nil && (strings.HasPrefix(line, "-") || strings.HasPrefix(line, "•")):
			point := strings.TrimPrefix(strings.TrimPrefix(line, "-"), "•")
			*current = append(*current, strings.TrimSpace(point))
		}
	}

	return result
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}
