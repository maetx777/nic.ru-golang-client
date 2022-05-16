package main

import (
	"fmt"
	"github.com/libdns/nicrudns"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func downloadZoneCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "download-zone",
		Short: `команда для скачивания зоны`,
		Run: func(cmd *cobra.Command, args []string) {
			client := nicrudns.NewClient(provider)
			if data, err := client.DownloadZone(zoneName); err != nil {
				logrus.Fatalln(err)
			} else {
				fmt.Println(data)
			}
		},
	}
	return cmd
}
