package model

import (
	"go-reading/initDB"
	"log"
)

type NoteModel struct {
	Id         int64  `form:"id"`
	Uid        string `form:"uid"  binding:"required"`
	Bookid     string `form:"bookid"  binding:"email,required"`
	Page       string `form:"page" binding:"required"`
	Content    string `form:"content"`
	Image      string `form:"image"`
	Status     string `form:"status"`
	CreateTime string `form:"create_time"`
	UpdateTime string `form:"update_time"`
}

func (user *UserModel) UploadImage() int64 {
	result, err := initDB.Db.Exec("insert into gr_user (username,email, password) values (?,?,?);", user.Username, user.Email, user.Password)
	if err != nil {
		panic(err.Error())
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	return lastId
}
