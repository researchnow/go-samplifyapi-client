package samplify

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"time"
)

// ErrIncorrectEnvironemt ...
var ErrIncorrectEnvironemt = errors.New("one of local/dev/uat/prod only are allowed")

// ClientOptions to use while creating a new Client
var (
	LocalClientOptions = &ClientOptions{
		APIBaseURL:  "http://localhost:8090/sample/v1",
		AuthURL:     "http://localhost:8090/auth/v1",
		InternalURL: "http://localhost:8090/internal/v1",
		StatusURL:   "http://localhost:8090/status",
		GatewayURL:  "http://localhost:8090/status/gateway",
	}
	DevClientOptions = &ClientOptions{
		APIBaseURL:  "https://api.dev.pe.dynata.com/sample/v1",
		AuthURL:     "https://api.dev.pe.dynata.com/auth/v1",
		InternalURL: "https://api.dev.pe.dynata.com/internal/v1",
		StatusURL:   "https://api.dev.pe.dynata.com/status",
		GatewayURL:  "https://api.dev.pe.dynata.com/status/gateway",
	}
	UATClientOptions = &ClientOptions{
		APIBaseURL:  "https://api.uat.pe.dynata.com/sample/v1",
		AuthURL:     "https://api.uat.pe.dynata.com/auth/v1",
		InternalURL: "https://api.uat.pe.dynata.com/internal/v1",
		StatusURL:   "https://api.uat.pe.dynata.com/status",
		GatewayURL:  "https://api.uat.pe.dynata.com/status/gateway",
	}
	ProdClientOptions = &ClientOptions{
		APIBaseURL:  "https://api.researchnow.com/sample/v1",
		AuthURL:     "https://api.researchnow.com/auth/v1",
		InternalURL: "https://api.researchnow.com/internal/v1",
		StatusURL:   "https://api.researchnow.com/status",
		GatewayURL:  "https://api.researchnow.com/status/gateway",
	}
)

// ErrSessionExpired ... Returns if both Access and Refresh tokens are expired
var ErrSessionExpired = errors.New("session expired")

const defaulttimeout = 20

// ClientOptions ...
type ClientOptions struct {
	APIBaseURL  string `conform:"trim"`
	AuthURL     string `conform:"trim"`
	InternalURL string `conform:"trim"`
	StatusURL   string `conform:"trim"`
	GatewayURL  string `conform:"trim"`
	Timeout     *int
}

// Client is used to make API requests to the Samplify API.
type Client struct {
	Credentials TokenRequest
	Auth        TokenResponse
	Options     *ClientOptions
}


// GetOrderDetailsWithContext ...
func (c *Client) GetOrderDetailsWithContext(ctx context.Context, ordNumber string) (*OrderDetailResponse, error) {
	path := fmt.Sprintf("/orderdetails/%s/", ordNumber)
	res := &OrderDetailResponse{}
	err := c.requestAndParseResponse(ctx, "GET", c.Options.InternalURL, path, res)
	return res , err
}

// GetOrderDetails ...
func (c *Client) GetOrderDetails(ordNumber string) (*OrderDetailResponse, error) {
	return c.GetOrderDetailsWithContext(context.Background(), ordNumber)
}
// CheckOrderNumberWithContext ...
func (c *Client) CheckOrderNumberWithContext(ctx context.Context, ordNumber string) (*OrderDetailResponse, error) {
	path := fmt.Sprintf("/orderdetails/check/%s", ordNumber)
	res := &OrderDetailResponse{}
	err := c.requestAndParseResponse(ctx, "GET", c.Options.InternalURL, path, res)
	return res , err
}

// CheckOrderNumber ...
func (c *Client) CheckOrderNumber(ordNumber string) (*OrderDetailResponse, error) {
	return c.CheckOrderNumberWithContext(context.Background(), ordNumber)
}

// GetInvoicesSummaryWithContext ...
func (c *Client) GetInvoicesSummaryWithContext(ctx context.Context, options *QueryOptions) (*APIResponse, error) {
	path := fmt.Sprintf("/projects/invoices/summary%s", query2String(options))
	return c.request(ctx, "GET", c.Options.APIBaseURL, path, nil)
}

// GetInvoicesSummary ...
func (c *Client) GetInvoicesSummary(options *QueryOptions) (*APIResponse, error) {
	return c.GetInvoicesSummaryWithContext(context.Background(), options)
}

// CreateProjectWithContext ...
func (c *Client) CreateProjectWithContext(ctx context.Context, project *CreateProjectCriteria) (*ProjectResponse, error) {
	err := Validate(project)
	if err != nil {
		return nil, err
	}
	res := &ProjectResponse{}
	err = c.requestAndParseResponse(ctx, "POST", "/projects", project, res)
	return res, err
}

