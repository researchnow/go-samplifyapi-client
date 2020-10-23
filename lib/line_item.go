package samplify

import (
	"strings"
)

const (
	// TierStandard is the constant defined for Standard Tier in Attribute
	TierStandard = "Standard"
)

// Action ...
type Action string

// QCellStatusType ...
type QCellStatusType string

// Action values for changing LineItem state
const (
	ActionLaunched        Action          = "launch"
	ActionPaused          Action          = "pause"
	ActionClosed          Action          = "close"
	QCellStatusTypeLaunch QCellStatusType = "LAUNCHED"
	QCellStatusTypePause  QCellStatusType = "PAUSED"
)

// FeasibilityStatus ...
type FeasibilityStatus string

// FeasibilityStatus values
const (
	FeasibilityStatusReady        FeasibilityStatus = "READY"
	FeasibilityStatusProcessing   FeasibilityStatus = "PROCESSING"
	FeasibilityStatusNotSupported FeasibilityStatus = "NOT_SUPPORTED"
	FeasibilityStatusFailed       FeasibilityStatus = "FAILED"
)

// Operator operator for the filters.
type Operator string

const (
	// OperatorInclude to include all the respondents that match the given options
	OperatorInclude Operator = "include"
	// OperatorExclude to exclude all the respondents that match the given options
	OperatorExclude Operator = "exclude"
)

// CostType ...
type CostType string

// CostType for the different types of cost
const (
	CostTypeBase    CostType = "BASE"
	CostTypePremium CostType = "PREMIUM"
)

// ToUpper converts the operator to upper case.
func (o Operator) ToUpper() string {
	return strings.ToUpper(string(o))
}

//QuotaPlan ...
type QuotaPlan struct {
	Filters     []*QuotaFilters `json:"filters,omitempty" valid:"optional"`
	QuotaGroups []*QuotaGroup   `json:"quotaGroups,omitempty" valid:"optional"`
}

// QuotaFilters ...
type QuotaFilters struct {
	AttributeID string    `json:"attributeId,omitempty"`
	Options     []string  `json:"options,omitempty"`
	Operator    *Operator `json:"operator"`
}

// QuotaGroup ...
type QuotaGroup struct {
	QuotaGroupID *string      `json:"quotaGroupId,omitempty" valid:"optional"`
	Name         *string      `json:"name"`
	QuotaCells   []*QuotaCell `json:"quotaCells"`
}

// QuotaCell ...
type QuotaCell struct {
	QuotaCellID *string          `json:"quotaCellId,omitempty" valid:"optional"`
	QuotaNodes  []*QuotaNode     `json:"quotaNodes"`
	Perc        *float64         `json:"perc,omitempty" valid:"optional"`
	Count       *uint32          `json:"count,omitempty" valid:"optional"`
	Status      *QCellStatusType `json:"status,omitempty" valid:"optional"`
}

// QuotaNode ...
type QuotaNode struct {
	AttributeID string   `json:"attributeId"`
	Options     []string `json:"options"`
}

// EndLinks ...
type EndLinks struct {
	Complete      string `json:"complete"`
	Screenout     string `json:"screenout"`
	OverQuota     string `json:"overquota"`
	SecurityKey1  string `json:"securityKey1"`
	SecurityKey2  string `json:"securityKey2"`
	SecurityLevel string `json:"securityLevel"`
}

// LineItemHeader ...
type LineItemHeader struct {
	Model
	ExtLineItemID string      `json:"extLineItemId"`
	State         State       `json:"state"`
	StateReason   string      `json:"stateReason"`
	LaunchedAt    *CustomTime `json:"launchedAt"`
}

