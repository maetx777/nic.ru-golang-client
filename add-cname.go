package main

import (
	"bufio"
	"bytes"
	"github.com/libdns/nicrudns"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

func addCnamesCmd() *cobra.Command {
	var (
		filePath string
		target   string
		ttl      string
	)
	cmd := &cobra.Command{
		Use:   `add-cnames`,
		Short: `позволяет массово создавать синеймы для имён, хранимых в файле построчно`,
		Run: func(cmd *cobra.Command, args []string) {
			client := nicrudns.NewClient(provider)

			if err := addCnamesFromFile(client, filePath, target, ttl); err != nil {
				logrus.Fatalln(err)
			}

		},
	}
	cmd.PersistentFlags().StringVar(&filePath, `file-path`, ``, `файл со списком имён записей`)
	cmd.PersistentFlags().StringVar(&target, `target`, ``, `куда отправить записи из списка`)
	cmd.PersistentFlags().StringVar(&ttl, `ttl`, `60`, `ttl для записи`)
	return cmd
}

func addCnameCmd() *cobra.Command {
	var (
		target string
		ttl    string
		name   string
	)
	cmd := &cobra.Command{
		Use:   `add-cname`,
		Short: `создать одну CNAME-запись`,
		Run: func(cmd *cobra.Command, args []string) {
			client := nicrudns.NewClient(provider)

			if response, err := client.AddCnames(zoneName, []string{name}, target, ttl); err != nil {
				logrus.Fatalln(err)
			} else {
				for _, rr := range response.Data.Zone[0].Rr {
					logrus.Infoln(`new record`, rr.ID, rr.Name, rr.Type, rr.Ttl, rr.Cname.Name)
				}
			}
		},
	}
	cmd.PersistentFlags().StringVar(&name, `name`, ``, `имя записи`)
	cmd.PersistentFlags().StringVar(&target, `target`, ``, `куда отправить запись`)
	cmd.PersistentFlags().StringVar(&ttl, `ttl`, `60`, `ttl для записи`)
	return cmd
}

func addCnamesFromFile(client nicrudns.IClient, filePath string, target string, ttl string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		logrus.Fatalln(`file error: %s`, err)
	}
	scanner := bufio.NewScanner(bytes.NewBuffer(data))
	var names []string
	for scanner.Scan() {
		names = append(names, scanner.Text())
	}

	if response, err := client.AddCnames(zoneName, names, target, ttl); err != nil {
		return err
	} else {
		for _, rr := range response.Data.Zone[0].Rr {
			logrus.Infoln(`new record`, rr.ID, rr.Name, rr.Type, rr.Ttl, rr.Cname.Name)
		}
		return nil
	}

}
