package namecheap

import (
	"fmt"
	"net/url"
	"strconv"
)

type DomainDNSGetHostsResult struct {
	Domain        string          `xml:"Domain,attr"`
	IsUsingOurDNS bool            `xml:"IsUsingOurDNS,attr"`
	Hosts         []DomainDNSHost `xml:"Host"`
}

type DomainDNSHost struct {
	ID      int    `xml:"HostId,attr"`
	Name    string `xml:"Name,attr"`
	Type    string `xml:"Type,attr"`
	Address string `xml:"Address,attr"`
	MXPref  int    `xml:"MXPref,attr"`
	TTL     int    `xml:"TTL,attr"`
}

type DomainDNSSetHostsResult struct {
	Domain    string `xml:"Domain,attr"`
	IsSuccess bool   `xml:"IsSuccess,attr"`
}

func (client *NamecheapClient) DomainsDNSGetHosts(sld string, tld string) (DomainDNSGetHostsResult, error) {
	resp := ApiResponse{}
	requestInfo := ApiRequest{
		command: "namecheap.domains.dns.getHosts",
		params:  url.Values{},
	}
	requestInfo.params.Set("SLD", sld)
	requestInfo.params.Set("TLD", tld)

	if err := client.get(requestInfo, &resp); err != nil {
		return DomainDNSGetHostsResult{}, err
	}
	return resp.DomainDNSHosts, nil
}

func (client *NamecheapClient) DomainDNSSetHosts(sld string, tld string, hosts []DomainDNSHost) (DomainDNSSetHostsResult, error) {
	resp := ApiResponse{}
	requestInfo := ApiRequest{
		command: "namecheap.domains.dns.setHosts",
		params:  url.Values{},
	}
	requestInfo.params.Set("SLD", sld)
	requestInfo.params.Set("TLD", tld)

	for i := range hosts {
		requestInfo.params.Set(fmt.Sprintf("HostName%v", i+1), hosts[i].Name)
		requestInfo.params.Set(fmt.Sprintf("RecordType%v", i+1), hosts[i].Type)
		requestInfo.params.Set(fmt.Sprintf("Address%v", i+1), hosts[i].Address)
		requestInfo.params.Set(fmt.Sprintf("TTL%v", i+1), strconv.Itoa(hosts[i].TTL))

	}

	if err := client.get(requestInfo, &resp); err != nil {
		return DomainDNSSetHostsResult{}, err
	}
	return resp.DomainDNSSetHosts, nil
}
