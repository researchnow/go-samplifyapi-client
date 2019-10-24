package samplify

const (
	// TierStandard is the constant defined for Standard Tier in Attribute
	TierStandard = "Standard"
)

// Common reasons for rejection.
const (
	WRONG_LANGUAGE Reason = "Wrong Language"
	INCORRECT_DEVICE_TYPE Reason = "Incorrect Device Type"
	TARGETING_UNDER_18 Reason = "Targeting under 18"
	INCORRECT_CATEGORY Reason = "Incorrect Category"
	SENSITIVE_CONTENT Reason = "Sensitive Content"
	ADULT_CONTENT Reason = "Adult Content"
	EXPLICIT_CONTENT Reason = "Explicit Content"
	PII_COLLECTION Reason = "PII Collection"
	SWEEPTAKES_INCENTIVES Reason = "Sweepstakes/Incentives"
	DOWNLOADS_WEBSITE Reason = "Downloads/Website"
	TESTING Reason = "Testing"
	RESEARCH_ON_RESEARCH Reason = "Research on Research"
	INCORRECT_EXIT_LINKS Reason = "Incorrect Exit Links"
)

// Type for the rejection.
const(
	PRE RejectionType = "Pre"
	POST RejectionType = "Post"
)

type Reason string
type RejectionType string

// All rejection resolution information together.
type RejectionResolution struct{
	ReasonTitle       Reason
	ReasonDescription string
	Comments          []string
	RejectionType     RejectionType
}

// Action ...
type Action string

