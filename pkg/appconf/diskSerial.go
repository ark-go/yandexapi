package appconf

import (
	"fmt"
	"github.com/jaypipes/ghw"
)

func init() {
}

// получаем серийный номер первого диска в системе
//	точнее первый доступный серийный номер.
func diskSerial() ([]byte, error) {
	block, err := ghw.Block()
	if err != nil {
		return nil, fmt.Errorf("%s", "ошибка при получении информации о блочном хранилище: "+err.Error())
	}
	for _, disk := range block.Disks {
		if disk.SerialNumber != "" {
			return []byte(disk.SerialNumber), nil
		}
	}
	return nil, fmt.Errorf("%s", "не найдено ни одного серийника.")
}
