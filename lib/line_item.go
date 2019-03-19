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
	Perc       *float64     `json:"perc,omitempty" valid:"optional"`
	Count      *uint32      `json:"count,omitempty" valid:"optional"`
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

// IsUpdateable returns false if the line item cannot be updated.
func (l *LineItem) IsUpdateable() bool {
	if l.State == StateProvisioned ||
		l.State == StateAwaitingApproval {
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
	ExtLineItemID       	string        	`json:"extLineItemId" valid:"required"`
	Title               	string        	`json:"title" valid:"required"`
	CountryISOCode      	string        	`json:"countryISOCode" valid:"required,ISO3166Alpha2"`
	LanguageISOCode     	string        	`json:"languageISOCode" valid:"required,languageISOCode"`
	SurveyURL           	*string       	`json:"surveyURL,omitempty" valid:"optional,surveyURL"`
	SurveyTestURL       	*string       	`json:"surveyTestURL,omitempty" valid:"optional"`
	IndicativeIncidence 	float64       	`json:"indicativeIncidence" valid:"required"`
	DaysInField         	int64         	`json:"daysInField" valid:"required"`
	LengthOfInterview   	int64         	`json:"lengthOfInterview" valid:"required"`
	DeliveryType        	*DeliveryType 	`json:"deliveryType" valid:"optional,DeliveryType"`
	RequiredCompletes   	int64         	`json:"requiredCompletes" valid:"required"`
	QuotaPlan           	*QuotaPlan    	`json:"quotaPlan" valid:"optional,quotaPlan"`
	SurveyUrlParams 		[]*URLParameter	`json:"surveyURLParams" valid:"optional"`
	SurveyTestUrlParams 	[]*URLParameter	`json:"surveyTestURLParams" valid:"optional"`
}

// UpdateLineItemCriteria has the fields to update a LineItem
type UpdateLineItemCriteria struct {
	ExtLineItemID       	string        	`json:"extLineItemId"`
	Title               	*string       	`json:"title,omitempty" valid:"optional"`
	CountryISOCode      	*string       	`json:"countryISOCode,omitempty" valid:"optional,ISO3166Alpha2"`
	LanguageISOCode     	*string       	`json:"languageISOCode,omitempty" valid:"optional,languageISOCode"`
	SurveyURL           	*string       	`json:"surveyURL,omitempty" valid:"optional,surveyURL"`
	SurveyTestURL       	*string       	`json:"surveyTestURL,omitempty" valid:"optional"`
	IndicativeIncidence 	*float64      	`json:"indicativeIncidence,omitempty" valid:"optional"`
	DaysInField         	*int64        	`json:"daysInField,omitempty" valid:"optional"`
	LengthOfInterview   	*int64        	`json:"lengthOfInterview,omitempty" valid:"optional"`
	DeliveryType        	*DeliveryType 	`json:"deliveryType" valid:"optional,DeliveryType"`
	RequiredCompletes   	*int64        	`json:"requiredCompletes,omitempty" valid:"optional"`
	QuotaPlan           	*QuotaPlan    	`json:"quotaPlan,omitempty" valid:"optional,quotaPlan"`
	SurveyUrlParams 		[]*URLParameter	`json:"surveyURLParams" valid:"optional"`
	SurveyTestUrlParams 	[]*URLParameter	`json:"surveyTestURLParams" valid:"optional"`
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
	Starts                int64   `json:"starts"`
	Conversion            float64 `json:"conversion"`
	CurrencyCode          string  `json:"currency"`
	RemainingCompletes    int64   `json:"remainingCompletes"`
	ActualMedianLOI       int64   `json:"actualMedianLOI"`
	IncurredCost          float64 `json:"incurredCost"`
	EstimatedCost         float64 `json:"estimatedCost"`
	LastAcceptedIncidence float64 `json:"lastAcceptedIncidenceRate"`
	LastAcceptedLOI       int64   `json:"lastAcceptedLOI"`
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
	Format             *string            `json:"format,omitempty"`
	LocalizedText      *string            `json:"localizedText,omitempty"`
}

// AttributeOption ...
type AttributeOption struct {
	ID            string  `json:"id"`
	Text          string  `json:"text"`
	LocalizedText *string `json:"localizedText,omitempty"`
}

type URLParameter struct {
	Key string 			`json:"key"`
	Values []string 	`json:"values"`
}