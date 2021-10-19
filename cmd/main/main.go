package main

import (
	"io/fs"
	"io/ioutil"
	"log"
	"path/filepath"

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
	//appconf.CriptoTest()
	//oauth.InitOAuth()

	if err := disk.InitAppDir(); err != nil {
		log.Println("ошибка:", err)
	}

	if appconf.Conf.Disk.IsAppDirYandex() { // не обязательная проверка, тест
		g := []byte("Привет мир!")
		disk.UploadByte("test.txt", &g)
		disk.UploadFile(filepath.Join(appconf.Conf.RootDir, "dddd.jpg"))
		disk.DownloadFile("dddd.jpg", filepath.Join(appconf.Conf.RootDir, "downl_1.jpg"))
		var fileByte []byte
		if err := disk.DownloadByte("dddd.jpg", &fileByte); err != nil {
			log.Println("error:", err.Error())
		} else {
			err := ioutil.WriteFile(filepath.Join(appconf.Conf.RootDir, "downl_2.jpg"), fileByte, fs.FileMode(0640))
			if err != nil {
				log.Println("err:", err)
			}
		}
	}
}
