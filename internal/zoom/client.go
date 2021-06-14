package zoom

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type Client struct {
	token            string
	secret           string
	defaultMeetingID string
}

func NewClient(token, secret, defaultMeetingID string) *Client {
	return &Client{token, secret, defaultMeetingID}
}

const apiPath = "https://api.zoom.us/v2"
const jwtExpiredIn time.Duration = 90 * time.Minute

func (c *Client) GetMeetingParticipants(meetingID string) ([]Participant, error) {
	mID := meetingID
	if meetingID == "" {
		mID = c.defaultMeetingID
	}
	if mID == "" {
		return nil, errors.New("Meeting id can't be blank")
	}
	url := apiPath + "/metrics/meetings/" + mID + "/participants"

	token, err := c.generateJwtToken()
	if err != nil {
		return nil, err
	}
	bearer := "Bearer " + token

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", bearer)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	pResp := &ParticipantsResp{}
	err = json.NewDecoder(resp.Body).Decode(pResp)
	if err != nil {
		return nil, err
	}

	return pResp.Partisipants, nil
}

func (c *Client) generateJwtToken() (string, error) {
	claims := jwt.StandardClaims{
		Issuer:    c.token,
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(jwtExpiredIn).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtString, err := token.SignedString([]byte(c.secret))

	if err != nil {
		return "", err
	}

	return jwtString, nil
}