// Action values for changing LineItem state
const (
	ActionLaunched Action = "launch"
	ActionPaused   Action = "pause"
	ActionClosed   Action = "close"
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

// DeliveryType ...
type DeliveryType string

// DeliveryType values
const (
	DeliveryTypeSlow     DeliveryType = "SLOW"
	DeliveryTypeBalanced DeliveryType = "BALANCED"
	DeliveryTypeFast     DeliveryType = "FAST"
)

//QuotaPlan ...
type QuotaPlan struct {
	Filters     []*QuotaFilters `json:"filters,omitempty" valid:"optional"`
	QuotaGroups []*QuotaGroup   `json:"quotaGroups,omitempty" valid:"optional"`
}

// QuotaFilters ...
type QuotaFilters struct {
	AttributeID string   `json:"attributeId" conform:"trim"`
	Options     []string `json:"options" conform:"trim"`
}

// QuotaGroup ...
type QuotaGroup struct {
	Name       *string      `json:"name" conform:"trim"`
	QuotaCells []*QuotaCell `json:"quotaCells"`
}

// QuotaCell ...
type QuotaCell struct {
	QuotaNodes []*QuotaNode `json:"quotaNodes"`
	Perc       *float64     `json:"perc,omitempty" valid:"optional"`
	Count      *uint32      `json:"count,omitempty" valid:"optional"`
}

// QuotaNode ...
type QuotaNode struct {
	AttributeID string   `json:"attributeId" conform:"trim"`
	Options     []string `json:"options"`
}

// EndLinks ...
type EndLinks struct {
	Complete      string `json:"complete" conform:"trim"`
	Screenout     string `json:"screenout" conform:"trim"`
	OverQuota     string `json:"overquota" conform:"trim"`
	SecurityKey1  string `json:"securityKey1" conform:"trim"`
	SecurityKey2  string `json:"securityKey2" conform:"trim"`
	SecurityLevel string `json:"securityLevel" conform:"trim"`
}

// LineItemHeader ...
type LineItemHeader struct {
	Model
	ExtLineItemID string      `json:"extLineItemId" conform:"trim"`
	State         State       `json:"state"`
	StateReason   string      `json:"stateReason" conform:"trim"`
	LaunchedAt    *CustomTime `json:"launchedAt"`
}

// LineItem ...
type LineItem struct {
	LineItemHeader
	Title               string            `json:"title" conform:"trim"`
	CountryISOCode      string            `json:"countryISOCode" conform:"trim"`
	LanguageISOCode     string            `json:"languageISOCode" conform:"trim"`
	SurveyURL           string            `json:"surveyURL" conform:"trim"`
	SurveyTestURL       string            `json:"surveyTestURL" conform:"trim"`
	IndicativeIncidence float64           `json:"indicativeIncidence"`
	DaysInField         int64             `json:"daysInField"`
	LengthOfInterview   int64             `json:"lengthOfInterview"`
	DeliveryType        DeliveryType      `json:"deliveryType"`
	RequiredCompletes   int64             `json:"requiredCompletes"`
	QuotaPlan           *QuotaPlan        `json:"quotaPlan"`
	EndLinks            *EndLinks         `json:"endLinks"`
	SurveyUrlParams     []*URLParameter   `json:"surveyURLParams"`
	Sources             []*LineItemSource `json:"sources"`
	RejectionResolution RejectionResolution   `json:"rejection"`
}

// IsUpdateable returns false if the line item cannot be updated.
func (l *LineItem) IsUpdateable() bool {
	if l.State == StateProvisioned ||
		l.State == StateRejected {
		return true
	}
	return false
}

// IsBuyable returns true if the lineitem can be bought or not
func (l *LineItem) IsBuyable() bool {
	if l.State == StateProvisioned ||
		l.State == StateRejected ||
		l.State == StateRejectedPaused {
		return true
	}
	return false
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
	ExtLineItemID       string            `json:"extLineItemId" valid:"required" conform:"trim"`
	Title               string            `json:"title" valid:"required" conform:"trim"`
	CountryISOCode      string            `json:"countryISOCode" valid:"required,ISO3166Alpha2" conform:"trim"`
	LanguageISOCode     string            `json:"languageISOCode" valid:"required,languageISOCode" conform:"trim"`
	SurveyURL           *string           `json:"surveyURL,omitempty" valid:"optional,surveyURL" conform:"trim"`
	SurveyTestURL       *string           `json:"surveyTestURL,omitempty" valid:"optional" conform:"trim"`
	IndicativeIncidence float64           `json:"indicativeIncidence" valid:"required"`
	DaysInField         int64             `json:"daysInField" valid:"required"`
	LengthOfInterview   int64             `json:"lengthOfInterview" valid:"required"`
	DeliveryType        *DeliveryType     `json:"deliveryType" valid:"optional,DeliveryType"`
	RequiredCompletes   int64             `json:"requiredCompletes" valid:"required"`
	QuotaPlan           *QuotaPlan        `json:"quotaPlan,omitempty" valid:"optional,quotaPlan"`
	SurveyUrlParams     []*URLParameter   `json:"surveyURLParams" valid:"optional"`
	SurveyTestUrlParams []*URLParameter   `json:"surveyTestURLParams" valid:"optional"`
	Sources             []*LineItemSource `json:"sources,omitempty" valid:"optional"`
}

// UpdateLineItemCriteria has the fields to update a LineItem
type UpdateLineItemCriteria struct {
	ExtLineItemID       string             `json:"extLineItemId" conform:"trim"`
	Title               *string            `json:"title,omitempty" valid:"optional" conform:"trim"`
	CountryISOCode      *string            `json:"countryISOCode,omitempty" valid:"optional,ISO3166Alpha2" conform:"trim"`
	LanguageISOCode     *string            `json:"languageISOCode,omitempty" valid:"optional,languageISOCode" conform:"trim"`
	SurveyURL           *string            `json:"surveyURL,omitempty" valid:"optional,surveyURL" conform:"trim"`
	SurveyTestURL       *string            `json:"surveyTestURL,omitempty" valid:"optional" conform:"trim"`
	IndicativeIncidence *float64           `json:"indicativeIncidence,omitempty" valid:"optional"`
	DaysInField         *int64             `json:"daysInField,omitempty" valid:"optional"`
	LengthOfInterview   *int64             `json:"lengthOfInterview,omitempty" valid:"optional"`
	DeliveryType        *DeliveryType      `json:"deliveryType" valid:"optional,DeliveryType"`
	RequiredCompletes   *int64             `json:"requiredCompletes,omitempty" valid:"optional"`
	QuotaPlan           *QuotaPlan         `json:"quotaPlan,omitempty" valid:"optional,quotaPlan"`
	SurveyUrlParams     []*URLParameter    `json:"surveyURLParams" valid:"optional"`
	SurveyTestUrlParams []*URLParameter    `json:"surveyTestURLParams" valid:"optional"`
	Sources             *[]*LineItemSource `json:"sources,omitempty" valid:"optional"`
}

// BuyProjectLineItem ...
type BuyProjectLineItem struct {
	ExtLineItemID string `json:"extLineItemId" conform:"trim"`
	State         State  `json:"state"`
}

// LineItemReport ...
type LineItemReport struct {
	ExtLineItemID         string  `json:"extLineItemId" conform:"trim"`
	Title                 string  `json:"title" conform:"trim"`
	CountryISOCode        string  `json:"countryISOCode" conform:"trim"`
	LanguageISOCode       string  `json:"languageISOCode" conform:"trim"`
	State                 State   `json:"state"`
	StateReason           string  `json:"stateReason" conform:"trim"`
	Attempts              int64   `json:"attempts"`
	Completes             int64   `json:"completes"`
	Overquotas            int64   `json:"overquotas"`
	Screenouts            int64   `json:"screenouts"`
	Incompletes           int64   `json:"incompletes"`
	Conversion            float64 `json:"conversion"`
	CurrencyCode          string  `json:"currency" conform:"trim"`
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
	ID                 string             `json:"id" conform:"trim"`
	Name               string             `json:"name" conform:"trim"`
	Text               string             `json:"text" conform:"trim"`
	IsAllowedInFilters bool               `json:"isAllowedInFilters"`
	IsAllowedInQuotas  bool               `json:"isAllowedInQuotas"`
	Type               string             `json:"type"`
	Options            []*AttributeOption `json:"options"`
	Format             *string            `json:"format,omitempty" conform:"trim"`
	LocalizedText      *string            `json:"localizedText,omitempty" conform:"trim"`
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
	ID            string  `json:"id" conform:"trim"`
	Text          string  `json:"text" conform:"trim"`
	LocalizedText *string `json:"localizedText,omitempty" conform:"trim"`
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
	ID            string  `json:"id" conform:"trim"`
	Text          string  `json:"text" conform:"trim"`
	LocalizedText *string `json:"localizedText,omitempty" conform:"trim"`
}

// URLParameter ...
type URLParameter struct {
	Key    string   `json:"key" conform:"trim"`
	Values []string `json:"values" conform:"trim"`
}

// LineItemSource source associated with the lineitem.
type LineItemSource struct {
	ID int64 `json:"id"`
}
