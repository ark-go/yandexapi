package oauth

import (
	"net/url"
)

// запрос страницы для пользователя, чтоб он получил код на сайте
func getCode() {
	baseUrl := "https://oauth.yandex.com"
	relativeUrl := "authorize"
	u, err := url.Parse(relativeUrl)
	if err != nil {
		log.Fatal(err)
	}
	queryString := u.Query()
	queryString.Set("response_type", "code")          //code
	queryString.Set("client_id", YaToken.AppId)       //идентификатор приложения
	queryString.Set("device_id", YaToken.DeviceId)    //идентификатор устройства
	queryString.Set("device_name", "moi test device") //имя устройства
	//queryString.Set("redirect_uri", "1")              //адрес перенаправления
	queryString.Set("login_hint", YaToken.UserName) //имя пользователя или электронный адрес
	//queryString.Set("scope", "1")                     //запрашиваемые необходимые права
	//queryString.Set("optional_scope", "1")  //запрашиваемые опциональные права
	queryString.Set("force_confirm", "yes") //yes
	queryString.Set("state", "test")        //произвольная строка
	//queryString.Set("display", "popup") // хз не нужно

	u.RawQuery = queryString.Encode()
	base, err := url.Parse(baseUrl)
	if err != nil {
		log.Fatal(err)
	}
	path := base.ResolveReference(u).String()
	//log.Println(path)
	openBrowser(path)
}
