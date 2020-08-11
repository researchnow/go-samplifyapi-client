package samplify

// RolesResponse ...
type RolesResponse struct {
	Roles []Role `json:"data"`
	Meta  Meta   `json:"meta"`
}

// Role holds the information about a user role and the actions that can be performed for that role
type Role struct {
	ID              string          `json:"id"`
	Name            string          `json:"name"`
	AssignableRoles []string        `json:"assignableRoles"`
	Description     string          `json:"description"`
	AllowedActions  []AllowedAction `json:"allowedActions"`
}

// AllowedAction ...
type AllowedAction struct {
	ID          string `json:"id"`
	Action      string `json:"action"`
	Description string `json:"description"`
}
