# Text Splitter for Copilot

A simple Go utility to split large text into chunks suitable for Microsoft Copilot's character limits.

## Problem

Microsoft Copilot has an 8,000 character limit per prompt. When you have large documents or text that exceed this limit, you need to split them into smaller chunks while maintaining readability and context.

## Solution

This tool intelligently splits text by trying to break at:
1. Sentence boundaries (. ! ?)
2. Paragraph boundaries (\n\n)
3. Line boundaries (\n)
4. Word boundaries (spaces)

This ensures your text chunks remain coherent and readable for AI processing.

## Installation

```bash
# Clone the repository
git clone https://github.com/[your-username]/textsplitter.git
cd textsplitter

# Build the application
go build -o textsplit

# Or install directly with Go
go install .
```

## Usage

### Command Line Options

- `-size`: Maximum characters per chunk (default: 3000)
- `-file`: Input file path (optional, reads from stdin if not provided)

### Examples

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

### Output Format

The tool outputs clearly formatted chunks:

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

## Use Case: Microsoft Copilot Workflow

1. Run your large text through this tool
2. Copy each part sequentially
3. Paste into Copilot with instructions like:
   - "This is part 1/3 of a document. Please process this and wait for the remaining parts."
   - "This is part 2/3, continuing from the previous part."
   - "This is the final part 3/3. Please provide your analysis of the complete document."

## Building from Source

Requires Go 1.16 or later:

```bash
go mod init textsplitter
go build -o textsplit
```

## License

MIT License - feel free to use and modify as needed.

## Contributing

Pull requests welcome! This is a simple utility that could benefit from additional features like:
- Save chunks to separate files
- Different output formats (JSON, numbered files, etc.)
- Configuration file support
- GUI version
