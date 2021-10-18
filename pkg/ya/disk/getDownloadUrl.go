package disk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func getDownloadUrl(filename string) (urldown string, err error) {
	if DiskConf.appDirYandex == "" {
		return "", fmt.Errorf("%s", "Не определен каталог приложения")
	}
	path := DiskConf.appDirYandex + "/" + filename
	baseUrl := "https://cloud-api.yandex.net/v1/disk/"
	relativeUrl := "resources/download"
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
	req.Header.Add("Authorization", *DiskConf.yaAccessToken)
	resp, reserr := http.DefaultClient.Do(req)
	if reserr != nil {
		log.Fatal(reserr)
	}
	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)
	resp.Body.Close()
	//log.Printf("Диск: %+v", res)
	//log.Println("Получено Status:", resp.StatusCode) //200
	if v, ok := res["error"]; ok {
		switch v {
		case "DiskPathDoesntExistsError":
			log.Println("Указанного пути не существует")
		case "DiskNotFoundError":
			log.Println("Не удалось найти ресурс:", path)
		default:
			log.Println("Ошибка: ", v)
			if e, ok := res["message"].(string); ok {
				log.Println("Не обработано, ошибка: ", e)
			}
		}
		return "", fmt.Errorf("%s", v)
	}
	// if v, ok := res["method"].(string); ok { //  GET ? templated:false
	// 	log.Println("DownloadFile method: ", v)
	// }
	if v, ok := res["href"].(string); ok {
		//DownloadFile(v)
		return v, nil
	} else {
		return "", fmt.Errorf("%s", "не получили ссылку на скачивание:")
	}

	//	return "", nil
}
