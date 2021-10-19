Чтение и запись файла на Yandex диск.
Это все что умеем.

Applicatio ID предоставляет доступ к API Yandex Disk, пользователь приложения должен будет разрешить, через Яндекс, доступ приложения к своему Yandex Disk. в данном примере доступ будет только к Disk:/Приложения/XXX XXX - название приложения указывается при регистрации (пункт 2)  
?? _Application ID сам по себе не дает доступа к диску, будет запрашиваться разрешение пользователя. вероятно что ID с паролем можно зашить в приложение_

1. получить Applicatio ID https://oauth.yandex.com/client/new
2. ввести имя приложения (будет каталогом)
3. выбрать Platform WebService
4. указать Calback https://oauth.yandex.com/verification_code
5. выбрать Permissions
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

- v0.0.7 шифруем конфиг серийным номером диска
- v0.0.5 шифруем конфиг
- v0.0.3 старт
