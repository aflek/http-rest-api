# http-rest-api - Простой REST API сервер
1. Находясь в коне проекта нужно выполнить команду

```bash
make
```

для создания исполняемого файла apiserver.

2. После воздания файла apiserver можно посмотреть доступные команды с помощью:

```bash
./apiserver -help
```

3. Файл конфигурации находится в configs/apiserver.toml
Для считывания его данных используется пакет https://github.com/burntsushi/toml

```bash
go get github.com/BurntSushi/toml
```
Запуск сервера с настройками из данного конфигурационным файла осуществляется командой:
./apiserver -config-path configs/apiserver.toml
