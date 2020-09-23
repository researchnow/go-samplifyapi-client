package samplify_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	samplify "github.com/morningconsult/go-samplifyapi-client/lib"
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

		client := samplify.NewClient("", "", "", nil)
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
		"/events",
		"/events/1",
		"/projects/test-report-id/detailedReport",
		"/projects/test-report-id/lineItems/test-lineitem-id/detailedReport",
		"/studyMetadata",
		"/projects/test-prj-id/lineItems/test-lineitem-id/quotaCells/1/pause",
		"/projects/test-prj-id/lineItems/test-lineitem-id/quotaCells/2/launch",
	}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		urls = append(urls, r.URL.String())
	}))

	client := samplify.NewClient("", "", "", nil)
	client.Options.APIBaseURL = ts.URL
	client.Options.AuthURL = ts.URL
	client.Auth = getAuth()

	client.CreateProject(getProjectCriteria())
	client.UpdateProject(&samplify.UpdateProjectCriteria{ExtProjectID: "update-test"})
	client.BuyProject("buy-test", getBuyProjectCriteria())
	client.CloseProject("close-test")
	client.GetAllProjects(nil)
	client.GetProjectBy("test-prj-id")
	client.GetProjectReport("test-report-id")
	client.AddLineItem("test", getLineItemCriteria())
	client.UpdateLineItem("test-prj-id", "test-lineitem-id", getUpdateLineItemCriteria())
	client.UpdateLineItemState("test-prj-id", "test-lineitem-id", samplify.ActionPaused)
	client.GetAllLineItems("test-prj-id", nil)
	client.GetLineItemBy("test-prj-id", "test-lineitem-id")
	client.GetFeasibility("test-prj-id", nil)
	client.GetCountries(nil)
	client.GetAttributes("GB", "en", nil)
	client.GetSurveyTopics(nil)
	client.GetEvents(nil)
	client.GetEventBy("1")
	client.GetDetailedProjectReport("test-report-id")
	client.GetDetailedLineItemReport("test-report-id", "test-lineitem-id")
	client.GetStudyMetadata()
	client.SetQuotaCellStatus("test-prj-id", "test-lineitem-id", "1", "pause")
	client.SetQuotaCellStatus("test-prj-id", "test-lineitem-id", "2", "launch")
	ts.Close()

	if len(urls) != len(tests) {
		t.Errorf("Validation failed on endpoint(s)\n")
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
		{
			expectedURL: "/projects?createdAt=2018/11/01,2019/01/01",
			query:       getQueryOptionsFour(),
		},
		{
			expectedURL: "/projects?startDate=2019-06-12&amp;endDate=2019-06-19&amp;extProjectId=test-project-id",
			query:       getQueryOptionsInvoicesSummary(),
		},
	}

	for _, tt := range tests {

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			url = r.URL.String()
		}))

		client := samplify.NewClient("", "", "", nil)
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
			&samplify.Filter{Field: samplify.QueryFieldTitle, Value: samplify.FilterValue{
				Value: "Samplify Client Test"}},
			&samplify.Filter{Field: samplify.QueryFieldState, Value: samplify.FilterValue{Value: samplify.StateProvisioned}},
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
			&samplify.Filter{Field: samplify.QueryFieldTitle, Value: samplify.FilterValue{Value: "Samplify Client Test"}},
			&samplify.Filter{Field: samplify.QueryFieldState, Value: samplify.FilterValue{Value: samplify.StateProvisioned}},
		},
		SortBy: []*samplify.Sort{
			&samplify.Sort{Field: samplify.QueryFieldCreatedAt, Direction: samplify.SortDirectionAsc},
			&samplify.Sort{Field: samplify.QueryFieldExtProjectID, Direction: samplify.SortDirectionDesc},
		},
	}
}

func getQueryOptionsFour() *samplify.QueryOptions {
	fromdate := time.Date(2018, time.November, 1, 0, 0, 0, 0, time.UTC)
	todate := time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)
	value := samplify.DateFilterValue{
		From: &fromdate,
		To:   &todate,
	}
	return &samplify.QueryOptions{
		FilterBy: []*samplify.Filter{
			&samplify.Filter{Field: samplify.QueryFieldCreatedAt, Value: value},
		},
	}
}

