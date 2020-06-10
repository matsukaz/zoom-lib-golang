package zoom

const (
	// ListMetricsMeetingsPath - v2 path for listing metrics meetings
	ListMetricsMeetingsPath = "/metrics/meetings"
)

type (
	// ListMetricsMeetingOptions contains options for ListMetricsMeetings
	ListMetricsMeetingsOptions struct {
		Type          *MetricsMeetingType `url:"type,omitempty"`
		PageSize      int                 `url:"page_size"`
		From          string              `url:"from,omitempty"`
		To            string              `url:"to,omitempty"`
		NextPageToken string              `url:"next_page_token,omitempty"`
	}

	// ListMetricsMeetingsResponse container the response from a call to ListMetricsMeetings
	ListMetricsMeetingsResponse struct {
		From          string           `url:"from,omitempty"`
		To            string           `url:"to,omitempty"`
		PageCount     int              `json:"page_count"`
		PageSize      int              `json:"page_size"`
		TotalRecords  int              `json:"total_records"`
		NextPageToken string           `url:"next_page_token,omitempty"`
		Meetings      []MetricsMeeting `json:"meetings"`
	}
)

// ListMetricsMeetings calls /metrics/meetings
// https://marketplace.zoom.us/docs/api-reference/zoom-api/dashboards/dashboardmeetings
func (c *Client) ListMetricsMeetings(opts ListMetricsMeetingsOptions) (ListMetricsMeetingsResponse, error) {
	var ret = ListMetricsMeetingsResponse{}
	return ret, c.requestV2(requestV2Opts{
		Method:        Get,
		Path:          ListMetricsMeetingsPath,
		URLParameters: &opts,
		Ret:           &ret,
	})
}
