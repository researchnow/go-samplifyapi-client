package samplify

import (
	"errors"
	"fmt"
	"time"
)


// ExclusionType ...
type ExclusionType string

var ExclusionTypes = []ExclusionType{
	ExclusionTypeProject,
	ExclusionTypeTag,
	ExclusionTypeThisMonth,
	ExclusionTypeLastMonth,
	ExclusionTypeLastThreeMonth,
	ExclusionTypeCustom,
}

// ExclusionType values
const (
	ExclusionTypeProject 		ExclusionType = "PROJECT"
	ExclusionTypeTag 			ExclusionType = "TAG"
	ExclusionTypeThisMonth     	ExclusionType = "THIS_MONTH"
	ExclusionTypeLastMonth     	ExclusionType = "LAST_MONTH"
	ExclusionTypeLastThreeMonth ExclusionType = "LAST_THREE_MONTHS"
	ExclusionTypeCustom     	ExclusionType = "CUSTOM"
)

func (p ExclusionType) String() string {
	return string(p)
}

// Exclusions ... Project's exclusions
type Exclusions struct {
	Type ExclusionType `json:"type" valid:"ExclusionType"`
	List []string      `json:"list"`
	StartDate *string  `json:"startDate"`
	EndDate   *string  `json:"endDate"`
}

func (e Exclusions) Compute() (*string, *string){
	now := time.Now()
	current := now.Format(TimeLayout)
	switch e.Type {
	case ExclusionTypeThisMonth:
		startDate := BeginningOfMonth(now).Format(TimeLayout)
		return &startDate, &current
	case ExclusionTypeLastMonth:
		startDate := DaysBeforeAfterMonth(now, -30).Format(TimeLayout)
		return &startDate, &current
	case ExclusionTypeLastThreeMonth:
		startDate := DaysBeforeAfterMonth(now, -90).Format(TimeLayout)
		return &startDate, &current
	case ExclusionTypeCustom:
		return e.StartDate,e.EndDate
	}
	return nil, nil
}

func BeginningOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}
func DaysBeforeAfterMonth(t time.Time, days int) time.Time {
	return t.AddDate(0, 0, days)
}

func EndOfMonth(t time.Time) time.Time {
	return BeginningOfMonth(t).AddDate(0, 1, 0).Add(-time.Second)
}

func (e Exclusions) Validate() error {
	switch e.Type {
	case ExclusionTypeCustom:
		if e.StartDate == nil{
			return errors.New("exclusion start date cannot be empty")
		}
		if e.EndDate == nil {
			return errors.New("exclusion end date cannot be empty")
		}
		start, err := time.Parse(TimeLayout, *e.StartDate)
		end, err := time.Parse(TimeLayout, *e.EndDate)
		if err != nil{
			return err
		}
		if start.After(end) || end.Before(start){
			return fmt.Errorf("invalid date ranges: %s and %s", *e.StartDate, *e.EndDate)
		}
	}
	return nil
}
