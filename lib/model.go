package samplify

// Model ...
type Model struct {
	StateLastUpdatedAt *CustomTime `json:"stateLastUpdatedAt"`
	CreatedAt          CustomTime  `json:"createdAt"`
	UpdatedAt          CustomTime  `json:"updatedAt"`
}
