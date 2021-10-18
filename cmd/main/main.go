package main

import (
	"io/fs"
	"io/ioutil"
	"log"

	"github.com/ark-go/yandexapi/pkg/appconf"
	"github.com/ark-go/yandexapi/pkg/ya/disk"
)

func init() {
}

var versionProg string

//var config map[string]string

func main() {

	log.Println("Версия:", versionProg)
	appconf.Conf.InitConfig()
	//oauth.InitOAuth()

	if err := disk.InitAppDir(); err != nil {
		log.Println("ошибка:", err)
	}

	if appconf.Conf.Disk.IsAppDirYandex() { // не обязательная проверка, тест
		g := []byte("Привет мир!")
		disk.UploadByte("test.txt", g)
		disk.UploadFile(`C:\Users\Zinaida\Pictures\2018-06-29 19.56.13.jpg`)
		disk.DownloadFile("2018-06-29 19.56.13.jpg", "c:\\temp2\\fileTesting1.jpg")
		var fileByte []byte
		if err := disk.DownloadByte("2018-06-29 19.56.13.jpg", &fileByte); err != nil {
			log.Println("error:", err.Error())
		} else {
			err := ioutil.WriteFile("c:\\temp2\\fileTesting2.jpg", fileByte, fs.FileMode(0640))
			if err != nil {
				log.Println("err:", err)
			}
		}
	}
}
