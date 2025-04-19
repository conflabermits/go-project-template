package test

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/chris/go-project-template/internal/common"
)

func TestHello(t *testing.T) {
	// Capture the output of the Hello function
	output := captureOutput(common.Hello)

	// Check if the output matches "Hello world"
	if !strings.Contains(output, "Hello world") {
		t.Errorf("Expected 'Hello world', but got '%s'", output)
	}
}

// Helper function to capture output from a function
func captureOutput(f func()) string {
	old := os.Stdout // keep backup of the real stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old // restoring the real stdout

	var buf bytes.Buffer
	io.Copy(&buf, r)

	return buf.String()
}
