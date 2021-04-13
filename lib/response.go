package samplify

// StatusType ...
type StatusType string

// StatusType values
const (
	StatusTypeSuccess StatusType = "success"
	StatusTypeFail    StatusType = "fail"
	StatusTypeUnknown StatusType = "unknown"
)

// ProjectResponse ...
type ProjectResponse struct {
	Project        *Project       `json:"data"`
	ResponseStatus ResponseStatus `json:"status"`
}

// BuyProjectResponse represents the response from Buy Project request
type BuyProjectResponse struct {
	List           []*BuyProjectLineItem `json:"data"`
	ResponseStatus ResponseStatus        `json:"status"`
}

// GetAllProjectsResponse ...
type GetAllProjectsResponse struct {
	Projects       []*ProjectHeader `json:"data"`
	ResponseStatus ResponseStatus   `json:"status"`
	Meta           Meta             `json:"meta"`
}

// ProjectReportResponse ...
type ProjectReportResponse struct {
	Report         *ProjectReport `json:"data"`
	ResponseStatus ResponseStatus `json:"status"`
}

// CloseProjectResponse ...
type CloseProjectResponse struct {
	Project *struct {
		ProjectHeader
		LineItems []*LineItemHeader `json:"lineItems"`
	} `json:"data"`
	ResponseStatus ResponseStatus `json:"status"`
}

// LineItemResponse ... Response returned by Add, Update and Get LineItem requests
type LineItemResponse struct {
	Item           *LineItem      `json:"data"`
	ResponseStatus ResponseStatus `json:"status"`
}

// UpdateLineItemStateResponse ...
type UpdateLineItemStateResponse struct {
	LineItem       *LineItem      `json:"data"`
	ResponseStatus ResponseStatus `json:"status"`
}

// LineItemListItem ...
type LineItemListItem struct {
	Model
	ExtLineItemID   string      `json:"extLineItemId"`
	State           State       `json:"state"`
	StateReason     string      `json:"stateReason"`
	LaunchedAt      *CustomTime `json:"launchedAt"`
	Title           string      `json:"title"`
	CountryISOCode  string      `json:"countryISOCode"`
	LanguageISOCode string      `json:"languageISOCode"`
}

// GetAllLineItemsResponse ...
type GetAllLineItemsResponse struct {
	List           []*LineItemListItem `json:"data"`
	ResponseStatus ResponseStatus      `json:"status"`
	Meta           Meta                `json:"meta"`
}

// GetFeasibilityResponse ...
type GetFeasibilityResponse struct {
	List []*struct {
		ExtLineItemID string       `json:"extLineItemId"`
		Feasibility   *Feasibility `json:"feasibility"`
		Quote         Quote        `json:"quote"`
	} `json:"data"`
	ResponseStatus ResponseStatus `json:"status"`
}

// GetCountriesResponse ...
type GetCountriesResponse struct {
	List           []*Country     `json:"data"`
	ResponseStatus ResponseStatus `json:"status"`
	Meta           Meta           `json:"meta"`
}

// GetAttributesResponse ...
type GetAttributesResponse struct {
	List           []*Attribute   `json:"data"`
	ResponseStatus ResponseStatus `json:"status"`
	Meta           Meta           `json:"meta"`
}

// GetSurveyTopicsResponse ...
type GetSurveyTopicsResponse struct {
	List           []*SurveyTopic `json:"data"`
	ResponseStatus ResponseStatus `json:"status"`
	Meta           Meta           `json:"meta"`
}

// GetEventListResponse ...
type GetEventListResponse struct {
	List           []*Event       `json:"data"`
	ResponseStatus ResponseStatus `json:"status"`
	Meta           Meta           `json:"meta"`
}

// GetEventResponse ...
type GetEventResponse struct {
	Event          *Event         `json:"data"`
	ResponseStatus ResponseStatus `json:"status"`
}

// DetailedProjectReportResponse ...
type DetailedProjectReportResponse struct {
	Report         DetailedProjectReport `json:"data"`
	ResponseStatus ResponseStatus        `json:"status"`
	Meta           Meta                  `json:"meta"`
}

// DetailedLineItemReportResponse ...
type DetailedLineItemReportResponse struct {
	Report         DetailedLineItemReport `json:"data"`
	ResponseStatus ResponseStatus         `json:"status"`
	Meta           Meta                   `json:"meta"`
}

// StudyMetadataResponse ...
type StudyMetadataResponse struct {
	StudyMetadata  StudyMetadata  `json:"data"`
	ResponseStatus ResponseStatus `json:"status"`
	Meta           Meta           `json:"meta"`
}

