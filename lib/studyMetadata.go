package samplify

// StudyMetadata ...
type StudyMetadata struct {
	Category      CategoryMetadata `json:"category"`
	DeliveryTypes []MetadataItem   `json:"deliveryTypes"`
}

// CategoryMetadata ...
type CategoryMetadata struct {
	StudyRequirements []MetadataItem `json:"studyRequirements"`
	StudyTypes        []MetadataItem `json:"studyTypes"`
	SurveyTopics      []MetadataItem `json:"surveyTopics"`
}

// MetadataItem ...
type MetadataItem struct {
	Allowed     bool   `json:"allowed"`
	Description string `json:"description"`
	ID          string `json:"id"`
	Name        string `json:"name"`
}
