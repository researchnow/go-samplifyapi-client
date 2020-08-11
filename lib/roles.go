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

// GetRolesCriteria has filters to get specific roles. If no filters are provided then the result would give out all the roles.
type GetRolesCriteria struct {
	ID   []string `json:"id" valid:"optional"`
	Name string   `json:"name" valid:"optional"`
}
