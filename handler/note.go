package handler

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

func Insert(c *gin.Context) {

}

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

func NoteInsert(c *gin.Context) {
	var note model.NoteModel
	if err := c.ShouldBind(&note); err != nil {
		log.Println("err ->", err.Error())
		c.String(http.StatusBadRequest, "输入的数据不合法")
	} else {
		note.CreateTime = time.Now().Unix()
		note.Status = 1

		id := note.Insert()
		data := gin.H{
			"message": "笔记插入成功" + strconv.FormatInt(id, 10),
			"success": "true",
		}
		c.JSON(200, data)
	}
}
