# TextSplitter

A web application for splitting large text into chunks suitable for Microsoft Copilot's 8,000 character limit.

## Problem

Microsoft Copilot has an 8,000 character limit per prompt. When you have large documents that exceed this limit, you need to split them into smaller chunks while maintaining readability and context.

## Solution

TextSplitter is a simple web application that intelligently splits text by trying to break at:
1. Sentence boundaries (. ! ?)
2. Paragraph boundaries (\n\n)
3. Line boundaries (\n)
4. Word boundaries (spaces)

This ensures your text chunks remain coherent and readable for AI processing.

## Installation

```bash
# Clone the repository
git clone https://github.com/nichecode/textsplitter.git
cd textsplitter

# One-command setup
task setup

# Or build manually
go build -o textsplitter ./cmd/textsplitter
```

## Usage

### Quick Start

```bash
# Run the application (automatically opens browser)
task run

# Or run the binary directly
./textsplitter
```

### Options

```bash
# Custom port
./textsplitter --port 8081

# Don't automatically open browser
./textsplitter --open=false

# Verbose output
./textsplitter -v

# Show version
./textsplitter --version
```

### Perfect Copilot Workflow

1. **Run TextSplitter**: `task run` or `./textsplitter`
2. **Paste your large document** in the text area
3. **Adjust chunk size** if needed (default 3000 is safe for 8000 limit)
4. **Click "Split Text"**
5. **Copy each numbered chunk** using the individual copy buttons
6. **Paste into Copilot** with context like:
   - "This is part 1/3 of a document. Please process and wait for remaining parts."
   - "This is part 2/3, continuing from the previous part."
   - "This is the final part 3/3. Please analyze the complete document."

## Features

- 🌐 **Self-contained web application** - No external dependencies
- 📱 **Responsive design** - Works on desktop and mobile
- ✂️ **Intelligent text splitting** - Preserves readability and context
- 📋 **Individual copy buttons** - Easy workflow for each chunk
- ⚡ **Real-time feedback** - Character counts and chunk estimates
- 🚀 **Auto-opens browser** - Ready to use immediately
- 🔧 **Configurable** - Custom ports and options

## Installation Options

### Global Installation

```bash
# Install to system PATH (requires sudo)
task install

# Or install to ~/bin (no sudo required)
task install-local

# Then run from anywhere
textsplitter
```

### Development

```bash
# See all available tasks
task --list

# Essential commands
task setup          # First time setup
task run            # Start the application
task build          # Build the binary
task clean          # Clean up build artifacts
task release        # Build for multiple platforms
```

## Building from Source

Requires Go 1.19 or later:

```bash
go mod tidy
go build -o textsplitter ./cmd/textsplitter
```

## Project Structure

```
textsplitter/
├── cmd/
│   └── textsplitter/
│       ├── main.go           # Main application
│       └── web/
│           └── index.html    # Embedded web interface
├── go.mod
├── Taskfile.yml
└── README.md
```

## How It Works

TextSplitter is a single Go binary with the HTML interface embedded using Go's `embed` package. When you run it:

1. Starts an HTTP server on localhost
2. Serves the embedded HTML interface
3. Automatically opens your default browser
4. Provides a clean interface for splitting and copying text chunks

Perfect for working with AI tools that have character limitations!

## Dependencies

- **None!** - Pure Go standard library
- **Self-contained binary** - HTML interface embedded
- **Cross-platform** - Works on macOS, Linux, and Windows

## License

MIT License - feel free to use and modify as needed.

## Contributing

Pull requests welcome! Ideas for enhancements:
- Save/load functionality
- Different output formats
- Batch processing
- Custom splitting rules