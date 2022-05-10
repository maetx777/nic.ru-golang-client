package examples

import (
	"fmt"
	"github.com/maetx777/nic.ru-golang-client"
)

func DownloadZone() {
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
	if zoneContent, err := client.DownloadZone(); err != nil {
		fmt.Printf(`Download zone error: %s`, err.Error())
		return
	} else {
		fmt.Println(zoneContent)
	}
}
