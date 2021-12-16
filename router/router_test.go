package router

import (
	"bytes"
	"encoding/json"
	"github.com/go-playground/assert/v2"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func performRequest(r http.Handler, method, path string, body io.Reader, headers map[string]string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, body)
	w := httptest.NewRecorder()
	for key, value := range headers {
		req.Header.Add(key, value)
	}
	r.ServeHTTP(w, req)
	return w
}

func TestSetupRouter(t *testing.T) {

	value := url.Values{}
	value.Add("username", "quanjw")
	value.Add("email", "quanjw@gmail.com")
	value.Add("password", "123456")
	value.Add("password-again", "123456")

	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded; param=value",
	}

	valueLogin := url.Values{}
	valueLogin.Add("email", "quanjw@gmail.com")
	valueLogin.Add("password", "123456")
	testCases := []struct {
		routerPath string
		httpMethod string
		body       io.Reader
		headers    map[string]string
	}{
		{"/ping", http.MethodGet, nil, nil},
		{"/user/login", http.MethodPost, bytes.NewBufferString(valueLogin.Encode()), headers},
		//{"/user/register", http.MethodPost},
		{"/user/register", http.MethodPost, bytes.NewBufferString(value.Encode()), headers},
	}

	router := SetupRouter()
	for _, tc := range testCases {
		t.Run(tc.routerPath, func(t *testing.T) {
			w := performRequest(router, tc.httpMethod, tc.routerPath, tc.body, tc.headers)
			assert.Equal(t, http.StatusOK, w.Code)

			var dat map[string]interface{}
			if err := json.Unmarshal([]byte(w.Body.String()), &dat); err == nil {
				t.Log(dat)
			}
			assert.Equal(t, "true", dat["success"])
		})
	}

}