// CreateProject ...
func (c *Client) CreateProject(project *CreateProjectCriteria) (*ProjectResponse, error) {
	return c.CreateProjectWithContext(context.Background(), project)
}

// UpdateProjectWithContext ...
func (c *Client) UpdateProjectWithContext(ctx context.Context, project *UpdateProjectCriteria) (*ProjectResponse, error) {

	err := Validate(project)
	if err != nil {
		return nil, err
	}
	res := &ProjectResponse{}
	path := fmt.Sprintf("/projects/%s", project.ExtProjectID)
	err = c.requestAndParseResponse(ctx, "POST", path, project, res)
	return res, err
}

// UpdateProject ...
func (c *Client) UpdateProject(project *UpdateProjectCriteria) (*ProjectResponse, error) {
	return c.UpdateProjectWithContext(context.Background(), project)
}

// BuyProjectWithContext ...
func (c *Client) BuyProjectWithContext(ctx context.Context, extProjectID string, buy []*BuyProjectCriteria) (*BuyProjectResponse, error) {
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
	err = c.requestAndParseResponse(ctx, "POST", path, buy, res)
	return res, err
}

// BuyProject ...
func (c *Client) BuyProject(extProjectID string, buy []*BuyProjectCriteria) (*BuyProjectResponse, error) {
	return c.BuyProjectWithContext(context.Background(), extProjectID, buy)
}

// CloseProjectWithContext ...
func (c *Client) CloseProjectWithContext(ctx context.Context, extProjectID string) (*CloseProjectResponse, error) {
	err := ValidateNotEmpty(extProjectID)
	if err != nil {
		return nil, err
	}
	res := &CloseProjectResponse{}
	path := fmt.Sprintf("/projects/%s/close", extProjectID)
	err = c.requestAndParseResponse(ctx, "POST", path, nil, res)
	return res, err
}

// CloseProject ...
func (c *Client) CloseProject(extProjectID string) (*CloseProjectResponse, error) {
	return c.CloseProjectWithContext(context.Background(), extProjectID)
}

// GetAllProjectsWithContext ...
func (c *Client) GetAllProjectsWithContext(ctx context.Context, options *QueryOptions) (*GetAllProjectsResponse, error) {
	res := &GetAllProjectsResponse{}
	query := query2String(options)
	path := fmt.Sprintf("/projects%s", query)
	err := c.requestAndParseResponse(ctx, "GET", path, nil, res)
	return res, err
}

// GetAllProjects ...
func (c *Client) GetAllProjects(options *QueryOptions) (*GetAllProjectsResponse, error) {
	return c.GetAllProjectsWithContext(context.Background(), options)
}

// GetProjectByWithContext returns project by id
func (c *Client) GetProjectByWithContext(ctx context.Context, extProjectID string) (*ProjectResponse, error) {
	err := ValidateNotEmpty(extProjectID)
	if err != nil {
		return nil, err
	}
	res := &ProjectResponse{}
	path := fmt.Sprintf("/projects/%s", extProjectID)
	err = c.requestAndParseResponse(ctx, "GET", path, nil, res)
	return res, err
}

// GetProjectBy returns project by id
func (c *Client) GetProjectBy(extProjectID string) (*ProjectResponse, error) {
	return c.GetProjectByWithContext(context.Background(), extProjectID)
}

// GetProjectReportWithContext returns a project's report based on observed data from actual panelists.
func (c *Client) GetProjectReportWithContext(ctx context.Context, extProjectID string) (*ProjectReportResponse, error) {
	err := ValidateNotEmpty(extProjectID)
	if err != nil {
		return nil, err
	}
	res := &ProjectReportResponse{}
	path := fmt.Sprintf("/projects/%s/report", extProjectID)
	err = c.requestAndParseResponse(ctx, "GET", path, nil, res)
	return res, err
}

// GetProjectReport returns a project's report based on observed data from actual panelists.
func (c *Client) GetProjectReport(extProjectID string) (*ProjectReportResponse, error) {
	return c.GetProjectReportWithContext(context.Background(), extProjectID)
}

// AddLineItemWithContext ...
func (c *Client) AddLineItemWithContext(ctx context.Context, extProjectID string, lineItem *CreateLineItemCriteria) (*LineItemResponse, error) {
	err := ValidateNotEmpty(extProjectID)
	if err != nil {
		return nil, err
	}
	err = Validate(lineItem)
	if err != nil {
		return nil, err
	}
	err = ValidateSchedule(&lineItem.DaysInField, lineItem.FieldSchedule)
	if err != nil {
		return nil, err
	}
	res := &LineItemResponse{}
	path := fmt.Sprintf("/projects/%s/lineItems", extProjectID)
	err = c.requestAndParseResponse(ctx, "POST", path, lineItem, res)
	return res, err
}

