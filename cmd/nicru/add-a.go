package main

import (
	"fmt"
	api "github.com/maetx777/nic.ru-golang-client"
	"github.com/spf13/cobra"
)

func addACmd() *cobra.Command {

	var (
		names  []string //здесь будем хранить список записей, которые нужно создать
		target string   //здесь хранится ipv4-адрес, на которые будет ссылаться каждое имя, описанное выше
		ttl    int      //TTL для каждой записи
	)

	cmd := &cobra.Command{
		Use:   `add-a`,
		Short: `добавление A-записей`,

		Run: func(cmd *cobra.Command, args []string) {
			//инициализируем клиента
			client := api.NewClient(config)
			if response, err := client.AddA(names, target, ttl); err != nil {
				//обрабатываем ошибку
				fmt.Printf("Add A-record error: %s\n", err.Error())
				return
			} else {
				//печатаем результат
				for _, record := range response.Data.Zone[0].Rr {
					fmt.Printf("Added A-record: %s -> %s\n", record.Name, record.A.String())
				}
			}
		},
	}
	//создаём флаги для указания имён создаваемых записей, ipv4-таргета и TTL
	cmd.PersistentFlags().StringSliceVar(&names, `names`, []string{}, `имена, которые нужно создать`)
	cmd.PersistentFlags().StringVar(&target, `target`, ``, `куда нужно отправить имена (например, 1.2.3.4)`)
	cmd.PersistentFlags().IntVar(&ttl, `ttl`, 600, `TTL для созданным записей`)

	return cmd
}
