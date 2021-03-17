package test

import (
	"bytes"
	"encoding/json"
	"inssa_club_waitlist_backend/cmd/server/errors"
	"inssa_club_waitlist_backend/cmd/server/forms"
	"inssa_club_waitlist_backend/cmd/server/routes"
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

func requestInterestWithDuplicateEmailTest(t *testing.T) {
	var response map[string]interface{}

	form := &forms.AddInterest{
		Email:  "example@example.com",
		UserID: "123",
	}

	w := performRequestWithForm(engine, "POST", "/interest", form)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	responseToMap(w.Body, &response)
	assert.Equal(t, errors.DuplicateEmailError, response["errorType"])
}

func deleteInterest(t *testing.T) {
	form := &forms.AddInterest{
		Email: "example@example.com",
	}

	w := performRequestWithForm(engine, "DELETE", "/interest", form)
	assert.Equal(t, http.StatusOK, w.Code)
}

// test interest

func TestSetupRouter(t *testing.T) {
	engine = gin.New()
	for _, controller := range routes.GetRoutes() {
		engine.Handle(controller.Method, controller.Path, controller.Handler)
	} // setup routes
}

func TestInterest(t *testing.T) {
	t.Run("Test interest feature", func(t *testing.T) {
		t.Run("Request interest without email should fail", requestInterestWithoutEmailTest)
		t.Run("Request interest without proper email should fail", requestInterestWithoutProperEmailTest)
		t.Run("Request with proper email, user id should success", requestInterest)
		t.Run("Request interest with duplicate email should fail", requestInterestWithDuplicateEmailTest)
		t.Run("Request delete an interest should success", deleteInterest)
	})
}
