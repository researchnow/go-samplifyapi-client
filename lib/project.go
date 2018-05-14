package samplify

// DeviceType ...
type DeviceType string

// DeviceType values
const (
	DeviceTypeMobile  DeviceType = "mobile"
	DeviceTypeDesktop DeviceType = "desktop"
	DeviceTypeTablet  DeviceType = "tablet"
)

// ExclusionType ...
type ExclusionType string

// ExclusionType values
const (
	ExclusionTypeProject ExclusionType = "PROJECT"
	ExclusionTypeTag     ExclusionType = "TAG"
)

// State ...
type State string

// State values for Projects and LineItems
const (
	StateProvisioned State = "PROVISIONED"
	StateLaunched    State = "LAUNCHED"
	StatePaused      State = "PAUSED"
	StateClosed      State = "CLOSED"

	SateAwaitingApproval State = "AWAITING_APPROVAL"
	StateInvoiced        State = "INVOICED"
)

// Category is Project's Category
type Category struct {
	SurveyTopic []string `json:"surveyTopic"`
}

// Exclusions ... Samplify project's exclusions
type Exclusions struct {
	Type ExclusionType `json:"type"`
	List []string      `json:"list"`
}

// ProjectHeader is a Samplify project header
type ProjectHeader struct {
	ExtProjectID       string      `json:"extProjectId"`
	Title              string      `json:"title"`
	State              State       `json:"state"`
	StateLastUpdatedAt *CustomTime `json:"stateLastUpdatedAt"`
	CreatedAt          CustomTime  `json:"createdAt"`
	UpdatedAt          CustomTime  `json:"updatedAt"`
}

// Project ...
type Project struct {
	ProjectHeader
	NotificationEmails []string     `json:"notificationEmails"`
	Devices            []DeviceType `json:"devices"`
	Category           *Category    `json:"category"`
	LineItems          []*LineItem  `json:"lineItems"`
	Exclusions         *Exclusions  `json:"exclusions"`
}

// CreateUpdateProjectCriteria has the fields to create or update a project
type CreateUpdateProjectCriteria struct {
	ExtProjectID       string       `json:"extProjectId"`
	Title              string       `json:"title"`
	NotificationEmails []string     `json:"notificationEmails"`
	Devices            []DeviceType `json:"devices"`
	Category           *Category    `json:"category"`
	LineItems          []*LineItem  `json:"lineItems"`
	Exclusions         *Exclusions  `json:"exclusions"`
}

// BuyProjectCriteria ...
type BuyProjectCriteria struct {
	ExtLineItemID string `json:"extLineItemId"`
	SurveyURL     string `json:"surveyURL"`
	SurveyTestURL string `json:"surveyTestURL"`
}

// ProjectReport ...
type ProjectReport struct {
	ExtProjectID       string            `json:"extProjectId"`
	State              State             `json:"state"`
	Attempts           int64             `json:"attempts"`
	Completes          int64             `json:"completes"`
	Screenouts         int64             `json:"screenouts"`
	Overquotas         int64             `json:"overquotas"`
	Starts             int64             `json:"starts"`
	Conversion         int64             `json:"conversion"`
	RemainingCompletes int64             `json:"remainingCompletes"`
	ActualMedianLOI    int64             `json:"actualMedianLOI"`
	IncurredCost       float64           `json:"incurredCost"`
	EstimatedCost      float64           `json:"estimatedCost"`
	LineItems          []*LineItemReport `json:"lineItems"`
}

// SurveyTopic ... Represents survey Topic for a project. Required to setup a project
type SurveyTopic struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
