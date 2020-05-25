package alert

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type TwilioAlerter struct {
	lastMsg  time.Time
	sendTo   string
	sendFrom string
	sid      string
	token    string
	url      string
}

func NewTwilioAlerter(sid, token, sendTo, sendFrom string) (Alerter, error) {
	if err := validation.Validate(&sid, validation.Required); err != nil {
		return nil, fmt.Errorf("sid %s", err)
	}
	if err := validation.Validate(&token, validation.Required); err != nil {
		return nil, fmt.Errorf("token %s", err)
	}
	if err := validation.Validate(&sendTo, validation.Required); err != nil {
		return nil, fmt.Errorf("sendTo %s", err)
	}
	if err := validation.Validate(&sendFrom, validation.Required); err != nil {
		return nil, fmt.Errorf("sisendFromd %s", err)
	}

	return &TwilioAlerter{
		lastMsg:  time.Time{},
		sid:      sid,
		token:    token,
		sendTo:   sendTo,
		sendFrom: sendFrom,
		url:      fmt.Sprintf("https://api.twilio.com/2010-04-01/Accounts/%s/Messages.json", sid),
	}, nil
}

func (a *TwilioAlerter) Alert(msg string) error {
	if time.Since(a.lastMsg) < time.Minute*10 {
		return errors.New("message not sent, throttled")
	}
	// setup msg
	msgData := url.Values{}
	msgData.Set("To", a.sendTo)
	msgData.Set("From", a.sendFrom)
	msgData.Set("Body", msg)
	msgDataReader := *strings.NewReader(msgData.Encode())

	// setup client
	client := &http.Client{}
	req, err := http.NewRequest("POST", a.url, &msgDataReader)
	if err != nil {
		return err
	}
	req.SetBasicAuth(a.sid, a.token)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// send
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		a.lastMsg = time.Now()
		return nil
	}

	return fmt.Errorf("failed to send message: %s", string(b))
}
