package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/sirupsen/logrus"
)

func deployIt(project, requestor string) error {

	qURL := "	https://sqs.eu-west-1.amazonaws.com/671762830139/deployit"
	svcSqs := sqs.New(session.New())

	_, err := svcSqs.SendMessage(&sqs.SendMessageInput{
		DelaySeconds: aws.Int64(10),
		MessageAttributes: map[string]*sqs.MessageAttributeValue{
			"project": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String(project),
			},
			"requestor": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String(requestor),
			},
		},
		MessageBody: aws.String("Deploy project"),
		QueueUrl:    &qURL,
	})
	if err != nil {
		logrus.Error(err)
	}

	return err
}
