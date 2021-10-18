package oauth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// обновление токена по refresh токену
func updateToken() {

	if YaToken.RefreshToken == "" {
		log.Println("Нет refresh-токена")
		return
	}
	if YaToken.ExpiresIn != 0 {
		d := time.Until(YaToken.StartDate) / time.Hour / 24
		//		log.Println("Осталось дней:", int(d))
		if int(d) > 35 {
			log.Printf("OAuth закончится через: %d дней, обновление не требуется", int(d))
			return
		}
	}

	log.Println("refresh -", YaToken.RefreshToken)
	data := url.Values{
		"grant_type":    {"refresh_token"},      //refresh_token
		"refresh_token": {YaToken.RefreshToken}, //код подтверждения
		"client_id":     {YaToken.AppId},        //идентификатор приложения
		"client_secret": {YaToken.AppPass},      //пароль приложения
		"device_id":     {YaToken.DeviceId},     //идентификатор устройства, Если идентификатор был указан в device_id параметра при запросе коды подтверждения, то device_idи device_nameпараметры игнорируются
		//"device_name":        {"gardener"},           //имя устройства

	}

	resp, err := http.PostForm("https://oauth.yandex.com/token", data)

	if err != nil {
		log.Fatal(err)
	}
	jsonErr := json.NewDecoder(resp.Body).Decode(&YaToken)
	resp.Body.Close()
	//jsonErr := json.Unmarshal(resp.Body, &resToken)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	YaToken.StartDate = time.Now().Add(time.Duration(YaToken.ExpiresIn) * time.Second)
	YaToken.SaveConfig()
	fmt.Printf("Обновление токена: %+v\n", YaToken)
}

/* update error
Коды ошибок:

invalid_client: Приложение с указанным идентификатором ( client_idпараметром) не найдено или заблокировано. Этот код также возвращается, если client_secretпараметр передал недопустимый пароль приложения.
invalid_grant: Недействительный или просроченный токен обновления. Этот код также возвращается, если токен обновления принадлежит другому приложению (не соответствует переданному client_id).
invalid_request: Неверный формат запроса (один из параметров не указан, указан дважды или не передан в теле запроса).
unauthorized_client: Приложение было отклонено во время модерации или ожидает модерации.
unsupported_grant_type: Недопустимое grant_typeзначение параметра.
Basic auth required: Тип авторизации, указанный в Authorizationзаголовке, не указан Basic.
Malformed Authorization header: AuthorizationЗаголовок не в <client_id>:<client_secret>формате, или эта строка не закодирована в Base64.
*/
