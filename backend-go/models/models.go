package models

// Response from PDF upload
type UploadResponse struct {
	Text string `json:"text"`
}

// Response from resume analysis
type AnalysisResponse struct {
	StrongPoints    []string `json:"strong_points"`
	WeakPoints      []string `json:"weak_points"`
	Conclusion      string   `json:"conclusion"`
	MatchPercentage string   `json:"match_percentage"`
}

// Generic error response
type ErrorResponse struct {
	Error string `json:"error"`
}
