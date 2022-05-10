package main

import (
	"fmt"
	api "github.com/maetx777/nic.ru-golang-client/client"
	"github.com/spf13/cobra"
)

func rollbackCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:   `rollback`,
		Short: `откат изменений`,
		Run: func(cmd *cobra.Command, args []string) {
			//инициализируем клиента
			client := api.NewClient(config)
			//коммитим результат
			if _, err := client.RollbackZone(); err != nil {
				fmt.Printf("Rollback error: %s\n", err.Error())
			} else {
				fmt.Printf("Zone reverted\n")
			}
		},
	}

	return cmd
}