// AddLineItem ...
func (c *Client) AddLineItem(extProjectID string, lineItem *CreateLineItemCriteria) (*LineItemResponse, error) {
	return c.AddLineItemWithContext(context.Background(), extProjectID, lineItem)
}

// UpdateLineItemWithContext ...
func (c *Client) UpdateLineItemWithContext(ctx context.Context, extProjectID, extLineItemID string,
	lineItem *UpdateLineItemCriteria) (*LineItemResponse, error) {

	err := ValidateNotEmpty(extProjectID, extLineItemID)
	if err != nil {
		return nil, err
	}
	err = Validate(lineItem)
	if err != nil {
		return nil, err
	}
	err = ValidateSchedule(lineItem.DaysInField, lineItem.FieldSchedule)
	if err != nil {
		return nil, err
	}
	res := &LineItemResponse{}
	path := fmt.Sprintf("/projects/%s/lineItems/%s", extProjectID, extLineItemID)
	err = c.requestAndParseResponse(ctx, "POST", path, lineItem, res)
	return res, err
}

// UpdateLineItem ...
func (c *Client) UpdateLineItem(extProjectID, extLineItemID string,
	lineItem *UpdateLineItemCriteria) (*LineItemResponse, error) {

	return c.UpdateLineItemWithContext(context.Background(), extProjectID, extLineItemID, lineItem)
}

// UpdateLineItemStateWithContext ... Changes the state of the line item based on provided action.
func (c *Client) UpdateLineItemStateWithContext(ctx context.Context, extProjectID, extLineItemID string, action Action) (
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
	err = c.requestAndParseResponse(ctx, "POST", path, nil, res)
	return res, err
}

// UpdateLineItemState ... Changes the state of the line item based on provided action.
func (c *Client) UpdateLineItemState(extProjectID, extLineItemID string, action Action) (
	*UpdateLineItemStateResponse, error) {
	return c.UpdateLineItemStateWithContext(context.Background(), extProjectID, extLineItemID, action)
}

// LaunchLineItemWithContext utility function to launch a line item
func (c *Client) LaunchLineItemWithContext(ctx context.Context, pid, lid string) (*UpdateLineItemStateResponse, error) {
	return c.UpdateLineItemStateWithContext(ctx, pid, lid, ActionLaunched)
}

// LaunchLineItem utility function to launch a line item
func (c *Client) LaunchLineItem(pid, lid string) (*UpdateLineItemStateResponse, error) {
	return c.LaunchLineItemWithContext(context.Background(), pid, lid)
}

// PauseLineItemWithContext utility function to pause a lineitem
func (c *Client) PauseLineItemWithContext(ctx context.Context, pid, lid string) (*UpdateLineItemStateResponse, error) {
	return c.UpdateLineItemStateWithContext(ctx, pid, lid, ActionPaused)
}

// PauseLineItem utility function to pause a lineitem
func (c *Client) PauseLineItem(pid, lid string) (*UpdateLineItemStateResponse, error) {
	return c.PauseLineItemWithContext(context.Background(), pid, lid)
}

// CloseLineItemWithContext utility function to close a lineitem
func (c *Client) CloseLineItemWithContext(ctx context.Context, pid, lid string) (*UpdateLineItemStateResponse, error) {
	return c.UpdateLineItemStateWithContext(ctx, pid, lid, ActionClosed)
}

// CloseLineItem utility function to close a lineitem
func (c *Client) CloseLineItem(pid, lid string) (*UpdateLineItemStateResponse, error) {
	return c.CloseLineItemWithContext(context.Background(), pid, lid)
}

// SetQuotaCellStatusWithContext ... Changes the state of the line item based on provided action.
func (c *Client) SetQuotaCellStatusWithContext(ctx context.Context, extProjectID, extLineItemID string, quotaCellID string, action Action) (
	*QuotaCellResponse, error) {
	err := ValidateNotEmpty(extProjectID, extLineItemID, quotaCellID)
	if err != nil {
		return nil, err
	}
	err = ValidateAction(action)
	if err != nil {
		return nil, err
	}
	res := &QuotaCellResponse{}
	path := fmt.Sprintf("/projects/%s/lineItems/%s/quotaCells/%s/%s", extProjectID, extLineItemID, quotaCellID, action)
	err = c.requestAndParseResponse(ctx, "POST", path, nil, res)
	return res, err
}

// SetQuotaCellStatus ... Changes the state of the line item based on provided action.
func (c *Client) SetQuotaCellStatus(extProjectID, extLineItemID string, quotaCellID string, action Action) (
	*QuotaCellResponse, error) {
	return c.SetQuotaCellStatusWithContext(context.Background(), extProjectID, extLineItemID, quotaCellID, action)
}

