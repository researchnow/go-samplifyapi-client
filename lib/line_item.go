package samplify

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
	FeasibilityStatusReady      FeasibilityStatus = "READY"
	FeasibilityStatusProcessing FeasibilityStatus = "PROCESSING"
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
	AttributeID string   `json:"attributeId"`
	Options     []string `json:"options"`
}

// QuotaGroup ...
type QuotaGroup struct {
	Name       *string      `json:"name"`
	QuotaCells []*QuotaCell `json:"quotaCells"`
}

// QuotaCell ...
type QuotaCell struct {
	QuotaNodes []*QuotaNode `json:"quotaNodes"`
	Perc       float64      `json:"perc"`
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

// LineItem ...
type LineItem struct {
	LineItemHeader
	Title               string       `json:"title"`
	CountryISOCode      string       `json:"countryISOCode"`
	LanguageISOCode     string       `json:"languageISOCode"`
	SurveyURL           string       `json:"surveyURL"`
	SurveyTestURL       string       `json:"surveyTestURL"`
	IndicativeIncidence float64      `json:"indicativeIncidence"`
	DaysInField         int64        `json:"daysInField"`
	LengthOfInterview   int64        `json:"lengthOfInterview"`
	DeliveryType        DeliveryType `json:"deliveryType"`
	RequiredCompletes   int64        `json:"requiredCompletes"`
	QuotaPlan           *QuotaPlan   `json:"quotaPlan"`
	EndLinks            *EndLinks    `json:"endLinks"`
}

// `CreateL`ineItemCriteria has the fields to create a LineItem
type CreateLineItemCriteria struct {
	ExtLineItemID       string        `json:"extLineItemId" valid:"required"`
	Title               string        `json:"title" valid:"required"`
	CountryISOCode      string        `json:"countryISOCode" valid:"required,ISO3166Alpha2"`
	LanguageISOCode     string        `json:"languageISOCode" valid:"required,languageISOCode"`
	SurveyURL           *string       `json:"surveyURL,omitempty" valid:"optional,surveyURL"`
	SurveyTestURL       *string       `json:"surveyTestURL,omitempty" valid:"optional,surveyURL"`
	IndicativeIncidence float64       `json:"indicativeIncidence" valid:"required"`
	DaysInField         int64         `json:"daysInField" valid:"required"`
	LengthOfInterview   int64         `json:"lengthOfInterview" valid:"required"`
	DeliveryType        *DeliveryType `json:"deliveryType" valid:"optional,DeliveryType"`
	RequiredCompletes   int64         `json:"requiredCompletes" valid:"required"`
	QuotaPlan           *QuotaPlan    `json:"quotaPlan" valid:"optional,quotaPlan"`
}

// UpdateLineItemCriteria has the fields to update a LineItem
type UpdateLineItemCriteria struct {
	ExtLineItemID       string        `json:"extLineItemId"`
	Title               *string       `json:"title,omitempty" valid:"optional"`
	CountryISOCode      *string       `json:"countryISOCode,omitempty" valid:"optional,ISO3166Alpha2"`
	LanguageISOCode     *string       `json:"languageISOCode,omitempty" valid:"optional,languageISOCode"`
	SurveyURL           *string       `json:"surveyURL,omitempty" valid:"optional,surveyURL"`
	SurveyTestURL       *string       `json:"surveyTestURL,omitempty" valid:"optional,surveyURL"`
	IndicativeIncidence *float64      `json:"indicativeIncidence,omitempty" valid:"optional"`
	DaysInField         *int64        `json:"daysInField,omitempty" valid:"optional"`
	LengthOfInterview   *int64        `json:"lengthOfInterview,omitempty" valid:"optional"`
	DeliveryType        *DeliveryType `json:"deliveryType" valid:"optional,DeliveryType"`
	RequiredCompletes   *int64        `json:"requiredCompletes,omitempty" valid:"optional"`
	QuotaPlan           *QuotaPlan    `json:"quotaPlan,omitempty" valid:"optional,quotaPlan"`
}

// BuyProjectLineItem ...
type BuyProjectLineItem struct {
	ExtLineItemID string `json:"extLineItemId"`
	State         State  `json:"state"`
}

// LineItemReport ...
type LineItemReport struct {
	ExtLineItemID      string  `json:"extLineItemId"`
	State              State   `json:"state"`
	Attempts           int64   `json:"attempts"`
	Completes          int64   `json:"completes"`
	Overquotas         int64   `json:"overquotas"`
	Screenouts         int64   `json:"screenouts"`
	Starts             int64   `json:"starts"`
	Conversion         float64 `json:"conversion"`
	CurrencyCode       string  `json:"currencyCode"`
	RemainingCompletes int64   `json:"remainingCompletes"`
	ActualMedianLOI    int64   `json:"actualMedianLOI"`
	IncurredCost       float64 `json:"incurredCost"`
	EstimatedCost      float64 `json:"estimatedCost"`
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
}

// AttributeOption ...
type AttributeOption struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}
