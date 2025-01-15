package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func testRequest(t *testing.T, ts *httptest.Server, method, path string) (*http.Response, string) {
	req, err := http.NewRequest(method, ts.URL+path, nil)
	require.NoError(t, err)
	resp, err := ts.Client().Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	return resp, string(respBody)
}

func TestRouter(t *testing.T) {
	ts := httptest.NewServer(CarRouter())
	defer ts.Close()

	var testTable = []struct {
		url    string
		want   string
		status int
	}{
		// success response - 200 OK
		{"/cars/renault/Logan", "Renault Logan", http.StatusOK},
		{"/cars/audi/a4", "Audi A4", http.StatusOK},
		{"/cars/BMW/M5", "BMW M5", http.StatusOK},
		{"/cars/bmw/X6", "BMW X6", http.StatusOK},
		{"/cars/Vw/Passat", "VW Passat", http.StatusOK},
		// failure response - 404 Not Found
		{"/cars/audi/a6", "unknown model: audi a6\n", http.StatusNotFound},
	}
	for _, test := range testTable {
		resp, getResp := testRequest(t, ts, http.MethodGet, test.url)
		assert.Equal(t, test.status, resp.StatusCode)
		assert.Equal(t, test.want, getResp)
	}
}