// GetAllLineItemsWithContext ...
func (c *Client) GetAllLineItemsWithContext(ctx context.Context, extProjectID string, options *QueryOptions) (*GetAllLineItemsResponse, error) {
	err := ValidateNotEmpty(extProjectID)
	if err != nil {
		return nil, err
	}
	res := &GetAllLineItemsResponse{}
	path := fmt.Sprintf("/projects/%s/lineItems%s", extProjectID, query2String(options))
	err = c.requestAndParseResponse(ctx, "GET", path, nil, res)
	return res, err
}

// GetAllLineItems ...
func (c *Client) GetAllLineItems(extProjectID string, options *QueryOptions) (*GetAllLineItemsResponse, error) {
	return c.GetAllLineItemsWithContext(context.Background(), extProjectID, options)
}

// GetLineItemByWithContext ...
func (c *Client) GetLineItemByWithContext(ctx context.Context, extProjectID, extLineItemID string) (*LineItemResponse, error) {
	err := ValidateNotEmpty(extProjectID, extLineItemID)
	if err != nil {
		return nil, err
	}
	res := &LineItemResponse{}
	path := fmt.Sprintf("/projects/%s/lineItems/%s", extProjectID, extLineItemID)
	err = c.requestAndParseResponse(ctx, "GET", path, nil, res)
	return res, err
}

// GetLineItemBy ...
func (c *Client) GetLineItemBy(extProjectID, extLineItemID string) (*LineItemResponse, error) {
	return c.GetLineItemByWithContext(context.Background(), extProjectID, extLineItemID)
}

// GetFeasibilityWithContext ... Returns the feasibility for all the line items of the requested project. Takes 20 - 120
// seconds to execute. Check the `GetFeasibilityResponse.Feasibility.Status` field value to see if it is
// FeasibilityStatusReady ("READY") or FeasibilityStatusProcessing ("PROCESSING")
// If GetFeasibilityResponse.Feasibility.Status == FeasibilityStatusProcessing, call this function again in 2 mins.
func (c *Client) GetFeasibilityWithContext(ctx context.Context, extProjectID string, options *QueryOptions) (*GetFeasibilityResponse, error) {
	err := ValidateNotEmpty(extProjectID)
	if err != nil {
		return nil, err
	}
	res := &GetFeasibilityResponse{}
	path := fmt.Sprintf("/projects/%s/feasibility%s", extProjectID, query2String(options))
	err = c.requestAndParseResponse(ctx, "GET", path, nil, res)

	return res, err
}

// GetFeasibility ... Returns the feasibility for all the line items of the requested project. Takes 20 - 120
// seconds to execute. Check the `GetFeasibilityResponse.Feasibility.Status` field value to see if it is
// FeasibilityStatusReady ("READY") or FeasibilityStatusProcessing ("PROCESSING")
// If GetFeasibilityResponse.Feasibility.Status == FeasibilityStatusProcessing, call this function again in 2 mins.
func (c *Client) GetFeasibility(extProjectID string, options *QueryOptions) (*GetFeasibilityResponse, error) {
	return c.GetFeasibilityWithContext(context.Background(), extProjectID, options)
}

// GetInvoiceWithContext ... Get the invoice of the requested project
func (c *Client) GetInvoiceWithContext(ctx context.Context, extProjectID string, options *QueryOptions) (*APIResponse, error) {
	path := fmt.Sprintf("/projects/%s/invoices", extProjectID)
	return c.request(ctx, "GET", c.Options.APIBaseURL, path, nil)
}

// GetInvoice ... Get the invoice of the requested project
func (c *Client) GetInvoice(extProjectID string, options *QueryOptions) (*APIResponse, error) {
	return c.GetInvoiceWithContext(context.Background(), extProjectID, options)
}

// UploadReconcileWithContext ...  Upload the Request correction file
func (c *Client) UploadReconcileWithContext(ctx context.Context, extProjectID string, file multipart.File, fileName string, message string, options *QueryOptions) (*APIResponse, error) {
	path := fmt.Sprintf("/projects/%s/reconcile", extProjectID)
	res, err := sendFormData(ctx, c.Options.APIBaseURL, "POST", path, c.Auth.AccessToken, file, fileName, message, *c.Options.Timeout)
	return res, err
}

// UploadReconcile ...  Upload the Request correction file
func (c *Client) UploadReconcile(extProjectID string, file multipart.File, fileName string, message string, options *QueryOptions) (*APIResponse, error) {
	return c.UploadReconcileWithContext(context.Background(), extProjectID, file, fileName, message, options)
}

