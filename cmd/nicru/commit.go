package main

import (
	"fmt"
	api "github.com/maetx777/nic.ru-golang-client/client"
	"github.com/spf13/cobra"
)

func commitCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:   `commit`,
		Short: `фиксация изменений`,
		Run: func(cmd *cobra.Command, args []string) {
			//инициализируем клиента
			client := api.NewClient(config)
			//коммитим результат
			if _, err := client.CommitZone(); err != nil {
				fmt.Printf("Commit error: %s\n", err.Error())
			} else {
				fmt.Printf("Zone committed\n")
			}
		},
	}

	return cmd
}
