package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/sanity-io/litter"
)

// VERSION of the application
const VERSION = "0.1.1"

var (
	// CommitHash is the revision hash of the build's git repository
	CommitHash string
	// BuildTime is the build's time
	BuildTime string
)

// Handler handles lex events
func Handler(ctx context.Context, event events.LexEvent) (*events.LexResponse, error) {
	fmt.Printf("Received an input from Amazon Lex. Current Intent: %s", event.CurrentIntent.Name)
	litter.Dump(event)

	messageContent := "Hello from AWS Lambda!"

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
