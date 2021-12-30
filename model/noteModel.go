package model

import (
	"go-reading/initDB"
	"log"
)

type NoteModel struct {
	Id         int64
	Uid        int64  `form:"uid" json:"uid" binding:"required"`
	Bookid     int64  `form:"bookid" json:"bookid" binding:"required"`
	Page       int64  `form:"page" json:"page" binding:"required"`
	Content    string `form:"content" json:"content"`
	Image      string `form:"image" json:"image"`
	Status     int8   `form:"status" json:"status"`
	CreateTime int64  `form:"create_time" json:"create_time"`
	UpdateTime int64  `form:"update_time" json:"update_time"`
}

func (note *NoteModel) Insert() int64 {
	sql := "INSERT INTO `reading`.`gr_note` (`uid`, `bookid`, `page`, `content`, `image`, `status`, `create_time`, `update_time`) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	result, err := initDB.Db.Exec(sql, note.Uid, note.Bookid, note.Page, note.Content, note.Image, note.Status, note.CreateTime, note.UpdateTime)
	if err != nil {
		panic(err.Error())
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	return lastId
}
