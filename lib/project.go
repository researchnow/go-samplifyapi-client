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

	StateAwaitingApproval       State = "AWAITING_APPROVAL"
	StateInvoiced               State = "INVOICED"
	StateQAApproved             State = "QA_APPROVED"
	StateRejected               State = "REJECTED"
	StateCancelled              State = "CANCELLED"
	StateAwaitingApprovalPaused State = "AWAITING_APPROVAL_PAUSED"
	StateAwaitingClientApproval State = "AWAITING_CLIENT_APPROVAL"
	StateRejectedPaused         State = "REJECTED_PAUSED"
)

// Category is a Project's category
type Category struct {
	SurveyTopic []string `json:"surveyTopic" valid:"required"`
}

// Exclusions ... Project's exclusions
type Exclusions struct {
	Type ExclusionType `json:"type" valid:"ExclusionType"`
	List []string      `json:"list"`
}

// Author ...
type Author struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Username string `json:"username"`
}

// ProjectHeader ...
type ProjectHeader struct {
	Model
	ExtProjectID string  `json:"extProjectId"`
	Title        string  `json:"title"`
	State        State   `json:"state"`
	Author       *Author `json:"author"`
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

// CreateProjectCriteria has the fields to create a project
type CreateProjectCriteria struct {
	ExtProjectID       string                    `json:"extProjectId" valid:"required"`
	Title              string                    `json:"title" valid:"required"`
	NotificationEmails []string                  `json:"notificationEmails" valid:"email,required"`
	Devices            []DeviceType              `json:"devices" valid:"optional,DeviceType"`
	Category           *Category                 `json:"category" valid:"required"`
	LineItems          []*CreateLineItemCriteria `json:"lineItems" valid:"required"`
	Exclusions         *Exclusions               `json:"exclusions,omitempty" valid:"optional"`
}

// UpdateProjectCriteria has the fields to update a project
type UpdateProjectCriteria struct {
	ExtProjectID       string                     `json:"extProjectId" valid:"required"`
	Title              *string                    `json:"title,omitempty" valid:"optional"`
	NotificationEmails *[]string                  `json:"notificationEmails,omitempty" valid:"email,optional"`
	Devices            *[]DeviceType              `json:"devices,omitempty" valid:"DeviceType,optional"`
	Category           *Category                  `json:"category,omitempty" valid:"optional"`
	LineItems          *[]*UpdateLineItemCriteria `json:"lineItems,omitempty" valid:"optional"`
	Exclusions         *Exclusions                `json:"exclusions,omitempty" valid:"optional"`
}

// BuyProjectCriteria ...
type BuyProjectCriteria struct {
	ExtLineItemID string `json:"extLineItemId" valid:"required"`
	SurveyURL     string `json:"surveyURL" valid:"required,surveyURL"`
	SurveyTestURL string `json:"surveyTestURL" valid:"required"`
}

// ProjectReport ...
type ProjectReport struct {
	ExtProjectID       string            `json:"extProjectId"`
	Title              string            `json:"title"`
	State              State             `json:"state"`
	Attempts           int64             `json:"attempts"`
	Completes          int64             `json:"completes"`
	Screenouts         int64             `json:"screenouts"`
	Overquotas         int64             `json:"overquotas"`
	Starts             int64             `json:"starts"`
	Conversion         float64           `json:"conversion"`
	CurrencyCode       string            `json:"currency"`
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
