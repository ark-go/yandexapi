package appconf

import (
	"bytes"
	"compress/gzip"
	"encoding/gob"
	"errors"
	"io/ioutil"

	"os"
)

func init() {

}

func (conf *conf) LoadConfig() error {

	fi, err := os.Open(conf.fileConfigPath)
	if err != nil {
		log.Printf("Инициализация конфигурационного файла\n%s\n", conf.fileConfigPath)
		if err := conf.SaveConfig(); err != nil { // ну если ошибка - создадим новый.
			return errors.New("ошибка при инициализации конфигурационного файла\n" + err.Error())
		} else {
			fi, err = os.Open(conf.fileConfigPath)
			if err != nil {
				return errors.New("ошибка при открытии конфигурационного файла\n" + err.Error())
			}
			defer fi.Close()
		}

	} else {
		defer fi.Close()
	}
	data, err := ioutil.ReadAll(fi) // Читаем файл в []byte
	if err != nil {
		return err
	}
	decr := decriptConfig(data)    // отправляем на расшифровку
	data2 := bytes.NewReader(decr) // создаем bytes.Reader из []byte

	//data2 := bytes.NewReader(data)
	fz, err := gzip.NewReader(data2) // zip Reader распакуем при чтении
	//fz, err := gzip.NewReader(fi)
	if err != nil {
		log.Println("Ошибка при чтении конфигурационного файла", err.Error())
		return err
	}
	defer fz.Close()

	decoder := gob.NewDecoder(fz) // подготовим Gob для чтения данных
	err = decoder.Decode(&conf)   // декодируем Gob в структуру
	//err = decoder.Decode(reqBodyBytes)
	if err != nil {
		log.Println("Ошибка при разборе конфигурационного файла", err.Error())
		return err
	}

	return nil
}
