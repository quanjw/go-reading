package utils

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func RootPath() string {
	path, err := exec.LookPath(os.Args[0])
	if err != nil {
		log.Panicln("发生错误", err.Error())
	}
	fmt.Print(path)
	i := strings.LastIndex(path, "\\")
	rootPath := path[0 : i+1]
	return rootPath
}
