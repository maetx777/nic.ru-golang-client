# NIC.ru API Golang Client

see examples

## CLI Utility
go build ./cmd/nicru

see help:
```
$ ./nicru  --help
oauth2 ключи нужно получить в ЛК nic.ru - https://www.nic.ru/manager/oauth.cgi?step=oauth.app_register
имя DNS-сервиса можно посмотреть по адресу https://www.nic.ru/manager/services.cgi?step=srv.my_dns&view.order_by=domain

Usage:
  nicru [command]

Available Commands:
  add-a       добавление A-записей
  add-cnames  добавление CNAME-записей
  commit      фиксация изменений
  completion  Generate the autocompletion script for the specified shell
  delete      удаление записей
  help        Help about any command
  list-a      просмотр A-записей
  list-cnames просмотр CNAME-записей
  rollback    откат изменений

Flags:
      --cache-path string         путь до файла, где будет храниться авторизация от API (default "/tmp/.nic.ru.token")
  -h, --help                      help for nicru
      --oauth2-client-id string   oauth2 client id
      --oauth2-secret-id string   oauth2 secret id
      --password string           пароль от nic.ru
      --service-name string       имя DNS-сервиса (default "EXAMPLE")
      --username string           логин от ЛК nic.ru (******/NIC-D)
      --zone-name string          имя DNS-зоны (default "example.com")

Use "nicru [command] --help" for more information about a command.
```
