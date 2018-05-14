# Samplify API client

A golang client library to connect with researchnow/ssi demand api

### Prerequisites

* Account credentials to access researchnow/ssi demand api
* Or, a test account to explore the API on dev server

## Usage examples

Default authentication endpoint: "https://api.researchnow.com/auth/v1/token/password"
Default API base url: "https://api.researchnow.com/sample/v1"

For test account, use the following:
Authentication endpoint: "https://api.dev.pe.researchnow.com/auth/v1/token/password"
API base url: "https://api.dev.pe.researchnow.com/sample/v1"

### Creating a client connection

Using default client:
```
client = samplify.NewClient("client_id", "username", "password", nil)
```

Or, with manually configured ClientOptions:
```
client := samplify.NewClient("client_id", "username", "password",
	&samplify.ClientOptions{AuthURL: devAuthURL, APIBaseURL: devAPIBaseURL})
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
r, err := client.GetAllProjects()
if err != nil {
	for _, p := range r.Projects {
		fmt.Println(p.Title)
	}
	fmt.Printf("Next page url: %s", r.Meta.Next)
}
```

### Example - Creating/Updating a Project

Create a new project:
```
p := &samplify.CreateUpdateProjectCriteria{
	ExtProjectID: "prj01",
	Title:        "Samplify Test Project 01",
	...
}
r, err := client.CreateProject(p)
```

Update an existing project:

```
p := &samplify.CreateUpdateProjectCriteria{
	ExtProjectID: "prj01",
	Title:        "Updated, samplify test project 01",
	...
}
r, err := client.UpdateProject(p)
```

The returned, `ProjectResponse` object contains:
* `r.Project` the newly created or updated project object.
* `r.ResponseStatus`

## Supported API functions

* CreateProject(project *CreateUpdateProjectCriteria) (*ProjectResponse, error)
* UpdateProject(project *CreateUpdateProjectCriteria) (*ProjectResponse, error)
* BuyProject(extProjectID string, buy []*BuyProjectCriteria) (*BuyProjectResponse, error)
* CloseProject(extProjectID string) (*CloseProjectResponse, error)
* GetAllProjects() (*GetAllProjectsResponse, error)
* GetProjectBy(extProjectID string) (*ProjectResponse, error)
* GetProjectReport(extProjectID string) (*ProjectReportResponse, error)
* AddLineItem(extProjectID string, lineItem *LineItem) (*LineItemResponse, error)
* UpdateLineItem(extProjectID, extLineItemID string, lineItem *LineItem) (*LineItemResponse, error)
* ChangeLineItemState(extProjectID, extLineItemID string, action Action) (*ChangeLineItemStateResponse, error)
* GetAllLineItems(extProjectID string) (*GetAllLineItemsResponse, error)
* GetLineItemBy(extProjectID, extLineItemID string) (*LineItemResponse, error)
* GetFeasibility(extProjectID string) (*GetFeasibilityResponse, error)
* GetCountries() (*GetCountriesResponse, error)
* GetAttributes(countryCode, languageCode string) (*GetAttributesResponse, error)
* GetSurveyTopics() (*GetSurveyTopicsResponse, error)

## Versioning

### 1.0
Supports API functionalities, such as:
* Authentication, including automatic re-authentication on token expire.
* Project and Line Items related requests. Such as, create, update, get-all, get-by-id requests etc.
* Pricing & Feasibility
* Data endpoint functions, serving attributes, categories and countries & languages data.

## Authors

* Maaz Nisar