/* Test query for field of: `QueryFieldStartDate` and `QueryFieldEndDate`*/
func getQueryOptionsInvoicesSummary() *samplify.QueryOptions {
	f1 := samplify.Filter{
		Field: samplify.QueryFieldStartDate,
		Value: samplify.FilterValue{Value: "2019-06-12"},
	}
	f2 := samplify.Filter{
		Field: samplify.QueryFieldEndDate,
		Value: samplify.FilterValue{Value: "2019-06-19"},
	}

	filters := []*samplify.Filter{
		&f1, &f2,
	}

	projectID := "test-project-id"

	option := samplify.QueryOptions{
		FilterBy:     filters,
		ExtProjectId: &projectID,
	}

	return &option
}

func getAuth() samplify.TokenResponse {
	now := time.Now()
	return samplify.TokenResponse{
		AccessToken: "test",
		Acquired:    &now,
		ExpiresIn:   1800,
	}
}

func getProjectCriteria() *samplify.CreateProjectCriteria {
	return &samplify.CreateProjectCriteria{
		ExtProjectID:       "project001",
		Title:              "Test Survey",
		NotificationEmails: []string{"api-test@researchnow.com"},
		Devices:            []samplify.DeviceType{samplify.DeviceTypeMobile, samplify.DeviceTypeDesktop},
		Category:           &samplify.Category{SurveyTopic: []string{"AUTOMOTIVE", "BUSINESS"}},
		LineItems:          []*samplify.CreateLineItemCriteria{getLineItemCriteria()},
	}
}

func getLineItemCriteria() *samplify.CreateLineItemCriteria {
	surveyURL := "www.mysurvey.com/live/survey?pid=<#DubKnowledge[1500/Entity id]>&k2=<#Project[Secure Key 2]>&psid=<#IdParameter[Value]>"
	surveyTestURL := "www.mysurvey.com/test/survey"
	percs := []float64{30.0, 70.0}

	surveyTestingNotesVal := "survey-testing-notes"

	return &samplify.CreateLineItemCriteria{
		ExtLineItemID:       "lineItem001",
		Title:               "US College",
		CountryISOCode:      "US",
		LanguageISOCode:     "en",
		SurveyURL:           &surveyURL,
		SurveyTestURL:       &surveyTestURL,
		IndicativeIncidence: 20.0,
		DaysInField:         20,
		LengthOfInterview:   10,
		SurveyTestingNotes:  &surveyTestingNotesVal,
		QuotaPlan: &samplify.QuotaPlan{
			Filters: []*samplify.QuotaFilters{
				&samplify.QuotaFilters{AttributeID: "4091", Options: []string{"3", "4"}},
			},
			QuotaGroups: []*samplify.QuotaGroup{
				&samplify.QuotaGroup{
					QuotaCells: []*samplify.QuotaCell{
						&samplify.QuotaCell{
							QuotaNodes: []*samplify.QuotaNode{
								&samplify.QuotaNode{AttributeID: "11", Options: []string{"1"}},
							},
							Perc: &percs[0],
						},
						&samplify.QuotaCell{
							QuotaNodes: []*samplify.QuotaNode{
								&samplify.QuotaNode{AttributeID: "11", Options: []string{"2"}},
							},
							Perc: &percs[1],
						},
					},
				},
			},
		},
	}
}

func getUpdateLineItemCriteria() *samplify.UpdateLineItemCriteria {
	var dif = int64(10)
	return &samplify.UpdateLineItemCriteria{DaysInField: &dif}
}

func getBuyProjectCriteria() []*samplify.BuyProjectCriteria {
	return []*samplify.BuyProjectCriteria{
		&samplify.BuyProjectCriteria{
			ExtLineItemID: "lineItem001",
			SurveyURL:     "www.mysurvey.com/live/survey?pid=<#DubKnowledge[1500/Entity id]>&k2=<#Project[Secure Key 2]>&psid=<#IdParameter[Value]>",
			SurveyTestURL: "www.mysurvey.com/test/survey",
		},
	}
}
