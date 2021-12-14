package model

import (
	"fmt"
	"go-reading/initDB"
	"log"
)

type UserModel struct {
	Username      string `form:"username" binding:"required"`
	Email         string `form:"email"  binding:"required,email"`
	Password      string `form:"password" binding:"required"`
	PasswordAgain string `form:"password-again" binding:"eqfield=Password"`
}

func (user *UserModel) Save() int64 {
	sql := "INSERT INTO students(email, first_name, last_name) VALUES ('admin@gmail.com', 'admin','admin')"

	fmt.Print(initDB.Db)

	res, err := initDB.Db.Exec(sql)

	if err != nil {
		panic(err.Error())
	}

	lastId, err := res.LastInsertId()

	if err != nil {
		log.Fatal(err)
	}
	return lastId
}
