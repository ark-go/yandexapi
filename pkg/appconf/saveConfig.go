package appconf

import (
	"bytes"
	"compress/gzip"
	"encoding/gob"

	"os"
)

func (conf *conf) SaveConfig() error {

	fi, err := os.Create(conf.fileConfigPath) // Откроем файл для записи/перезаписи
	if err != nil {
		return err
	}
	defer fi.Close() // при выходе закроем
	//
	w := new(bytes.Buffer) // подготовим буфер, здесь будут упакованные в Zip данные
	//fz := gzip.NewWriter(fi)
	fz := gzip.NewWriter(w)       // настроим zip Writer для записи в буфер
	defer fz.Close()              // при выходе закроем Writer
	encoder := gob.NewEncoder(fz) // настроим GobEncoder на использование Zip Writer
	err = encoder.Encode(&conf)   // Упаковываем Gob encoder в zip-writer
	//err = encoder.Encode(criptConf)
	if err != nil {
		return err
	}
	fz.Flush() // затолкнем все что могло быть в gzip буфере, не успело записаться в наш буфер.

	cript := criptConfig(w.Bytes()) // отправим полученное на шифровку
	//data2 := bytes.NewReader(cript)
	//	fi.Write(w.Bytes()) // пишем в файл
	fi.Write(cript) // запишем зашифрованное в файл

	log.Println("Параметры сохранены.")
	return nil
}
