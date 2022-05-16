# NIC.ru API Golang Client Utility
- Консольная утилита для работы с апишкой NIC.ru
- Использует библиотеку https://github.com/libdns/nicrudns

## Сборка
```
make build
```

## Help
```$ nicru --help
oauth2 ключи нужно получить в ЛК nic.ru - https://www.nic.ru/manager/oauth.cgi?step=oauth.app_register
имя DNS-сервиса можно посмотреть по адресу https://www.nic.ru/manager/services.cgi?step=srv.my_dns&view.order_by=domain

Usage:
  nicru [command]

Available Commands:
  add-a              создать одну A-запись
  add-cname          создать одну CNAME-запись
  add-cnames         позволяет массово создавать синеймы для имён, хранимых в файле построчно
  commit             фиксирует изменения
  completion         Generate the autocompletion script for the specified shell
  delete-record      удалять одну запись
  delete-records     позволяет массово удалять записи с числовыми айдишниками, хранимыми в файле построчно
  download-zone      команда для скачивания зоны
  help               Help about any command
  list-a-records     ищет A-записи, с возможностью фильтрации по имени и таргету
  list-cname-records ищет CNAME-записи, с возможностью фильтрации по имени и таргету
  list-records       ищет записи, с возможностью фильтрации по типа, имени и таргету
  rollback           откатывает изменения

Flags:
      --dns-service-name string   имя DNS-сервиса в ЛК nic.ru (default "FOO")
  -h, --help                      help for nicru
      --oauth2-client-id string   nic.ru oauth2 client-id
      --oauth2-secret-id string   nic.ru oauth2 secret-id
      --password string           пароль от ЛК nic.ru (default "*******")
      --token-cache-path string   файл для хранения токена авторизации (default "/tmp/.nic.ru.token")
      --username string           логин от ЛК nic.ru (default "123456/NIC-D")
      --zone-name string          имя зоны (default "example.com")

Use "nicru [command] --help" for more information about a command.
```
