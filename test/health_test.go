// +build e2e

package test

import (
	"fmt"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestHealthzEndpoint(t *testing.T) {
	fmt.Println("Running E2E test for health check endpoint")

	client := resty.New()
	resp, err := client.R().Get("http://localhost:8080/api/healthz")
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, 200, resp.StatusCode())
}
