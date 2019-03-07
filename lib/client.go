package samplify

import (
	"encoding/json"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

// ClientOptions to use while creating a new Client
var (
	UATClientOptions = &ClientOptions{
		APIBaseURL: "https://api.uat.pe.researchnow.com/sample/v1",
		AuthURL:    "https://api.uat.pe.researchnow.com/auth/v1",
	}
	ProdClientOptions = &ClientOptions{
		APIBaseURL: "https://api.researchnow.com/sample/v1",
		AuthURL:    "https://api.researchnow.com/auth/v1",
	}
)

// ErrSessionExpired ... Returns if both Access and Refresh tokens are expired
var ErrSessionExpired = errors.New("session expired")

const defaulttimeout = 20

// ClientOptions ...
type ClientOptions struct {
	APIBaseURL string
	AuthURL    string
	Timeout    *int
}

// Client is used to make API requests to the Samplify API.
type Client struct {
	Credentials TokenRequest
	Auth        TokenResponse
	Options     ClientOptions
}

// CreateProject ...
func (c *Client) CreateProject(project *CreateProjectCriteria) (*ProjectResponse, error) {
	err := Validate(project)
	if err != nil {
		return nil, err
	}
	res := &ProjectResponse{}
	err = c.requestAndParseResponse("POST", "/projects", project, res)
	return res, err
}

// UpdateProject ...
func (c *Client) UpdateProject(project *UpdateProjectCriteria) (*ProjectResponse, error) {
	err := Validate(project)
	if err != nil {
		return nil, err
	}
	res := &ProjectResponse{}
	path := fmt.Sprintf("/projects/%s", project.ExtProjectID)
	err = c.requestAndParseResponse("POST", path, project, res)
	return res, err
}

// BuyProject ...
func (c *Client) BuyProject(extProjectID string, buy []*BuyProjectCriteria) (*BuyProjectResponse, error) {
	err := ValidateNotEmpty(extProjectID)
	if err != nil {
		return nil, err
	}
	err = Validate(buy)
	if err != nil {
		return nil, err
	}
	res := &BuyProjectResponse{}
	path := fmt.Sprintf("/projects/%s/buy", extProjectID)
	err = c.requestAndParseResponse("POST", path, buy, res)
	return res, err
}

// CloseProject ...
func (c *Client) CloseProject(extProjectID string) (*CloseProjectResponse, error) {
	err := ValidateNotEmpty(extProjectID)
	if err != nil {
		return nil, err
	}
	res := &CloseProjectResponse{}
	path := fmt.Sprintf("/projects/%s/close", extProjectID)
	err = c.requestAndParseResponse("POST", path, nil, res)
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
	err := ValidateNotEmpty(extProjectID)
	if err != nil {
		return nil, err
	}
	res := &ProjectResponse{}
	path := fmt.Sprintf("/projects/%s", extProjectID)
	err = c.requestAndParseResponse("GET", path, nil, res)
	return res, err
}

// GetProjectReport returns a project's report based on observed data from actual panelists.
func (c *Client) GetProjectReport(extProjectID string) (*ProjectReportResponse, error) {
	err := ValidateNotEmpty(extProjectID)
	if err != nil {
		return nil, err
	}
	res := &ProjectReportResponse{}
	path := fmt.Sprintf("/projects/%s/report", extProjectID)
	err = c.requestAndParseResponse("GET", path, nil, res)
	return res, err
}

// AddLineItem ...
func (c *Client) AddLineItem(extProjectID string, lineItem *CreateLineItemCriteria) (*LineItemResponse, error) {
	err := ValidateNotEmpty(extProjectID)
	if err != nil {
		return nil, err
	}
	err = Validate(lineItem)
	if err != nil {
		return nil, err
	}
	res := &LineItemResponse{}
	path := fmt.Sprintf("/projects/%s/lineItems", extProjectID)
	err = c.requestAndParseResponse("POST", path, lineItem, res)
	return res, err
}

// UpdateLineItem ...
func (c *Client) UpdateLineItem(extProjectID, extLineItemID string,
	lineItem *UpdateLineItemCriteria) (*LineItemResponse, error) {

	err := ValidateNotEmpty(extProjectID, extLineItemID)
	if err != nil {
		return nil, err
	}
	err = Validate(lineItem)
	if err != nil {
		return nil, err
	}
	res := &LineItemResponse{}
	path := fmt.Sprintf("/projects/%s/lineItems/%s", extProjectID, extLineItemID)
	err = c.requestAndParseResponse("POST", path, lineItem, res)
	return res, err
}

// UpdateLineItemState ... Changes the state of the line item based on provided action.
func (c *Client) UpdateLineItemState(extProjectID, extLineItemID string, action Action) (
	*UpdateLineItemStateResponse, error) {
	err := ValidateNotEmpty(extProjectID, extLineItemID)
	if err != nil {
		return nil, err
	}
	err = ValidateAction(action)
	if err != nil {
		return nil, err
	}
	res := &UpdateLineItemStateResponse{}
	path := fmt.Sprintf("/projects/%s/lineItems/%s/%s", extProjectID, extLineItemID, action)
	err = c.requestAndParseResponse("POST", path, nil, res)
	return res, err
}

// GetAllLineItems ...
func (c *Client) GetAllLineItems(extProjectID string, options *QueryOptions) (*GetAllLineItemsResponse, error) {
	err := ValidateNotEmpty(extProjectID)
	if err != nil {
		return nil, err
	}
	res := &GetAllLineItemsResponse{}
	path := fmt.Sprintf("/projects/%s/lineItems%s", extProjectID, query2String(options))
	err = c.requestAndParseResponse("GET", path, nil, res)
	return res, err
}

// GetLineItemBy ...
func (c *Client) GetLineItemBy(extProjectID, extLineItemID string) (*LineItemResponse, error) {
	err := ValidateNotEmpty(extProjectID, extLineItemID)
	if err != nil {
		return nil, err
	}
	res := &LineItemResponse{}
	path := fmt.Sprintf("/projects/%s/lineItems/%s", extProjectID, extLineItemID)
	err = c.requestAndParseResponse("GET", path, nil, res)
	return res, err
}

// GetFeasibility ... Returns the feasibility for all the line items of the requested project. Takes 20 - 120
// seconds to execute. Check the `GetFeasibilityResponse.Feasibility.Status` field value to see if it is
// FeasibilityStatusReady ("READY") or FeasibilityStatusProcessing ("PROCESSING")
// If GetFeasibilityResponse.Feasibility.Status == FeasibilityStatusProcessing, call this function again in 2 mins.
func (c *Client) GetFeasibility(extProjectID string, options *QueryOptions) (*GetFeasibilityResponse, error) {
	err := ValidateNotEmpty(extProjectID)
	if err != nil {
		return nil, err
	}
	res := &GetFeasibilityResponse{}
	path := fmt.Sprintf("/projects/%s/feasibility%s", extProjectID, query2String(options))
	err = c.requestAndParseResponse("GET", path, nil, res)
	return res, err
}

// GetInvoice ... Get the invoice of the requested project
func (c *Client) GetInvoice(extProjectID string, options *QueryOptions) (*APIResponse, error) {
	path := fmt.Sprintf("/projects/%s/invoices", extProjectID)
	return c.request("GET", c.Options.APIBaseURL, path, nil)
}

// Reconcile ...  Upload the Request correction file
func (c *Client) UploadReconcile(extProjectID string, file multipart.File, fileName string, message string, options *QueryOptions) (*APIResponse, error) {
	//res := &APIResponse{}
	path := fmt.Sprintf("/projects/%s/reconcile", extProjectID)
	res, err := sendFormData(c.Options.APIBaseURL, "POST", path, c.Auth.AccessToken, file, fileName, message, *c.Options.Timeout)
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
	err := ValidateNotEmpty(countryCode, languageCode)
	if err != nil {
		return nil, err
	}
	err = IsCountryCodeOrEmpty(countryCode)
	if err != nil {
		return nil, err
	}
	err = IsLanguageCodeOrEmpty(languageCode)
	if err != nil {
		return nil, err
	}
	res := &GetAttributesResponse{}
	path := fmt.Sprintf("/attributes/%s/%s%s", countryCode, languageCode, query2String(options))
	err = c.requestAndParseResponse("GET", path, nil, res)
	return res, err
}

// GetSurveyTopics ... Get the list of supported Survey Topics for a project. This data is required to setup a project.
func (c *Client) GetSurveyTopics(options *QueryOptions) (*GetSurveyTopicsResponse, error) {
	res := &GetSurveyTopicsResponse{}
	path := fmt.Sprintf("/categories/surveyTopics%s", query2String(options))
	err := c.requestAndParseResponse("GET", path, nil, res)
	return res, err
}

// GetEvents ... Returns the list of all events that have occurred for your company account. Most recent events occur at the top of the list.
func (c *Client) GetEvents(options *QueryOptions) (*GetEventListResponse, error) {
	res := &GetEventListResponse{}
	path := fmt.Sprintf("/events%s", query2String(options))
	err := c.requestAndParseResponse("GET", path, nil, res)
	return res, err
}

// GetEventBy ... Returns the requested event based on the eventID
func (c *Client) GetEventBy(eventID string) (*GetEventResponse, error) {
	res := &GetEventResponse{}
	path := fmt.Sprintf("/events/%s", eventID)
	err := c.requestAndParseResponse("GET", path, nil, res)
	return res, err
}

// AcceptEvent ...
func (c *Client) AcceptEvent(event *Event) error {
	if event.Actions == nil || len(event.Actions.AcceptURL) == 0 {
		return ErrEventActionNotApplicable
	}
	_, err := c.request("POST", event.Actions.AcceptURL, "", nil)
	return err
}

// RejectEvent ...
func (c *Client) RejectEvent(event *Event) error {
	if event.Actions == nil || len(event.Actions.RejectURL) == 0 {
		return ErrEventActionNotApplicable
	}
	_, err := c.request("POST", event.Actions.RejectURL, "", nil)
	return err
}

// RefreshToken ...
func (c *Client) RefreshToken() error {
	if c.Auth.RefreshTokenExpired() {
		return ErrSessionExpired
	}
	t := time.Now()
	req := struct {
		ClientID     string `json:"clientId"`
		RefreshToken string `json:"refreshToken"`
	}{
		ClientID:     c.Credentials.ClientID,
		RefreshToken: c.Auth.RefreshToken,
	}
	ar, err := sendRequest(c.Options.AuthURL, "POST", "/token/refresh", "", req, *c.Options.Timeout)
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
	_, err := sendRequest(c.Options.AuthURL, "POST", "/logout", "", req, *c.Options.Timeout)
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
	ar, err := c.request(method, c.Options.APIBaseURL, url, body)
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

func (c *Client) request(method, host, url string, body interface{}) (*APIResponse, error) {
	err := c.validateTokens()
	if err != nil {
		return nil, err
	}
	ar, err := sendRequest(host, method, url, c.Auth.AccessToken, body, *c.Options.Timeout)
	errResp, ok := err.(*ErrorResponse)
	if ok && errResp.HTTPCode == http.StatusUnauthorized {
		err := c.requestAndParseToken()
		if err != nil {
			return nil, err
		}
		return sendRequest(host, method, url, c.Auth.AccessToken, body, *c.Options.Timeout)
	}
	return ar, err
}

func (c *Client) requestAndParseToken() error {
	log.WithFields(log.Fields{"module": "go-samplifyapi-client", "function": "requestAndParseToken", "ClientID": c.Credentials.ClientID}).Info()
	t := time.Now()
	ar, err := sendRequest(c.Options.AuthURL, "POST", "/token/password", "", c.Credentials, *c.Options.Timeout)
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

// ValidateTokens ...
func (c *Client) validateTokens() error {
	if c.Auth.AccessTokenExpired() {
		err := c.RefreshToken()
		if err != nil {
			err := c.requestAndParseToken()
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// NewClient returns an API client.
// If options is nil, UATClientOptions will be used.
func NewClient(clientID, username, passsword string, options *ClientOptions) *Client {
	if options == nil {
		options = UATClientOptions
	}
	if options != nil {
		if options.Timeout == nil {
			timeout := defaulttimeout
			options.Timeout = &timeout
		}
	}
	return &Client{
		Credentials: TokenRequest{
			ClientID: clientID,
			Username: username,
			Password: passsword,
		},
		Options: *options,
	}
}
