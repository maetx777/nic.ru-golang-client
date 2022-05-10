package nic_ru_golang_client

type IClient interface {
	AddA(names []string, target string, ttl int) (*Response, error)
	AddCnames(names []string, target string, ttl int) (*Response, error)
	CommitZone() (*Response, error)
	DeleteRecord(id int) (*Response, error)
	DownloadZone() (string, error)
	GetARecords(nameFilter string, targetFilter string) ([]*RR, error)
	GetCnameRecords(nameFilter string, targetFilter string) ([]*RR, error)
	RollbackZone() (*Response, error)
	GetServices() ([]*Service, error)
}
