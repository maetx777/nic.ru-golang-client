package main

import (
	"github.com/libdns/nicrudns"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func rollbackCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   `rollback`,
		Short: `откатывает изменения`,
		Run: func(cmd *cobra.Command, args []string) {
			client := nicrudns.NewClient(provider)

			if _, err := client.RollbackZone(zoneName); err != nil {
				logrus.Fatalln(err)
			} else {
				logrus.Infoln(`zone`, zoneName, `reverted`)
			}
		},
	}
	return cmd
}
