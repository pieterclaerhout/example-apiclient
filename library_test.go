package exampleapiclient_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	exampleapiclient "github.com/pieterclaerhout/example-apiclient"
	"github.com/stretchr/testify/assert"
)

func TestValid(t *testing.T) {

	input := "expected"

	s := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.Write([]byte(strings.ToUpper(input)))
		}),
	)
	defer s.Close()

	apiClient := exampleapiclient.NewAPIClient(s.URL, 5*time.Second)

	actual, err := apiClient.ToUpper(input)
	assert.NoError(t, err, "error")
	assert.Equal(t, strings.ToUpper(input), actual, "actual")

}

func TestInvalidURL(t *testing.T) {

	apiClient := exampleapiclient.NewAPIClient("ht&@-tp://:aa", 5*time.Second)

	actual, err := apiClient.ToUpper("hello")
	assert.Error(t, err)
	assert.Empty(t, actual)

}

func TestTimeout(t *testing.T) {

	s := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			time.Sleep(50 * time.Millisecond)
			w.Write([]byte("actual"))
		}),
	)
	defer s.Close()

	apiClient := exampleapiclient.NewAPIClient(s.URL, 25*time.Millisecond)

	actual, err := apiClient.ToUpper("hello")
	assert.Error(t, err)
	assert.Empty(t, actual)

}

func TestBodyReadError(t *testing.T) {

	s := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.Header().Set("Content-Length", "1")
		}),
	)
	defer s.Close()

	apiClient := exampleapiclient.NewAPIClient(s.URL, 25*time.Millisecond)

	actual, err := apiClient.ToUpper("hello")
	assert.Error(t, err)
	assert.Empty(t, actual)

}
