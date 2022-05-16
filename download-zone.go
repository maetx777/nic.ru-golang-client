package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func downloadZoneCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "download-zone",
		Short: `команда для скачивания зоны`,
		Run: func(cmd *cobra.Command, args []string) {
			if data, err := apiClient.DownloadZone(zoneName); err != nil {
				logrus.Fatalln(err)
			} else {
				fmt.Println(data)
			}
		},
	}
	return cmd
}
