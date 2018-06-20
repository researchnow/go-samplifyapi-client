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

	StateAwaitingApproval State = "AWAITING_APPROVAL"
	StateInvoiced         State = "INVOICED"
)

// Category is a Project's category
type Category struct {
	SurveyTopic []string `json:"surveyTopic"`
}

// Exclusions ... Project's exclusions
type Exclusions struct {
	Type ExclusionType `json:"type"`
	List []string      `json:"list"`
}

// ProjectHeader ...
type ProjectHeader struct {
	Model
	ExtProjectID string `json:"extProjectId"`
	Title        string `json:"title"`
	State        State  `json:"state"`
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

// CreateProjectCriteria holds fields to create a project.
type CreateProjectCriteria struct {
	ExtProjectID       string                    `json:"extProjectId"`
	Title              string                    `json:"title"`
	NotificationEmails []string                  `json:"notificationEmails"`
	Devices            []DeviceType              `json:"devices"`
	Category           *Category                 `json:"category"`
	LineItems          []*CreateLineItemCriteria `json:"lineItems"`
	Exclusions         *Exclusions               `json:"exclusions"`
}

// UpdateProjectCriteria has the fields to update a project.
// Excluding ExtProjectID, its fields are optional.
type UpdateProjectCriteria struct {
	ExtProjectID       string                     `json:"extProjectId"`
	Title              *string                    `json:"title,omitempty"`
	NotificationEmails *[]string                  `json:"notificationEmails,omitempty"`
	Devices            *[]DeviceType              `json:"devices,omitempty"`
	Category           *Category                  `json:"category,omitempty"`
	LineItems          *[]*CreateLineItemCriteria `json:"lineItems,omitempty"`
	Exclusions         *Exclusions                `json:"exclusions,omitempty"`
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

// SurveyTopic ... Represents Survey Topic for a project. Required to setup a project
type SurveyTopic struct {
	Topic       string `json:"topic"`
	Description string `json:"description"`
}
