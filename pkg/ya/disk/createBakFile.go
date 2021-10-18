package disk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func createBakFile(pathfile string) error {
	baseUrl := "https://cloud-api.yandex.net/v1/disk/"
	relativeUrl := "resources/move"
	u, err := url.Parse(relativeUrl)
	if err != nil {
		return err
	}
	queryString := u.Query()
	queryString.Set("from", pathfile) //code
	queryString.Set("path", pathfile+".bak")

	u.RawQuery = queryString.Encode()
	base, err := url.Parse(baseUrl)
	if err != nil {
		return err
	}
	url := base.ResolveReference(u).String()
	req, reqerr := http.NewRequest("POST", url, nil) //strings.NewReader(form.Encode())
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
			//log.Println("файла .bak не существует.")
			return nil
		case "DiskResourceAlreadyExistsError": // если bak уже есть - удаляем и пробуеи снова
			//log.Println("Удаляем bak.")
			if err := DeleteFile(pathfile + ".bak"); err != nil {
				return err
			}
			return createBakFile(pathfile)
		case "DiskPathDoesntExistsError":
			log.Println("Указанного пути не существует")
		case "UnauthorizedError":
			if err := InitAppDir(); err != nil {
				return err
			}
			return createBakFile(pathfile)
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
	//log.Println("move:", res) //:: 2021/10/17 17:53:38 move: map[href:https://cloud-api.yandex.net/v1/disk/resources?path=disk%3A%2F%D0%9F%D1%80%D0%B8%D0%BB%D0%BE%D0%B6%D0%B5%D0%BD%D0%B8%D1%8F%2FANI+YandexDisk%2Ftestik8.jpg.bak method:GET templated:false]
	return nil
}
