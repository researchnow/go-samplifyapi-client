package samplify

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	prodAuthBaseURL = "https://api.researchnow.com/auth/v1"
	prodAPIBaseURL  = "https://api.researchnow.com/sample/v1"
	uatAuthBaseURL  = "https://api.uat.pe.researchnow.com/auth/v1"
	uatAPIBaseURL   = "https://api.uat.pe.researchnow.com/sample/v1"
)

// ClientOptions ...
type ClientOptions struct {
	APIBaseURL string
	AuthURL    string
}

// Client is used to make API requests to the Samplify API.
type Client struct {
	Credentials TokenRequest
	Auth        TokenResponse
	Options     ClientOptions
}

// CreateProject ...
func (c *Client) CreateProject(project *ProjectCriteria) (*ProjectResponse, error) {
	res := &ProjectResponse{}
	err := c.requestAndParseResponse("POST", "/projects", project, res)
	return res, err
}

// UpdateProject ...
func (c *Client) UpdateProject(project *ProjectCriteria) (*ProjectResponse, error) {
	res := &ProjectResponse{}
	path := fmt.Sprintf("/projects/%s", project.ExtProjectID)
	err := c.requestAndParseResponse("POST", path, project, res)
	return res, err
}

// BuyProject ...
func (c *Client) BuyProject(extProjectID string, buy []*BuyProjectCriteria) (*BuyProjectResponse, error) {
	res := &BuyProjectResponse{}
	path := fmt.Sprintf("/projects/%s/buy", extProjectID)
	err := c.requestAndParseResponse("POST", path, buy, res)
	return res, err
}

// CloseProject ...
func (c *Client) CloseProject(extProjectID string) (*CloseProjectResponse, error) {
	res := &CloseProjectResponse{}
	path := fmt.Sprintf("/projects/%s/close", extProjectID)
	err := c.requestAndParseResponse("POST", path, nil, res)
	return res, err
}

// GetAllProjects ...
func (c *Client) GetAllProjects(options *QueryOptions) (*GetAllProjectsResponse, error) {
	res := &GetAllProjectsResponse{}
	query := query2String(options)
	path := fmt.Sprintf("/projects%s", query)
	err := c.requestAndParseResponse("GET", path, nil, res)
	return res, err
}

// GetProjectBy returns project by id
func (c *Client) GetProjectBy(extProjectID string) (*ProjectResponse, error) {
	res := &ProjectResponse{}
	path := fmt.Sprintf("/projects/%s", extProjectID)
	err := c.requestAndParseResponse("GET", path, nil, res)
	return res, err
}

// GetProjectReport returns a project's report based on observed data from actual panelists.
func (c *Client) GetProjectReport(extProjectID string) (*ProjectReportResponse, error) {
	res := &ProjectReportResponse{}
	path := fmt.Sprintf("/projects/%s/report", extProjectID)
	err := c.requestAndParseResponse("GET", path, nil, res)
	return res, err
}

// AddLineItem ...
func (c *Client) AddLineItem(extProjectID string, lineItem *LineItemCriteria) (*LineItemResponse, error) {
	res := &LineItemResponse{}
	path := fmt.Sprintf("/projects/%s/lineItems", extProjectID)
	err := c.requestAndParseResponse("POST", path, lineItem, res)
	return res, err
}

// UpdateLineItem ...
func (c *Client) UpdateLineItem(extProjectID, extLineItemID string, lineItem *LineItemCriteria) (*LineItemResponse, error) {
	res := &LineItemResponse{}
	path := fmt.Sprintf("/projects/%s/lineItems/%s", extProjectID, extLineItemID)
	err := c.requestAndParseResponse("POST", path, lineItem, res)
	return res, err
}

// UpdateLineItemState ... Changes the state of the line item based on provided action.
func (c *Client) UpdateLineItemState(extProjectID, extLineItemID string, action Action) (
	*UpdateLineItemStateResponse, error) {

	res := &UpdateLineItemStateResponse{}
	path := fmt.Sprintf("/projects/%s/lineItems/%s/%s", extProjectID, extLineItemID, action)
	err := c.requestAndParseResponse("POST", path, nil, res)
	return res, err
}

// GetAllLineItems ...
func (c *Client) GetAllLineItems(extProjectID string, options *QueryOptions) (*GetAllLineItemsResponse, error) {
	res := &GetAllLineItemsResponse{}
	path := fmt.Sprintf("/projects/%s/lineItems%s", extProjectID, query2String(options))
	err := c.requestAndParseResponse("GET", path, nil, res)
	return res, err
}

// GetLineItemBy ...
func (c *Client) GetLineItemBy(extProjectID, extLineItemID string) (*LineItemResponse, error) {
	res := &LineItemResponse{}
	path := fmt.Sprintf("/projects/%s/lineItems/%s", extProjectID, extLineItemID)
	err := c.requestAndParseResponse("GET", path, nil, res)
	return res, err
}

// GetFeasibility ... Returns the feasibility for all the line items of the requested project. Takes 20 - 120
// seconds to execute. Check the `GetFeasibilityResponse.Feasibility.Status` field value to see if it is
// FeasibilityStatusReady ("READY") or FeasibilityStatusProcessing ("PROCESSING")
// If GetFeasibilityResponse.Feasibility.Status == FeasibilityStatusProcessing, call this function again in 2 mins.
func (c *Client) GetFeasibility(extProjectID string, options *QueryOptions) (*GetFeasibilityResponse, error) {
	res := &GetFeasibilityResponse{}
	path := fmt.Sprintf("/projects/%s/feasibility%s", extProjectID, query2String(options))
	err := c.requestAndParseResponse("GET", path, nil, res)
	return res, err
}

