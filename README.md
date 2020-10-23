# Samplify API client

A golang client library to connect with researchnow/ssi demand api
<br /><a href="https://developers.dynata.com/demand-api-reference/" target="_blank">See complete API reference here</a>

### Prerequisites

* Account credentials to access researchnow/ssi demand api
* Or, a test account to explore the API on UAT server

## Usage examples

Predefined `ClientOptions` are available for creating a new client.
Use `UATClientOptions` or `ProdClientOptions` for uat or prod environment, respectively.

### Creating a client connection

The new client is initialized based on `ClientOptions` parameter, described above. If `ClientOptions` is nil, `UATClientOptions` will be used.

```
client = samplify.NewClient("client_id", "username", "password", samplify.UATClientOptions)
```

The session expires after some time but the client will automatically acquire one by making an authentication request before sending out the actual request, again.

### Basic request structure

All the request functions return their respective response object, along with an error object.
Generally, all response objects consist of:
* Requested data object.
* ResponseStatus, which consists of the "status" part of the json, returned. It is basically the API's custom status messages related to the request execution.

Some of the response objects (such as those that return a list) also contain a "Meta" field.
* Meta, contains metadata such as page navigation links etc.

```
r, err := client.GetAllProjects(nil)
if err == nil {
	for _, p := range r.Projects {
		fmt.Println(p.Title)
	}
	fmt.Printf("Next page url: %s", r.Meta.Next)
}
```

### Example - Creating/Updating a Project

Create a new project:
```
p := &samplify.CreateProjectCriteria{
	ExtProjectID: "prj01",
	Title:        "Samplify Test Project 01",
	...
}
r, err := client.CreateProject(p)
```

Update an existing project:

```
title := "Updated, samplify test project 01"
p := &samplify.UpdateProjectCriteria{
	ExtProjectID: "prj01",
	Title:        &title,
	...
}
r, err := client.UpdateProject(p)
```

The returned, `ProjectResponse` object contains:
* `r.Project` the newly created or updated project object.
* `r.ResponseStatus`

## Filtering & Sorting

All client functions that take `*QueryOptions` parameter, support filtering/sorting & pagination. Nested fields are not supported for filtering and sorting operations. Default `limit` value is set to 10 but value up to 1000 is permitted.

```
options := &samplify.QueryOptions{
	FilterBy: []*samplify.Filter{
		&samplify.Filter{Field: samplify.QueryFieldTitle, Value: "Test Survey"},
		&samplify.Filter{Field: samplify.QueryFieldState, Value: samplify.StateProvisioned},
	},
	SortBy: []*samplify.Sort{
		&samplify.Sort{Field: samplify.QueryFieldCreatedAt, Direction: samplify.SortDirectionAsc},
		&samplify.Sort{Field: samplify.QueryFieldExtProjectID, Direction: samplify.SortDirectionDesc},
	},
	Offset: 10,
	Limit: 5,
}

r, err := client.GetAllProjects(options)
if err == nil {
	for _, p := range r.Projects {
		fmt.Println(p.Title)
	}
}
```

If multiple sort objects are provided, the order in which they are added in the slice, is followed.

## Supported API functions

* CreateProject(project *CreateProjectCriteria) (*ProjectResponse, error)
* CreateProjectWithContext(ctx context.Context, project *CreateProjectCriteria) (*ProjectResponse, error)
* UpdateProject(project *UpdateProjectCriteria) (*ProjectResponse, error)
* UpdateProjectWithContext(ctx context.Context, project *UpdateProjectCriteria) (*ProjectResponse, error)
* BuyProject(extProjectID string, buy []*BuyProjectCriteria) (*BuyProjectResponse, error)
* BuyProjectWithContext(ctx context.Context, extProjectID string, buy []*BuyProjectCriteria) (*BuyProjectResponse, error)
* CloseProject(extProjectID string) (*CloseProjectResponse, error)
* CloseProjectWithContext(ctx context.Context, extProjectID string) (*CloseProjectResponse, error)
* GetAllProjects(options *QueryOptions) (*GetAllProjectsResponse, error)
* GetAllProjectsWithContext(ctx context.Context, options *QueryOptions) (*GetAllProjectsResponse, error)
* GetProjectBy(extProjectID string) (*ProjectResponse, error)
* GetProjectByWithContext(ctx context.Context, extProjectID string) (*ProjectResponse, error)
* GetProjectReport(extProjectID string) (*ProjectReportResponse, error)
* GetProjectReportWithContext(ctx context.Context, extProjectID string) (*ProjectReportResponse, error)
* AddLineItem(extProjectID string, lineItem *CreateLineItemCriteria) (*LineItemResponse, error)
* AddLineItemWithContext(ctx context.Context, extProjectID string, lineItem *CreateLineItemCriteria) (*LineItemResponse, error)
* UpdateLineItem(extProjectID, extLineItemID string, lineItem *UpdateLineItemCriteria) (*LineItemResponse, error)
* UpdateLineItemWithContext(ctx context.Context, extProjectID, extLineItemID string, lineItem *UpdateLineItemCriteria) (*LineItemResponse, error)
* UpdateLineItemState(extProjectID, extLineItemID string, action Action) (*ChangeLineItemStateResponse, error)
* UpdateLineItemStateWithContext(ctx context.Context, extProjectID, extLineItemID string, action Action) (*ChangeLineItemStateResponse, error)
* GetAllLineItems(extProjectID string, options *QueryOptions) (*GetAllLineItemsResponse, error)
* GetAllLineItemsWithContext(ctx context.Context, extProjectID string, options *QueryOptions) (*GetAllLineItemsResponse, error)
* GetLineItemBy(extProjectID, extLineItemID string) (*LineItemResponse, error)
* GetLineItemByWithContext(ctx context.Context, extProjectID, extLineItemID string) (*LineItemResponse, error)
* GetFeasibility(extProjectID string, options *QueryOptions) (*GetFeasibilityResponse, error)
* GetFeasibilityWithContext(ctx context.Context, extProjectID string, options *QueryOptions) (*GetFeasibilityResponse, error)
* GetCountries(options *QueryOptions) (*GetCountriesResponse, error)
* GetCountriesWithContext(ctx context.Context, options *QueryOptions) (*GetCountriesResponse, error)
* GetAttributes(countryCode, languageCode string, options *QueryOptions) (*GetAttributesResponse, error)
* GetAttributesWithContext(ctx context.Context, countryCode, languageCode string, options *QueryOptions) (*GetAttributesResponse, error)
* GetSurveyTopics(options *QueryOptions) (*GetSurveyTopicsResponse, error)
* GetSurveyTopicsWithContext(ctx context.Context, options *QueryOptions) (*GetSurveyTopicsResponse, error)
* RefreshToken() error
* RefreshTokenWithContext(ctx context.Context, ) error
* Logout() error
* LogoutWithContext(ctx context.Context, ) error


## Versioning

### 1.0
Supports API functionalities, such as:
* Authentication, including automatic re-authentication on token expire.
* Project and Line Items related requests. Such as, create, update, get-all, get-by-id requests etc.
* Filtering & Sorting
* Pricing & Feasibility
* Data endpoint functions, serving attributes, categories and countries & languages data.

## Authors

* Maaz Nisar
