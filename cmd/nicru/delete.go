package main

import (
	"fmt"
	api "github.com/maetx777/nic.ru-golang-client"
	"github.com/spf13/cobra"
)

func deleteCmd() *cobra.Command {

	var (
		id int
	)

	cmd := &cobra.Command{
		Use:   `delete`,
		Short: `удаление записей`,

		Run: func(cmd *cobra.Command, args []string) {
			//инициализируем клиента
			client := api.NewClient(config)
			if _, err := client.DeleteRecord(id); err != nil {
				//обрабатываем ошибку
				fmt.Printf("Delete record error: %s\n", err.Error())
				return
			} else {
				//печатаем результат
				fmt.Printf("Deleted record id %d\n", id)
			}
		},
	}
	cmd.PersistentFlags().IntVar(&id, `id`, 0, `ID записи`)

	return cmd
}
