package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// VERSION of the application
const VERSION = "0.1.3"

var (
	// CommitHash is the revision hash of the build's git repository
	CommitHash string
	// BuildTime is the build's time
	BuildTime string
)

// Handler handles lex events
func Handler(ctx context.Context, event events.LexEvent) (*events.LexResponse, error) {

	bs, err := json.Marshal(event)
	if err != nil {
		return &events.LexResponse{}, err
	}

	fmt.Printf("Received an input from Amazon Lex. Current Intent: %s json: %s", event.CurrentIntent.Name, string(bs))

	if event.CurrentIntent.Name != "Deployit" {
		return &events.LexResponse{}, fmt.Errorf("Unsupported intent %s", event.CurrentIntent.Name)
	}

	project := ""
	for _, s := range event.CurrentIntent.Slots {
		fmt.Printf("slot: %v", s)
	}

	err = deployIt(project, "tester")
	if err != nil {
		return &events.LexResponse{}, err
	}

	messageContent := "Deployment of " + project + " started"

	return &events.LexResponse{
		SessionAttributes: event.SessionAttributes,
		DialogAction: events.LexDialogAction{
			Type: "Close",
			Message: map[string]string{
				"content":     messageContent,
				"contentType": "PlainText",
			},
			FulfillmentState: "Fulfilled",
		},
	}, nil
}

func main() {
	lambda.Start(Handler)
}
