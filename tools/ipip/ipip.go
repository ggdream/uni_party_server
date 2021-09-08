package ipip

import (
	"github.com/ipipdotnet/ipdb-go"
)

type IpIp struct {
	db   *ipdb.City
	lang string
}

func New() (*IpIp, error) {
	db, err := ipdb.NewCity("./free.ipdb")
	if err != nil {
		return nil, err
	}

	return &IpIp{
		db:   db,
		lang: "CN",
	}, nil
}

// Query 查询省份和城市名称
func (i *IpIp) Query(ip string) (string, string, error) {
	info, err := i.query(ip)
	if err != nil {
		return "", "", err
	}

	return info.RegionName, info.CityName, nil
}

func (i *IpIp) query(ip string) (*ipdb.CityInfo, error) {
	return i.db.FindInfo(ip, i.lang)
}
