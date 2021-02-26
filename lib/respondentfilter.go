package samplify

import (
	"errors"
	"github.com/researchnow/tareekh"
	"time"
)

const ( // TimeLayout ...
	TimeLayout = "2006/01/02 15:04:05"
	// DateLayout ...
	DateLayout = "2006/01/02"
)

// RespondentFilterType ...
type RespondentFilterType string

// RespondentFilterType values
const (
	RespondentFilterTypeProject  RespondentFilterType = "PROJECT"
	RespondentFilterTypeTag      RespondentFilterType = "TAG"
	RespondentFilterTypeCategory RespondentFilterType = "CATEGORY"
)

// RespondentFilterTypes ...
var RespondentFilterTypes = []RespondentFilterType{
	RespondentFilterTypeProject,
	RespondentFilterTypeTag,
	RespondentFilterTypeCategory,
}

type RespondentStatus string

const (
	RespondentStatusCompleted RespondentStatus = "COMPLETED"
	RespondentStatusAttempted RespondentStatus = "ATTEMPTED"
	RespondentStatusOverQuota RespondentStatus = "OVERQUOTA"
	RespondentStatusScreenOut RespondentStatus = "SCREENOUT"
)

// RespondentStatuses ...
var RespondentStatuses = []RespondentStatus{
	RespondentStatusCompleted,
	RespondentStatusAttempted,
	RespondentStatusOverQuota,
	RespondentStatusScreenOut,
}

var (
	// ErrInvalidRespondentFilterStartDate ...
	ErrInvalidRespondentFilterStartDate = errors.New("respondent filter start date cannot be empty")
	// ErrInvalidRespondentFilterEndDate ...
	ErrInvalidRespondentFilterEndDate = errors.New("respondent filter end date cannot be empty")
	// ErrInvalidRespondentFilterDateRanges ...
	ErrInvalidRespondentFilterDateRanges = errors.New("respondent filter date ranges are invalid")
	ErrInvalidRelativeType               = errors.New("relative type is invalid")
	ErrInvalidRelativeValue              = errors.New("relative value is invalid")
)

// RelativeType ...
type RelativeType int

// RelativeType values
const (
	RelativeTypeDays   RelativeType = 1
	RelativeTypeWeeks  RelativeType = 7
	RelativeTypeMonths RelativeType = 30
)

// RespondentScheduleType ...
type RespondentScheduleType string

// RespondentScheduleType values
const (
	RespondentScheduleTypeAllDates   RespondentScheduleType = "ALL_DATES"
	RespondentScheduleTypeThisMonth  RespondentScheduleType = "THIS_MONTH"
	RespondentScheduleTypeLastDays   RespondentScheduleType = "LAST_DAYS"
	RespondentScheduleTypeLastMonths RespondentScheduleType = "LAST_MONTHS"
	RespondentScheduleTypeCustom     RespondentScheduleType = "CUSTOM"
)

// RespondentScheduleTypes ...
var RespondentScheduleTypes = []RespondentScheduleType{
	RespondentScheduleTypeAllDates,
	RespondentScheduleTypeThisMonth,
	RespondentScheduleTypeLastDays,
	RespondentScheduleTypeLastMonths,
	RespondentScheduleTypeCustom,
}

// RespondentFilter ... Project's respondent filter
type RespondentFilter struct {
	Type         string                   `json:"type"`
	List         []string                 `json:"list"`
	Dispositions []string                 `json:"dispositions"`
	Schedule     RespondentFilterSchedule `json:"schedule"`
}

type RespondentFilterSchedule struct {
	Type      string  `json:"type"`
	Value     int32   `json:"value"`
	StartDate *string `json:"startDate"`
	EndDate   *string `json:"endDate"`
}

