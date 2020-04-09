package samplify

// AppError ...
type AppError struct {
	Data   interface{} `json:"data"`
	Meta   *Meta       `json:"meta"`
	Status *Status     `json:"status"`
}

// TemplateCriteria ...
type TemplateCriteria struct {
	CountryISOCode  string     `json:"countryISOCode"`
	Description     string     `json:"description"`
	LanguageISOCode string     `json:"languageISOCode"`
	Name            string     `json:"name"`
	QuotaPlan       *QuotaPlan `json:"quotaPlan"`
	Tags            []string   `json:"tags"`
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
	CountryISOCode  *string    `json:"countryISOCode,omitempty"`
	CreatedAt       *string    `json:"createdAt,omitempty"`
	Description     *string    `json:"description,omitempty"`
	Editable        bool       `json:"editable"`
	ID              int        `json:"id"`
	LanguageISOCode *string    `json:"languageISOCode,omitempty"`
	Name            string     `json:"name"`
	State           string     `json:"state"`
	Tags            []string   `json:"tags"`
	UpdatedAt       *string    `json:"updatedAt,omitempty"`
	QuotaPlan       *QuotaPlan `json:"quotaPlan"`
}
