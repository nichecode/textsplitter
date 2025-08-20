package main

import (
	"embed"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os/exec"
	"runtime"
	"time"
)

//go:embed web/index.html
var webFiles embed.FS

var (
	version   = "dev"
	buildTime = "unknown"
)

func main() {
	var (
		port    = flag.Int("port", 8080, "Port to run the web server on")
		open    = flag.Bool("open", true, "Automatically open browser")
		verbose = flag.Bool("v", false, "Verbose output")
		ver     = flag.Bool("version", false, "Show version information")
	)
	flag.Parse()

	if *ver {
		fmt.Printf("TextSplitter %s (built %s)\n", version, buildTime)
		return
	}

	// Find an available port
	actualPort := findAvailablePort(*port)
	
	// Serve the embedded HTML file
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" && r.URL.Path != "/index.html" {
			http.NotFound(w, r)
			return
		}
		
		content, err := webFiles.ReadFile("web/index.html")
		if err != nil {
			http.Error(w, "Could not read index.html", http.StatusInternalServerError)
			return
		}
		
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write(content)
	})

	serverURL := fmt.Sprintf("http://localhost:%d", actualPort)
	
	fmt.Printf("🌐 TextSplitter for Copilot\n")
	fmt.Printf("📱 Starting web interface at %s\n", serverURL)
	fmt.Printf("💡 Perfect for splitting large text for Microsoft Copilot!\n")
	
	if *open {
		fmt.Printf("🚀 Opening browser...\n")
		go func() {
			time.Sleep(500 * time.Millisecond) // Give server time to start
			openBrowser(serverURL)
		}()
	}
	
	fmt.Printf("🛑 Press Ctrl+C to stop\n\n")

	if *verbose {
		log.Printf("Starting server on %s", serverURL)
	}
	
	if err := http.ListenAndServe(fmt.Sprintf(":%d", actualPort), nil); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}

// findAvailablePort finds an available port starting from the given port
func findAvailablePort(startPort int) int {
	for port := startPort; port < startPort+100; port++ {
		listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
		if err == nil {
			listener.Close()
			if port != startPort {
				fmt.Printf("⚠️  Port %d was busy, using port %d instead\n", startPort, port)
			}
			return port
		}
	}
	// If we can't find a port in the range, just return the original and let it fail gracefully
	return startPort
}

// openBrowser opens the default browser to the given URL
func openBrowser(url string) {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		fmt.Printf("Please open %s in your browser\n", url)
		return
	}
	if err != nil {
		fmt.Printf("Could not automatically open browser. Please open %s manually\n", url)
	}
}