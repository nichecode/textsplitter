package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	version   = "dev"
	buildTime = "unknown"
)

// Styles
var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#7C3AED")).
			MarginBottom(1)

	inputBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#7C3AED")).
			Padding(1).
			Width(80).
			Height(8)

	chunkStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#059669")).
			Padding(1).
			Width(80).
			MarginBottom(1)

	chunkHeaderStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#059669")).
			Background(lipgloss.Color("#F0FDF4")).
			Padding(0, 1).
			MarginBottom(1)

	buttonStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#7C3AED")).
			Foreground(lipgloss.Color("#FFFFFF")).
			Padding(0, 2).
			Margin(0, 1)

	activeButtonStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#A855F7")).
			Foreground(lipgloss.Color("#FFFFFF")).
			Padding(0, 2).
			Margin(0, 1).
			Bold(true)

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#6B7280")).
			MarginTop(1)

	separatorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#D1D5DB")).
			Bold(true)

	instructionStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#059669")).
			Background(lipgloss.Color("#F0FDF4")).
			Padding(0, 1).
			MarginBottom(1)
)

type model struct {
	inputText   string
	chunks      []string
	chunkSize   int
	currentView int // 0 = input, 1 = output
	width       int
	height      int
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "tab":
			if m.currentView == 0 && len(m.chunks) > 0 {
				m.currentView = 1
			} else if m.currentView == 1 {
				m.currentView = 0
			}

		case "enter":
			if m.currentView == 0 {
				// Split the text
				m.chunks = splitText(m.inputText, m.chunkSize)
				if len(m.chunks) > 0 {
					m.currentView = 1
				}
			}

		case "up":
			if m.chunkSize < 8000 {
				m.chunkSize += 500
				if len(m.inputText) > 0 {
					m.chunks = splitText(m.inputText, m.chunkSize)
				}
			}

		case "down":
			if m.chunkSize > 500 {
				m.chunkSize -= 500
				if len(m.inputText) > 0 {
					m.chunks = splitText(m.inputText, m.chunkSize)
				}
			}

		case "r":
			// Reset
			m.inputText = ""
			m.chunks = []string{}
			m.currentView = 0

		case "backspace":
			if m.currentView == 0 && len(m.inputText) > 0 {
				m.inputText = m.inputText[:len(m.inputText)-1]
			}

		default:
			// Add character to input
			if m.currentView == 0 && len(msg.String()) == 1 {
				m.inputText += msg.String()
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	var sections []string

	// Title
	title := titleStyle.Render("📝 Text Splitter for Copilot")
	sections = append(sections, title)

	// Chunk size info
	sizeInfo := fmt.Sprintf("Chunk Size: %d characters (↑/↓ to adjust)", m.chunkSize)
	sections = append(sections, helpStyle.Render(sizeInfo))

	if m.currentView == 0 {
		// Input view
		sections = append(sections, m.renderInputView())
	} else {
		// Output view with all chunks
		sections = append(sections, m.renderAllChunksView())
	}

	// Help text
	help := m.renderHelp()
	sections = append(sections, help)

	return lipgloss.JoinVertical(lipgloss.Left, sections...)
}

func (m model) renderInputView() string {
	var content []string

	// Input box
	inputLabel := "📝 Paste your text here:"
	content = append(content, inputLabel)

	inputContent := m.inputText
	if len(inputContent) == 0 {
		inputContent = "Start typing or paste your text..."
	}

	inputBox := inputBoxStyle.Render(inputContent)
	content = append(content, inputBox)

	// Character count
	charCount := fmt.Sprintf("Characters: %d", len(m.inputText))
	if len(m.inputText) > 0 {
		estimatedChunks := (len(m.inputText) + m.chunkSize - 1) / m.chunkSize
		charCount += fmt.Sprintf(" | Estimated chunks: %d", estimatedChunks)
	}
	content = append(content, helpStyle.Render(charCount))

	// Split button
	var splitButton string
	if len(m.inputText) > 0 {
		splitButton = activeButtonStyle.Render("Press ENTER to Split")
	} else {
		splitButton = buttonStyle.Render("Enter text to split")
	}
	content = append(content, splitButton)

	return lipgloss.JoinVertical(lipgloss.Left, content...)
}

func (m model) renderAllChunksView() string {
	if len(m.chunks) == 0 {
		return "No chunks available"
	}

	var content []string

	// Instructions
	instructions := "💡 Select and copy text from any section below (Cmd+A selects all text in a section)"
	content = append(content, instructionStyle.Render(instructions))

	// Summary header
	summary := fmt.Sprintf("📊 Split into %d chunks | Total: %d characters", len(m.chunks), len(m.inputText))
	content = append(content, chunkHeaderStyle.Render(summary))

	// Render all chunks
	for i, chunk := range m.chunks {
		// Chunk header
		header := fmt.Sprintf("📄 PART %d/%d (%d characters)", i+1, len(m.chunks), len(chunk))
		chunkHeader := chunkHeaderStyle.Render(header)
		content = append(content, chunkHeader)

		// Chunk content in a styled box
		chunkBox := chunkStyle.Render(chunk)
		content = append(content, chunkBox)

		// Add some spacing between chunks (except for the last one)
		if i < len(m.chunks)-1 {
			separator := separatorStyle.Render(strings.Repeat("─", 80))
			content = append(content, separator)
		}
	}

	// Final instructions
	finalInstructions := "💾 Copy each section separately and paste into Copilot with context like: 'This is part X/Y...'"
	content = append(content, instructionStyle.Render(finalInstructions))

	return lipgloss.JoinVertical(lipgloss.Left, content...)
}

func (m model) renderHelp() string {
	var helpLines []string

	if m.currentView == 0 {
		helpLines = []string{
			"📋 Paste text and press ENTER to split",
			"🔄 R to reset | ↑/↓ to adjust chunk size",
			"❌ Q or Ctrl+C to quit",
		}
	} else {
		helpLines = []string{
			"📄 All chunks displayed | Select and copy any section",
			"🔄 TAB back to input | R to reset | ↑/↓ adjust size",
			"❌ Q or Ctrl+C to quit",
		}
	}

	return helpStyle.Render(strings.Join(helpLines, " | "))
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

func main() {
	// Check for command line chunk size
	chunkSize := 3000
	if len(os.Args) > 1 {
		if size, err := strconv.Atoi(os.Args[1]); err == nil && size > 0 {
			chunkSize = size
		}
	}

	m := model{
		chunkSize:   chunkSize,
		currentView: 0,
	}

	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program: %v", err)
		os.Exit(1)
	}
}
