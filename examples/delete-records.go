package examples

import (
	"fmt"
	"github.com/maetx777/nic.ru-golang-client"
)

func DeleteRecords() {
	config := &nic_ru_golang_client.Config{
		Credentials: &nic_ru_golang_client.Credentials{
			OAuth2: &nic_ru_golang_client.OAuth2Creds{
				ClientID: "qwe",
				SecretID: "xyz",
			},
			Username: "foo",
			Password: "bar",
		},
		ZoneName:       "example.com",
		DnsServiceName: "EXAMPLE",
		CachePath:      "/tmp/.nic.ru.token",
	}
	client := nic_ru_golang_client.NewClient(config)
	if response, err := client.DeleteRecord(123456); err != nil {
		fmt.Printf(`Delete record error: %s`, err.Error())
		return
	} else {
		for _, record := range response.Data.Zone[0].Rr {
			fmt.Printf(`Deleted record: %s CNAME %s`, record.Name, record.A.String())
		}
	}
	if _, err := client.CommitZone(); err != nil {
		fmt.Printf(`Commit error: %s`, err.Error())
	} else {
		fmt.Printf(`Zone committed`)
	}
}
