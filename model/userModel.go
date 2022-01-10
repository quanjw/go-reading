package model

import (
	"errors"
	"fmt"
	"go-reading/initDB"
	"gorm.io/gorm"
	"log"
)

type UserModel struct {
	Id            int64  `form:"id" gorm:"primary_key"`
	Username      string `form:"username"  binding:"required"`
	Email         string `form:"email"  binding:"email,required"`
	Password      string `form:"password" binding:"required"`
	PasswordAgain string `form:"password-again" binding:"eqfield=Password" gorm:"-"`
}

func (user *UserModel) Save() int64 {
	result := initDB.Db.Table("gr_user").Create(&user)
	//result, err := initDB.Db.Exec("insert into gr_user (username,email, password) values (?,?,?);", user.Username, user.Email, user.Password)
	if result.Error != nil {
		fmt.Print(result)
		log.Panicln("用户胡插入失败", result.Error)
	}
	return user.Id
}

func (user *UserModel) ExistUser() bool {

	//sqlStatement := `SELECT id FROM gr_user WHERE username=? OR email =? ;`
	//var id int64
	//row := initDB.Db.QueryRow(sqlStatement, user.Username, user.Email)
	//err := row.Scan(&id)
	err := initDB.Db.Table("gr_user").Where("username = ? OR email = ?", user.Username, user.Email).First(&user).Error
	isError := errors.Is(err, gorm.ErrRecordNotFound)
	flag := true
	if isError {
		flag = false
	}

	return flag
}

type LoginUser struct {
	Email    string `form:"email"  binding:"email,required"`
	Password string `form:"password" binding:"required"`
}

func (loginUser *LoginUser) QueryByEmail() UserModel {
	user := UserModel{}
	//row := initDB.Db.QueryRow("select id,username,email,password from gr_user where email = ?;", loginUser.Email)
	//e := row.Scan(&user.Id, &user.Username, &user.Email, &user.Password)
	//if e != nil {
	//	return user
	//}
	initDB.Db.Table("gr_user").Where("email = ?", loginUser.Email).First(&user)
	return user
}
