package api

import (
	"encoding/xml"
)

type Request struct {
	XMLName xml.Name `xml:"request"`
	Text    string   `xml:",chardata"`
	RrList  *RrList  `xml:"rr-list"`
}

type RrList struct {
	Text string `xml:",chardata"`
	Rr   []*RR  `xml:"rr"`
}

type RR struct {
	Text    string `xml:",chardata"`
	ID      string `xml:"id,attr,omitempty"`
	Name    string `xml:"name"`
	IdnName string `xml:"idn-name,omitempty"`
	Ttl     string `xml:"ttl"`
	Type    string `xml:"type"`
	Soa     *Soa   `xml:"soa"`
	Ns      *Ns    `xml:"ns"`
	A       *A     `xml:"a"`
	Txt     *Txt   `xml:"txt"`
	Cname   *Cname `xml:"cname"`
}

type A string

func (a *A) String() string {
	return string(*a)
}

type Service struct {
	Text         string `xml:",chardata"`
	Admin        string `xml:"admin,attr"`
	DomainsLimit string `xml:"domains-limit,attr"`
	DomainsNum   string `xml:"domains-num,attr"`
	Enable       string `xml:"enable,attr"`
	HasPrimary   string `xml:"has-primary,attr"`
	Name         string `xml:"name,attr"`
	Payer        string `xml:"payer,attr"`
	Tariff       string `xml:"tariff,attr"`
	RrLimit      string `xml:"rr-limit,attr"`
	RrNum        string `xml:"rr-num,attr"`
}

type Soa struct {
	Text  string `xml:",chardata"`
	Mname struct {
		Text    string `xml:",chardata"`
		Name    string `xml:"name"`
		IdnName string `xml:"idn-name,omitempty"`
	} `xml:"mname"`
	Rname struct {
		Text    string `xml:",chardata"`
		Name    string `xml:"name"`
		IdnName string `xml:"idn-name,omitempty"`
	} `xml:"rname"`
	Serial  string `xml:"serial"`
	Refresh string `xml:"refresh"`
	Retry   string `xml:"retry"`
	Expire  string `xml:"expire"`
	Minimum string `xml:"minimum"`
}

type Ns struct {
	Text    string `xml:",chardata"`
	Name    string `xml:"name"`
	IdnName string `xml:"idn-name,omitempty"`
}

type Cname struct {
	Text    string `xml:",chardata"`
	Name    string `xml:"name"`
	IdnName string `xml:"idn-name,omitempty"`
}

type Txt struct {
	Text   string `xml:",chardata"`
	String string `xml:"string"`
}

type Zone struct {
	Text       string `xml:",chardata"`
	Admin      string `xml:"admin,attr"`
	Enable     string `xml:"enable,attr"`
	HasChanges string `xml:"has-changes,attr"`
	HasPrimary string `xml:"has-primary,attr"`
	ID         string `xml:"id,attr"`
	IdnName    string `xml:"idn-name,attr"`
	Name       string `xml:"name,attr"`
	Payer      string `xml:"payer,attr"`
	Service    string `xml:"service,attr"`
	Rr         []*RR  `xml:"rr"`
}

type Revision struct {
	Text   string `xml:",chardata"`
	Date   string `xml:"date,attr"`
	Ip     string `xml:"ip,attr"`
	Number string `xml:"number,attr"`
}

type Error struct {
	Text string `xml:",chardata"`
	Code string `xml:"code,attr"`
}

type Response struct {
	XMLName xml.Name `xml:"response"`
	Text    string   `xml:",chardata"`
	Status  string   `xml:"status"`
	Errors  struct {
		Text  string `xml:",chardata"`
		Error *error `xml:"error"`
	} `xml:"errors"`
	Data struct {
		Text     string `xml:",chardata"`
		Service  []*Service
		Zone     []*Zone     `xml:"zone"`
		Address  []*A        `xml:"address"`
		Revision []*Revision `xml:"revision"`
	} `xml:"data"`
}
