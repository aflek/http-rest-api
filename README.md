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

7. Для работы с БД используется пакет database-sql. Учебник здесь: http://go-database-sql.org
   В качестве альтернативы можно рассмотреть пакет https://github.com/jmoiron/sqlx

8. Posgres драйвер здесь: https://github.com/lib/pq
   Устанавливается командой:
   ```bash
   go get github.com/lib/pq 
   ```
9. Создаем базу данных командой:
    ```bash
      createdb restapi_dev
    ```
    Запускаем сервер:
    ```bash
      make; ./apiserver
    ```
10. Для применения миграций используем пакет https://github.com/golang-migrate/migrate
    Установка под Ubuntu из 4 команд: 
    ```bash
    curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | sudo apt-key add -
    echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" > sudo /etc/apt/sources.list.d/migrate.list
    sudo apt-get update
    sudo apt-get install -y migrate
    ```
    Успешность установки монжо провреть командой:
    ```bash
    migrate -help
    ```
    Создаем миграции. Находясь в каталоге проекта выполняем команду:
    ```bash
    migrate create -ext sql -dir migrations create_users
    ```
    В результате в каталоге migrations будут созданы два пустых файла с расширением sql, вида:
    20200917182500_create_users.up.sql
    20200917182500_create_users.down.sql

11. Запуск миграций для dev БД:
    Под текущим пользователем:
    ```bash
    migrate -path migrations -database "postgres://localhost:5432/restapi_dev?sslmode=disable" up
    ```
    Под пользователем "postgres" и паролем "123456":
    ```bash
    migrate -path migrations -database "postgres://postgres:123456@localhost:5432/restapi_dev?sslmode=disable" up
    ```
    Проверяем резузльтат.
    Подключаемся к базе: psql -d restapi_dev
    Выводим списко таблиц: \dt
    Смотрим таблицу users: \d users;
    Выходим из базы данных: \q

12. Запуск миграций для тестовой БД:
    createdb restapi_test
    migrate -path migrations -database "postgres://postgres:123456@localhost:5432/restapi_test?sslmode=disable" up