// GetCountriesWithContext ... Get the list of supported countries and languages in each country.
func (c *Client) GetCountriesWithContext(ctx context.Context, options *QueryOptions) (*GetCountriesResponse, error) {
	res := &GetCountriesResponse{}
	path := fmt.Sprintf("/countries%s", query2String(options))
	err := c.requestAndParseResponse(ctx, "GET", path, nil, res)
	return res, err
}

// GetCountries ... Get the list of supported countries and languages in each country.
func (c *Client) GetCountries(options *QueryOptions) (*GetCountriesResponse, error) {
	return c.GetCountriesWithContext(context.Background(), options)
}

// GetAttributesWithContext ... Get the list of supported attributes for a country and language. This data is required to build up the Quota Plan.
func (c *Client) GetAttributesWithContext(ctx context.Context, countryCode, languageCode string, options *QueryOptions) (*GetAttributesResponse, error) {
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
	err = c.requestAndParseResponse(ctx, "GET", path, nil, res)
	return res, err
}

// GetAttributes ... Get the list of supported attributes for a country and language. This data is required to build up the Quota Plan.
func (c *Client) GetAttributes(countryCode, languageCode string, options *QueryOptions) (*GetAttributesResponse, error) {
	return c.GetAttributesWithContext(context.Background(), countryCode, languageCode, options)
}

// GetSurveyTopicsWithContext ... Get the list of supported Survey Topics for a project. This data is required to setup a project.
func (c *Client) GetSurveyTopicsWithContext(ctx context.Context, options *QueryOptions) (*GetSurveyTopicsResponse, error) {
	res := &GetSurveyTopicsResponse{}
	path := fmt.Sprintf("/categories/surveyTopics%s", query2String(options))
	err := c.requestAndParseResponse(ctx, "GET", path, nil, res)
	return res, err
}

// GetSurveyTopics ... Get the list of supported Survey Topics for a project. This data is required to setup a project.
func (c *Client) GetSurveyTopics(options *QueryOptions) (*GetSurveyTopicsResponse, error) {
	return c.GetSurveyTopicsWithContext(context.Background(), options)
}

// GetSourcesWithContext ... Get the list of all the Sample sources
func (c *Client) GetSourcesWithContext(ctx context.Context, options *QueryOptions) (*GetSampleSourceResponse, error) {
	res := &GetSampleSourceResponse{}
	path := fmt.Sprintf("/sources%s", query2String(options))
	err := c.requestAndParseResponse(ctx, "GET", path, nil, res)
	return res, err
}

// GetSources ... Get the list of all the Sample sources
func (c *Client) GetSources(options *QueryOptions) (*GetSampleSourceResponse, error) {
	return c.GetSourcesWithContext(context.Background(), options)
}

// GetEventsWithContext ... Returns the list of all events that have occurred for your company account. Most recent events occur at the top of the list.
func (c *Client) GetEventsWithContext(ctx context.Context, options *QueryOptions) (*GetEventListResponse, error) {
	res := &GetEventListResponse{}
	path := fmt.Sprintf("/events%s", query2String(options))
	err := c.requestAndParseResponse(ctx, "GET", path, nil, res)
	return res, err
}

// GetEvents ... Returns the list of all events that have occurred for your company account. Most recent events occur at the top of the list.
func (c *Client) GetEvents(options *QueryOptions) (*GetEventListResponse, error) {
	return c.GetEventsWithContext(context.Background(), options)
}

// GetEventByWithContext ... Returns the requested event based on the eventID
func (c *Client) GetEventByWithContext(ctx context.Context, eventID string) (*GetEventResponse, error) {
	res := &GetEventResponse{}
	path := fmt.Sprintf("/events/%s", eventID)
	err := c.requestAndParseResponse(ctx, "GET", path, nil, res)
	return res, err
}

// GetEventBy ... Returns the requested event based on the eventID
func (c *Client) GetEventBy(eventID string) (*GetEventResponse, error) {
	return c.GetEventByWithContext(context.Background(), eventID)
}

// AcceptEventWithContext ...
func (c *Client) AcceptEventWithContext(ctx context.Context, event *Event) error {
	if event.Actions == nil || len(event.Actions.AcceptURL) == 0 {
		return ErrEventActionNotApplicable
	}
	_, err := c.request(ctx, "POST", event.Actions.AcceptURL, "", nil)
	return err
}

// AcceptEvent ...
func (c *Client) AcceptEvent(event *Event) error {
	return c.AcceptEventWithContext(context.Background(), event)
}

// RejectEventWithContext ...
func (c *Client) RejectEventWithContext(ctx context.Context, event *Event) error {
	if event.Actions == nil || len(event.Actions.RejectURL) == 0 {
		return ErrEventActionNotApplicable
	}
	_, err := c.request(ctx, "POST", event.Actions.RejectURL, "", nil)
	return err
}

