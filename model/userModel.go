package model

import (
	"database/sql"
	"go-reading/initDB"
	"log"
)

type UserModel struct {
	Id            int64  `form:"id"`
	Username      string `form:"username"  binding:"required"`
	Email         string `form:"email"  binding:"email,required"`
	Password      string `form:"password" binding:"required"`
	PasswordAgain string `form:"password-again" binding:"eqfield=Password"`
}

func (user *UserModel) Save() int64 {
	result := initDB.Db.Create(&user)
	//result, err := initDB.Db.Exec("insert into gr_user (username,email, password) values (?,?,?);", user.Username, user.Email, user.Password)
	if result.Error != nil {
		panic("用户胡插入失败")
	}
	return user.Id
}

func (user *UserModel) ExistUser() bool {

	sqlStatement := `SELECT id FROM gr_user WHERE username=? OR email =? ;`
	var id int64
	row := initDB.Db.QueryRow(sqlStatement, user.Username, user.Email)
	err := row.Scan(&id)

	flag := false
	switch err {
	case sql.ErrNoRows:
		flag = false
	case nil:
		flag = true
	default:
		log.Fatal(err)
	}
	return flag
}

type LoginUser struct {
	Email    string `form:"email"  binding:"email,required"`
	Password string `form:"password" binding:"required"`
}

func (loginUser *LoginUser) QueryByEmail() UserModel {
	user := UserModel{}
	row := initDB.Db.QueryRow("select id,username,email,password from gr_user where email = ?;", loginUser.Email)
	e := row.Scan(&user.Id, &user.Username, &user.Email, &user.Password)
	if e != nil {
		return user
	}
	return user
}
