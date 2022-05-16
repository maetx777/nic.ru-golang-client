package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func rollbackCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   `rollback`,
		Short: `откатывает изменения`,
		Run: func(cmd *cobra.Command, args []string) {
			if _, err := apiClient.RollbackZone(zoneName); err != nil {
				logrus.Fatalln(err)
			} else {
				logrus.Infoln(`zone`, zoneName, `reverted`)
			}
		},
	}
	return cmd
}
