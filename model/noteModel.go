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
		log.Panicln("数据发生错误，无法插入", err.Error())
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	return lastId
}

func (note *NoteModel) DeleteOne() {
	sql := "DELETE FROM `reading`.`gr_note` WHERE `id` = ?"
	_, err := initDB.Db.Exec(sql, note.Id)
	if err != nil {
		log.Panicln("数据发生错误，无法删除", err.Error())
	}
}

func (note *NoteModel) FindById() *NoteModel {
	sql := "SELECT * FROM `reading`.`gr_note` WHERE `id` = ? AND status = 1"
	row := initDB.Db.QueryRow(sql, note.Id)
	log.Print(note)
	if err := row.Scan(&note.Id, &note.Uid, &note.Bookid, &note.Page, &note.Content, &note.Image, &note.Status, &note.CreateTime, &note.UpdateTime); err != nil {
		log.Panicln("绑定发生错误", err.Error())
	}
	return note
}

func (note *NoteModel) GetAll() []NoteModel {
	sql := "SELECT * FROM `reading`.`gr_note` WHERE status = 1"
	rows, err := initDB.Db.Query(sql)
	if err != nil {
		log.Panicln("查询错误", err.Error())
	}
	var notes []NoteModel
	for rows.Next() {
		var n NoteModel
		if err := rows.Scan(&n.Id, &n.Uid, &n.Bookid, &n.Page, &n.Content, &n.Image, &n.Status, &n.CreateTime, &n.UpdateTime); err == nil {
			notes = append(notes, n)
		}
	}
	return notes
}
