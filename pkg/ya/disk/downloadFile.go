package disk

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/ark-go/yandexapi/pkg/appconf"
)

func DownloadByte(fileName string, fileByte *[]byte) error {
	//body, err := ioutil.ReadAll(r.Body)
	return downloadFile(fileName, "", fileByte)
}

func DownloadFile(fileName string, filePath string) error {
	return downloadFile(fileName, filePath, nil)
}

func downloadFile(fileName string, filePath string, fileByte *[]byte) error {
	pathYafile, err := getDownloadUrl(fileName)
	if err != nil {
		return err
	}
	if pathYafile == "" {
		return fmt.Errorf("%s", "не получили ссылку для скачивания ыайла")
	}

	// Create blank file
	var file *os.File
	if fileByte == nil {
		file, err = os.Create(filePath)
		if err != nil {
			return err
			//	log.Fatal(err)
		}
		defer file.Close()
	}
	req, reqerr := http.NewRequest("GET", pathYafile, nil) //strings.NewReader(form.Encode())
	if reqerr != nil {
		return reqerr
		//log.Fatal(reqerr)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", appconf.Conf.YaToken.AccessToken)
	resp, reserr := http.DefaultClient.Do(req)
	if reserr != nil {
		return reserr
		//	log.Fatal(reserr)
	}
	defer resp.Body.Close()
	if fileByte == nil {
		size, copyerr := io.Copy(file, resp.Body)
		if copyerr != nil {
			return copyerr
			//log.Fatal(copyerr)
		}
		log.Println("Получено:", size)
	} else {
		*fileByte, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("byte down:", err.Error())
			return err
		}

	}
	return nil
}

// ошибки стандартные вероятно
