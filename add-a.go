package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func addACmd() *cobra.Command {
	var (
		target string
		ttl    string
		name   string
	)
	cmd := &cobra.Command{
		Use:   `add-a`,
		Short: `создать одну A-запись`,
		Run: func(cmd *cobra.Command, args []string) {
			if response, err := apiClient.AddA(zoneName, []string{name}, target, ttl); err != nil {
				logrus.Fatalln(err)
			} else {
				for _, rr := range response.Data.Zone[0].Rr {
					logrus.Infoln(`new record`, rr.ID, rr.Name, rr.Type, rr.Ttl, rr.A.String())
				}
			}
		},
	}
	cmd.PersistentFlags().StringVar(&name, `name`, ``, `имя записи`)
	cmd.PersistentFlags().StringVar(&target, `target`, ``, `куда отправить запись`)
	cmd.PersistentFlags().StringVar(&ttl, `ttl`, `60`, `ttl для записи`)
	return cmd
}
