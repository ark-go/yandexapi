package disk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type diskInfo struct {
	Max_file_size                int                    `json:"max_file_size"`                // 54760833024,
	Unlimited_autoupload_enabled bool                   `json:"unlimited_autoupload_enabled"` // true,
	Total_space                  int                    `json:"total_space"`                  // 66571993088,
	Trash_size                   int                    `json:"trash_size"`                   // 16404899,
	Is_paid                      bool                   `json:"is_paid"`                      // false,
	Used_space                   int                    `json:"used_space"`                   // 11424322597,
	System_folders               map[string]interface{} `json:"system_folders"`
	// каталог Приложения, в зависимости от языка
	appRoot string
}

var DiskInfo *diskInfo

func GetDiskInfo() error {
	DiskInfo = &diskInfo{}
	baseUrl := "https://cloud-api.yandex.net/v1/disk/"
	url, err := url.Parse(baseUrl)
	if err != nil {
		return err
	}
	req, reqerr := http.NewRequest("GET", url.String(), nil) //strings.NewReader(form.Encode())
	if reqerr != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", *DiskConf.yaAccessToken)
	resp, reserr := http.DefaultClient.Do(req)
	if reserr != nil {
		return reserr
	}
	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&DiskInfo)
	resp.Body.Close()

	if v, ok := res["error"]; ok {
		switch v {
		// case "DiskPathPointsToExistentDirectoryError":
		// 	log.Println("Каталог существует")
		// case "DiskPathDoesntExistsError":
		// 	log.Println("Указанного пути не существует")
		// case "UnauthorizedError": // не происходит,
		// 	log.Println("1)Нет авторизации, токен сброшен получите заново, перезагрузите программу")
		// 	return fmt.Errorf("%s", v)
		default:
			log.Println("Ошибка: ", v)
			if e, ok := res["message"].(string); ok {
				log.Println("Не обработано, ошибка: ", e)
			}
		}
		return fmt.Errorf("%s", v)
	}
	//log.Printf("diskinfo %+v", DiskInfo)
	// без авторизации просто выдаются нули, ошибки нет
	if v, ok := DiskInfo.System_folders["applications"].(string); ok {
		if DiskConf.DirOnlyApp {
			// если права ограничены только каталогом этого приложения (разрешенным OAuthID), тут получим его
			if err := сreateAppRoot(); err != nil {
				return err
			}
		} else {
			// раз ограничений нет, получим каталог с названием "Приложения"
			DiskInfo.appRoot = v
		}
		//	log.Println("Путь к Приложениям:", v)
	} else {
		return fmt.Errorf("%s", "Не получить корневой диск приложения Yandex")
	}

	//	log.Printf("INFO2:%+v\n", DiskInfo)
	return nil
}
