package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"greenlight.github.com/ScreamingArrow/internal/assert"
	"greenlight.github.com/ScreamingArrow/internal/jsonlog"
)

func TestHealthcheck(t *testing.T) {
	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)

	app := &application{
		logger: logger,
	}

	ts := httptest.NewTLSServer(app.routes())
	defer ts.Close()

	rs, err := ts.Client().Get(ts.URL + "/v1/healthcheck")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, rs.StatusCode, http.StatusOK)
	defer rs.Body.Close()

	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, strings.TrimSpace(string(body)), `{"status":"available","system_info":{"environment":"","version":"1.0.0"}}`)
}
