package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-reading/utils"
	"log"
	"net/http"
	"os"
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
	path := filepath.Join(utils.RootPath(), "upload/", formatted)
	e = os.MkdirAll(path, os.ModePerm)
	if e != nil {
		log.Panicln("创建文件夹失败", e.Error())
	}
	fileName := strconv.FormatInt(time.Now().Unix(), 10) + file.Filename

	dst := filepath.Join(path, file.Filename)
	e = c.SaveUploadedFile(file, dst)
	if e != nil {
		log.Panicln("无法保存文件", e.Error())
	}
	log.Println(path + fileName)

}
