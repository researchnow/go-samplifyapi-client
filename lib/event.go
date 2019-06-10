package samplify

import "errors"

// Event errors
var (
	ErrEventActionNotApplicable = errors.New("requested action is not applicable for this event")
)

// EventType ...
type EventType string

// EventType values
const (
	EventLineItemRepriceTriggered EventType = "LineItem:RepriceTriggered"
	EventLineItemRepriceAccepted  EventType = "LineItem:RepriceAccepted"
	EventLineItemRepriceRejected  EventType = "LineItem:RepriceRejected"
)

// EventStatus ...
type EventStatus string

// EventStatus values
const (
	EventStatusLaunched         EventStatus = "LAUNCHED"
	EventStatusPaused           EventStatus = "PAUSED"
	EventStatusAwaitingApproval EventStatus = "AWAITING_CLIENT_APPROVAL"
	EventStatusClosed           EventStatus = "CLOSED"
)

// EventActions ...
type EventActions struct {
	AcceptURL string `json:"acceptURL" conform:"trim"`
	RejectURL string `json:"rejectURL" conform:"trim"`
}

// EventValues ...
type EventValues struct {
	NewValue      float64 `json:"newValue"`
	PreviousValue float64 `json:"previousValue"`
}

// EventStatusValues ...
type EventStatusValues struct {
	NewValue      EventStatus `json:"newValue"`
	PreviousValue EventStatus `json:"previousValue"`
}

// EventResource ...
type EventResource struct {
	CostPerInterview    *EventValues       `json:"costPerInterview"`
	EstimatedCost       *EventValues       `json:"estimatedCost"`
	LengthOfInterview   *EventValues       `json:"lengthOfInterview"`
	IndicativeIncidence *EventValues       `json:"incidenceRate"`
	Currency            string             `json:"currency" conform:"trim"`
	Status              *EventStatusValues `json:"status"`
	Reason              string             `json:"reason" conform:"trim"`
}

// Event ...
type Event struct {
	EventID       int64          `json:"eventId"`
	EventType     EventType      `json:"eventType"`
	ExtProjectID  string         `json:"extProjectId"`
	ExtLineItemID string         `json:"extLineItemId"`
	Resource      *EventResource `json:"resource"`
	Actions       *EventActions  `json:"actions"`
	CreatedAt     CustomTime     `json:"createdAt"`
	ParentEventID *int64         `json:"parentEventId,omitempty"`
}
