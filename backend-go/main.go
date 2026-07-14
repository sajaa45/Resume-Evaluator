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
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
