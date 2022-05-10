package main

import (
	api "github.com/maetx777/nic.ru-golang-client"
)

var config = &api.Config{
	Credentials: &api.Credentials{
		OAuth2: &api.OAuth2Creds{
			ClientID: "",
			SecretID: "",
		},
		Username: "",
		Password: "",
	},
	ZoneName:       "",
	DnsServiceName: "",
	CachePath:      "",
}
