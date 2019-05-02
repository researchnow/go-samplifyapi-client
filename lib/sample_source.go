package samplify

// SampleSource ...
type SampleSource struct {
	CountryISOCode  string    `json:"countryISOCode"`
	LanguageISOCode string    `json:languageISOCode`
	Sources         []Sources `json:sources`
	Default         bool      `json:default`
}

// Sources ...
type Sources struct {
	ID       int        `json:"id"`
	Name     string     `json:name`
	Category []Category `json:category`
}

// GetSampleSourceResponse ...
type GetSampleSourceResponse struct {
	List           []*SampleSource `json:"data"`
	ResponseStatus ResponseStatus  `json:"status"`
	Meta           Meta            `json:"meta"`
}
