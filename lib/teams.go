package samplify

// TeamsResponse holds api response object and returns a list of teams associated to a company.
type TeamsResponse struct {
	List           []*CompanyTeam `json:"data"`
	ResponseStatus ResponseStatus `json:"status"`
	Meta           Meta           `json:"meta"`
}

// CompanyTeam holds info about a company team
type CompanyTeam struct {
	ID          int32       `json:"id"`
	Name        string      `json:"name"`
	Status      string      `json:"status"`
	Default     bool        `json:"default"`
	Description string      `json:"description"`
	CreatedAt   *CustomTime `json:"createdAt"`
	UpdatedAt   *CustomTime `json:"updatedAt"`
}
