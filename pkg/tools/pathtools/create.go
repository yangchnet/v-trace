package pathtools

import (
	"log"
	"os"
)

func MkdirIfNotExist(dirPath string) {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.MkdirAll(dirPath, 0o777)
		if err != nil {
			log.Printf("创建目录失败：%s", err.Error())
		}
	}
}
