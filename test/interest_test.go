package test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/gin-gonic/gin"
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
