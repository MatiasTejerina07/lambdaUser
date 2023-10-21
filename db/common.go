package db

import (
	"lambdaUser/models"
	"lambdaUser/secret"
	"os"
)

var SecretModel models.SecretRDSJson
var err error

func ReadSecret() error {
	SecretModel, err = secret.GetSecret(os.Getenv("SecretName"))
	return err
}
