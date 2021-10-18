package appconf

import (
	"compress/gzip"
	"encoding/gob"

	"os"
)

func (conf *conf) SaveConfig() error {

	fi, err := os.Create(conf.fileConfigPath)
	if err != nil {
		return err
	}
	defer fi.Close()

	fz := gzip.NewWriter(fi)
	defer fz.Close()

	encoder := gob.NewEncoder(fz)

	err = encoder.Encode(&conf)
	if err != nil {
		return err
	}
	log.Println("Параметры сохранены.")
	return nil
}
