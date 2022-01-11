package model

import (
	"fmt"
	"go-reading/initDB"
	"log"
)

type NoteModel struct {
	Id         int64  `gorm:"primary_key"`
	Uid        int64  `form:"uid" json:"uid" binding:"required"`
	Bookid     int64  `form:"bookid" json:"bookid" binding:"required"`
	Page       int64  `form:"page" json:"page" binding:"required"`
	Content    string `form:"content" json:"content"`
	Image      string `form:"image" json:"image"`
	Status     int8   `form:"status" json:"status"`
	CreateTime int64  `form:"create_time" json:"create_time" gorm:"autoCreateTime"`
	UpdateTime int64  `form:"update_time" json:"update_time"  gorm:"autoUpdateTime"`
}

func (note *NoteModel) Insert() int64 {
	result := initDB.Db.Table("gr_note").Create(&note)
	if result.Error != nil {
		fmt.Print(result)
		log.Panicln("用户胡插入失败", result.Error)
	}
	return note.Id
}

func (note *NoteModel) DeleteOne() {
	initDB.Db.Table("gr_note").Delete(&NoteModel{},note.Id)
}

func (note *NoteModel) FindById( id int64) {
	initDB.Db.Table("gr_note").Where("status = 1").First(&note, id)
}

func (note *NoteModel) GetAll() []NoteModel {

	var notes []NoteModel
	initDB.Db.Table("gr_note").Where("status = 1").Find(&notes)
	return notes
}
