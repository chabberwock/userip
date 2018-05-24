package ipstack

import (
	"github.com/chabberwock/userip/app/geoip"
)

type Provider struct {
	hitCounter geoip.Counter
	RequestLimit int
	accessKey string
}

func Create(requestLimit int, accessKey string) *Provider {
	provider := new(Provider)
	provider.RequestLimit = requestLimit
	provider.accessKey = accessKey
	return provider
}

// b2a856f222be417cb767f831cbca347d
func (provider *Provider) CountryByIp(ip string) string {
	provider.hitCounter.RecordHit()
	resp := new(Response)
	getJson("http://api.ipstack.com/"+ip+"?access_key="+provider.accessKey, resp)
	return resp.Country_name
}

func (provider *Provider) Hits() int {
	return provider.hitCounter.Count()
}

func (provider *Provider) Name() string {
	return "Ipstack"
}

func (provider *Provider) IsAvailable() bool {
	provider.hitCounter.Cleanup()
	return provider.hitCounter.Count() < provider.RequestLimit
}

