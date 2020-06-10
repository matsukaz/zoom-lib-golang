package zoom

import "fmt"

const (
	// ListMetricsMeetingParticipantsPath - v2 path for listing metrics meeting participants
	ListMetricsMeetingParticipantsPath = "/metrics/meetings/%d/participants"
)

type (
	// ListMetricsMeetingParticipantsOptions contains options for ListMetricsMeetingParticipants
	ListMetricsMeetingParticipantsOptions struct {
		MeetingID     int64               `url:"-"`
		Type          *MetricsMeetingType `url:"type,omitempty"`
		PageSize      int                 `url:"page_size"`
		NextPageToken string              `url:"next_page_token,omitempty"`
	}

	// ListMetricsMeetingParticipantsResponse container the response from a call to ListMetricsMeetingParticipants
	ListMetricsMeetingParticipantsResponse struct {
		PageCount     int                         `json:"page_count"`
		PageSize      int                         `json:"page_size"`
		TotalRecords  int                         `json:"total_records"`
		NextPageToken string                      `url:"next_page_token,omitempty"`
		Participants  []MetricsMeetingParticipant `json:"participants"`
	}
)

// ListMetricsMeetingParticipants calls /metrics/meetings/ID/participants
// https://marketplace.zoom.us/docs/api-reference/zoom-api/dashboards/dashboardmeetingparticipants
func (c *Client) ListMetricsMeetingParticipants(opts ListMetricsMeetingParticipantsOptions) (ListMetricsMeetingParticipantsResponse, error) {
	var ret = ListMetricsMeetingParticipantsResponse{}
	return ret, c.requestV2(requestV2Opts{
		Method:        Get,
		Path:          fmt.Sprintf(ListMetricsMeetingParticipantsPath, opts.MeetingID),
		URLParameters: &opts,
		Ret:           &ret,
	})
}
