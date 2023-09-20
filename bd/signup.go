package bd

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/yonatan/gambituser/models"
	"github.com/yonatan/gambituser/tools"
)

func SignUp(sig models.SignUp) error {
	fmt.Println("Comienza Registro")

	err := DbConnect()

	if err != nil {
		return err
	}

	defer Db.Close()

	sentecia := "INSERT INTO users (User_Email, User_UUID, User_DateAdd) VALUES('" + sig.UserEmail + "', '" + sig.UserUUID + "','" + tools.FechaaMySQL() + "')"

	fmt.Println(sentecia)

	_, err = Db.Exec(sentecia)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("SignUp > Ejecición Exitosa")

	return nil
}
