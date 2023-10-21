package db

import (
	"database/sql"
	"fmt"
	"lambdaUser/models"
	"lambdaUser/secret"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var SecretModel models.SecretRDSJson
var err error
var Db *sql.DB

func ReadSecret() error {
	SecretModel, err = secret.GetSecret(os.Getenv("SecretName"))
	return err
}

func DbConnect() error {
	Db, err = sql.Open("mysql", ConnStr(SecretModel))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	err = Db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("Connected to database")
	return nil
}

func ConnStr(claves models.SecretRDSJson) string {
	var dbUser, authToken, dbEnpoints, dbName string
	dbUser = claves.Username
	authToken = claves.Password
	dbEnpoints = claves.Host
	dbName = "gambit"
	dsn := fmt.Sprintf("%s:%s:@tcp(%s)/%s?allowCleartextPasswords=true", dbUser, authToken, dbEnpoints, dbName)
	fmt.Println(dsn)
	return dsn
}
