package main

import (
	"fmt"
	api "github.com/maetx777/nic.ru-golang-client/client"
	"github.com/spf13/cobra"
)

func listACmd() *cobra.Command {

	var (
		nameFilter   string //переменная хранит регулярку фильтра по имени
		targetFilter string //переменная хранит регулярку фильтра по таргету
	)

	cmd := &cobra.Command{
		Use:   `list-a`,
		Short: `просмотр A-записей`,

		Run: func(cmd *cobra.Command, args []string) {
			//инициализируем клиента
			client := api.NewClient(config)
			if records, err := client.GetARecords(nameFilter, targetFilter); err != nil {
				//обрабатываем ошибку
				fmt.Printf("Get A-records error: %s\n", err.Error())
				return
			} else {
				//печатаем результат
				for _, record := range records {
					fmt.Printf("%s %s IN A %s %s\n", record.ID, record.Name, record.Ttl, record.A)
				}
			}
		},
	}
	cmd.PersistentFlags().StringVar(&nameFilter, `name-filter`, `^.*$`, `фильтр по имени записи`)
	cmd.PersistentFlags().StringVar(&targetFilter, `target-filter`, `^.*$`, `фильтр по таргету`)

	return cmd
}