// RejectEvent ...
func (c *Client) RejectEvent(event *Event) error {
	return c.RejectEventWithContext(context.Background(), event)
}

// GetDetailedProjectReportWithContext returns a project's detailed report based on observed data from actual panelists.
func (c *Client) GetDetailedProjectReportWithContext(ctx context.Context, extProjectID string) (*DetailedProjectReportResponse, error) {
	err := ValidateNotEmpty(extProjectID)
	if err != nil {
		return nil, err
	}
	res := &DetailedProjectReportResponse{}
	path := fmt.Sprintf("/projects/%s/detailedReport", extProjectID)
	err = c.requestAndParseResponse(ctx, "GET", path, nil, res)
	return res, err
}

// GetDetailedProjectReport returns a project's detailed report based on observed data from actual panelists.
func (c *Client) GetDetailedProjectReport(extProjectID string) (*DetailedProjectReportResponse, error) {
	return c.GetDetailedProjectReportWithContext(context.Background(), extProjectID)
}

// GetDetailedLineItemReportWithContext returns a lineitems's report with quota cell level stats based on observed data from actual panelists.
func (c *Client) GetDetailedLineItemReportWithContext(ctx context.Context, extProjectID, extLineItemID string) (*DetailedLineItemReportResponse, error) {
	err := ValidateNotEmpty(extProjectID)
	if err != nil {
		return nil, err
	}
	res := &DetailedLineItemReportResponse{}
	path := fmt.Sprintf("/projects/%s/lineItems/%s/detailedReport", extProjectID, extLineItemID)
	err = c.requestAndParseResponse(ctx, "GET", path, nil, res)
	return res, err
}

// GetDetailedLineItemReport returns a lineitems's report with quota cell level stats based on observed data from actual panelists.
func (c *Client) GetDetailedLineItemReport(extProjectID, extLineItemID string) (*DetailedLineItemReportResponse, error) {
	return c.GetDetailedLineItemReportWithContext(context.Background(), extProjectID, extLineItemID)
}

// GetUserInfoWithContext gives information about the user that is currently logged in.
func (c *Client) GetUserInfoWithContext(ctx context.Context) (*UserResponse, error) {
	res := &UserResponse{}
	path := "/users/info"
	err := c.requestAndParseResponse(ctx, "GET", path, nil, res)
	return res, err
}

// GetUserInfo gives information about the user that is currently logged in.
func (c *Client) GetUserInfo() (*UserResponse, error) {
	return c.GetUserInfoWithContext(context.Background())
}

// GetUserDetailsWithContext ...
func (c *Client) GetUserDetailsWithContext(ctx context.Context) (*UserDetailsResponse, error) {
	res := &UserDetailsResponse{}
	path := "/user"
	err := c.requestAndParseResponse(ctx, "GET", path, nil, res)
	return res, err
}

// GetUserDetails ...
func (c *Client) GetUserDetails() (*UserDetailsResponse, error) {
	return c.GetUserDetailsWithContext(context.Background())
}

// CompanyUsersWithContext gives information about the user that is currently logged in.
func (c *Client) CompanyUsersWithContext(ctx context.Context) (*CompanyUsersResponse, error) {
	res := &CompanyUsersResponse{}
	path := "/users"
	err := c.requestAndParseResponse(ctx, "GET", path, nil, res)
	return res, err
}

// CompanyUsers gives information about the user that is currently logged in.
func (c *Client) CompanyUsers() (*CompanyUsersResponse, error) {
	return c.CompanyUsersWithContext(context.Background())
}

// TeamsInfoWithContext gives information about the user that is currently logged in.
func (c *Client) TeamsInfoWithContext(ctx context.Context) (*TeamsResponse, error) {
	res := &TeamsResponse{}
	path := "/teams"
	err := c.requestAndParseResponse(ctx, "GET", path, nil, res)
	return res, err
}

// TeamsInfo gives information about the user that is currently logged in.
func (c *Client) TeamsInfo() (*TeamsResponse, error) {
	return c.TeamsInfoWithContext(context.Background())
}

// RolesWithContext returns the roles specified in the filter.
func (c *Client) RolesWithContext(ctx context.Context, options *QueryOptions) (*RolesResponse, error) {
	res := &RolesResponse{}
	path := fmt.Sprintf("/roles%s", query2String(options))
	err := c.requestAndParseResponse(ctx, "GET", path, nil, res)
	return res, err
}

// Roles returns the roles specified in the filter.
func (c *Client) Roles(options *QueryOptions) (*RolesResponse, error) {
	return c.RolesWithContext(context.Background(), options)
}

