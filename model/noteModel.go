package model

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