type Schedule struct {
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

// LineItem ...
type LineItem struct {
	LineItemHeader
	Title               string            `json:"title"`
	CountryISOCode      string            `json:"countryISOCode"`
	LanguageISOCode     string            `json:"languageISOCode"`
	SurveyURL           string            `json:"surveyURL"`
	SurveyTestURL       string            `json:"surveyTestURL"`
	IndicativeIncidence float64           `json:"indicativeIncidence"`
	DaysInField         int64             `json:"daysInField"`
	FieldSchedule       *Schedule         `json:"fieldSchedule"`
	LengthOfInterview   int64             `json:"lengthOfInterview"`
	DeliveryType        *string           `json:"deliveryType"`
	RequiredCompletes   int64             `json:"requiredCompletes"`
	QuotaPlan           *QuotaPlan        `json:"quotaPlan"`
	EndLinks            *EndLinks         `json:"endLinks"`
	SurveyURLParams     []*URLParameter   `json:"surveyURLParams"`
	Sources             []*LineItemSource `json:"sources"`
	Targets             []*LineItemTarget `json:"targets"`
	SurveyTestingNotes  string            `json:"surveyTestingNotes"`
}

// IsUpdateable returns false if the line item cannot be updated.
func (l *LineItem) IsUpdateable() bool {
	return l.State == StateProvisioned ||
		l.State == StateRejected
}

// IsBuyable returns true if the lineitem can be bought or not
func (l *LineItem) IsBuyable() bool {
	return l.State == StateProvisioned ||
		l.State == StateRejected ||
		l.State == StateRejectedPaused
}

// IsRebalanceable returns false if the line item cannot be updated.
func (l *LineItem) IsRebalanceable() bool {
	if l.State == StateClosed ||
		l.State == StateCancelled ||
		l.State == StateInvoiced ||
		l.State == StateAwaitingClientApproval {
		return false
	}
	return true
}

// IsCloseable returns false if the line item cannot be updated.
func (l *LineItem) IsCloseable() bool {
	if l.State == StateClosed ||
		l.State == StateCancelled ||
		l.State == StateInvoiced {
		return false
	}
	return true
}

// CreateLineItemCriteria has the fields to create a LineItem
type CreateLineItemCriteria struct {
	ExtLineItemID       string            `json:"extLineItemId" valid:"required"`
	Title               string            `json:"title" valid:"required"`
	CountryISOCode      string            `json:"countryISOCode" valid:"required,ISO3166Alpha2"`
	LanguageISOCode     string            `json:"languageISOCode" valid:"required,languageISOCode"`
	SurveyURL           *string           `json:"surveyURL,omitempty" valid:"optional,surveyURL"`
	SurveyTestURL       *string           `json:"surveyTestURL,omitempty" valid:"optional"`
	IndicativeIncidence float64           `json:"indicativeIncidence" valid:"required"`
	DaysInField         int64             `json:"daysInField" valid:"optional"`
	FieldSchedule       *Schedule         `json:"fieldSchedule" valid:"optional"`
	LengthOfInterview   int64             `json:"lengthOfInterview" valid:"required"`
	DeliveryType        *string           `json:"deliveryType" valid:"optional"`
	RequiredCompletes   int64             `json:"requiredCompletes" valid:"required"`
	QuotaPlan           *QuotaPlan        `json:"quotaPlan,omitempty" valid:"optional,quotaPlan"`
	SurveyURLParams     []*URLParameter   `json:"surveyURLParams" valid:"optional"`
	SurveyTestURLParams []*URLParameter   `json:"surveyTestURLParams" valid:"optional"`
	Sources             []*LineItemSource `json:"sources,omitempty" valid:"optional"`
	Targets             []*LineItemTarget `json:"targets"`
	SurveyTestingNotes  *string           `json:"surveyTestingNotes,omitempty" valid:"optional"`
}

// UpdateLineItemCriteria has the fields to update a LineItem
type UpdateLineItemCriteria struct {
	ExtLineItemID       string             `json:"extLineItemId"`
	Title               *string            `json:"title,omitempty" valid:"optional"`
	CountryISOCode      *string            `json:"countryISOCode,omitempty" valid:"optional,ISO3166Alpha2"`
	LanguageISOCode     *string            `json:"languageISOCode,omitempty" valid:"optional,languageISOCode"`
	SurveyURL           *string            `json:"surveyURL,omitempty" valid:"optional,surveyURL"`
	SurveyTestURL       *string            `json:"surveyTestURL,omitempty" valid:"optional"`
	IndicativeIncidence *float64           `json:"indicativeIncidence,omitempty" valid:"optional"`
	DaysInField         *int64             `json:"daysInField,omitempty" valid:"optional"`
	FieldSchedule       *Schedule          `json:"fieldSchedule" valid:"optional"`
	LengthOfInterview   *int64             `json:"lengthOfInterview,omitempty" valid:"optional"`
	DeliveryType        *string            `json:"deliveryType" valid:"optional"`
	RequiredCompletes   *int64             `json:"requiredCompletes,omitempty" valid:"optional"`
	QuotaPlan           *QuotaPlan         `json:"quotaPlan,omitempty" valid:"optional,quotaPlan"`
	SurveyURLParams     []*URLParameter    `json:"surveyURLParams" valid:"optional"`
	SurveyTestURLParams []*URLParameter    `json:"surveyTestURLParams" valid:"optional"`
	Sources             *[]*LineItemSource `json:"sources,omitempty" valid:"optional"`
	Targets             []*LineItemTarget  `json:"targets"`
	SurveyTestingNotes  *string            `json:"surveyTestingNotes,omitempty" valid:"optional"`
}

// BuyProjectLineItem ...
type BuyProjectLineItem struct {
	ExtLineItemID string `json:"extLineItemId"`
	State         State  `json:"state"`
}

// LineItemReport ...
type LineItemReport struct {
	ExtLineItemID         string  `json:"extLineItemId"`
	Title                 string  `json:"title"`
	CountryISOCode        string  `json:"countryISOCode"`
	LanguageISOCode       string  `json:"languageISOCode"`
	State                 State   `json:"state"`
	StateReason           string  `json:"stateReason"`
	Attempts              int64   `json:"attempts"`
	Completes             int64   `json:"completes"`
	Overquotas            int64   `json:"overquotas"`
	Screenouts            int64   `json:"screenouts"`
	Incompletes           int64   `json:"incompletes"`
	Conversion            float64 `json:"conversion"`
	CurrencyCode          string  `json:"currency"`
	RemainingCompletes    int64   `json:"remainingCompletes"`
	ActualMedianLOI       int64   `json:"actualMedianLOI"`
	IncurredCost          float64 `json:"incurredCost"`
	EstimatedCost         float64 `json:"estimatedCost"`
	LastAcceptedIncidence float64 `json:"lastAcceptedIncidenceRate"`
	LastAcceptedLOI       int64   `json:"lastAcceptedLOI"`
	CompletesRefused      int64   `json:"completesRefused"`
}

// Feasibility ...
type Feasibility struct {
	Status           FeasibilityStatus `json:"status"`
	CostPerInterview float64           `json:"costPerInterview"`
	Currency         string            `json:"currency"`
	Feasible         bool              `json:"feasible"`
	TotalCount       int64             `json:"totalCount"`
	ValueCounts      []*ValueCount     `json:"valueCounts"`
}

// Quote holds the information for premium pricing
type Quote struct {
	CostPerUnit   float64         `json:"costPerUnit"`
	Currency      string          `json:"currency"`
	DetailedQuote []DetailedQuote `json:"detailedQuote"`
	EstimatedCost float64         `json:"estimatedCost"`
}

// QuoteType ...
type QuoteType string

const (
	// TypeBase ...
	TypeBase QuoteType = "BASE"
	// TypePremium ...
	TypePremium QuoteType = "PREMIUM"
)

// DetailedQuote ...
type DetailedQuote struct {
	CostPerUnit   float64   `json:"costPerUnit"`
	EstimatedCost float64   `json:"estimatedCost"`
	Title         string    `json:"title"`
	Type          QuoteType `json:"type"`
	Units         int64     `json:"units"`
}

// ValueCount ...
type ValueCount struct {
	QuotaCells []*FeasibilityQuotaCell `json:"quotaCells"`
}

// FeasibilityQuotaCell ...
type FeasibilityQuotaCell struct {
	FeasibilityCount int64        `json:"feasibilityCount"`
	QuotaNodes       []*QuotaNode `json:"quotaNodes"`
}

// Attribute ... Supported attribute for a country and language. Required to build up the Quota Plan
type Attribute struct {
	ID                 string             `json:"id"`
	Name               string             `json:"name"`
	Text               string             `json:"text"`
	IsAllowedInFilters bool               `json:"isAllowedInFilters"`
	IsAllowedInQuotas  bool               `json:"isAllowedInQuotas"`
	Type               string             `json:"type"`
	Options            []*AttributeOption `json:"options"`
	Format             *string            `json:"format,omitempty"`
	LocalizedText      *string            `json:"localizedText,omitempty"`
	State              AttributeState     `json:"state"`
	Tier               string             `json:"tier"`
	AttributeCategory  AttributeCategory  `json:"category"`
	Exclusions         []*string          `json:"exclusions,omitempty"`
}

// AttributeCategory ...
type AttributeCategory struct {
	MainCategory AttrCategory `json:"mainCategory"`
	SubCategory  AttrCategory `json:"subCategory"`
}

// AttrCategory ...
type AttrCategory struct {
	ID            string  `json:"id"`
	Text          string  `json:"text"`
	LocalizedText *string `json:"localizedText,omitempty"`
}

// AttributeState defines the state of an attribute
type AttributeState string

const (
	// StateActive is the active state for attribute
	StateActive AttributeState = "ACTIVE"
	// StateDeprecated is the deprecated state for attribute
	StateDeprecated AttributeState = "DEPRECATED"
	// StateInactive is the inactive state for attribute
	StateInactive AttributeState = "INACTIVE"
)

// AttributeOption ...
type AttributeOption struct {
	ID            string  `json:"id"`
	Text          string  `json:"text"`
	LocalizedText *string `json:"localizedText,omitempty"`
}

// URLParameter ...
type URLParameter struct {
	Key    string   `json:"key"`
	Values []string `json:"values"`
}

// LineItemSource source associated with the lineitem.
type LineItemSource struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// TargetType ...
type TargetType string

// TargetType values
const (
	TargetTypeComplete TargetType = "COMPLETE"
)

// LineItemTarget target associated with the line item.
type LineItemTarget struct {
	Count      *uint32    `json:"count,omitempty"`
	DailyLimit *uint32    `json:"dailyLimit,omitempty"`
	Type       TargetType `json:"type,omitempty"`
}

// DetailedLineItemReport ...
type DetailedLineItemReport struct {
	ExtLineItemID   string            `json:"extLineItemId"`
	Title           string            `json:"title"`
	State           State             `json:"state"`
	StateReason     string            `json:"stateReason"`
	CountryISOCode  *string           `json:"countryISOCode"`
	LanguageISOCode *string           `json:"languageISOCode"`
	Sources         []*LineItemSource `json:"sources"`
	Cost            Cost              `json:"cost"`
	Stats           DetailedStats     `json:"stats"`
	// QuotaGroups is applicable only for detailed lineitem report, not applicable for lineitems in detailed project report
	QuotaGroups []*DetailedQuotaGroupReport `json:"quotaGroups,omitempty"`
}

// DetailedQuotaGroupReport ...
type DetailedQuotaGroupReport struct {
	QuotaGroupID string                     `json:"quotaGroupId"`
	Stats        DetailedStats              `json:"stats"`
	QuotaCells   []*DetailedQuotaCellReport `json:"quotaCells"`
}

// DetailedQuotaCellReport ...
type DetailedQuotaCellReport struct {
	QuotaCellID string        `json:"quotaCellId"`
	QuotaNodes  []*QuotaNode  `json:"quotaNodes"`
	Stats       DetailedStats `json:"stats"`
}

// Cost ...
type Cost struct {
	// CostPerUnit and DetailedCost are only applicable at lineitem level, not applicable at project level
	CostPerUnit   float64         `json:"costPerUnit,omitempty"`
	Currency      string          `json:"currency"`
	EstimatedCost float64         `json:"estimatedCost"`
	IncurredCost  float64         `json:"incurredCost"`
	DetailedCost  []*DetailedCost `json:"detailedCost,omitempty"`
}

// DetailedCost ...
type DetailedCost struct {
	Title          string   `json:"title"`
	Type           CostType `json:"type"`
	CostPerUnit    float64  `json:"costPerUnit"`
	EstimatedCost  float64  `json:"estimatedCost"`
	IncurredCost   float64  `json:"incurredCost"`
	DeliveredUnits int64    `json:"deliveredUnits"`
	RequestedUnits int64    `json:"requestedUnits"`
}

// Allocation enum for allocation type.
type Allocation string

const (
	// AllocationPercentage percentage allocation
	AllocationPercentage Allocation = "perc"
	// AllocationCount count allocation
	AllocationCount Allocation = "count"
)

// AllocationType retuns the type of allocation in the quota cell.
func (c *QuotaCell) AllocationType() Allocation {
	if c.Count != nil {
		return AllocationCount
	}
	// what if both are not present
	return AllocationPercentage
}