// ComputeDates for Respondent Schedule Types
// Last 2 days (start date: now - 2 days before, end date: now)
// Last 2 weeks (start date: now - 14 days before, end date: now)
// Last 2 months (start date: now - 60 days before, end date: now)
// This month (start date: beginning of month, end date - now)
func (rf *RespondentFilter) ComputeDates() {
	now := time.Now()
	current := now.Format(DateLayout)
	switch rf.Schedule.Type {
	case string(RespondentScheduleTypeAllDates):
		rf.Schedule.StartDate = nil
		rf.Schedule.EndDate = nil
		rf.Schedule.Value = 0
		return
	case string(RespondentScheduleTypeThisMonth):
		startDate := tareekh.BeginningOfMonth().Format(DateLayout)
		rf.Schedule.StartDate = &startDate
		rf.Schedule.EndDate = &current
		rf.Schedule.Value = 0
		return
	case string(RespondentScheduleTypeLastDays):
		multiplier := int(rf.Schedule.Value)
		startDate := tareekh.DaysAgo(multiplier * int(RelativeTypeDays)).Format(DateLayout)
		rf.Schedule.StartDate = &startDate
		rf.Schedule.EndDate = &current
		return
	case string(RespondentScheduleTypeLastMonths):
		multiplier := int(rf.Schedule.Value)
		startDate := tareekh.DaysAgo(multiplier * int(RelativeTypeMonths)).Format(DateLayout)
		rf.Schedule.StartDate = &startDate
		rf.Schedule.EndDate = &current
		return
	case string(RespondentScheduleTypeCustom):
		rf.Schedule.Value = 0
	}
}

// ValidateDates ranges when user provides start and end dates.
// (valid start date, valid end date, valid relation between start and end date)
func (rf *RespondentFilter) ValidateDates() error {
	switch rf.Schedule.Type {
	case string(RespondentScheduleTypeLastDays):
		if rf.Schedule.Value == 0 {
			return ErrInvalidRelativeValue
		}
		return nil
	case string(RespondentScheduleTypeLastMonths):
		if rf.Schedule.Value == 0 {
			return ErrInvalidRelativeValue
		}
		return nil
	case string(RespondentScheduleTypeCustom):
		if rf.Schedule.StartDate == nil || len(*rf.Schedule.StartDate) == 0 {
			return ErrInvalidRespondentFilterStartDate
		}
		if rf.Schedule.EndDate == nil || len(*rf.Schedule.EndDate) == 0 {
			return ErrInvalidRespondentFilterEndDate
		}
		start, err := time.Parse(DateLayout, *rf.Schedule.StartDate)
		if err != nil {
			return err
		}
		end, err := time.Parse(DateLayout, *rf.Schedule.EndDate)
		if err != nil {
			return err
		}
		if start.After(end) || end.Before(start) || start.Equal(end) {
			return ErrInvalidRespondentFilterDateRanges
		}
	}
	return nil
}

// PopulateProjects is appending projects to the respondent filter list, existing projects are ignored
func (rf *RespondentFilter) PopulateProjects(extProjectIDs []string) error {
	m := make(map[string]bool)

	for _, item := range rf.List {
		m[item] = true
	}
	//Existing projects will be ignored, avoid duplicating
	for _, item := range extProjectIDs {
		if _, ok := m[item]; !ok {
			rf.List = append(rf.List, item)
		}
	}
	return nil
}

// SetDates sets start and end dates
func (rf *RespondentFilter) SetDates(start *time.Time, end *time.Time) {
	if start != nil {
		s := start.Format(DateLayout)
		rf.Schedule.StartDate = &s
	}
	if end != nil {
		ed := end.Format(DateLayout)
		rf.Schedule.EndDate = &ed
	}
}

// ExclusionList checks if the list in respondent filter object is nil or not.
// If the list is nil then the function returns an empty array of string so that we do not pass nil to the API which causes an error.
func (rf *RespondentFilter) ExclusionList() []string {
	if rf.List == nil {
		return []string{}
	}
	return rf.List
}
