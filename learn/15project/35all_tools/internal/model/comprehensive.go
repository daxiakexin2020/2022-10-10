package model

type IpInfo struct {
	Ip string `json:"ip" binding:required`
}

type DomainMapIp struct {
	Domain string `json:"domain" binding:required`
}

func NewIpInfo() *IpInfo {
	return &IpInfo{}
}

func NewDomainMapIp() *DomainMapIp {
	return &DomainMapIp{}
}

type ComprehensiveRepo interface {
	IpInfo(ip string) (interface{}, error)
	DomainMapIp(domain string) (interface{}, error)
}
