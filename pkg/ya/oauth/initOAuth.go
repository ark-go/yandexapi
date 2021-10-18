/*
Импортируется структура
*/
package oauth

import (
	"fmt"
	"time"
)

type SYaToken struct {
	// получаем данные с yandex

	TokenType        string    `json:"token_type"`   // тип (не интересно)
	AccessToken      string    `json:"access_token"` // токен
	ExpiresIn        int       `json:"expires_in"`   // время в секундах жизни токена
	StartDate        time.Time // дата когда получили токен
	RefreshToken     string    `json:"refresh_token"`     // токен восстановления
	Scope            string    `json:"scope"`             // просто текст
	ErrorDescription string    `json:"error_description"` // развернутая ошибка
	Error            string    `json:"error"`             // ошибки запросов
	// требуемые данные для получения токена

	UserName   string       // можно задать пользователя Яндекс, чтоб ему свое имя не набирать
	DeviceId   string       // я выдаю случайный ID,
	OAuthCode  string       // код полученный пользователем, одноразовый
	AppId      string       // *ID приложения, получается на сайте Яндекс при регистрации приложения
	AppPass    string       // *пароль приложения, получается на сайте Яндекс при регистрации приложения
	SaveConfig func() error // *функция для записи полученного токена, т.е. всей этой структуры
}

var YaToken *SYaToken

func init() {
	YaToken = &SYaToken{}
}

// Запрос токена
//	false - (без параметра) будет проверятся время истечения токена и его обновление за 30 дней
//	true - безусловный запрос токена.
func InitOAuth(force ...bool) (string, error) {
	if YaToken.AppId == "" || YaToken.AppPass == "" {
		log.Fatalln("Не указаны Application ID или пароль")
	}
	if len(force) > 0 && force[0] {
		fmt.Println("Требуется подтверждение на сайте:")
		getCode()
		var kb string
		fmt.Print("Введите полученный код:")
		fmt.Scanf("%s\n", &kb)
		YaToken.OAuthCode = kb
		token, err := getToken()
		if err != nil {
			return "", err
		}
		return token, nil
	} else {
		token, err := updateToken()
		if err != nil {
			return InitOAuth(true)
		}
		return token, nil
	}

}
