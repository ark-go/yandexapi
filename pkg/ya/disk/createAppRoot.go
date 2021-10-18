package disk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// вызывается из GetDiskInfo  и по другому не надо
//	вызов создаст и вернет каталог приложения на yandexDisk который прописан  на сайте при создании приложения ()
func сreateAppRoot() error {
	baseUrl := "https://cloud-api.yandex.net/v1/disk/resources?path=app:/"
	url, err := url.Parse(baseUrl)
	if err != nil {
		return err
	}
	//log.Println("path:", url.String())
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

	json.NewDecoder(resp.Body).Decode(&res)
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

	if v, ok := res["path"].(string); ok {
		if v != "" {
			DiskInfo.appRoot = v
		} else {
			return fmt.Errorf("не правильный каталог для приложения")
		}
	} else {
		return fmt.Errorf("нет каталога для приложения")
	}

	//log.Println("app dir", res)
	//	log.Printf("INFO2:%+v\n", DiskInfo)
	return nil
}

/* res
 app dir map[
	  _embedded:map[
		  items:[
			  map[comment_ids: map[
				  private_resource:183882:2cd0e58f947240dfc81f22ffd917d6669ec8dafeb6c3491c1796ab57244e470e
				  public_resource:183882:2cd0e58f947240dfc81f22ffd917d6669ec8dafeb6c3491c1796ab57244e470e]
				  created:2021-10-17T18:38:14+00:00 exif:map[]
				  modified:2021-10-17T18:38:14+00:00
				  name:ANI YandexDisk
				  path:disk:/Приложения/Disk/ANI YandexDisk
				  resource_id:183882:2cd0e58f947240dfc81f22ffd917d6669ec8dafeb6c3491c1796ab57244e470e
				  revision:1.63449589492827e+15 type:dir]]
				  limit:20
				  offset:0
				  path:disk:/Приложения/Disk
				  sort: total:1]
				  comment_ids:map[
					  private_resource:183882:82bb2c192275623aea54a0a45190db7e5787ce04d54306f5aafbe8e131bc4cf6
					  public_resource:183882:82bb2c192275623aea54a0a45190db7e5787ce04d54306f5aafbe8e131bc4cf6]
					  created:2021-10-17T18:38:13+00:00 exif:map[]
					  modified:2021-10-17T18:38:13+00:00
					  name:Disk path:disk:/Приложения/Disk
					  resource_id:183882:82bb2c192275623aea54a0a45190db7e5787ce04d54306f5aafbe8e131bc4cf6
					  revision:1.634495893921299e+15 type:dir]
*/
