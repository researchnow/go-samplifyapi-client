package samplify

import (
	"fmt"
	"net/url"
)

// QueryField ... Supported fields for filtering and sorting
type QueryField string

// QueryField values
const (
	QueryFieldID                 QueryField = "id"
	QueryFieldExtProjectID       QueryField = "extProjectId"
	QueryFieldExtLineItemID      QueryField = "extLineItemId"
	QueryFieldCreatedAt          QueryField = "createdAt"
	QueryFieldUpdatedAt          QueryField = "updatedAt"
	QueryFieldTitle              QueryField = "title"
	QueryFieldName               QueryField = "name"
	QueryFieldText               QueryField = "text"
	QueryFieldType               QueryField = "type"
	QueryFieldState              QueryField = "state"
	QueryFieldStateReason        QueryField = "stateReason"
	QueryFieldStateLastUpdatedAt QueryField = "stateLastUpdatedAt"
	QueryFieldIsoCode            QueryField = "isoCode"
	QueryFieldCountryName        QueryField = "countryName"
	QueryFieldCountryISOCode     QueryField = "countryISOCode"
	QueryFieldLanguageISOCode    QueryField = "languageISOCode"
	QueryFieldLaunchedAt         QueryField = "launchedAt"
)

// SortDirection (asc, desc)
type SortDirection string

// SortDirection values
const (
	SortDirectionAsc  SortDirection = "asc"
	SortDirectionDesc SortDirection = "desc"
)

// Sort by top level fields only. Nested fields are not supported for sorting.
type Sort struct {
	Field     QueryField
	Direction SortDirection `json:"direction"`
}

// Filter by top level fields only. Nested fields are not supported for filtering.
type Filter struct {
	Field QueryField
	Value interface{}
}

// QueryOptions ... Filtering/Sorting for GET endpoints that return an object list
type QueryOptions struct {
	FilterBy []*Filter
	SortBy   []*Sort
}

func query2String(options *QueryOptions) string {
	query := ""
	if options != nil {
		query = "?"
		sep := ""
		if len(options.FilterBy) > 0 {
			for _, f := range options.FilterBy {
				query = fmt.Sprintf("%s%s%s=%s", query, sep, f.Field, url.QueryEscape(fmt.Sprintf("%s", f.Value)))
				sep = "&amp;"
			}
		}
		if len(options.SortBy) > 0 {
			query = fmt.Sprintf("%s%ssort=", query, sep)
			sep = ""
			for _, s := range options.SortBy {
				query = fmt.Sprintf("%s%s%s:%s", query, sep, s.Field, s.Direction)
				sep = ","
			}
		}
	}
	return query
}
