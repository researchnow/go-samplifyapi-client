package samplify

// UserResponse to hold the api response object.
type UserResponse struct {
	User           *User          `json:"data"`
	ResponseStatus ResponseStatus `json:"status"`
}

// CompanyUsersResponse holds api response object and returns a list of company users.
type CompanyUsersResponse struct {
	List           []*CompanyUser `json:"data"`
	ResponseStatus ResponseStatus `json:"status"`
	Meta           Meta           `json:"meta"`
}

// User to hold any information related to the user.
type User struct {
	ID        int32     `json:"id,omitempty"`
	Email     string    `json:"email"`
	Username  string    `json:"userName"`
	FullName  string    `json:"fullName"`
	Companies []Company `json:"companies"`
}

// Company holds the information of a company associated to the user.
type Company struct {
	ID          int32  `json:"id"`
	Name        string `json:"name"`
	DefaultRole string `json:"defaultRole"`
	Default     bool   `json:"default"`
	Teams       []Team `json:"teams"`
}

// Team holds the information about a team associated to a company.
type Team struct {
	ID      int32  `json:"id"`
	Name    string `json:"name"`
	Role    string `json:"role"`
	Default bool   `json:"default"`
	Status  string `json:"status"`
}

// CompanyUser holds the information of users associated to a company.
type CompanyUser struct {
	Email    string `json:"email"`
	Username string `json:"userName"`
	FullName string `json:"name"`
}

// SwitchCompanyCriteria ...
type SwitchCompanyCriteria struct {
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	RefreshToken string `json:"refreshToken"`
	CompanyID    int32  `json:"companyId"`
}
