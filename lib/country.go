package samplify

// Country ...
type Country struct {
	ID                 string      `json:"id" conform:"trim"`
	IsoCode            string      `json:"isoCode" conform:"trim"`
	CountryName        string      `json:"countryName" conform:"trim"`
	SupportedLanguages []*Language `json:"supportedLanguages"`
}

// Language ...
type Language struct {
	ID           string `json:"id" conform:"trim"`
	IsoCode      string `json:"isoCode" conform:"trim"`
	LanguageName string `json:"languageName" conform:"trim"`
}
