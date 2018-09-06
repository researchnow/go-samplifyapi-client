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
	Direction SortDirection
}

// Filter by top level fields only. Nested fields are not supported for filtering.
type Filter struct {
	Field QueryField
	Value interface{}
}

const maxLimit uint = 1000

// QueryOptions ... Filtering/Sorting and pagination params for GET endpoints that return an object list
// Default limit = 10, maximum limit value = 1000
type QueryOptions struct {
	FilterBy []*Filter
	SortBy   []*Sort
	Offset   uint
	Limit    uint
	Scope    string
}

func query2String(options *QueryOptions) string {
	query := ""
	if options != nil {
		query = "?"
		sep := ""
		if len(options.Scope) > 0 {
			query = fmt.Sprintf("?scope=%s", options.Scope)
			sep = "&amp;"
		}
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
		if len(sep) > 0 {
			sep = "&amp;"
		}
		if options.Offset > 0 {
			query = fmt.Sprintf("%s%soffset=%d", query, sep, options.Offset)
			sep = "&amp;"
		}
		if options.Limit > 0 {
			if options.Limit > maxLimit {
				options.Limit = maxLimit
			}
			query = fmt.Sprintf("%s%slimit=%d", query, sep, options.Limit)
		}
	}
	return query
}
