package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use:   `nicru`,
		Short: `утилита для управления записями в DNS NIC.ru`,
		Long: `oauth2 ключи нужно получить в ЛК nic.ru - https://www.nic.ru/manager/oauth.cgi?step=oauth.app_register
имя DNS-сервиса можно посмотреть по адресу https://www.nic.ru/manager/services.cgi?step=srv.my_dns&view.order_by=domain
`,
	}
	//создаём флаги для заполнения конфига
	cmd.PersistentFlags().StringVar(&config.Credentials.OAuth2.ClientID, `oauth2-client-id`, ``, `oauth2 client id`)
	cmd.PersistentFlags().StringVar(&config.Credentials.OAuth2.SecretID, `oauth2-secret-id`, ``, `oauth2 secret id`)
	cmd.PersistentFlags().StringVar(&config.Credentials.Username, `username`, ``, `логин от ЛК nic.ru (******/NIC-D)`)
	cmd.PersistentFlags().StringVar(&config.Credentials.Password, `password`, ``, `пароль от nic.ru`)
	cmd.PersistentFlags().StringVar(&config.ZoneName, `zone-name`, `example.com`, `имя DNS-зоны`)
	cmd.PersistentFlags().StringVar(&config.DnsServiceName, `service-name`, `EXAMPLE`, `имя DNS-сервиса`)
	cmd.PersistentFlags().StringVar(&config.CachePath, `cache-path`, `/tmp/.nic.ru.token`, `путь до файла, где будет храниться авторизация от API`)

	cmd.AddCommand(addACmd())   // подключаем команду add-a
	cmd.AddCommand(commitCmd()) // подключаем команду add-a
	if err := cmd.Execute(); err != nil {
		logrus.Infoln(err.Error())
	}
}
