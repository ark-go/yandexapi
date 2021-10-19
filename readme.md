Чтение и запись файла на Yandex диск.
Это все что умеем.

1. получить Applicatio ID https://oauth.yandex.com/client/new
2. выбрать Platform WebService
3. указать Calback https://oauth.yandex.com/verification_code
4. выбрать Permissions
   - Access to app folder in Yandex.Disk
   - Access to information about Yandex.Disk

- сборка: смотреть makefile
- запуск примера xxx.exe -AppId=XXXX -AppPass=YYYY параметры для ввода данных о приложении yandex disk, записывается в конфиг-файл

```go
    g := []byte("Привет мир!")
	disk.UploadByte("test.txt", &g)
	disk.UploadFile(filepath.Join(appconf.Conf.RootDir, "dddd.jpg"))
	disk.DownloadFile("dddd.jpg", filepath.Join(appconf.Conf.RootDir, "downl_1.jpg"))
	var fileByte []byte
	if err := disk.DownloadByte("dddd.jpg", &fileByte); err != nil {
		log.Println("error:", err.Error())
	} else {
		err := ioutil.WriteFile(filepath.Join(appconf.Conf.RootDir, "downl_2.jpg"), fileByte, fs.FileMode(0640))
		if err != nil {
			log.Println("err:", err)
		}
	}
```

- v0.0.5
  шифруем конфиг
- v0.0.3 старт
