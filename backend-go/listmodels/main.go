package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/bedrock"
	"github.com/joho/godotenv"
)

func main() {
	// Try both relative paths
	if err := godotenv.Load("../.env"); err != nil {
		godotenv.Load(".env")
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(os.Getenv("AWS_REGION")),
	)
	if err != nil {
		log.Fatal(err)
	}

	client := bedrock.NewFromConfig(cfg)
	resp, err := client.ListFoundationModels(context.TODO(), &bedrock.ListFoundationModelsInput{})
	if err != nil {
		log.Fatal(err)
	}

	for _, m := range resp.ModelSummaries {
		if m.ModelLifecycle != nil && string(m.ModelLifecycle.Status) == "ACTIVE" {
			fmt.Println(*m.ModelId)
		}
	}
}