// QuotaCellResponse ...
type QuotaCellResponse struct {
	QuotaCell      QuotaCell      `json:"data"`
	ResponseStatus ResponseStatus `json:"status"`
	Meta           Meta           `json:"meta"`
}

// OrderDetailResponseData ...
type OrderDetailResponse struct {
	OrderDetail    OrderDetail    `json:"data"`
	ResponseStatus ResponseStatus `json:"status"`
	Meta           Meta           `json:"meta"`
}

// CheckOrderNumber ...
type CheckOrderNumberResponse struct {
	CheckOrderNumber CheckOrderNumber `json:"data"`
	ResponseStatus   ResponseStatus   `json:"status"`
	Meta             Meta             `json:"meta"`
}

type CheckOrderNumber struct {
	Availability bool `json:"availability"`
}

type OrderDetail struct {
	// sales order
	SalesOrder *SalesOrder `form:"salesOrder" json:"salesOrder" yaml:"salesOrder" xml:"salesOrder"`
	// sales details
	SalesOrderDetails []*SalesOrderDetails `form:"salesOrderDetails" json:"salesOrderDetails" yaml:"salesOrderDetails" xml:"salesOrderDetails"`
}

// SalesOrder user type.
type SalesOrder struct {
	// basicSecurityKey
	BasicSecurityKey string `form:"basicSecurityKey" json:"basicSecurityKey" yaml:"basicSecurityKey" xml:"basicSecurityKey"`
	// Unique guid
	GUID string `form:"guid" json:"guid" yaml:"guid" xml:"guid"`
	// highSecurityKey
	HighSecurityKey string `form:"highSecurityKey" json:"highSecurityKey" yaml:"highSecurityKey" xml:"highSecurityKey"`
	// name
	Name string `form:"name" json:"name" yaml:"name" xml:"name"`
	// noCharge
	NoCharge bool `form:"noCharge" json:"noCharge" yaml:"noCharge" xml:"noCharge"`
	// orderType
	OrderType string `form:"orderType" json:"orderType" yaml:"orderType" xml:"orderType"`
	// ordernumber
	Ordernumber string `form:"ordernumber" json:"ordernumber" yaml:"ordernumber" xml:"ordernumber"`
	// relatedOrderCpi
	RelatedOrderCpi float64 `form:"relatedOrderCpi" json:"relatedOrderCpi" yaml:"relatedOrderCpi" xml:"relatedOrderCpi"`
	// secureEndLinkLevel
	SecureEndLinkLevel int `form:"secureEndLinkLevel" json:"secureEndLinkLevel" yaml:"secureEndLinkLevel" xml:"secureEndLinkLevel"`
	// secureEndLinkLevelName
	SecureEndLinkLevelName string `form:"secureEndLinkLevelName" json:"secureEndLinkLevelName" yaml:"secureEndLinkLevelName" xml:"secureEndLinkLevelName"`
}

