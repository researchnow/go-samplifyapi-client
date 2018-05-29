package samplify_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	samplify "github.com/researchnow/go-samplifyapi-client/lib"
)

func TestAuth(t *testing.T) {
	var auth string
	tests := []struct {
		accessToken  string
		expectedAuth string
	}{
		{
			accessToken:  "test-token",
			expectedAuth: "Bearer test-token",
		},
	}

	for _, tt := range tests {

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			auth = r.Header.Get("Authorization")
		}))

		client := samplify.NewClient("", "", "")
		client.Options.APIBaseURL = ts.URL
		client.Options.AuthURL = ts.URL
		client.Auth = getAuth()
		client.Auth.AccessToken = tt.accessToken
		client.GetAllProjects(nil)
		ts.Close()
		if auth != tt.expectedAuth {
			t.FailNow()
		}
	}
}

func TestClientFunctions(t *testing.T) {
	var urls []string
	tests := []string{
		"/projects",
		"/projects/update-test",
		"/projects/buy-test/buy",
		"/projects/close-test/close",
		"/projects",
		"/projects/test-prj-id",
		"/projects/test-report-id/report",
		"/projects/test/lineItems",
		"/projects/test-prj-id/lineItems/test-lineitem-id",
		"/projects/test-prj-id/lineItems/test-lineitem-id/pause",
		"/projects/test-prj-id/lineItems",
		"/projects/test-prj-id/lineItems/test-lineitem-id",
		"/projects/test-prj-id/feasibility",
		"/countries",
		"/attributes/GB/en",
		"/categories/surveyTopics",
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		urls = append(urls, r.URL.String())
	}))

	client := samplify.NewClient("", "", "")
	client.Options.APIBaseURL = ts.URL
	client.Options.AuthURL = ts.URL
	client.Auth = getAuth()

	client.CreateProject(nil)
	client.UpdateProject(&samplify.CreateUpdateProjectCriteria{ExtProjectID: "update-test"})
	client.BuyProject("buy-test", nil)
	client.CloseProject("close-test")
	client.GetAllProjects(nil)
	client.GetProjectBy("test-prj-id")
	client.GetProjectReport("test-report-id")
	client.AddLineItem("test", nil)
	client.UpdateLineItem("test-prj-id", "test-lineitem-id", nil)
	client.UpdateLineItemState("test-prj-id", "test-lineitem-id", samplify.ActionPaused)
	client.GetAllLineItems("test-prj-id", nil)
	client.GetLineItemBy("test-prj-id", "test-lineitem-id")
	client.GetFeasibility("test-prj-id", nil)
	client.GetCountries(nil)
	client.GetAttributes("GB", "en", nil)
	client.GetSurveyTopics(nil)
	ts.Close()

	if len(urls) != len(tests) {
		t.FailNow()
	}
	for i, tt := range tests {
		if urls[i] != tt {
			t.Errorf("Expected API URL: %s\n Instead, got: %s\n", tt, urls[i])
			t.FailNow()
		}
	}
}

func TestQueryString(t *testing.T) {
	url := ""
	tests := []struct {
		expectedURL string
		query       *samplify.QueryOptions
	}{
		{
			expectedURL: "/projects?title=Samplify+Client+Test&amp;state=PROVISIONED",
			query:       getQueryOptionsOne(),
		},
		{
			expectedURL: "/projects?sort=createdAt:asc,extProjectId:desc",
			query:       getQueryOptionsTwo(),
		},
		{
			expectedURL: "/projects?title=Samplify+Client+Test&amp;state=PROVISIONED&amp;sort=createdAt:asc,extProjectId:desc",
			query:       getQueryOptionsThree(),
		},
	}
	for _, tt := range tests {

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			url = r.URL.String()
		}))

		client := samplify.NewClient("", "", "")
		client.Options.APIBaseURL = ts.URL
		client.Options.AuthURL = ts.URL
		client.Auth = getAuth()
		client.GetAllProjects(tt.query)
		ts.Close()
		if url != tt.expectedURL {
			t.FailNow()
		}
	}
}

func getQueryOptionsOne() *samplify.QueryOptions {
	return &samplify.QueryOptions{
		FilterBy: []*samplify.Filter{
			&samplify.Filter{Field: samplify.QueryFieldTitle, Value: "Samplify Client Test"},
			&samplify.Filter{Field: samplify.QueryFieldState, Value: samplify.StateProvisioned},
		},
	}
}

func getQueryOptionsTwo() *samplify.QueryOptions {
	return &samplify.QueryOptions{
		SortBy: []*samplify.Sort{
			&samplify.Sort{Field: samplify.QueryFieldCreatedAt, Direction: samplify.SortDirectionAsc},
			&samplify.Sort{Field: samplify.QueryFieldExtProjectID, Direction: samplify.SortDirectionDesc},
		},
	}
}

func getQueryOptionsThree() *samplify.QueryOptions {
	return &samplify.QueryOptions{
		FilterBy: []*samplify.Filter{
			&samplify.Filter{Field: samplify.QueryFieldTitle, Value: "Samplify Client Test"},
			&samplify.Filter{Field: samplify.QueryFieldState, Value: samplify.StateProvisioned},
		},
		SortBy: []*samplify.Sort{
			&samplify.Sort{Field: samplify.QueryFieldCreatedAt, Direction: samplify.SortDirectionAsc},
			&samplify.Sort{Field: samplify.QueryFieldExtProjectID, Direction: samplify.SortDirectionDesc},
		},
	}
}

func getAuth() samplify.TokenResponse {
	now := time.Now()
	return samplify.TokenResponse{
		AccessToken: "test",
		Acquired:    &now,
		ExpiresIn:   1800,
	}
}
