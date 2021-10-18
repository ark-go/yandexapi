package disk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/ark-go/yandexapi/pkg/appconf"
)

func CreateDirectory(path string) error {
	baseUrl := "https://cloud-api.yandex.net/v1/disk/"
	relativeUrl := "resources"
	u, err := url.Parse(relativeUrl)
	if err != nil {
		return err
	}
	queryString := u.Query()
	queryString.Set("path", path) //code

	u.RawQuery = queryString.Encode()
	base, err := url.Parse(baseUrl)
	if err != nil {
		return err
	}
	url := base.ResolveReference(u).String()
	req, reqerr := http.NewRequest("PUT", url, nil) //strings.NewReader(form.Encode())
	if reqerr != nil {
		return reqerr
	}
	req.Header.Set("Content-Type", "application/json") // "application/json; charset=utf-8"
	req.Header.Add("Authorization", appconf.Conf.YaToken.AccessToken)
	resp, reserr := http.DefaultClient.Do(req)
	if reserr != nil {
		return reserr
	}
	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)
	resp.Body.Close()
	//log.Printf("Диск: %+v", res)

	if v, ok := res["error"]; ok {
		switch v {
		case "DiskPathPointsToExistentDirectoryError":
			//log.Println("Каталог существует")
			return nil
		case "DiskPathDoesntExistsError":
			log.Println("Указанного пути не существует")
		case "UnauthorizedError":
			//	log.Println("2)Нет авторизации, токен сброшен получите заново, перезагрузите программу")
			if err := InitAppDir(); err != nil {
				return err
			}
			return CreateDirectory(path)
		case "FieldValidationError":
			// не заполнено обязательное поле
		case "ForbiddenError":
			log.Println("Не достаточно прав на этот путь:", path)
		default:
			log.Println("Ошибка: ", v)
			if e, ok := res["message"].(string); ok {
				log.Println("Не обработано, ошибка: ", e)
			}
		}
		//log.Println("testww", fmt.Errorf("%s", v))
		return fmt.Errorf("%s", v)
	}
	// if v, ok := res["href"].(string); ok {
	// 	return nil
	// }
	return nil
}
