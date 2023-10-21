package secret

import (
	"fmt"
	user "lambdaUser/aws"
	"lambdaUser/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func GetSecret(secretName string) (models.SecretRDSJson, error) {
	var dataSecret models.SecretRDSJson
	fmt.Println(">Pido secret : " + secretName)
	svc := secretsmanager.NewFromConfig(user.Cfg)
	clave, err := svc.GetSecretValue(user.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	return dataSecret, err
}
