package samplify

// AppError ...
type AppError struct {
	Data   interface{} `json:"data"`
	Meta   *Meta       `json:"meta"`
	Status *Status     `json:"status"`
}

// TargetingAttribute user type.
type TargetingAttribute struct {
	AttributeID string   `json:"attributeId"`
	Operator    *string  `json:"operator,omitempty"`
	Options     []string `json:"options"`
}

// QuotaCellTemplate user type.
type QuotaCellTemplate struct {
	Perc       *float64              `json:"perc,omitempty"`
	QuotaNodes []*TargetingAttribute `json:"quotaNodes,omitempty"`
}

// QuotaGroupTemplate user type.
type QuotaGroupTemplate struct {
	Name       *string              `json:"name,omitempty"`
	QuotaCells []*QuotaCellTemplate `json:"quotaCells"`
}

// QuotaPlanTemplate user type.
type QuotaPlanTemplate struct {
	Filters     []*TargetingAttribute `json:"filters,omitempty"`
	QuotaGroups []*QuotaGroupTemplate `json:"quotaGroups,omitempty"`
}

// TemplateCriteria ...
type TemplateCriteria struct {
	CountryISOCode  string             `json:"countryISOCode"`
	Description     string             `json:"description"`
	LanguageISOCode string             `json:"languageISOCode"`
	Name            string             `json:"name"`
	QuotaPlan       *QuotaPlanTemplate `json:"quotaPlan"`
	Tags            *string            `json:"tags,omitempty"`
}

// ErrorType user type.
type ErrorType struct {
	Code     string    `json:"code"`
	Message  string    `json:"message"`
	Resource *Resource `json:"resource"`
}

// Status user type.
type Status struct {
	Errors  []*ErrorType `json:"errors"`
	Message string       `json:"message"`
}

// TemplateResponse response
type TemplateResponse struct {
	Data   *TemplateData `json:"data"`
	Meta   *Meta         `json:"meta"`
	Status *Status       `json:"status"`
}

// TemplatesResponse response
type TemplatesResponse struct {
	Data   []*TemplateData `json:"data"`
	Meta   *Meta           `json:"meta"`
	Status *Status         `json:"status"`
}

// TemplateData ...
type TemplateData struct {
	CountryISOCode  *string            `json:"countryISOCode,omitempty"`
	CreatedAt       *string            `json:"createdAt,omitempty"`
	Description     *string            `json:"description,omitempty"`
	Editable        bool               `json:"editable"`
	ID              int                `json:"id"`
	LanguageISOCode *string            `json:"languageISOCode,omitempty"`
	Name            string             `json:"name"`
	State           string             `json:"state"`
	Tags            *string            `json:"tags,omitempty"`
	UpdatedAt       *string            `json:"updatedAt,omitempty"`
	QuotaPlan       *QuotaPlanTemplate `json:"quotaPlan"`
}
