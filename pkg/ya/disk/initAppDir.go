package disk

import (
	"fmt"

	"github.com/ark-go/yandexapi/pkg/appconf"
	"github.com/ark-go/yandexapi/pkg/ya/oauth"
)

// Создает каталог для приложения на Yandex Disk
//	проверяется доступ к диску, при необходимости запрашивается код для получения токена
//	зациклено до входа, TODO придумать выход по ошибке N-раз
//	сейчас выход только по Ctrl+C
func InitAppDir() error {
	GetDiskInfo()
	if DiskInfo.AppDirPath == "" {
		if err := oauth.InitOAuth(true); err != nil {
			return err
		}
		return InitAppDir()
		//return fmt.Errorf("%s", "не получили инфо по диску, авторизация ?")
	}
	appconf.Conf.AppDirYandex = ""
	appDir := DiskInfo.AppDirPath + "/" + appconf.Conf.AppName
	if appDir != "" {
		if err := CreateDirectory(DiskInfo.AppDirPath); err == nil { // на системный каталог не говорит что уже существует??
			if err := CreateDirectory(appDir); err != nil { // пропускаем ошибку существования
				log.Println("Каталог приложения не создан, или другая ошибка", appDir) //! Выход ?
				return err
			}
			appconf.Conf.AppDirYandex = appDir
			return nil
		}
	}
	return fmt.Errorf("%s", "не известная ошибка при создании каталога приложения ?")
}
