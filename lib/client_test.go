package samplify_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/rn/samplify/lib"
)

func TestAuth(t *testing.T) {
	var auth string
	now := time.Now()
	tests := []struct {
		clientID     string
		username     string
		password     string
		AccessToken  string
		expectedAuth string
	}{
		{
			username:     "test",
			password:     "test",
			AccessToken:  "test-token",
			expectedAuth: "Bearer test-token",
		},
	}

	for _, tt := range tests {

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			auth = r.Header.Get("Authorization")
		}))

		client := samplify.NewClient(
			tt.clientID,
			tt.username,
			tt.password,
			&samplify.ClientOptions{AuthURL: ts.URL, APIBaseURL: ts.URL},
		)
		client.Auth = samplify.TokenResponse{
			AccessToken: tt.AccessToken,
			Acquired:    &now,
			ExpiresIn:   1800,
		}
		client.GetAllProjects()
		ts.Close()
		if auth != tt.expectedAuth {
			t.FailNow()
		}
	}
}
