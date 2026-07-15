package services

import (
	"strings"
	"testing"
)

func TestExtractText(t *testing.T) {
	text, err := ExtractText("testdata/sample.pdf")
	if err != nil {
		t.Fatalf("ExtractText returned error: %v", err)
	}

	want := "Hello Resume Screener"
	if !strings.Contains(text, want) {
		t.Errorf("ExtractText() = %q, want it to contain %q", text, want)
	}
}

func TestExtractText_MissingFile(t *testing.T) {
	_, err := ExtractText("testdata/does-not-exist.pdf")
	if err == nil {
		t.Fatal("ExtractText() expected an error for a missing file, got nil")
	}
}
