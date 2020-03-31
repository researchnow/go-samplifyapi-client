package samplify

import (
	"fmt"
	"net/url"
	"strings"
	"time"
)

// QueryField ... Supported fields for filtering and sorting
type QueryField string

// QueryField values
const (
	QueryFieldID                       QueryField = "id"
	QueryFieldExtProjectID             QueryField = "extProjectId"
	QueryFieldExtLineItemID            QueryField = "extLineItemId"
	QueryFieldCreatedAt                QueryField = "createdAt"
	QueryFieldUpdatedAt                QueryField = "updatedAt"
	QueryFieldTitle                    QueryField = "title"
	QueryFieldJobNumber                QueryField = "jobNumber"
	QueryFieldName                     QueryField = "name"
	QueryFieldText                     QueryField = "text"
	QueryFieldType                     QueryField = "type"
	QueryFieldState                    QueryField = "state"
	QueryFieldStateReason              QueryField = "stateReason"
	QueryFieldStateLastUpdatedAt       QueryField = "stateLastUpdatedAt"
	QueryFieldIsoCode                  QueryField = "isoCode"
	QueryFieldCountryName              QueryField = "countryName"
	QueryFieldCountryISOCode           QueryField = "countryISOCode"
	QueryFieldLanguageISOCode          QueryField = "languageISOCode"
	QueryFieldLaunchedAt               QueryField = "launchedAt"
	QueryFieldSurveyTopic              QueryField = "surveyTopic"
	QueryFieldStartDate                QueryField = "startDate"
	QueryFieldEndDate                  QueryField = "endDate"
	QueryFieldBillingDate              QueryField = "billingDate"
	QueryFieldIsAllowedInSurveyAppends QueryField = "isAllowedInSurveyAppends"
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
	Value Value
}

// Value ...
type Value interface {
	String() string
}

// DateFilterValue ...
type DateFilterValue struct {
	From *time.Time
	To   *time.Time
}

// FilterValue ...
type FilterValue struct {
	Value interface{}
}

func (filtervalue FilterValue) String() string {
	return url.QueryEscape(fmt.Sprintf("%s", filtervalue.Value))
}

func (datefilter DateFilterValue) String() string {
	fromdate := ""
	todate := ""
	if datefilter.From != nil {
		fromdate = datefilter.From.Format("2006/01/02")
	}
	if datefilter.To != nil {
		todate = datefilter.To.Format("2006/01/02")
	}
	value := fmt.Sprintf("%s,%s", fromdate, todate)
	return value
}

// StringSlice ..
type StringSlice []string

func (ss StringSlice) String() string {
	return strings.Join(ss, ",")
}

const maxLimit uint = 1000

// QueryOptions ... Filtering/Sorting and pagination params for GET endpoints that return an object list
// Default limit = 10, maximum limit value = 1000
type QueryOptions struct {
	FilterBy      []*Filter
	SortBy        []*Sort
	Offset        uint
	Limit         uint
	Scope         string `conform:"trim"`
	ExtProjectId  *string
	ExtLineItemId *string
	EventType     *string `conform:"trim"`
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
				query = fmt.Sprintf("%s%s%s=%s", query, sep, f.Field, f.Value.String())
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
		if options.ExtProjectId != nil {
			query = fmt.Sprintf("%s%sextProjectId=%s", query, sep, *options.ExtProjectId)
		}
		if options.ExtLineItemId != nil {
			query = fmt.Sprintf("%s%sextLineItemId=%s", query, sep, *options.ExtLineItemId)
		}
		if options.EventType != nil {
			query = fmt.Sprintf("%s%seventType=%s", query, sep, *options.EventType)
		}
	}
	return query
}
