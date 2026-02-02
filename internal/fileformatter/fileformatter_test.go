package fileformatter

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestOutputAsJSON(t *testing.T) {
	var buf bytes.Buffer

	data := map[string]any{
		"title": "Test",
		"count": 2,
	}

	err := OutputAsJSON(&buf, data)
	if err != nil {
		t.Fatal(err)
	}

	var decoded map[string]any
	err = json.Unmarshal(buf.Bytes(), &decoded)
	if err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}

	if decoded["title"] != "Test" {
		t.Fatalf("expected title 'Test', got %v", decoded["title"])
	}

	if int(decoded["count"].(float64)) != 2 {
		t.Fatalf("expected count 2, got %v", decoded["count"])
	}
}

func testOutputAsCSV(t *testing.T) {
	var buf bytes.Buffer

	records := []string{"one", "two"}

	err := OutputAsCSV(&buf, records)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	out := buf.String()
	if !strings.Contains(out, "one") || !strings.Contains(out, "two") {
		t.Fatalf("unexpected csv output: %s", out)
	}
}
