package handlers

import (
	"reflect"
	"testing"

	"backend-go/models"
)

func TestParseAnalysis(t *testing.T) {
	raw := `STRONG POINTS:
- 5 years of Go experience
- Kubernetes and Docker proficiency

WEAK POINTS:
- No mention of CI/CD pipelines
- Limited cloud infrastructure exposure

CONCLUSION:
Yes

MATCH PERCENTAGE:
78%`

	want := models.AnalysisResponse{
		StrongPoints:    []string{"5 years of Go experience", "Kubernetes and Docker proficiency"},
		WeakPoints:      []string{"No mention of CI/CD pipelines", "Limited cloud infrastructure exposure"},
		Conclusion:      "Yes",
		MatchPercentage: "78%",
	}

	got := parseAnalysis(raw)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("parseAnalysis() = %+v, want %+v", got, want)
	}
}

func TestParseAnalysis_BulletVariant(t *testing.T) {
	raw := `STRONG POINTS:
• Strong Go background

WEAK POINTS:
• No tests

CONCLUSION:
Maybe

MATCH PERCENTAGE:
50%`

	got := parseAnalysis(raw)

	if len(got.StrongPoints) != 1 || got.StrongPoints[0] != "Strong Go background" {
		t.Errorf("StrongPoints = %v, want [\"Strong Go background\"]", got.StrongPoints)
	}
	if got.Conclusion != "Maybe" {
		t.Errorf("Conclusion = %q, want %q", got.Conclusion, "Maybe")
	}
	if got.MatchPercentage != "50%" {
		t.Errorf("MatchPercentage = %q, want %q", got.MatchPercentage, "50%")
	}
}

func TestParseAnalysis_Empty(t *testing.T) {
	got := parseAnalysis("")

	want := models.AnalysisResponse{}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("parseAnalysis(\"\") = %+v, want zero value %+v", got, want)
	}
}
