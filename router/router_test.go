package router

import (
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestServer(t *testing.T) {
	port := os.Getenv("PORT")
	os.Setenv("PORT", "8081")
	defer os.Setenv("PORT", port)

	server := NewServer()
	go func() {
		err := server.ListenAndServe()
		if err != http.ErrServerClosed {
			t.Error(err)
		}
	}()
	defer server.Close()

	// Wait 50 milliseconds for server to start listening to requests
	time.Sleep(50 * time.Millisecond)

	resp, err := http.Get("http://localhost:8081/exoplanets/read/name")

	assert.NoError(t, err)
	defer resp.Body.Close()

	expectedContentType := "application/json"
	actualContentType := resp.Header.Get("Content-Type")
	assert.Equalf(t, expectedContentType, actualContentType, "handler returned wrong content-type: got %v want %v", actualContentType, expectedContentType)
}