// ProjectPermissionsWithContext gives information about the user that is currently logged in.
func (c *Client) ProjectPermissionsWithContext(ctx context.Context, extProjectID string) (*ProjectPermissionsResponse, error) {
	err := ValidateNotEmpty(extProjectID)
	if err != nil {
		return nil, err
	}
	res := &ProjectPermissionsResponse{}
	path := fmt.Sprintf("/projects/%s/permissions", extProjectID)
	err = c.requestAndParseResponse(ctx, "GET", path, nil, res)
	return res, err
}

// ProjectPermissions gives information about the user that is currently logged in.
func (c *Client) ProjectPermissions(extProjectID string) (*ProjectPermissionsResponse, error) {
	return c.ProjectPermissionsWithContext(context.Background(), extProjectID)
}

// UpsertProjectPermissionsWithContext gives information about the user that is currently logged in.
func (c *Client) UpsertProjectPermissionsWithContext(ctx context.Context, permissions *UpsertPermissionsCriteria) (*ProjectPermissionsResponse, error) {
	err := Validate(permissions)
	if err != nil {
		return nil, err
	}
	res := &ProjectPermissionsResponse{}
	path := fmt.Sprintf("/projects/%s/permissions", permissions.ExtProjectID)
	err = c.requestAndParseResponse(ctx, "POST", path, permissions, res)
	return res, err
}

// UpsertProjectPermissions gives information about the user that is currently logged in.
func (c *Client) UpsertProjectPermissions(permissions *UpsertPermissionsCriteria) (*ProjectPermissionsResponse, error) {
	return c.UpsertProjectPermissionsWithContext(context.Background(), permissions)
}

// GetStudyMetadataWithContext returns study metadata property info
func (c *Client) GetStudyMetadataWithContext(ctx context.Context) (*StudyMetadataResponse, error) {
	res := &StudyMetadataResponse{}
	path := "/studyMetadata"
	err := c.requestAndParseResponse(ctx, "GET", path, nil, res)
	return res, err
}

// GetStudyMetadata returns study metadata property info
func (c *Client) GetStudyMetadata() (*StudyMetadataResponse, error) {
	return c.GetStudyMetadataWithContext(context.Background())
}

// CreateTemplateWithContext ...
func (c *Client) CreateTemplateWithContext(ctx context.Context, template *TemplateCriteria) (*TemplateResponse, error) {
	err := Validate(template)
	if err != nil {
		return nil, err
	}
	res := &TemplateResponse{}
	err = c.requestAndParseResponse(ctx, "POST", "/templates/quotaPlan", template, res)
	return res, err
}

// CreateTemplate ...
func (c *Client) CreateTemplate(template *TemplateCriteria) (*TemplateResponse, error) {
	return c.CreateTemplateWithContext(context.Background(), template)
}

// UpdateTemplateWithContext ...
func (c *Client) UpdateTemplateWithContext(ctx context.Context, id int, template *TemplateCriteria) (*TemplateResponse, error) {
	err := Validate(template)
	if err != nil {
		return nil, err
	}
	res := &TemplateResponse{}
	path := fmt.Sprintf("/templates/quotaPlan/%d", id)
	err = c.requestAndParseResponse(ctx, "POST", path, template, res)
	return res, err
}

// UpdateTemplate ...
func (c *Client) UpdateTemplate(id int, template *TemplateCriteria) (*TemplateResponse, error) {
	return c.UpdateTemplateWithContext(context.Background(), id, template)
}

// GetTemplateListWithContext ...
func (c *Client) GetTemplateListWithContext(ctx context.Context, country string, lang string, options *QueryOptions) (*TemplatesResponse, error) {
	res := &TemplatesResponse{}
	query := query2String(options)
	path := fmt.Sprintf("/templates/quotaPlan/%s/%s%s", country, lang, query)
	err := c.requestAndParseResponse(ctx, "GET", path, nil, res)
	return res, err
}

// GetTemplateList ...
func (c *Client) GetTemplateList(country string, lang string, options *QueryOptions) (*TemplatesResponse, error) {
	return c.GetTemplateListWithContext(context.Background(), country, lang, options)
}

// DeleteTemplateWithContext ...
func (c *Client) DeleteTemplateWithContext(ctx context.Context, id int) (*AppError, error) {
	res := &AppError{}
	path := fmt.Sprintf("/templates/quotaPlan/%d", id)
	err := c.requestAndParseResponse(ctx, "DELETE", path, nil, res)
	return res, err
}

// DeleteTemplate ...
func (c *Client) DeleteTemplate(id int) (*AppError, error) {
	return c.DeleteTemplateWithContext(context.Background(), id)
}

