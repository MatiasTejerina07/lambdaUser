package db

import (
	"fmt"
	"lambdaUser/models"
	"lambdaUser/tools"

	_ "github.com/go-sql-driver/mysql"
)

func SignUp(sig models.SignUp) error {
	fmt.Println("Comienza registro")

	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	sentencia := "INSERT INTO users (User_Email, User_UUID, User_DateAdd) VALUES ('" + sig.UserEmail + "','" + sig.UserUUID + "','" + tools.DateMySQL() + "')"

	fmt.Println(sentencia)

	_, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println("Hay error en la sentencia:" + err.Error())
		return err
	}
	fmt.Println(">Registro exitoso")
	return nil
}
