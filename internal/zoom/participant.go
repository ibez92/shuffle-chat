package zoom

type ParticipantsResp struct {
	Partisipants []Participant `json:"participants"`
}

type Participant struct {
	UserName string `json:"user_name"`
}
