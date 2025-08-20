package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	version   = "dev"
	buildTime = "unknown"
)

func main() {
	// Command line flags
	chunkSize := flag.Int("size", 3000, "Maximum characters per chunk")
	inputFile := flag.String("file", "", "Input file (optional, reads from stdin if not provided)")
	flag.Parse()

	var input io.Reader

	// Determine input source
	if *inputFile != "" {
		file, err := os.Open(*inputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		input = file
	} else {
		input = os.Stdin
	}

	// Read all input
	scanner := bufio.NewScanner(input)
	var text strings.Builder
	for scanner.Scan() {
		text.WriteString(scanner.Text())
		text.WriteString("\n")
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	content := strings.TrimSpace(text.String())
	if len(content) == 0 {
		fmt.Println("No content to split")
		return
	}

	// Split into chunks
	chunks := splitText(content, *chunkSize)

	// Output chunks with headers
	for i, chunk := range chunks {
		fmt.Printf("=== PART %d/%d ===\n", i+1, len(chunks))
		fmt.Printf("Characters: %d\n", len(chunk))
		fmt.Println("---")
		fmt.Println(chunk)
		if i < len(chunks)-1 {
			fmt.Println("\n" + strings.Repeat("=", 50) + "\n")
		}
	}

	fmt.Printf("\n\nSummary: Split into %d parts\n", len(chunks))
}

// splitText splits text into chunks, trying to break at word boundaries
func splitText(text string, maxSize int) []string {
	if len(text) <= maxSize {
		return []string{text}
	}

	var chunks []string
	remaining := text

	for len(remaining) > 0 {
		if len(remaining) <= maxSize {
			chunks = append(chunks, remaining)
			break
		}

		// Find the best break point within maxSize
		breakPoint := maxSize
		chunk := remaining[:breakPoint]

		// Try to break at a sentence boundary (. ! ?)
		if lastSentence := strings.LastIndexAny(chunk, ".!?"); lastSentence != -1 && lastSentence > maxSize/2 {
			breakPoint = lastSentence + 1
		} else if lastParagraph := strings.LastIndex(chunk, "\n\n"); lastParagraph != -1 && lastParagraph > maxSize/3 {
			// Try to break at paragraph boundary
			breakPoint = lastParagraph + 2
		} else if lastNewline := strings.LastIndex(chunk, "\n"); lastNewline != -1 && lastNewline > maxSize/3 {
			// Try to break at line boundary
			breakPoint = lastNewline + 1
		} else if lastSpace := strings.LastIndex(chunk, " "); lastSpace != -1 && lastSpace > maxSize/2 {
			// Break at word boundary
			breakPoint = lastSpace + 1
		}
		// If no good break point found, just cut at maxSize

		chunks = append(chunks, strings.TrimSpace(remaining[:breakPoint]))
		remaining = strings.TrimSpace(remaining[breakPoint:])
	}

	return chunks
}
