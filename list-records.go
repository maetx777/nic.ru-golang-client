package main

import (
	"fmt"
	"github.com/libdns/nicrudns"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"regexp"
)

func listRecordsCmd() *cobra.Command {
	var (
		nameFilter    string
		targetFilter  string
		recTypeFilter string
	)
	cmd := &cobra.Command{
		Use:   `list-records`,
		Short: `ищет записи, с возможностью фильтрации по типа, имени и таргету`,
		Run: func(cmd *cobra.Command, args []string) {
			client := nicrudns.NewClient(provider)
			if err := doListRecords(client, recTypeFilter, nameFilter, targetFilter); err != nil {
				logrus.Fatalln(err)
			}
		},
	}
	cmd.PersistentFlags().StringVar(&recTypeFilter, `type`, `^(A|AAAA|CNAME|TXT|MX)$`, `типа записи, регулярка`)
	cmd.PersistentFlags().StringVar(&targetFilter, `target-filter`, ``, `фильтр по таргету, регулярка`)
	cmd.PersistentFlags().StringVar(&nameFilter, `name-filter`, ``, `фильтр по имени записи, регулярка`)
	return cmd
}

func listARecordsCmd() *cobra.Command {
	var (
		nameFilter   string
		targetFilter string
	)
	cmd := &cobra.Command{
		Use:   `list-a-records`,
		Short: `ищет A-записи, с возможностью фильтрации по имени и таргету`,
		Run: func(cmd *cobra.Command, args []string) {
			client := nicrudns.NewClient(provider)
			if err := doListRecords(client, `^A$`, nameFilter, targetFilter); err != nil {
				logrus.Fatalln(err)
			}
		},
	}
	cmd.PersistentFlags().StringVar(&targetFilter, `target-filter`, ``, `фильтр по таргету, регулярка`)
	cmd.PersistentFlags().StringVar(&nameFilter, `name-filter`, ``, `фильтр по имени записи, регулярка`)
	return cmd
}

func listCnameRecordsCmd() *cobra.Command {
	var (
		nameFilter   string
		targetFilter string
	)
	cmd := &cobra.Command{
		Use:   `list-cname-records`,
		Short: `ищет CNAME-записи, с возможностью фильтрации по имени и таргету`,
		Run: func(cmd *cobra.Command, args []string) {
			client := nicrudns.NewClient(provider)
			if err := doListRecords(client, `^CNAME$`, nameFilter, targetFilter); err != nil {
				logrus.Fatalln(err)
			}
		},
	}
	cmd.PersistentFlags().StringVar(&targetFilter, `target-filter`, ``, `фильтр по таргету, регулярка`)
	cmd.PersistentFlags().StringVar(&nameFilter, `name-filter`, ``, `фильтр по имени записи, регулярка`)
	return cmd
}

func doListRecords(client nicrudns.IClient, typeFilter string, nameFilter string, targetFilter string) error {
	if rrs, err := client.GetRecords(zoneName); err != nil {
		return err
	} else {
		nameRegexp, err := regexp.Compile(nameFilter)
		if err != nil {
			return err
		}
		targetRegexp, err := regexp.Compile(targetFilter)
		if err != nil {
			return err
		}
		typeRegexp, err := regexp.Compile(typeFilter)
		if err != nil {
			return err
		}
		for _, rr := range rrs {
			if typeFilter != `` && !typeRegexp.MatchString(rr.Type) {
				continue
			}
			if nameFilter != `` && !nameRegexp.MatchString(rr.Name) {
				continue
			}
			if targetFilter != `` {
				switch rr.Type {
				case `A`:
					if !targetRegexp.MatchString(rr.A.String()) {
						continue
					}
				case `AAAA`:
					if !targetRegexp.MatchString(rr.AAAA.String()) {
						continue
					}
				case `CNAME`:
					if !targetRegexp.MatchString(rr.Cname.Name) {
						continue
					}
				case `MX`:
					if !targetRegexp.MatchString(rr.Mx.Exchange.Name) {
						continue
					}
				case `TXT`:
					if !targetRegexp.MatchString(rr.Txt.String) {
						continue
					}
				default:
					continue
				}
			}
			printRecord(rr)
		}
		return nil
	}
}

func printRecord(rr *nicrudns.RR) {
	switch rr.Type {
	case `A`:
		fmt.Printf("%s %s %s %s %s\n", rr.ID, rr.Name, rr.Ttl, rr.Type, rr.A.String())
	case `AAAA`:
		fmt.Printf("%s %s %s %s %s\n", rr.ID, rr.Name, rr.Ttl, rr.Type, rr.AAAA.String())
	case `CNAME`:
		fmt.Printf("%s %s %s %s %s\n", rr.ID, rr.Name, rr.Ttl, rr.Type, rr.Cname.Name)
	case `MX`:
		fmt.Printf("%s %s %s %s %s %s\n", rr.ID, rr.Name, rr.Ttl, rr.Type, rr.Mx.Preference, rr.Mx.Exchange.Name)
	case `TXT`:
		fmt.Printf("%s %s %s %s \"%s\"\n", rr.ID, rr.Name, rr.Ttl, rr.Type, rr.Txt.String)
	}
}
