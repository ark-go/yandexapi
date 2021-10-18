package disk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func DeleteFile(pathfile string) error {
	baseUrl := "https://cloud-api.yandex.net/v1/disk/"
	relativeUrl := "resources"
	u, err := url.Parse(relativeUrl)
	if err != nil {
		return err
	}
	queryString := u.Query()
	queryString.Set("permanently", "true") // Без корзины
	queryString.Set("path", pathfile)

	u.RawQuery = queryString.Encode()
	base, err := url.Parse(baseUrl)
	if err != nil {
		return err
	}
	url := base.ResolveReference(u).String()
	req, reqerr := http.NewRequest("DELETE", url, nil) //strings.NewReader(form.Encode())
	if reqerr != nil {
		return reqerr
	}
	req.Header.Set("Content-Type", "application/json") // "application/json; charset=utf-8"
	req.Header.Add("Authorization", *DiskConf.yaAccessToken)
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
		case "DiskNotFoundError":
			log.Println("Не удалось найти запрошенный ресурс.")
			return nil
		case "DiskResourceAlreadyExistsError":
			log.Println("ресурс уже существует.")
		case "DiskPathDoesntExistsError":
			log.Println("Указанного пути не существует")
		case "MethodNotAllowedError":
			log.Println("Метод не поддерживается")
		case "UnauthorizedError":
			if err := InitAppDir(); err != nil {
				return err
			}
			return DeleteFile(pathfile)
		case "FieldValidationError":
			// не заполнено обязательное поле
		default:
			log.Println("Ошибка: ", v)
			if e, ok := res["message"].(string); ok {
				log.Println("Не обработано, ошибка: ", e)
			}
		}
		log.Println("testww", fmt.Errorf("%s", v))
		return fmt.Errorf("%s", v)
	}
	// if v, ok := res["href"].(string); ok {
	// 	return nil
	// }
	return nil
}
