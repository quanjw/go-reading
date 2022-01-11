package note

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-reading/conf"
	"go-reading/model"
	"go-reading/utils"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"time"
)

func UploadNote(c *gin.Context) {

	file, e := c.FormFile("upload-file")
	fmt.Print("root path:", utils.RootPath())
	if e != nil {
		c.String(http.StatusBadRequest, "输入的数据不合法")
		log.Panicln("文件上传错误", e.Error())
	}

	t := time.Now()
	formatted := fmt.Sprintf("%d/%02d/%02d/",
		t.Year(), t.Month(), t.Day())
	uploadPath := filepath.Join(utils.RootPath(), "upload/", formatted)
	e = os.MkdirAll(uploadPath, os.ModePerm)
	if e != nil {
		log.Panicln("创建文件夹失败", e.Error())
	}

	extString := path.Ext(file.Filename)
	fileName := strconv.FormatInt(time.Now().Unix(), 10) + extString

	dst := filepath.Join(uploadPath, fileName)
	e = c.SaveUploadedFile(file, dst)
	if e != nil {
		log.Panicln("无法保存文件", e.Error())
	}

	uri := conf.Baseurl + "/upload/" + formatted + fileName
	data := gin.H{
		"success": "true",
		"message": "上传成功！",
		"uri":     uri,
	}
	c.JSON(http.StatusOK, data)
}

func Insert(c *gin.Context) {
	var note model.NoteModel
	if err := c.ShouldBind(&note); err != nil {
		log.Println("err ->", err.Error())
		c.String(http.StatusBadRequest, "输入的数据不合法")
	} else {
		note.Status = 1
		id := note.Insert()
		data := gin.H{
			"message": "笔记插入成功" + strconv.FormatInt(id, 10),
			"success": "true",
		}
		c.JSON(200, data)
	}
}

func Delete(c *gin.Context) {
	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		log.Panicln("id 不是 int 类型, id 转换失败", err.Error())
	}
	note := model.NoteModel{Id: int64(intId)}
	note.DeleteOne()
	data := gin.H{
		"message": "删除成功",
		"success": "true",
	}
	c.JSON(200, data)

}

func Get(c *gin.Context) {
	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		log.Panicln("id 不是 int 类型, id 转换失败", err.Error())
	}
	note := model.NoteModel{}
	note.FindById(int64(intId))
	data := gin.H{
		"message": "查询成功",
		"success": "true",
		"note":    note,
	}
	if note == (model.NoteModel{}){
		data = gin.H{
			"message": "数据为空",
			"success": "true",
		}
	}
	c.JSON(200, data)

}

func GetAll(c *gin.Context) {
	note := model.NoteModel{}
	notes := note.GetAll()
	data := gin.H{
		"message": "查询成功",
		"success": "true",
		"notes":   notes,
	}
	c.JSON(200, data)

}
