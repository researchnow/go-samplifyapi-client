package samplify

import "os"

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

func (s State) String() string {
	return string(s)
}

// State values for Projects and LineItems
const (
	StateProvisioned State = "PROVISIONED"
	StateLaunched    State = "LAUNCHED"
	StatePaused      State = "PAUSED"
	StateClosed      State = "CLOSED"
	StateCompleted   State = "COMPLETED"

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
	ExtProjectID string      `json:"extProjectId"`
	Title        string      `json:"title"`
	JobNumber    string      `json:"jobNumber"`
	State        State       `json:"state"`
	Author       *Author     `json:"author"`
	Billing      *Billing    `json:"billing"`
	LaunchedAt   *CustomTime `json:"launchedAt"`
	ClosedAt     *CustomTime `json:"closedAt"` 
}

// Billing ...
type Billing struct {
	ID   string      `json:"billingID"`
	Type BillingType `json:"type"`
	Date CustomTime  `json:"billingDate"`
}

// BillingType determines whether the invoiced project is monthly or single project
type BillingType string

const (
	// BillingTypeMonthly determines that the projects are being billed Monthly
	BillingTypeMonthly BillingType = "AGGREGATED_MONTHLY"
	// BillingTypePerProject determines the project is billed per project
	BillingTypePerProject BillingType = "PER_PROJECT"
)

// Project ...
type Project struct {
	ProjectHeader
	NotificationEmails []string     `json:"notificationEmails"`
	Devices            []DeviceType `json:"devices"`
	Category           *Category    `json:"category"`
	LineItems          []*LineItem  `json:"lineItems"`
	Exclusions         *Exclusions  `json:"exclusions"`
	Invoice            os.File      `json:"invoice"`
}

// CreateProjectCriteria has the fields to create a project
type CreateProjectCriteria struct {
	ExtProjectID       string                    `json:"extProjectId" valid:"required"`
	Title              string                    `json:"title" valid:"required"`
	NotificationEmails []string                  `json:"notificationEmails" valid:"email,required"`
	JobNumber          string                    `json:"jobNumber,omitempty" valid:"optional"`
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
	JobNumber          *string                    `json:"jobNumber,omitempty" valid:"optional"`
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
	JobNumber          string            `json:"jobNumber"`
	State              State             `json:"state"`
	Attempts           int64             `json:"attempts"`
	Completes          int64             `json:"completes"`
	Screenouts         int64             `json:"screenouts"`
	Overquotas         int64             `json:"overquotas"`
	Incompletes        int64             `json:"incompletes"`
	Conversion         float64           `json:"conversion"`
	CurrencyCode       string            `json:"currency"`
	RemainingCompletes int64             `json:"remainingCompletes"`
	ActualMedianLOI    int64             `json:"actualMedianLOI"`
	IncurredCost       float64           `json:"incurredCost"`
	EstimatedCost      float64           `json:"estimatedCost"`
	LineItems          []*LineItemReport `json:"lineItems"`
	CompletesRefused   int64             `json:"completesRefused"`
}

// DetailedProjectReport ...
type DetailedProjectReport struct {
	ExtProjectID string                    `json:"extProjectId"`
	JobNumber    string                    `json:"jobNumber"`
	Title        string                    `json:"title"`
	State        State                     `json:"state"`
	Stats        DetailedStats             `json:"stats"`
	Cost         Cost                      `json:"cost"`
	LineItems    []*DetailedLineItemReport `json:"lineItems"`
}

// DetailedStats ...
type DetailedStats struct {
	Attempts                   int64   `json:"attempts"`
	Completes                  int64   `json:"completes"`
	CompletesRefused           int64   `json:"completesRefused"`
	CompletesRefusedPercentage float64 `json:"completesRefusedPercentage"`
	Screenouts                 int64   `json:"screenouts"`
	ScreenoutsPercentage       float64 `json:"screenoutsPercentage"`
	Overquotas                 int64   `json:"overquotas"`
	OverquotasPercentage       float64 `json:"overquotasPercentage"`
	Incompletes                int64   `json:"incompletes"`
	IncompletesPercentage      float64 `json:"incompletesPercentage"`
	IncidenceRate              float64 `json:"incidenceRate"`
	// LastAcceptedIncidenceRate, LastAcceptedLOI and  ActualMedianLOI are applicable for lineitem stats, not applicable for quota group level or quota cell level or project level stats.
	LastAcceptedIncidenceRate float64 `json:"lastAcceptedIncidenceRate,omitempty"`
	LastAcceptedLOI           float64 `json:"lastAcceptedLOI,omitempty"`
	RemainingCompletes        int64   `json:"remainingCompletes"`
	ActualMedianLOI           int64   `json:"actualMedianLOI,omitempty"`
	Conversion                float64 `json:"conversion"`
}

// Invoice ... Represents Invoice for a project.
type Invoice struct {
	File []byte `json:"data"`
}

// Reconcile ... Represents Request correction file
type Reconcile struct {
	File        []byte `json:"data"`
	Description string `json:"description"`
}

// SurveyTopic ... Represents Survey Topic for a project. Required to setup a project
type SurveyTopic struct {
	Topic       string `json:"topic"`
	Description string `json:"description"`
}
