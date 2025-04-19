package test

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/chris/go-project-template/internal/common"
	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	// Capture the output of the Hello function
	output := captureOutput(common.Hello)

	// Use testify's assert to check the output
	assert.Contains(t, output, "Hello world", "Output should contain 'Hello world'")
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
