package oauth

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

// получение токена по коду, который получит пользователь
func getToken() {
	YaToken.Error = ""
	YaToken.ErrorDescription = ""
	data := url.Values{
		"grant_type":    {"authorization_code"}, //authorization_code
		"code":          {YaToken.OAuthCode},    //код подтверждения
		"client_id":     {YaToken.AppId},        //идентификатор приложения
		"client_secret": {YaToken.AppPass},      //пароль приложения
		"device_id":     {YaToken.DeviceId},     //идентификатор устройства, Если идентификатор был указан в device_id параметра при запросе коды подтверждения, то device_idи device_nameпараметры игнорируются
		//"device_name":        {"gardener"},           //имя устройства

	}
	resp, err := http.PostForm("https://oauth.yandex.com/token", data)

	if err != nil {
		log.Fatal(err)
	}
	//	body, err := ioutil.ReadAll(r.Body) // []byte
	jsonErr := json.NewDecoder(resp.Body).Decode(&YaToken)
	resp.Body.Close()
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	// bad_verification_code  - при вводе не правильного кода
	if YaToken.Error != "" {
		log.Println("token error:", YaToken.Error)
	}
	YaToken.StartDate = time.Now().Add(time.Duration(YaToken.ExpiresIn) * time.Second)
	YaToken.SaveConfig()

}

/*
Коды ошибок:

authorization_pending: Пользователь еще не ввел код подтверждения.
bad_verification_code: codeПереданное значение параметра не является семизначным числом.
invalid_client: Приложение с указанным идентификатором ( client_idпараметром) не найдено или заблокировано. Этот код также возвращается, если client_secretпараметр передал недопустимый пароль приложения.
invalid_grant: Недействительный или просроченный код подтверждения.
invalid_request: Неверный формат запроса (один из параметров не указан, указан дважды или не передан в теле запроса).
invalid_scope: Права приложения изменились после того, как был сгенерирован код подтверждения.
unauthorized_client: Приложение было отклонено во время модерации или ожидает модерации.
unsupported_grant_type: Недопустимое grant_typeзначение параметра.
Basic auth required: Тип авторизации, указанный в Authorizationзаголовке, не указан Basic.
Malformed Authorization header: AuthorizationЗаголовок не в <client_id>:<client_secret>формате, или эта строка не закодирована в Base64.
*/
