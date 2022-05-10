package nic_ru_golang_client

type Config struct {
	Credentials    *Credentials // структура с данными авторизации в апишке
	ZoneName       string       // имя зоны
	DnsServiceName string       // имя DNS-услуги
	CachePath      string       // путь для хранения токена авторизации
}

type Credentials struct {
	OAuth2   *OAuth2Creds
	Username string
	Password string
}

type OAuth2Creds struct {
	ClientID string
	SecretID string
}
