package disk

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/ark-go/yandexapi/pkg/appconf"
)

func UploadByte(fileName string, buffer []byte) error {
	buf := bytes.NewReader(buffer)
	return uploadFile(fileName, buf)
}

func UploadFile(filePath string) error {

	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("не открыть файл %s", err.Error())
	}
	defer file.Close()
	filename := filepath.Base(filePath)
	return uploadFile(filename, file)

}

func uploadFile(filename string, file io.Reader) error {
	if appconf.Conf.AppDirYandex == "" {
		return fmt.Errorf("%s", "Не определен каталог приложения")
	}
	pathfile, err := GetUploadUrl(appconf.Conf.AppDirYandex + "/" + filename)
	if err != nil {
		return err
	}

	// file, err := os.Open(`C:\Users\Zinaida\Pictures\2018-06-29 19.56.13.jpg`)
	// if err != nil {
	// 	return fmt.Errorf("не открыть файл %s", err.Error())
	// }
	// defer file.Close()
	//r := bytes.NewReader(byteData) // []byte
	//	response, err := http.Post(url, "binary/octet-stream", file)
	request, err1 := http.NewRequest("PUT", pathfile, file)
	if err1 != nil {
		return fmt.Errorf("сборка запроса не удалась: %s", err1.Error())
	}
	request.Header.Add("Content-Type", "binary/octet-stream")
	request.Header.Add("Authorization", appconf.Conf.YaToken.AccessToken)

	resp, reserr := http.DefaultClient.Do(request)
	if reserr != nil {
		return fmt.Errorf("запрос не удался: %s", reserr.Error())
	}

	defer resp.Body.Close()

	//log.Printf("con: %+v", resp)
	if resp.StatusCode == 201 {
		url, _ := resp.Location()
		log.Printf("Записан файл: %s", url.Path)
		return nil
	}
	return fmt.Errorf("передача файла завершилось ошибкой, status: %s", resp.Status)
}
