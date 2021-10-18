package disk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/ark-go/yandexapi/pkg/appconf"
)

func GetUploadUrl(path string) (href string, err error) {
	CreateBakFile(path)
	baseUrl := "https://cloud-api.yandex.net/v1/disk/"
	relativeUrl := "resources/upload"
	u, err := url.Parse(relativeUrl)
	if err != nil {
		log.Fatal(err)
	}
	queryString := u.Query()
	queryString.Set("path", path) //code

	u.RawQuery = queryString.Encode()
	base, err := url.Parse(baseUrl)
	if err != nil {
		log.Fatal(err)
	}
	url := base.ResolveReference(u).String()
	//log.Println("path: ", url)
	req, reqerr := http.NewRequest("GET", url, nil) //strings.NewReader(form.Encode())
	if reqerr != nil {
		log.Fatal(reqerr)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", appconf.Conf.YaToken.AccessToken)
	resp, reserr := http.DefaultClient.Do(req)
	if reserr != nil {
		log.Fatal(reserr)
	}
	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)
	resp.Body.Close()
	//log.Printf("Диск: %+v", res)

	if v, ok := res["error"]; ok {
		switch v {
		case "DiskResourceAlreadyExistsError":
			log.Println("Ресурс уже существует")
		case "DiskPathDoesntExistsError":
			log.Println("Указанного пути не существует")
		default:
			log.Println("Ошибка: ", v)
			if e, ok := res["message"].(string); ok {
				log.Println("Не обработано, ошибка: ", e)
			}
		}
		return "", fmt.Errorf("%s", v)
	}
	if v, ok := res["href"].(string); ok {
		//	log.Println("UploadFile:", v)
		return v, nil
	}
	return "", nil
}

/*
Коды ответов при загрузке файла
API отвечает 201 Createdкодом, если файл был загружен без ошибок.

Другие коды ответа HTTP:

202 Accepted - Файл был получен сервером, но еще не перенесен на Яндекс.Диск.

412 Precondition Failed- Content-RangeПри загрузке файла в заголовке был передан неверный диапазон .

413 Payload Too Large - Размер файла превышает 10 ГБ.

500 Internal Server Errorили 503 Service Unavailable- Ошибка сервера. Попробуйте повторить загрузку.

507 Insufficient Storage - На Диске пользователя недостаточно свободного места для загруженного файла.
*/
