# Text Splitter for Copilot

A Go utility to split large text into chunks suitable for Microsoft Copilot's character limits. Available in both command-line and interactive terminal UI versions.

## Problem

Microsoft Copilot has an 8,000 character limit per prompt. When you have large documents or text that exceed this limit, you need to split them into smaller chunks while maintaining readability and context.

## Solution

This tool intelligently splits text by trying to break at:
1. Sentence boundaries (. ! ?)
2. Paragraph boundaries (\n\n)
3. Line boundaries (\n)
4. Word boundaries (spaces)

This ensures your text chunks remain coherent and readable for AI processing.

## Two Versions Available

### 🎨 Terminal UI Version (Recommended)
Interactive terminal interface for easy text splitting:
- Paste text directly into the interface
- Visual chunk navigation
- Real-time character counting
- Adjustable chunk sizes
- Beautiful, intuitive interface

### ⚡ Command Line Version
Traditional CLI tool for scripting and automation:
- Pipe text or read from files
- Scriptable and automatable
- Perfect for integration with other tools

## Installation

```bash
# Clone the repository
git clone https://github.com/nichecode/textsplitter.git
cd textsplitter

# One-command setup (recommended)
task setup

# Or install dependencies and build manually
go mod tidy
task build

# Or build individually
task build-tui  # Terminal UI version
task build-cli  # Command line version
```

## Usage

### 🎨 Terminal UI Version

```bash
# Run the interactive version
./textsplit-tui

# Or run with custom default chunk size
./textsplit-tui 7500
```

**TUI Controls:**
- **Type/Paste** text in the input area
- **ENTER** to split text into chunks
- **← →** to navigate between chunks
- **↑ ↓** to adjust chunk size
- **TAB** to switch between input and output views
- **R** to reset and start over
- **Q** or **Ctrl+C** to quit

The TUI provides a clean, visual interface where you can:
1. Paste your large text
2. See real-time character counts and chunk estimates
3. Adjust chunk size with arrow keys (↑/↓)
4. Press ENTER to split into chunks
5. **View all chunks simultaneously** in separate, copyable sections
6. **Select and copy any section individually** - perfect for Copilot workflow!

### ⚡ Command Line Version

```bash
# Split a file with default 3000 character chunks
./textsplit -file document.txt

# Split with custom size (leave room under 8000 limit)
./textsplit -size 7500 -file large-document.txt

# Pipe text directly
echo "Your long text here..." | ./textsplit

# Read from clipboard (macOS)
pbpaste | ./textsplit

# Read from clipboard (Windows)
Get-Clipboard | ./textsplit

# Read from clipboard (Linux with xclip)
xclip -selection clipboard -o | ./textsplit
```

**CLI Options:**
- `-size`: Maximum characters per chunk (default: 3000)
- `-file`: Input file path (optional, reads from stdin if not provided)

### CLI Output Format

The CLI outputs clearly formatted chunks:

```
=== PART 1/3 ===
Characters: 2847
---
[Your text content here...]

==================================================

=== PART 2/3 ===
Characters: 2943
---
[Next chunk of text...]

==================================================

=== PART 3/3 ===
Characters: 1205
---
[Final chunk...]


Summary: Split into 3 parts
```

## Workflow: Microsoft Copilot Integration

### Using the TUI (Recommended)
1. Run `./textsplit-tui` or `task run-tui`
2. Paste your large document in the input area
3. Press ENTER to split into chunks
4. **All chunks are displayed simultaneously** in separate sections
5. **Select and copy any chunk** by clicking and dragging or using Cmd+A within a section
6. Paste each chunk into Copilot with context like:
   - "This is part 1/3 of a document. Please process and wait for remaining parts."
   - "This is part 2/3, continuing from the previous part."
   - "This is the final part 3/3. Please analyze the complete document."

**Key Advantage:** No more navigating between chunks! All sections are visible and copyable at once.

### Using the CLI
1. Run your text through: `pbpaste | ./textsplit`
2. Copy each part sequentially
3. Paste into Copilot with appropriate context

## Development

```bash
# See all available tasks
task --list

# Essential commands
task setup      # First time setup
task run-tui    # Run the interactive UI
task run-cli    # Run command line version  
task demo       # Quick demonstration

# Development workflow
task test       # Run tests
task fmt        # Format code
task check      # Run all checks (fmt, lint, test)
task clean      # Clean up build artifacts
task release    # Build for multiple platforms
```



## Building from Source

Requires Go 1.19 or later:

```bash
# Initialize module and install dependencies
go mod tidy

# Build both versions
go build -o textsplit main.go      # CLI version
go build -o textsplit-tui tui.go   # TUI version
```

## Dependencies

- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - TUI framework
- [Lip Gloss](https://github.com/charmbracelet/lipgloss) - Terminal styling

## License

MIT License - feel free to use and modify as needed.

## Contributing

Pull requests welcome! Some ideas for additional features:
- Clipboard integration for automatic copying
- Save chunks to separate files
- Different output formats (JSON, numbered files, etc.)
- Configuration file support
- Web interface version
- Integration with various AI platforms

## Screenshots

The TUI provides a beautiful, interactive experience:
- Clean input area for pasting text
- **Multiple chunk sections displayed simultaneously**
- **Each chunk in its own copyable box** with headers
- Real-time character counts and feedback
- Visual separators between chunks
- Copy instructions and guidance
- Intuitive keyboard controls

**Perfect Copilot Workflow:** Paste text → Press ENTER → Copy each section individually → Paste into Copilot with context!
