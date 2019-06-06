package samplify

// UserResponse to hold the api response object.
type UserResponse struct {
	User           *User          `json:"data"`
	ResponseStatus ResponseStatus `json:"status"`
}

// User to hold any information related to the user.
type User struct {
	Applications []Application `json:"applications"`
	CompanyID    int32         `json:"companyId"`
	CompanyName  string        `json:"companyName" conform:"trim"`
	Email        string        `json:"email" conform:"trim"`
	Username     string        `json:"username" conform:"trim"`
	FullName     string        `json:"fullName" conform:"trim"`
}

// Application to hold the app level information of the user.
type Application struct {
	ID      int32  `json:"appId"`
	Current bool   `json:"current"`
	Default bool   `json:"default"`
	Name    string `json:"name" conform:"trim"`
}
