package samplify

// StatusType ...
type StatusType string

// StatusType values
const (
	StatusTypeSuccess StatusType = "success"
	StatusTypeFail    StatusType = "fail"
	StatusTypeUnknown StatusType = "unknown"
)

// ProjectResponse ...
type ProjectResponse struct {
	Project        *Project       `json:"data"`
	ResponseStatus ResponseStatus `json:"status"`
}

// BuyProjectResponse represents the response from Buy Project request
type BuyProjectResponse struct {
	List           []*BuyProjectLineItem `json:"data"`
	ResponseStatus ResponseStatus        `json:"status"`
}

// GetAllProjectsResponse ...
type GetAllProjectsResponse struct {
	Projects       []*ProjectHeader `json:"data"`
	ResponseStatus ResponseStatus   `json:"status"`
	Meta           Meta             `json:"meta"`
}

// ProjectReportResponse ...
type ProjectReportResponse struct {
	Report         *ProjectReport `json:"data"`
	ResponseStatus ResponseStatus `json:"status"`
}

// CloseProjectResponse ...
type CloseProjectResponse struct {
	Project *struct {
		ProjectHeader
		LineItems []*LineItemHeader `json:"lineItems"`
	} `json:"data"`
	ResponseStatus ResponseStatus `json:"status"`
}

// LineItemResponse ... Response returned by Add, Update and Get LineItem requests
type LineItemResponse struct {
	Item           *LineItem      `json:"data"`
	ResponseStatus ResponseStatus `json:"status"`
}

// UpdateLineItemStateResponse ...
type UpdateLineItemStateResponse struct {
	LineItem       *LineItem      `json:"data"`
	ResponseStatus ResponseStatus `json:"status"`
}

// LineItemListItem ...
type LineItemListItem struct {
	Model
	ExtLineItemID   string      `json:"extLineItemId"`
	State           State       `json:"state"`
	StateReason     string      `json:"stateReason" conform:"trim"`
	LaunchedAt      *CustomTime `json:"launchedAt"`
	Title           string      `json:"title" conform:"trim"`
	CountryISOCode  string      `json:"countryISOCode" conform:"trim"`
	LanguageISOCode string      `json:"languageISOCode" conform:"trim"`
}

// GetAllLineItemsResponse ...
type GetAllLineItemsResponse struct {
	List           []*LineItemListItem `json:"data"`
	ResponseStatus ResponseStatus      `json:"status"`
	Meta           Meta                `json:"meta"`
}

// GetFeasibilityResponse ...
type GetFeasibilityResponse struct {
	List []*struct {
		ExtLineItemID string       `json:"extLineItemId"`
		Feasibility   *Feasibility `json:"feasibility"`
	} `json:"data"`
	ResponseStatus ResponseStatus `json:"status"`
}

// GetCountriesResponse ...
type GetCountriesResponse struct {
	List           []*Country     `json:"data"`
	ResponseStatus ResponseStatus `json:"status"`
	Meta           Meta           `json:"meta"`
}

// GetAttributesResponse ...
type GetAttributesResponse struct {
	List           []*Attribute   `json:"data"`
	ResponseStatus ResponseStatus `json:"status"`
	Meta           Meta           `json:"meta"`
}

// GetSurveyTopicsResponse ...
type GetSurveyTopicsResponse struct {
	List           []*SurveyTopic `json:"data"`
	ResponseStatus ResponseStatus `json:"status"`
	Meta           Meta           `json:"meta"`
}

// GetEventListResponse ...
type GetEventListResponse struct {
	List           []*Event       `json:"data"`
	ResponseStatus ResponseStatus `json:"status"`
	Meta           Meta           `json:"meta"`
}

// GetEventResponse ...
type GetEventResponse struct {
	Event          *Event         `json:"data"`
	ResponseStatus ResponseStatus `json:"status"`
}

// ResponseStatus is the custom status part in API response. (Optional in some endpoints)
type ResponseStatus struct {
	Message string      `json:"message" conform:"trim"`
	Errors  []ErrorInfo `json:"errors"`
}

// ErrorInfo ... Custom API errors
type ErrorInfo struct {
	Code    string `json:"code" conform:"trim"`
	Message string `json:"message" conform:"trim"`
}

// Meta ...
type Meta struct {
	Links    `json:"links"`
	Total    int64 `json:"total"`
	PageSize int64 `json:"pageSize"`
}

// Links for page navigation
type Links struct {
	First string `json:"first" conform:"trim"`
	Last  string `json:"last" conform:"trim"`
	Next  string `json:"next" conform:"trim"`
	Prev  string `json:"prev" conform:"trim"`
	Self  string `json:"self" conform:"trim"`
}

// Get ... Reads "message" from API's custom success/error response and interprets the status
func (s *ResponseStatus) Get() StatusType {
	switch s.Message {
	case "success":
		return StatusTypeSuccess
	case "fail":
		return StatusTypeFail
	default:
		return StatusTypeUnknown
	}
}
