package samplify

// SampleSource ...
type SampleSource struct {
	CountryISOCode  string    `json:"countryISOCode"`
	LanguageISOCode string    `json:"languageISOCode"`
	Sources         []Sources `json:"sources"`
}

// Sources ...
type Sources struct {
	ID       int                  `json:"id"`
	Name     string               `json:"name"`
	Category SampleSourceCategory `json:"category"`
	Default  bool                 `json:"default"`
}

// SampleSourceCategory is a Sample source's allowed list of surveytopics
type SampleSourceCategory struct {
	SurveyTopic []string `json:"surveyTopics" valid:"required"`
}

// GetSampleSourceResponse ...
type GetSampleSourceResponse struct {
	List           []*SampleSource `json:"data"`
	ResponseStatus ResponseStatus  `json:"status"`
	Meta           Meta            `json:"meta"`
}
