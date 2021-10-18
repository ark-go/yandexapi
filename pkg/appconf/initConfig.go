package appconf

import (
	"flag"
	"os"
	"path/filepath"

	"github.com/ark-go/yandexapi/pkg/ya/disk"
	"github.com/ark-go/yandexapi/pkg/ya/oauth"
	uuid "github.com/nu7hatch/gouuid"
)

type conf struct {
	//	AppName        string
	//	AppDirYandex   string // каталог на яндекс для нашего приложения
	RootDir        string
	fileConfigPath string
	YaToken        *oauth.SYaToken // это единственная связь с модулем oauth
	Disk           *disk.SDiskConf
}

var Conf conf

func init() {
	rootDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatalln("не определить рабочий каталог")

	}
	Conf = conf{
		fileConfigPath: filepath.Join(rootDir, "appconfig.cfg"),
		YaToken:        oauth.YaToken,
		Disk:           disk.DiskConf,
	}
	Conf.YaToken.SaveConfig = Conf.SaveConfig // мы хотим сами (в oauth) сохранять конфиг при получении нового токена
	Conf.YaToken.Scope = "Раз Два Три"
	Conf.Disk.AppName = "ANI Yandex Config" // каталог будет создан в каталоге приложения (хз зачем) т.е. Disk:/Приложения/Имя приложения/этот каталог
	Conf.Disk.DirOnlyApp = true             // true - ограничено только каталогом приложения, если при выдаче прав мы не разрешаем весь диск
	// будет создан каталог по названию приложения которое указываем при получении API пароля на yandex.ru

}

func (c *conf) InitConfig() {
	//newConf := &conf{}
	Conf.LoadConfig()
	flag.StringVar(&Conf.YaToken.AppId, "AppId", Conf.YaToken.AppId, "Application ID")
	flag.StringVar(&Conf.YaToken.AppPass, "AppPass", Conf.YaToken.AppPass, "Application Password")
	flag.StringVar(&Conf.YaToken.UserName, "UserName", Conf.YaToken.UserName, "Имя или почта пользователя yandex.ru")
	flag.StringVar(&Conf.YaToken.OAuthCode, "OAutCode", Conf.YaToken.OAuthCode, "Запрашивается приложением")
	//	log.Println(">>>", Conf)
	if Conf.YaToken.DeviceId == "" {
		if ui, err := uuid.NewV4(); err != nil {
			log.Fatalln("Не создать Device ID ")
		} else {
			Conf.YaToken.DeviceId = ui.String()
		}
	}
	flag.Parse()
	Conf.SaveConfig()

}
