package main

import (
	"github.com/libdns/nicrudns"
	"github.com/spf13/cobra"
)

var (
	provider  = &nicrudns.Provider{}
	apiClient = nicrudns.NewClient(provider)
	zoneName  string
)

func main() {
	cmd := &cobra.Command{
		Use:   "nicru",
		Short: `утилита для работы с апишкой nic.ru`,
		Long: `oauth2 ключи нужно получить в ЛК nic.ru - https://www.nic.ru/manager/oauth.cgi?step=oauth.app_register
имя DNS-сервиса можно посмотреть по адресу https://www.nic.ru/manager/services.cgi?step=srv.my_dns&view.order_by=domain`,
	}

	cmd.PersistentFlags().StringVar(&provider.OAuth2ClientID, "oauth2-client-id", ``, `nic.ru oauth2 client-id`)
	cmd.PersistentFlags().StringVar(&provider.OAuth2SecretID, "oauth2-secret-id", ``, `nic.ru oauth2 secret-id`)
	cmd.PersistentFlags().StringVar(&provider.Username, "username", `123456/NIC-D`, `логин от ЛК nic.ru`)
	cmd.PersistentFlags().StringVar(&provider.Password, "password", `*******`, `пароль от ЛК nic.ru`)
	cmd.PersistentFlags().StringVar(&zoneName, `zone-name`, `example.com`, `имя зоны`)
	cmd.PersistentFlags().StringVar(&provider.DnsServiceName, `dns-service-name`, `FOO`, `имя DNS-сервиса в ЛК nic.ru`)
	cmd.PersistentFlags().StringVar(&provider.CachePath, `token-cache-path`, `/tmp/.nic.ru.token`, `файл для хранения токена авторизации`)

	cmd.AddCommand(downloadZoneCmd())
	cmd.AddCommand(listRecordsCmd())
	cmd.AddCommand(listARecordsCmd())
	cmd.AddCommand(listCnameRecordsCmd())
	cmd.AddCommand(deleteRecordsCmd())
	cmd.AddCommand(deleteRecord())
	cmd.AddCommand(addCnamesCmd())
	cmd.AddCommand(addCnameCmd())
	cmd.AddCommand(addACmd())
	cmd.AddCommand(commitCmd())
	cmd.AddCommand(rollbackCmd())

	_ = cmd.Execute()
}