// SalesOrderDetails ...
type SalesOrderDetails struct {
	// costPerInterview
	CostPerInterview float64 `form:"costPerInterview" json:"costPerInterview" yaml:"costPerInterview" xml:"costPerInterview"`
	// costPerInterviewWithCurrency
	CostPerInterviewWithCurrency string `form:"costPerInterviewWithCurrency" json:"costPerInterviewWithCurrency" yaml:"costPerInterviewWithCurrency" xml:"costPerInterviewWithCurrency"`
	// countryIsoCode
	CountryIsoCode string `form:"countryIsoCode" json:"countryIsoCode" yaml:"countryIsoCode" xml:"countryIsoCode"`
	// extendedamount
	Extendedamount string `form:"extendedamount" json:"extendedamount" yaml:"extendedamount" xml:"extendedamount"`
	// guid
	GUID string `form:"guid" json:"guid" yaml:"guid" xml:"guid"`
	// labelForMobile
	LabelForMobile string `form:"labelForMobile" json:"labelForMobile" yaml:"labelForMobile" xml:"labelForMobile"`
	// productIdGuid
	ProductIDGUID string `form:"productIdGuid" json:"productIdGuid" yaml:"productIdGuid" xml:"productIdGuid"`
	// productIdName
	ProductIDName string `form:"productIdName" json:"productIdName" yaml:"productIdName" xml:"productIdName"`
	// quantity
	Quantity int `form:"quantity" json:"quantity" yaml:"quantity" xml:"quantity"`
	// ssiAdditionalPoints
	SsiAdditionalPoints int `form:"ssiAdditionalPoints" json:"ssiAdditionalPoints" yaml:"ssiAdditionalPoints" xml:"ssiAdditionalPoints"`
	// ssiCalculatedIr
	SsiCalculatedIr string `form:"ssiCalculatedIr" json:"ssiCalculatedIr" yaml:"ssiCalculatedIr" xml:"ssiCalculatedIr"`
	// ssiCalculatedLoi
	SsiCalculatedLoi string `form:"ssiCalculatedLoi" json:"ssiCalculatedLoi" yaml:"ssiCalculatedLoi" xml:"ssiCalculatedLoi"`
	// ssiChartsNum
	SsiChartsNum int `form:"ssiChartsNum" json:"ssiChartsNum" yaml:"ssiChartsNum" xml:"ssiChartsNum"`
	// ssiFamilyId
	SsiFamilyID string `form:"ssiFamilyId" json:"ssiFamilyId" yaml:"ssiFamilyId" xml:"ssiFamilyId"`
	// ssiImagesNum
	SsiImagesNum int `form:"ssiImagesNum" json:"ssiImagesNum" yaml:"ssiImagesNum" xml:"ssiImagesNum"`
	// ssiImagesSpecialNum
	SsiImagesSpecialNum int `form:"ssiImagesSpecialNum" json:"ssiImagesSpecialNum" yaml:"ssiImagesSpecialNum" xml:"ssiImagesSpecialNum"`
	// ssiInputPrice
	SsiInputPrice float64 `form:"ssiInputPrice" json:"ssiInputPrice" yaml:"ssiInputPrice" xml:"ssiInputPrice"`
	// ssiIr
	SsiIr int `form:"ssiIr" json:"ssiIr" yaml:"ssiIr" xml:"ssiIr"`
	// ssiLabel
	SsiLabel string `form:"ssiLabel" json:"ssiLabel" yaml:"ssiLabel" xml:"ssiLabel"`
	// ssiProductType
	SsiProductType string `form:"ssiProductType" json:"ssiProductType" yaml:"ssiProductType" xml:"ssiProductType"`
	// ssiProductTypeId
	SsiProductTypeID int `form:"ssiProductTypeId" json:"ssiProductTypeId" yaml:"ssiProductTypeId" xml:"ssiProductTypeId"`
	// ssiSampleCountryCode
	SsiSampleCountryCode string `form:"ssiSampleCountryCode" json:"ssiSampleCountryCode" yaml:"ssiSampleCountryCode" xml:"ssiSampleCountryCode"`
	// ssiSampleCountryId
	SsiSampleCountryID string `form:"ssiSampleCountryId" json:"ssiSampleCountryId" yaml:"ssiSampleCountryId" xml:"ssiSampleCountryId"`
	// ssiTitle
	SsiTitle string `form:"ssiTitle" json:"ssiTitle" yaml:"ssiTitle" xml:"ssiTitle"`
	// ssiVendorUsed
	SsiVendorUsed string `form:"ssiVendorUsed" json:"ssiVendorUsed" yaml:"ssiVendorUsed" xml:"ssiVendorUsed"`
	// ssiVideosNum
	SsiVideosNum int `form:"ssiVideosNum" json:"ssiVideosNum" yaml:"ssiVideosNum" xml:"ssiVideosNum"`
	// vendorLine
	VendorLine bool `form:"vendorLine" json:"vendorLine" yaml:"vendorLine" xml:"vendorLine"`
}

// ResponseStatus is the custom status part in API response. (Optional in some endpoints)
type ResponseStatus struct {
	Message string      `json:"message"`
	Errors  []ErrorInfo `json:"errors"`
}

// ErrorInfo ... Custom API errors
type ErrorInfo struct {
	Code     string   `json:"code"`
	Message  string   `json:"message"`
	Resource Resource `json:"resource"`
}

// Resource ...
type Resource struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

// Meta ...
type Meta struct {
	Links    `json:"links"`
	Total    int64 `json:"total"`
	PageSize int64 `json:"pageSize"`
}

// Links for page navigation
type Links struct {
	First string `json:"first"`
	Last  string `json:"last"`
	Next  string `json:"next"`
	Prev  string `json:"prev"`
	Self  string `json:"self"`
}

// Get ... Reads "message" from API's custom success/error response and interprets the status
func (s *ResponseStatus) Get() StatusType {
	switch s.Message {
	case "success":
		return StatusTypeSuccess
	case "fail":
		return StatusTypeFail
	default:
		return StatusTypeUnknown
	}
}
