package proxy

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestIsLocalhost_non(t *testing.T) {
	fn := IsLocalhost()

	data := []string{
		"https://twitter.com",
		"http://example.com:8081",
		"http://subosito.com",
		"http://202.67.40.25",
	}

	for i := range data {
		req, _ := http.NewRequest("GET", data[i], nil)
		assert.False(t, fn(req, nil))
	}
}

func TestIsLocalhost(t *testing.T) {
	fn := IsLocalhost()

	data := []string{
		"http://localhost:4567",
		"http://127.0.0.1:3000",
		"http://lvh.me",
		"http://www.127.0.0.1.xip.io",
	}

	for i := range data {
		req, _ := http.NewRequest("GET", data[i], nil)
		assert.True(t, fn(req, nil))
	}
}
