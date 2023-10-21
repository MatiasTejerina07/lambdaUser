package main

import (
	"context"
	"errors"
	"fmt"
	"lambdaUser/aws"
	"lambdaUser/db"
	"lambdaUser/models"
	"os"

	event "github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(RunLambdaUser)
}

func RunLambdaUser(ctx context.Context, event event.CognitoEventUserPoolsPostConfirmation) (event.CognitoEventUserPoolsPostConfirmation, error) {
	aws.InitialAws()

	if !ValidateParameters() {
		fmt.Println("Error en los parámetros. debe enviar 'SecretName'")
		err := errors.New("Error en los parámetros. debe enviar 'SecretName'")
		return event, err
	}
	var datos models.SignUp

	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			datos.UserEmail = att
			fmt.Print("Email ==" + datos.UserEmail)
		case "sub":
			datos.UserUUID = att
			fmt.Print("Email ==" + datos.UserUUID)
		}
	}

	err := db.ReadSecret()

	if err != nil {
		fmt.Println("Error al leer el secret" + err.Error())
		return event, err
	}

}

func ValidateParameters() bool {
	var Parameters bool
	_, Parameters = os.LookupEnv("SecretName")
	return Parameters
}
