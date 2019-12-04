package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(ctx context.Context, event events.LexEvent) (*events.LexResponse, error) {
	fmt.Printf("Received an input from Amazon Lex. Current Intent: %s", event.CurrentIntent.Name)

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
