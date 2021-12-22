package varification

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

func SendSMS(to string, msg string) bool {
	accountSid := os.Getenv("ACCOUNT_SID")
	authToken := os.Getenv("AUTH_TOKEN")
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"

	// Set up rand
	rand.Seed(time.Now().Unix())

	msgData := url.Values{}
	msgData.Set("To", "+91"+to)
	msgData.Set("From", os.Getenv("TWILIO_NUMBER"))
	msgData.Set("Body", msg)
	msgDataReader := *strings.NewReader(msgData.Encode())

	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &msgDataReader)
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, _ := client.Do(req)
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var data map[string]interface{}
		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(&data)
		if err == nil {
			return true
		}
	}
	return false
}
