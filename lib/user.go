package samplify

// UserResponse to hold the api response object.
type UserResponse struct {
	User           *User          `json:"data"`
	ResponseStatus ResponseStatus `json:"status"`
}

// User to hold any information related to the user.
type User struct {
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	FullName  string    `json:"fullName"`
	ID        int32     `json:"id"`
	Companies []Company `json:"companies"`
}

// Company to hold the company information of the user.
type Company struct {
	ID          int32  `json:"id"`
	Default     bool   `json:"default"`
	DefaultRole string `json:"defaultRole"`
	Name        string `json:"name"`
	Teams       []Team `json:"teams"`
}

// Team to hold the team information of the company.
type Team struct {
	ID      int32  `json:"id"`
	Default bool   `json:"default"`
	Role    string `json:"role"`
	Name    string `json:"name"`
	Status  string `json:"status"`
}
