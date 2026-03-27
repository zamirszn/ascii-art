package ascii

import (
	"fmt"
	"os"
	"strings"
)

// LoadBanner reads a banner file and returns a slice where
// each index corresponds to a character's 8 lines of art.
// Index 0 = Space (ASCII 32), Index 1 = '!' (ASCII 33), etc.
func LoadBanner(bannerName string) ([][]string, error) {
	// Build path to banner file
	filePath := fmt.Sprintf("banners/%s.txt", bannerName)

	// Read the entire file as bytes, then convert to string
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("could not read banner file '%s': %w", filePath, err)
	}

	// Banner files use \n line endings.
	// We split the entire file content by newline to get individual lines.
	lines := strings.Split(string(data), "\n")

	// Each character block = 8 lines of art + 1 blank separator line = 9 lines
	// The file starts with a blank line at the top, so we skip index 0.
	// We'll store each character as a slice of 8 strings.
	var characters [][]string

	// Start from line 1 (skip the very first blank line in the file)
	i := 1
	for i < len(lines) {
		// Collect 8 lines for this character
		charLines := make([]string, 8)
		for row := 0; row < 8; row++ {
			if i+row < len(lines) {
				charLines[row] = lines[i+row]
			} else {
				charLines[row] = ""
			}
		}
		characters = append(characters, charLines)
		i += 9 // Move past 8 art lines + 1 blank separator line
	}

	return characters, nil
}

// Render takes an input string and a banner name, and returns
// the full ASCII art as a single string ready to be printed.
func Render(input string, bannerName string) (string, error) {
	// Load the banner character map
	banner, err := LoadBanner(bannerName)
	if err != nil {
		return "", err
	}

	// Handle the \n escape sequence in the input.
	// os.Args gives us the literal characters \n (backslash + n),
	// NOT a real newline. We split on the literal "\n" string.
	lines := strings.Split(input, "\\n")

	var result strings.Builder

	for lineIndex, line := range lines {
		// An empty line (from "Hello\n\nThere" — the middle segment is "")
		// means we just print a real newline and move on.
		if line == "" {
			// Only print the blank line if it's NOT the very last segment,
			// because a trailing \n means "end the output with a newline".
			if lineIndex < len(lines)-1 {
				result.WriteString("\n")
			}
			continue
		}

		// For each of the 8 rows of character height...
		for row := 0; row < 8; row++ {
			// ...go through every character in this line of input
			for _, ch := range line {
				// Calculate which index in our banner slice this char belongs to.
				// Space = ASCII 32 = index 0
				// '!' = ASCII 33 = index 1
				// 'A' = ASCII 65 = index 33
				index := int(ch) - 32

				// Safety check: only handle printable ASCII (32–126)
				if index < 0 || index >= len(banner) {
					continue
				}

				// Append this character's art for the current row
				result.WriteString(banner[index][row])
			}
			// After writing all characters for this row, add a newline
			result.WriteString("\n")
		}
	}

	return result.String(), nil
}