// GetCountries ... Get the list of supported countries and languages in each country.
func (c *Client) GetCountries(options *QueryOptions) (*GetCountriesResponse, error) {
	res := &GetCountriesResponse{}
	path := fmt.Sprintf("/countries%s", query2String(options))
	err := c.requestAndParseResponse("GET", path, nil, res)
	return res, err
}

// GetAttributes ... Get the list of supported attributes for a country and language. This data is required to build up the Quota Plan.
func (c *Client) GetAttributes(countryCode, languageCode string, options *QueryOptions) (*GetAttributesResponse, error) {
	res := &GetAttributesResponse{}
	path := fmt.Sprintf("/attributes/%s/%s%s", countryCode, languageCode, query2String(options))
	err := c.requestAndParseResponse("GET", path, nil, res)
	return res, err
}

// GetSurveyTopics ... Get the list of supported Survey Topics for a project. This data is required to setup a project.
func (c *Client) GetSurveyTopics(options *QueryOptions) (*GetSurveyTopicsResponse, error) {
	res := &GetSurveyTopicsResponse{}
	path := fmt.Sprintf("/categories/surveyTopics%s", query2String(options))
	err := c.requestAndParseResponse("GET", path, nil, res)
	return res, err
}

// RefreshToken ...
func (c *Client) RefreshToken() error {
	if c.Auth.AccessTokenExpired() {
		auth, err := c.GetAuth()
		if err != nil {
			return err
		}
		c.Auth = auth
		return nil
	}
	t := time.Now()
	req := struct {
		ClientID     string `json:"clientId"`
		RefreshToken string `json:"refreshToken"`
	}{
		ClientID:     c.Credentials.ClientID,
		RefreshToken: c.Auth.RefreshToken,
	}
	ar, err := sendRequest(c.Options.AuthURL, "POST", "/token/refresh", "", req)
	if err != nil {
		return err
	}
	err = json.Unmarshal(ar.Body, &c.Auth)
	if err != nil {
		return err
	}
	c.Auth.Acquired = &t
	return nil
}

// Logout ...
func (c *Client) Logout() error {
	if c.Auth.AccessTokenExpired() {
		return nil
	}
	req := struct {
		ClientID     string `json:"clientId"`
		RefreshToken string `json:"refreshToken"`
		AccessToken  string `json:"accessToken"`
	}{
		ClientID:     c.Credentials.ClientID,
		RefreshToken: c.Auth.RefreshToken,
		AccessToken:  c.Auth.AccessToken,
	}
	_, err := sendRequest(c.Options.AuthURL, "POST", "/logout", "", req)
	return err
}

// GetAuth ...
func (c *Client) GetAuth() (TokenResponse, error) {
	err := c.requestAndParseToken()
	if err != nil {
		return TokenResponse{}, err
	}
	return c.Auth, err
}

func (c *Client) requestAndParseResponse(method, url string, body interface{}, resObj interface{}) error {
	ar, err := c.request(method, url, body)
	if err != nil {
		if ar != nil {
			json.Unmarshal(ar.Body, &resObj)
		}
		return err
	}

	err = json.Unmarshal(ar.Body, &resObj)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) request(method, url string, body interface{}) (*APIResponse, error) {
	if c.Auth.AccessTokenExpired() {
		err := c.requestAndParseToken()
		if err != nil {
			return nil, err
		}
	}
	ar, err := sendRequest(c.Options.APIBaseURL, method, url, c.Auth.AccessToken, body)
	errResp, ok := err.(*ErrorResponse)
	if ok && errResp.HTTPCode == http.StatusUnauthorized {
		err := c.requestAndParseToken()
		if err != nil {
			return nil, err
		}
		return sendRequest(c.Options.APIBaseURL, method, url, c.Auth.AccessToken, body)
	}
	return ar, err
}

func (c *Client) requestAndParseToken() error {
	log.Printf("Acquiring access token for %v", c.Credentials.ClientID)
	t := time.Now()
	ar, err := sendRequest(c.Options.AuthURL, "POST", "/token/password", "", c.Credentials)
	if err != nil {
		return err
	}
	err = json.Unmarshal(ar.Body, &c.Auth)
	if err != nil {
		return err
	}
	c.Auth.Acquired = &t
	return nil
}

// NewClient returns an API client.
// Uses environment variable `env` = {uat|prod} to select host. If none provided, "uat" is used.
func NewClient(clientID, username, passsword string) *Client {
	options := &ClientOptions{APIBaseURL: uatAPIBaseURL, AuthURL: uatAuthBaseURL}
	if isProdEnv() {
		options = &ClientOptions{APIBaseURL: prodAPIBaseURL, AuthURL: prodAuthBaseURL}
	}
	log.Printf("using env: %s\n", getEnvironment())
	return &Client{
		Credentials: TokenRequest{
			ClientID: clientID,
			Username: username,
			Password: passsword,
		},
		Options: *options,
	}
}
