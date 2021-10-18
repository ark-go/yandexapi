package appconf

import (
	"flag"
	"os"
	"path/filepath"

	"github.com/ark-go/yandexapi/pkg/ya/oauth"
	uuid "github.com/nu7hatch/gouuid"
)

type conf struct {
	AppName        string
	AppDirYandex   string // каталог на яндекс для нашего приложения
	RootDir        string
	DirOnlyApp     bool // ограничено только каталогом приложения
	fileConfigPath string
	YaToken        *oauth.SYaToken // это единственная связь с модулем oauth
}

var Conf conf

func init() {
	rootDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatalln("не определить рабочий каталог")

	}
	Conf = conf{
		AppName:        "ANI Yandex Config",
		fileConfigPath: filepath.Join(rootDir, "appconfig.cfg"),
		YaToken:        oauth.YaToken,
		DirOnlyApp:     true, // true - ограничено только каталогом приложения
	}
	Conf.YaToken.SaveConfig = Conf.SaveConfig
	Conf.YaToken.Scope = "Раз Два Три"
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
