package zoom

type ParticipantsResp struct {
	Partisipants []Participant `json:"participants"`
	Code         int32         `json:"code"`
	Message      string        `json:"message"`
}

type Participant struct {
	UserName string `json:"user_name"`
}
