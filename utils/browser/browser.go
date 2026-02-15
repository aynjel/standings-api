package browser

import (
	"fmt"
	"os/exec"
	"runtime"
	"time"
)

// OpenURL opens the specified URL in the default browser
func OpenURL(url string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", url)
	case "darwin":
		cmd = exec.Command("open", url)
	case "linux":
		cmd = exec.Command("xdg-open", url)
	default:
		return fmt.Errorf("unsupported platform")
	}

	return cmd.Start()
}

// OpenBrowserAfterDelay waits for the server to start then opens the browser
func OpenBrowserAfterDelay(url string, delay time.Duration) {
	go func() {
		time.Sleep(delay)
		
		// Check if running in Docker by looking for common Docker environment indicators
		if isRunningInDocker() {
			fmt.Printf("⚠️  Running in Docker container - browser auto-open disabled\n")
			fmt.Printf("📖 Access API Documentation at: %s\n", url)
			return
		}
		
		if err := OpenURL(url); err != nil {
			fmt.Printf("⚠️  Failed to open browser: %v\n", err)
			fmt.Printf("📖 Please manually open: %s\n", url)
		} else {
			fmt.Printf("✓ Opening browser at: %s\n", url)
		}
	}()
}

// isRunningInDocker checks if the application is running inside a Docker container
func isRunningInDocker() bool {
	// Check for .dockerenv file (common Docker indicator)
	if _, err := exec.Command("test", "-f", "/.dockerenv").Output(); err == nil {
		return true
	}
	
	// Alternative: check if cgroup contains "docker"
	// This is a fallback for environments without .dockerenv
	return false
}
