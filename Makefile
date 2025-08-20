APP_NAME = textsplitter
CLI_NAME = textsplit
TUI_NAME = textsplit-tui

.PHONY: all build-cli build-tui clean install deps test

all: build-cli build-tui

deps:
	@echo "📦 Installing dependencies..."
	go mod tidy
	go mod download

build-cli:
	@echo "🔨 Building CLI version..."
	go build -o $(CLI_NAME) main.go

build-tui:
	@echo "🎨 Building TUI version..."
	go build -o $(TUI_NAME) tui.go

install: all
	@echo "📲 Installing binaries..."
	cp $(CLI_NAME) /usr/local/bin/
	cp $(TUI_NAME) /usr/local/bin/

clean:
	@echo "🧹 Cleaning up..."
	rm -f $(CLI_NAME) $(TUI_NAME)

test:
	@echo "🧪 Running tests..."
	go test ./...

# Development helpers
run-cli:
	@echo "🚀 Running CLI version..."
	go run main.go

run-tui:
	@echo "🎨 Running TUI version..."
	go run tui.go

# Release builds
release:
	@echo "📦 Building release versions..."
	GOOS=darwin GOARCH=amd64 go build -o dist/$(CLI_NAME)-darwin-amd64 main.go
	GOOS=darwin GOARCH=arm64 go build -o dist/$(CLI_NAME)-darwin-arm64 main.go
	GOOS=linux GOARCH=amd64 go build -o dist/$(CLI_NAME)-linux-amd64 main.go
	GOOS=windows GOARCH=amd64 go build -o dist/$(CLI_NAME)-windows-amd64.exe main.go
	GOOS=darwin GOARCH=amd64 go build -o dist/$(TUI_NAME)-darwin-amd64 tui.go
	GOOS=darwin GOARCH=arm64 go build -o dist/$(TUI_NAME)-darwin-arm64 tui.go
	GOOS=linux GOARCH=amd64 go build -o dist/$(TUI_NAME)-linux-amd64 tui.go
	GOOS=windows GOARCH=amd64 go build -o dist/$(TUI_NAME)-windows-amd64.exe tui.go

help:
	@echo "Available commands:"
	@echo "  all        - Build both CLI and TUI versions"
	@echo "  build-cli  - Build CLI version only"
	@echo "  build-tui  - Build TUI version only"
	@echo "  deps       - Install dependencies"
	@echo "  install    - Install binaries to /usr/local/bin"
	@echo "  clean      - Remove built binaries"
	@echo "  test       - Run tests"
	@echo "  run-cli    - Run CLI version directly"
	@echo "  run-tui    - Run TUI version directly"
	@echo "  release    - Build release binaries for multiple platforms"
