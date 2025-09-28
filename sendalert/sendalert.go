package sendalert

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const alertChannel = "monitoring"

//var slackWebhookURL = os.Getenv("SLACK_WEBHOOK_URL")

// Slack webhook URL 
const slackWebhookURL = "XOXOX"

func SendSlackAlert(message string) error {
	//slackWebhookURL := os.Getenv("SLACK_WEBHOOK_URL")
	payload := map[string]string{"text": message}
	data, _ := json.Marshal(payload)

	resp, err := http.Post(slackWebhookURL, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("Slack return non-200 status: %d", resp.Status)
	}
	return nil
}

func sendAlert(message string) {
	if alertChannel == "monitoring" {
		if err := SendSlackAlert(message); err != nil {
			fmt.Println("Slack error:", err)
		}

	}
}
