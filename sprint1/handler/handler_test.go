package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStatusHandler(t *testing.T) {
	type wantResponse struct {
		statusCode   int
		responseBody string
		contentType  string
	}
	statusTests := []struct {
		name string
		wantResponse
	}{
		{
			name: "positive test #1",
			wantResponse: wantResponse{
				statusCode:   200,
				responseBody: `{"status":"ok"}`,
				contentType:  "application/json",
			},
		},
	}
	for _, test := range statusTests {
		t.Run(test.name, func(t *testing.T) {
			// Step 1: We've create new request by method NewRequest with parameters {Sending Method, Endpoint, Response Body}
			request := httptest.NewRequest(http.MethodGet, "/status", nil)
			// Step 2: Create new Recorder, it's our Response Writer
			recorder := httptest.NewRecorder()
			// Step 3: Call(invoke) the function(HTTP Handler)
			StatusHandler(recorder, request)
			// Step 4: Initiate new variable with Response Content
			res := recorder.Result()
			// Step 5: Compare expected Status Code with actual Status Code
			assert.Equal(t, test.wantResponse.statusCode, res.StatusCode)

			// Step 6: defer for close Response Body stream
			defer res.Body.Close()
			// Step 7: Get and check the Response Body
			resBody, err := io.ReadAll(res.Body)
			// Step 8: Guarantee the there is no error
			require.NoError(t, err)
			assert.JSONEq(t, test.wantResponse.responseBody, string(resBody))
			assert.Equal(t, test.wantResponse.contentType, res.Header.Get("Content-Type"))
		})
	}
}

func TestUserViewHandler(t *testing.T) {
	type wantResponse struct {
		statusCode  int
		contentType string
		user        User
	}
	testRequests := []struct {
		name    string
		request string
		users   map[string]User
		wantResponse
	}{
		{
			name:    "successful response",
			request: "/users?ID=user3",
			users: map[string]User{
				"user3": {
					ID:        "u3",
					FirstName: "Alem",
					LastName:  "Salem",
				},
			},
			wantResponse: wantResponse{
				statusCode:  200,
				contentType: "application/json",
				user: User{
					ID:        "u3",
					FirstName: "Alem",
					LastName:  "Salem",
				},
			},
		},
	}
	for _, testRequest := range testRequests {
		t.Run(testRequest.name, func(ft *testing.T) {
			request := httptest.NewRequest(http.MethodPost, testRequest.request, nil)
			recorder := httptest.NewRecorder()
			handler := http.HandlerFunc(UserViewHandler(testRequest.users))
			handler(recorder, request)
			result := recorder.Result()

			assert.Equal(ft, testRequest.wantResponse.statusCode, result.StatusCode)
			assert.Equal(ft, testRequest.wantResponse.contentType, result.Header.Get("Content-Type"))

			userResult, err := io.ReadAll(result.Body)
			require.NoError(ft, err)
			err = result.Body.Close()
			require.NoError(ft, err)

			user := User{}

			err = json.Unmarshal(userResult, &user)
			require.NoError(ft, err)

			assert.Equal(ft, testRequest.wantResponse.user, user)
		})
	}
}
