package main

import (
	"github.com/libdns/nicrudns"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func commitCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   `commit`,
		Short: `фиксирует изменения`,
		Run: func(cmd *cobra.Command, args []string) {
			client := nicrudns.NewClient(provider)

			if _, err := client.CommitZone(zoneName); err != nil {
				logrus.Fatalln(err)
			} else {
				logrus.Infoln(`zone`, zoneName, `committed`)
			}
		},
	}
	return cmd
}
