package zoom // types for /metrics endpoints

type MetricsMeetingType string

const (
	Live    MetricsMeetingType = "live"
	Past    MetricsMeetingType = "past"
	PastOne MetricsMeetingType = "pastOne"
)

type MetricsMeeting struct {
	UUID               string `json:"uuid,omitempty"`
	ID                 int64  `json:"id,omitempty"`
	Topic              string `json:"topic"`
	Host               string `json:"host,omitempty"`
	Email              string `json:"email"`
	UserType           string `json:"user_type"`
	StartTime          *Time  `json:"start_time"`
	EndTime            *Time  `json:"end_time"`
	Duration           string `json:"duration"`
	Participants       int    `json:"participants"`
	HasPstn            bool   `json:"has_pstn"`
	HasVoip            bool   `json:"has_voip"`
	Has3rdPartyAudio   bool   `json:"has_3rd_party_audio"`
	HasVideo           bool   `json:"has_video"`
	HasScreenShare     bool   `json:"has_screen_share"`
	HasRecording       bool   `json:"has_recording"`
	HasSip             bool   `json:"has_sip"`
	InRoomParticipants int    `json:"in_room_participants,omitempty"`
	Dept               string `json:"dept,omitempty"`
}

type MetricsMeetingParticipant struct {
	ID                 string `json:"id"`
	UserID             string `json:"user_id"`
	UserName           string `json:"user_name"`
	Device             string `json:"device"`
	IpAddress          string `json:"ip_address"`
	Location           string `json:"location"`
	NetworkType        string `json:"network_type"`
	Microphone         string `json:"microphone"`
	Speaker            string `json:"speaker"`
	DataCenter         string `json:"data_center"`
	ConnectionType     string `json:"connection_type"`
	JoinTime           *Time  `json:"join_time"`
	LeaveTime          *Time  `json:"leave_time"`
	ShareApplication   bool   `json:"share_application"`
	ShareDesktop       bool   `json:"share_desktop"`
	ShareWhiteboard    bool   `json:"share_whiteboard"`
	Recording          bool   `json:"recording"`
	PcName             string `json:"pc_name"`
	Domain             string `json:"domain"`
	MacAddr            string `json:"mac_addr"`
	HarddiskID         string `json:"harddisk_id"`
	Version            string `json:"version"`
	InRoomParticipants int    `json:"in_room_participants"`
	LeaveReason        string `json:"leave_reason"`
}
