package local

import (
	"35all_tools/internal/model"
	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
	"net"
	"os"
	"path/filepath"
)

type ComprehensiveRepository struct{}

var _ (model.ComprehensiveRepo) = (*ComprehensiveRepository)(nil)

func NewComprehensiveRepository() model.ComprehensiveRepo {
	return &ComprehensiveRepository{}
}

func (cr *ComprehensiveRepository) IpInfo(ip string) (interface{}, error) {

	currDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	dbPathDir := filepath.Join(currDir, "../")
	dbPath := dbPathDir + "/util_files/ip2region.xdb"
	searcher, err := xdb.NewWithFileOnly(dbPath)
	if err != nil {
		return nil, err
	}

	defer searcher.Close()
	region, err := searcher.SearchByStr(ip)
	if err != nil {
		return nil, err
	}
	return region, nil
}

func (cr *ComprehensiveRepository) DomainMapIp(domain string) (interface{}, error) {

	currDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	dbPathDir := filepath.Join(currDir, "../")
	dbPath := dbPathDir + "/util_files/ip2region.xdb"
	searcher, err := xdb.NewWithFileOnly(dbPath)
	if err != nil {
		return nil, err
	}

	defer searcher.Close()
	ips, err := net.LookupIP(domain)
	if err != nil {
		return nil, err
	}
	ipres := ips[0].String()
	return ipres, nil
}
