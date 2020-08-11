package samplify

// ProjectPermissionsResponse returns the permissions that a user has on a resource. Here the resource is Project.
type ProjectPermissionsResponse struct {
	ProjectPermissions *ProjectPermissions `json:"data"`
	ResponseStatus     ResponseStatus      `json:"status"`
}

// ProjectPermissions ...
type ProjectPermissions struct {
	ExtProjectID string      `json:"id"`
	CurrentUser  CurrentUser `json:"currentUser"`
	Users        []UserData  `json:"users"`
	Teams        []TeamData  `json:"teams"`
}

// CurrentUser holds a list of roles of the current user
type CurrentUser struct {
	Roles []string `json:"roles"`
}

// UserData holds the information of the users associated to a project
type UserData struct {
	ID       int32  `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

// TeamData holds the information of the teams associated to a project
type TeamData struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

// UpsertPermissionsCriteria has the fields to update and insert project permissions
type UpsertPermissionsCriteria struct {
	ExtProjectID    string             `json:"extProjectId" valid:"required"`
	UserPermissions *[]*UserPermission `json:"users" valid:"optional"`
	TeamPermissions *[]*TeamPermission `json:"teams" valid:"optional"`
}

// TeamPermission holds the team input to the upsert Project permissions.
type TeamPermission struct {
	ID []int32 `json:"id" valid:"required"`
}

// UserPermission holds the user input to the upsert Project permissions.
type UserPermission struct {
	ID   int32  `json:"id" valid:"required"`
	Role string `json:"role" valid:"required"`
}
