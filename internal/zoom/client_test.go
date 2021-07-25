package zoom

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testToken     = "token"
	testSecret    = "secret"
	testMeetingID = "meetingID"
)

var validMessage = []byte(`
{
	"page_count": 1,
	"page_size": 30,
	"total_records": 2,
	"next_page_token": "",
	"participants": [
		{
			"id": "d52f19c548b88490b5d16fcbd38",
			"user_id": "32dsfsd4g5gd",
			"user_name": "dojo",
			"device": "WIN",
			"ip_address": "127.0.0.1",
			"location": "New York",
			"network_type": "Wired",
			"microphone": "Plantronics BT600",
			"camera": "FaceTime HD Camera",
			"speaker": "Plantronics BT600",
			"data_center": "SC",
			"connection_type": "P2P",
			"join_time": "2019-09-07T13:15:02.837Z",
			"leave_time": "2019-09-07T13:15:09.837Z",
			"share_application": false,
			"share_desktop": true,
			"share_whiteboard": true,
			"recording": false,
			"status": "in_waiting_room",
			"pc_name": "dojo's pc",
			"domain": "Dojo-workspace",
			"mac_addr": " 00:0a:95:9d:68:16",
			"harddisk_id": "sed proident in",
			"version": "4.4.55383.0716",
			"leave_reason": "Dojo left the meeting.<br>Reason: Host ended the meeting."
		},
		{
			"id": "z8aaaaaaCfp8uQ",
			"user_id": "1670000000",
			"user_name": "Rea",
			"device": "Android",
			"ip_address": "120.000.000",
			"location": "San Jose (US)",
			"network_type": "Wifi",
			"data_center": "SC",
			"connection_type": "UDP",
			"join_time": "2019-08-02T15:31:48Z",
			"leave_time": "2019-08-02T16:04:12Z",
			"share_application": false,
			"share_desktop": false,
			"share_whiteboard": false,
			"recording": false,
			"pc_name": "Rea's PC",
			"domain": "Rea-workspace",
			"mac_addr": "",
			"harddisk_id": "",
			"version": "4.4.55383.0716",
			"leave_reason": "Rea left the meeting.<br>Reason: Host closed the meeting."
		}
	]
}`)

type TestCase struct {
	Name             string
	MeetingID        string
	DefaultMeetingID string
	Error            error
	HttpStatus       int
	HttpBody         []byte
	Result           []Participant
}

func TestGetMeetingParticipants(t *testing.T) {
	testCases := []TestCase{
		{
			Name:             "Bad request",
			MeetingID:        "error",
			DefaultMeetingID: testMeetingID,
			Error:            errors.New("invalid request"),
			Result:           nil,
			HttpBody:         []byte(`{"code": 500, "message": "invalid request"}`),
			HttpStatus:       500,
		},
		{
			Name:             "Empty meeting id",
			MeetingID:        "",
			DefaultMeetingID: "",
			Error:            errors.New("Meeting id can't be blank"),
			Result:           nil,
		},
		{
			Name:             "Valid request",
			MeetingID:        "custom",
			DefaultMeetingID: testMeetingID,
			Result: []Participant{
				{"dojo"},
				{"Rea"},
			},
			HttpBody:   validMessage,
			HttpStatus: 200,
		},
		{
			Name:             "Empty meeting id but valid default",
			MeetingID:        "",
			DefaultMeetingID: testMeetingID,
			Result: []Participant{
				{"dojo"},
				{"Rea"},
			},
			HttpBody:   validMessage,
			HttpStatus: 200,
		},
	}

	for _, tc := range testCases {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(tc.HttpStatus)
			w.Write(tc.HttpBody)
		}))
		client := NewClient(testToken, testSecret, tc.DefaultMeetingID, ts.URL)

		participants, err := client.GetMeetingParticipants(tc.MeetingID)
		if tc.Error != nil {
			assert.Equal(t, tc.Error, err, "%s test failed! Expected error: %+v. Received error: %+v", tc.Name, tc.Error, err)
		} else {
			assert.Equal(t, nil, err, "%s test failed! Expect no error got %+v", tc.Name, err)
		}
		assert.Equal(t, tc.Result, participants, "%s test failed! Expected result %+v got %+v", tc.Name, tc.Result, participants)
	}
}
