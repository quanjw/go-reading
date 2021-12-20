package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
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

	data := gin.H{
		"success":   "true",
		"message":   "上传成功！",
		"file_path": dst,
	}
	c.JSON(http.StatusOK, data)
}
