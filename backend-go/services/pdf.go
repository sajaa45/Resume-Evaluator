package services

import (
	"fmt"
	"strings"

	"github.com/gen2brain/go-fitz"
)

// ExtractText opens a PDF file and returns all text content as a single string
func ExtractText(filePath string) (string, error) {
	doc, err := fitz.New(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to open PDF: %w", err)
	}
	defer doc.Close()

	var sb strings.Builder
	for i := 0; i < doc.NumPage(); i++ {
		text, err := doc.Text(i)
		if err != nil {
			return "", fmt.Errorf("failed to read page %d: %w", i+1, err)
		}
		sb.WriteString(text)
		sb.WriteString("\n")
	}

	return sb.String(), nil
}
