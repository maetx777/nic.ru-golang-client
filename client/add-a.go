package api

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
)

func (client *Client) AddA(names []string, target string, ttl int) (*Response, error) {
	payload := &Request{
		RrList: &RrList{
			Rr: []*RR{},
		},
	}
	t := A(target)
	for _, name := range names {
		payload.RrList.Rr = append(payload.RrList.Rr, &RR{
			Name: name,
			Type: `A`,
			Ttl:  strconv.Itoa(ttl),
			A:    &t,
		})
	}

	buf := bytes.NewBuffer(nil)
	if err := xml.NewEncoder(buf).Encode(payload); err != nil {
		return nil, errors.Wrap(err, XmlEncodeError.Error())
	}

	url := fmt.Sprintf(AddRecordsUrlPattern, client.config.DnsServiceName, client.config.ZoneName)

	req, err := http.NewRequest(http.MethodPut, url, buf)
	if err != nil {
		return nil, errors.Wrap(err, RequestError.Error())
	}

	response, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, ResponseError.Error())
	}

	if response.StatusCode != http.StatusOK {
		return nil, errors.Wrap(err, strconv.Itoa(response.StatusCode))
	} else {
		buf = bytes.NewBuffer(nil)
		if _, err := buf.ReadFrom(response.Body); err != nil {
			return nil, errors.Wrap(err, BufferReadError.Error())
		}
		s := buf.String()

		buf = bytes.NewBuffer(nil)
		buf.WriteString(s)

		response := &Response{}
		if err := xml.NewDecoder(buf).Decode(&response); err != nil {
			return nil, errors.Wrap(err, XmlDecodeError.Error())
		}

		if response.Status != SuccessStatus {
			return nil, ApiNonSuccessError
		} else {
			return response, nil
		}
	}
}
