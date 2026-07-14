package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"backend-go/handlers"
	"backend-go/services"

	"github.com/joho/godotenv"
)

// corsMiddleware allows requests from any localhost origin (dev-friendly).
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Handle preflight
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	godotenv.Load()

	region := os.Getenv("AWS_REGION")
	if region == "" {
		region = "us-east-1"
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	llm, err := services.NewLLMService(region)
	if err != nil {
		log.Fatalf("Failed to init LLM service: %v", err)
	}

	app := &handlers.AppHandler{LLM: llm}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", handlers.HealthCheck)
	mux.HandleFunc("POST /analyze", app.Analyze)          // upload resume + job → analyze → save JSON
	mux.HandleFunc("GET /analysis", handlers.GetAnalysis) // get last saved analysis

	fmt.Printf("Server running on http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, corsMiddleware(mux)))
}