// SwitchCompanyWithContext ...
func (c *Client) SwitchCompanyWithContext(ctx context.Context, criteria *SwitchCompanyCriteria) error {
	response, err := c.request(ctx, "POST", c.Options.AuthURL, "/switchCompany", criteria)
	if err != nil {
		return err
	}
	err = json.Unmarshal(response.Body, &c.Auth)
	if err != nil {
		return err
	}
	now := time.Now()
	c.Auth.Acquired = &now
	return nil
}

// SwitchCompany ...
func (c *Client) SwitchCompany(criteria *SwitchCompanyCriteria) error {
	return c.SwitchCompanyWithContext(context.Background(), criteria)
}

// RefreshTokenWithContext ...
func (c *Client) RefreshTokenWithContext(ctx context.Context) error {
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
	ar, err := sendRequest(ctx, c.Options.AuthURL, "POST", "/token/refresh", "", req, *c.Options.Timeout)
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

// RefreshToken ...
func (c *Client) RefreshToken() error {
	return c.RefreshTokenWithContext(context.Background())
}

// LogoutWithContext ...
func (c *Client) LogoutWithContext(ctx context.Context) error {
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
	_, err := sendRequest(ctx, c.Options.AuthURL, "POST", "/logout", "", req, *c.Options.Timeout)
	return err
}

// Logout ...
func (c *Client) Logout() error {
	return c.LogoutWithContext(context.Background())
}

// GetAuthWithContext ...
func (c *Client) GetAuthWithContext(ctx context.Context) (TokenResponse, error) {
	err := c.requestAndParseToken(ctx)
	if err != nil {
		return TokenResponse{}, err
	}
	return c.Auth, err
}

// GetAuth ...
func (c *Client) GetAuth() (TokenResponse, error) {
	return c.GetAuthWithContext(context.Background())
}

func (c *Client) requestAndParseResponse(ctx context.Context, method, url string, body interface{}, resObj interface{}) error {
	ar, err := c.request(ctx, method, c.Options.APIBaseURL, url, body)
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

func (c *Client) request(ctx context.Context, method, host, url string, body interface{}) (*APIResponse, error) {
	err := c.validateTokens(ctx)
	if err != nil {
		return nil, err
	}
	ar, err := sendRequest(ctx, host, method, url, c.Auth.AccessToken, body, *c.Options.Timeout)
	errResp, ok := err.(*ErrorResponse)
	if ok && errResp.HTTPCode == http.StatusUnauthorized {
		err := c.requestAndParseToken(ctx)
		if err != nil {
			return nil, err
		}
		return sendRequest(ctx, host, method, url, c.Auth.AccessToken, body, *c.Options.Timeout)
	}
	return ar, err
}

func (c *Client) requestAndParseToken(ctx context.Context) error {
	// log.WithFields(log.Fields{"module": "go-samplifyapi-client", "function": "requestAndParseToken", "ClientID": c.Credentials.ClientID}).Info()
	t := time.Now()
	ar, err := sendRequest(ctx, c.Options.AuthURL, "POST", "/token/password", "", c.Credentials, *c.Options.Timeout)
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
func (c *Client) validateTokens(ctx context.Context) error {
	if c.Auth.AccessTokenExpired() {
		err := c.RefreshTokenWithContext(ctx)
		if err != nil {
			err := c.requestAndParseToken(ctx)
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

	if options != nil && options.Timeout == nil {
		timeout := defaulttimeout
		options.Timeout = &timeout
	}

	return &Client{
		Credentials: TokenRequest{
			ClientID: clientID,
			Username: username,
			Password: passsword,
		},
		Options: options,
	}
}

// SetOptions ...
func (c *Client) SetOptions(env string, timeout int) error {
	switch env {
	case "local":
		c.Options = LocalClientOptions
	case "dev":
		c.Options = DevClientOptions
	case "uat":
		c.Options = UATClientOptions
	case "prod":
		c.Options = ProdClientOptions
	}

	if c.Options == nil {
		return ErrIncorrectEnvironemt
	}

	if timeout == 0 {
		timeout = defaulttimeout
	}

	c.Options.Timeout = &timeout

	return nil
}

// NewClientFromEnv returns an API client.
func NewClientFromEnv(clientID, username, passsword string, env string, timeout int) (*Client, error) {
	client := &Client{
		Credentials: TokenRequest{
			ClientID: clientID,
			Username: username,
			Password: passsword,
		},
	}
	err := client.SetOptions(env, timeout)
	return client, err
}

// GetHealthyStatus ... Get the healthy status on API
func (c *Client) GetHealthyStatus() (*APIResponse, error) {
	return c.GetHealthyStatusWithContext(context.Background())
}

func (c *Client) GetHealthyStatusWithContext(ctx context.Context) (*APIResponse, error) {
	return c.request(ctx, "GET", c.Options.GatewayURL, "", nil)
}
