package samplify

// Country ...
type Country struct { 
	ID                 string      `json:"id"`
	IsoCode            string      `json:"isoCode"`
	CountryName        string      `json:"countryName"`
	SupportedLanguages []*Language `json:"supportedLanguages"`
}

// Language ...
type Language struct {
	ID           string `json:"id"`
	IsoCode      string `json:"isoCode"`
	LanguageName string `json:"languageName"`
}
