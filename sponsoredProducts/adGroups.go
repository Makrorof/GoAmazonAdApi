package sponsoredProducts

type IListAdGroupsRequest interface {
	getNewResponse() IListAdGroupsResponse
	getVersion() int
}
type IListAdGroupsResponse interface {
	getVersion() int
}

type ICreateAdGroupsRequest interface {
	getNewResponse() ICreateAdGroupsResponse
	getVersion() int
}
type ICreateAdGroupsResponse interface {
	getVersion() int
}

type IDeleteAdGroupsRequest interface {
	getNewResponse() IDeleteAdGroupsResponse
	getVersion() int
}
type IDeleteAdGroupsResponse interface {
	getVersion() int
}

type IUpdateAdGroupsRequest interface {
	getNewResponse() IUpdateAdGroupsResponse
	getVersion() int
}
type IUpdateAdGroupsResponse interface {
	getVersion() int
}

// region CreateV3
type CreateAdGroupsRequestV3Data struct {
	CampaignID string  `json:"campaignId,omitempty"`
	Name       string  `json:"name,omitempty"`
	State      string  `json:"state,omitempty"`
	DefaultBid float64 `json:"defaultBid,omitempty"`
}
type CreateAdGroupsRequestV3 struct {
	AdGroups []CreateAdGroupsRequestV3Data `json:"adGroups,omitempty"`
}
type CreateAdGroupsResponseV3Success struct {
	AdGroup struct {
		CampaignID   string        `json:"campaignId,omitempty"`
		Name         string        `json:"name,omitempty"`
		State        string        `json:"state,omitempty"`
		AdGroupID    string        `json:"adGroupId,omitempty"`
		DefaultBid   float64       `json:"defaultBid,omitempty"`
		ExtendedData *ExtendedData `json:"extendedData,omitempty"`
	} `json:"adGroup,omitempty"`
	Index     int    `json:"index,omitempty"`
	AdGroupID string `json:"adGroupId,omitempty"`
}
type CreateAdGroupsResponseV3 struct {
	AdGroups struct {
		Success []CreateAdGroupsResponseV3Success `json:"success,omitempty"`
		Error   []Error                           `json:"error,omitempty"`
	} `json:"adGroups,omitempty"`
}

func (r *CreateAdGroupsRequestV3) getNewResponse() ICreateAdGroupsResponse {
	return new(CreateAdGroupsResponseV3)
}

func (r *CreateAdGroupsRequestV3) getVersion() int {
	return 3
}

func (r *CreateAdGroupsResponseV3) getVersion() int {
	return 3
}

//endregion CreateV3

// region UpdateV3
type UpdateAdGroupsRequestV3Data struct {
	Name       string  `json:"name,omitempty"`
	State      string  `json:"state,omitempty"`
	AdGroupID  string  `json:"adGroupId,omitempty"`
	DefaultBid float64 `json:"defaultBid,omitempty"`
}
type UpdateAdGroupsRequestV3 struct {
	AdGroups []UpdateAdGroupsRequestV3Data `json:"adGroups,omitempty"`
}
type UpdateAdGroupsResponseV3Success struct {
	AdGroup struct {
		CampaignID   string        `json:"campaignId,omitempty"`
		Name         string        `json:"name,omitempty"`
		State        string        `json:"state,omitempty"`
		AdGroupID    string        `json:"adGroupId,omitempty"`
		DefaultBid   float64       `json:"defaultBid,omitempty"`
		ExtendedData *ExtendedData `json:"extendedData,omitempty"`
	} `json:"adGroup,omitempty"`
	Index     int    `json:"index,omitempty"`
	AdGroupID string `json:"adGroupId,omitempty"`
}
type UpdateAdGroupsResponseV3 struct {
	AdGroups struct {
		Success []UpdateAdGroupsResponseV3Success `json:"success,omitempty"`
		Error   []Error                           `json:"error,omitempty"`
	} `json:"adGroups,omitempty"`
}

func (r *UpdateAdGroupsRequestV3) getNewResponse() IUpdateAdGroupsResponse {
	return new(UpdateAdGroupsResponseV3)
}

func (r *UpdateAdGroupsRequestV3) getVersion() int {
	return 3
}

func (r *UpdateAdGroupsResponseV3) getVersion() int {
	return 3
}

//endregion UpdateV3

// region DeleteV3
type DeleteAdGroupsRequestV3 struct {
	AdGroupIDFilter *IDFilter `json:"adGroupIdFilter,omitempty"`
}
type DeleteAdGroupsResponseV3Success struct {
	AdGroup struct {
		CampaignID   string        `json:"campaignId,omitempty"`
		Name         string        `json:"name,omitempty"`
		State        string        `json:"state,omitempty"`
		AdGroupID    string        `json:"adGroupId,omitempty"`
		DefaultBid   float64       `json:"defaultBid,omitempty"`
		ExtendedData *ExtendedData `json:"extendedData,omitempty"`
	} `json:"adGroup,omitempty"`
	Index     int    `json:"index,omitempty"`
	AdGroupID string `json:"adGroupId,omitempty"`
}
type DeleteAdGroupsResponseV3 struct {
	AdGroups struct {
		Success []DeleteAdGroupsResponseV3Success `json:"success,omitempty"`
		Error   []Error                           `json:"error,omitempty"`
	} `json:"adGroups,omitempty"`
}

func (r *DeleteAdGroupsRequestV3) getNewResponse() IDeleteAdGroupsResponse {
	return new(DeleteAdGroupsResponseV3)
}

func (r *DeleteAdGroupsRequestV3) getVersion() int {
	return 3
}

func (r *DeleteAdGroupsResponseV3) getVersion() int {
	return 3
}

//endregion DeleteV3

// region ListV3
type ListAdGroupsRequestV3 struct {
	CampaignIDFilter            *IDFilter   `json:"campaignIdFilter,omitempty"`
	StateFilter                 *IDFilter   `json:"stateFilter,omitempty"`
	MaxResults                  int         `json:"maxResults,omitempty"`
	NextToken                   string      `json:"nextToken,omitempty"`
	AdGroupIDFilter             *IDFilter   `json:"adGroupIdFilter,omitempty"`
	IncludeExtendedDataFields   bool        `json:"includeExtendedDataFields,omitempty"`
	NameFilter                  *NameFilter `json:"nameFilter,omitempty"`
	CampaignTargetingTypeFilter string      `json:"campaignTargetingTypeFilter,omitempty"`
}
type ListAdGroupsResponseV3Data struct {
	CampaignID   string        `json:"campaignId,omitempty"`
	Name         string        `json:"name,omitempty"`
	State        string        `json:"state,omitempty"`
	AdGroupID    string        `json:"adGroupId,omitempty"`
	DefaultBid   float64       `json:"defaultBid,omitempty"`
	ExtendedData *ExtendedData `json:"extendedData,omitempty"`
}
type ListAdGroupsResponseV3 struct {
	TotalResults int                          `json:"totalResults,omitempty"`
	AdGroups     []ListAdGroupsResponseV3Data `json:"adGroups,omitempty"`
	NextToken    string                       `json:"nextToken,omitempty"`
}

func (r *ListAdGroupsRequestV3) getNewResponse() IListAdGroupsResponse {
	return new(ListAdGroupsResponseV3)
}

func (r *ListAdGroupsRequestV3) getVersion() int {
	return 3
}

func (r *ListAdGroupsResponseV3) getVersion() int {
	return 3
}

//endregion ListV3
