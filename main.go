package main

import (
	"context"

	event "github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(RunLambdaUser)
}

func RunLambdaUser(ctx context.Context, event event.CognitoEventUserPoolsPostConfirmation) (event.CognitoEventUserPoolsPostConfirmation, error) {
	return event, nil
}
