# Samplify API client

A golang client library to connect with researchnow/ssi demand api
<br /><a href="https://developers.researchnow.com/samplifyapi-docs" target="_blank">See complete API reference here</a>

### Prerequisites

* Account credentials to access researchnow/ssi demand api
* Or, a test account to explore the API on UAT server

## Usage examples

The following host URLs are configured based on the environment variable setting.

Prod settings:
* Use environment `env=prod`
* Authentication endpoint: "https://api.researchnow.com/auth/v1/token/password"
* API base url: "https://api.researchnow.com/sample/v1"

UAT (default) settings:
* Use environment `env=uat`
* Authentication endpoint: "https://api.uat.pe.researchnow.com/auth/v1/token/password"
* API base url: "https://api.uat.pe.researchnow.com/sample/v1"

### Creating a client connection

The new client is initialized based on environment variable setting described above. Default `env` will be considered "uat" if not provided.

```
client = samplify.NewClient("client_id", "username", "password")
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
p := &samplify.ProjectCriteria{
	ExtProjectID: "prj01",
	Title:        "Samplify Test Project 01",
	...
}
r, err := client.CreateProject(p)
```

Update an existing project:

```
p := &samplify.ProjectCriteria{
	ExtProjectID: "prj01",
	Title:        "Updated, samplify test project 01",
	...
}
r, err := client.UpdateProject(p)
```

The returned, `ProjectResponse` object contains:
* `r.Project` the newly created or updated project object.
* `r.ResponseStatus`

## Filtering & Sorting

All client functions that take `*QueryOptions` parameter, support filtering/sorting & pagination. Nested fields are not supported for filtering and sorting operations. Default `limit` value is set to 10 but values up to 50 are permitted.

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

* CreateProject(project *ProjectCriteria) (*ProjectResponse, error)
* UpdateProject(project *ProjectCriteria) (*ProjectResponse, error)
* BuyProject(extProjectID string, buy []*BuyProjectCriteria) (*BuyProjectResponse, error)
* CloseProject(extProjectID string) (*CloseProjectResponse, error)
* GetAllProjects(options *QueryOptions) (*GetAllProjectsResponse, error)
* GetProjectBy(extProjectID string) (*ProjectResponse, error)
* GetProjectReport(extProjectID string) (*ProjectReportResponse, error)
* AddLineItem(extProjectID string, lineItem *LineItemCriteria) (*LineItemResponse, error)
* UpdateLineItem(extProjectID, extLineItemID string, lineItem *LineItemCriteria) (*LineItemResponse, error)
* UpdateLineItemState(extProjectID, extLineItemID string, action Action) (*ChangeLineItemStateResponse, error)
* GetAllLineItems(extProjectID string, options *QueryOptions) (*GetAllLineItemsResponse, error)
* GetLineItemBy(extProjectID, extLineItemID string) (*LineItemResponse, error)
* GetFeasibility(extProjectID string, options *QueryOptions) (*GetFeasibilityResponse, error)
* GetCountries(options *QueryOptions) (*GetCountriesResponse, error)
* GetAttributes(countryCode, languageCode string, options *QueryOptions) (*GetAttributesResponse, error)
* GetSurveyTopics(options *QueryOptions) (*GetSurveyTopicsResponse, error)
* RefreshToken() error
* Logout() error


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