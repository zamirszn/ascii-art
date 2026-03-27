package ascii

import (
	"strings"
	"testing"
)

// TestLoadBanner checks that the banner file loads without error
// and contains enough characters (at least 95: from space to ~)
func TestLoadBanner(t *testing.T) {
	banner, err := LoadBanner("standard")
	if err != nil {
		t.Fatalf("Expected no error loading banner, got: %v", err)
	}

	// ASCII printable characters go from 32 (space) to 126 (~) = 95 characters
	if len(banner) < 95 {
		t.Errorf("Expected at least 95 characters in banner, got %d", len(banner))
	}

	// Each character should have exactly 8 lines
	for i, char := range banner {
		if len(char) != 8 {
			t.Errorf("Character at index %d has %d lines, expected 8", i, len(char))
		}
	}
}

// TestRenderEmpty checks that an empty string produces no output
func TestRenderEmpty(t *testing.T) {
	result, err := Render("", "standard")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if result != "" {
		t.Errorf("Expected empty output for empty input, got: %q", result)
	}
}

// TestRenderNewlineOnly checks that "\n" alone produces a single newline
func TestRenderNewlineOnly(t *testing.T) {
	result, err := Render("\\n", "standard")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if result != "\n" {
		t.Errorf("Expected single newline for input '\\n', got: %q", result)
	}
}

// TestRenderHello checks that "Hello" renders 8 lines of output
func TestRenderHello(t *testing.T) {
	result, err := Render("Hello", "standard")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	outputLines := strings.Split(strings.TrimRight(result, "\n"), "\n")
	if len(outputLines) != 8 {
		t.Errorf("Expected 8 lines for 'Hello', got %d", len(outputLines))
	}
}

// TestRenderNewlineSplit checks that "Hello\nThere" renders 16 lines (8 + 8)
func TestRenderNewlineSplit(t *testing.T) {
	result, err := Render("Hello\\nThere", "standard")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	outputLines := strings.Split(strings.TrimRight(result, "\n"), "\n")
	if len(outputLines) != 16 {
		t.Errorf("Expected 16 lines for 'Hello\\nThere', got %d", len(outputLines))
	}
}

// TestRenderDoubleNewline checks that "Hello\n\nThere" renders 8 + 1 + 8 = 17 lines
func TestRenderDoubleNewline(t *testing.T) {
	result, err := Render("Hello\\n\\nThere", "standard")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	outputLines := strings.Split(strings.TrimRight(result, "\n"), "\n")
	if len(outputLines) != 17 {
		t.Errorf("Expected 17 lines for 'Hello\\n\\nThere', got %d", len(outputLines))
	}
}

// TestRenderSpecialChars checks that special characters don't cause a crash
func TestRenderSpecialChars(t *testing.T) {
	_, err := Render("{Hello There}", "standard")
	if err != nil {
		t.Fatalf("Unexpected error rendering special characters: %v", err)
	}
}