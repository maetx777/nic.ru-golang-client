package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use:   `nicru`,
		Short: `утилита для управления записями в DNS NIC.ru`,
	}
	cmd.AddCommand(addACmd())   // подключаем команду add-a
	cmd.AddCommand(commitCmd()) // подключаем команду add-a
	if err := cmd.Execute(); err != nil {
		logrus.Infoln(err.Error())
	}
}
