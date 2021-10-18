package appconf

import (
	"compress/gzip"
	"encoding/gob"
	"errors"

	"os"
)

func init() {

}

func (conf *conf) LoadConfig() error {

	fi, err := os.Open(conf.fileConfigPath)
	if err != nil {
		log.Printf("Инициализация конфигурационного файла\n%s\n", conf.fileConfigPath)
		if err := conf.SaveConfig(); err != nil {
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
	fz, err := gzip.NewReader(fi)
	if err != nil {
		log.Println("Ошибка при чтении конфигурационного файла", err.Error())
		return err
	}
	defer fz.Close()

	decoder := gob.NewDecoder(fz)
	err = decoder.Decode(&conf)
	if err != nil {
		log.Println("Ошибка при разборе конфигурационного файла", err.Error())
		return err
	}
	return nil
}
