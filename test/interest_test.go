package test

import (
	"bytes"
	"encoding/json"
	"inssa_club_waitlist_backend/cmd/server/errors"
	"inssa_club_waitlist_backend/cmd/server/forms"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/tj/assert"
)

var engine *gin.Engine

func performRequest(h http.Handler, method string, path string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, body)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, req)
	return w
}

func performRequestWithForm(h http.Handler, method string, path string, form interface{}) *httptest.ResponseRecorder {
	formJSON, _ := json.Marshal(form)
	body := strings.NewReader(string(formJSON))
	return performRequest(h, method, path, body)
}

func responseToMap(body *bytes.Buffer, resultMap *map[string]interface{}) {
	stringBody := body.String()
	bytesBody := []byte(stringBody)
	json.Unmarshal(bytesBody, resultMap)
}

// helper functions for easier request

func requestInterestWithoutEmailTest(t *testing.T) {
	var response map[string]interface{}

	form := &forms.AddInterest{
		UserID: "123",
	}

	w := performRequestWithForm(engine, "POST", "/interest", form)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	responseToMap(w.Body, &response)
	assert.Equal(t, errors.ValidationError, response["errorType"])
}

func requestInterestWithoutProperEmailTest(t *testing.T) {
	var response map[string]interface{}

	form := &forms.AddInterest{
		Email:  "example",
		UserID: "123",
	}

	w := performRequestWithForm(engine, "POST", "/interest", form)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	responseToMap(w.Body, &response)
	assert.Equal(t, errors.ValidationError, response["errorType"])
}

func requestInterest(t *testing.T) {
	form := &forms.AddInterest{
		Email:  "example",
		UserID: "123",
	}

	w := performRequestWithForm(engine, "POST", "/interest", form)
	assert.Equal(t, http.StatusCreated, w.Code)
}
