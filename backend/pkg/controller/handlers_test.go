package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestHelloWorld(t *testing.T) {
	// Build our expected body
	body := map[string]string{
		"message": "Hello World!",
	}

	router := SetupRouter()

	w := performRequest(router, "GET", "/")

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal((w.Body.Bytes()), &response)
	if err != nil {
		fmt.Println(err)
	}

	// Make some assertions on the correctness of the response.
	value, exists := response["message"]
	fmt.Println((w.Body.String()))
	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, body["message"], value)

}
