package main

import (
	"fmt"
	api "github.com/maetx777/nic.ru-golang-client"
	"github.com/spf13/cobra"
)

func listCnamescmd() *cobra.Command {

	var (
		nameFilter   string //переменная хранит регулярку фильтра по имени
		targetFilter string //переменная хранит регулярку фильтра по таргету
	)

	cmd := &cobra.Command{
		Use:   `list-cnames`,
		Short: `просмотр CNAME-записей`,

		Run: func(cmd *cobra.Command, args []string) {
			//инициализируем клиента
			client := api.NewClient(config)
			if records, err := client.GetCnameRecords(nameFilter, targetFilter); err != nil {
				//обрабатываем ошибку
				fmt.Printf("Get CNAME-records error: %s\n", err.Error())
				return
			} else {
				//печатаем результат
				for _, record := range records {
					fmt.Printf("%s %s IN CNAME %s %s\n", record.ID, record.Name, record.Ttl, record.Cname.Name)
				}
			}
		},
	}
	cmd.PersistentFlags().StringVar(&nameFilter, `name-filter`, `^.*$`, `фильтр по имени записи`)
	cmd.PersistentFlags().StringVar(&targetFilter, `target-filter`, `^.*$`, `фильтр по таргету`)

	return cmd
}
