package disk

import (
	"fmt"

	"github.com/ark-go/yandexapi/pkg/ya/oauth"
)

type SDiskConf struct {
	// сам токен доступа
	yaAccessToken *string
	// созданный дополнительный каталог к выданному (имя каталога в AppName)
	appDirYandex string
	// это будет названием дополнительного каталога (путь будет в appDirYandex)
	AppName string
	// true - ограничено только каталогом приложения, если при выдаче прав мы не разрешаем весь диск,
	// будет создан каталог по названию приложения которое указываем при получении API пароля на yandex.ru
	DirOnlyApp bool
}

// создан ли каталог для приложения
func (sd *SDiskConf) IsAppDirYandex() bool {
	return sd.appDirYandex != ""
}

var DiskConf *SDiskConf

func init() {
	DiskConf = &SDiskConf{
		yaAccessToken: &oauth.YaToken.AccessToken,
	}
}

// Создает каталог для приложения на Yandex Disk
//	проверяется доступ к диску, при необходимости запрашивается код для получения токена
//	зациклено до входа, TODO придумать выход по ошибке N-раз
//	сейчас выход только по Ctrl+C
func InitAppDir() error {
	if err := GetDiskInfo(); err != nil {
		// нет данных, значит нет авторизации, запросим токен
		_, err := oauth.InitOAuth(true) // true - плевать на обновления токена, запрашивать
		if err != nil {
			return err
		}
		// Повторим после получения токена, если так и не получилось выходим
		if err := GetDiskInfo(); err != nil {
			return err
		}
	}
	_, err := oauth.InitOAuth() //  false - будет проверять время токена
	if err != nil {
		return err
	}
	DiskConf.appDirYandex = ""
	appDir := DiskInfo.appRoot + "/" + DiskConf.AppName
	if appDir != "" {
		if err := CreateDirectory(DiskInfo.appRoot); err == nil { // на системный каталог не говорит что уже существует??
			if err := CreateDirectory(appDir); err != nil { // пропускаем ошибку существования
				log.Println("Каталог приложения не создан, или другая ошибка", appDir) //! Выход ?
				return err
			}
			DiskConf.appDirYandex = appDir
			return nil
		} else {
			return err // FieldValidationError
		}
	}
	return fmt.Errorf("%s", "не известная ошибка при создании каталога приложения ?")
}
