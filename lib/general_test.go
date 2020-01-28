package samplify

import (
	"fmt"
	"testing"
)

func TestGeneral(t *testing.T) {
	fmt.Println("hi")

	fv1 := FilterValue{Value: "2019-06-12"}

	f1 := Filter{
		Field: QueryFieldStartDate,
		Value: fv1,
	}

	fv2 := FilterValue{Value: "2019-06-19"}

	f2 := Filter{
		Field: QueryFieldEndDate,
		Value: fv2,
	}

	filters := []*Filter{
		&f1, &f2,
	}

	projectID := "010528ef-8984-48c1-a06d-4dae730da027"

	option := QueryOptions{
		FilterBy:      filters,
		ExtProjectId:  &projectID,
	}

	val := query2String(&option)

	fmt.Println(val)
}