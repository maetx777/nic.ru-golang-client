package examples

import (
	"fmt"
	api "github.com/maetx777/nic.ru-golang-client/client"
)

func DownloadZone() {
	config := &api.Config{
		Credentials: &api.Credentials{
			OAuth2: &api.OAuth2Creds{
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
	client := api.NewClient(config)
	if zoneContent, err := client.DownloadZone(); err != nil {
		fmt.Printf(`Download zone error: %s`, err.Error())
		return
	} else {
		fmt.Println(zoneContent)
	}
}
