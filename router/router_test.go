package router

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestSetupRouter(t *testing.T) {
	testCases := []struct {
		routerPath string
		httpMethod string
	}{
		{"/ping", http.MethodGet},
		{"/user/login", http.MethodPost},
		{"/user/register", http.MethodPost},
	}

	router := SetupRouter()
	for _, tc := range testCases {
		t.Run(tc.routerPath, func(t *testing.T) {
			w := performRequest(router, tc.httpMethod, tc.routerPath)
			assert.Equal(t, http.StatusOK, w.Code)

			var dat map[string]interface{}
			if err := json.Unmarshal([]byte(w.Body.String()), &dat); err == nil {
				t.Log(dat)
				t.Log(dat["success"])
			}
			assert.Equal(t, "true", dat["success"])
		})
	}
}
