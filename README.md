# http-rest-api - Простой REST API сервер
1. Находясь в корне проекта, нужно выполнить команду:

```bash
make
```
для создания исполняемого файла apiserver.

2. После создания файла apiserver, можно посмотреть доступные команды с помощью:

```bash
./apiserver -help
```

3. Файл конфигурации находится в configs/apiserver.toml
Для считывания его данных используется пакет https://github.com/burntsushi/toml

Пакет устанавливается командой:
```bash
go get github.com/BurntSushi/toml
```
Запуск сервера с настройками из данного конфигурационным файла осуществляется командой:
```bash
./apiserver -config-path configs/apiserver.toml
```
4. Логирование apiserver реализовано с использованием пакета Logrus: https://github.com/sirupsen/logrus
Уровень логирования можно изменить в файле nfigs/apiserver.toml, в переменной: log_level.
По-умолчанию (если в файле apiserver.toml не указан уровень) log_level = "debug".
Возможно указание следующих уровней:
  - debug
  - info
  - warning
  - panic
  - fatal 

5. Роутер реализован на базе пакета https://github.com/gorilla/mux
Установка:
```bash
go get -u github.com/gorilla/mux 
```

6. Для тестов используется пакет https://github.com/stretchr/testify
  Установка:
  ```bash
  go get github.com/stretchr/testify
  ```
