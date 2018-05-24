package nekudo

import "github.com/chabberwock/userip/app/geoip"

type Provider struct {
	hitCounter geoip.Counter
	RequestLimit int
}

func Create(requestLimit int) *Provider {
	provider := new(Provider)
	provider.RequestLimit = requestLimit
	return provider
}

// b2a856f222be417cb767f831cbca347d
func (provider *Provider) CountryByIp(ip string) string {
	provider.hitCounter.RecordHit()
	resp := new(Response)
	getJson("http://geoip.nekudo.com/api/" + ip, resp)
	return resp.Country.Name
}

func (provider *Provider) Hits() int {
	return provider.hitCounter.Count()
}

func (provider *Provider) Name() string {
	return "Nekudo"
}

func (provider *Provider) IsAvailable() bool {
	provider.hitCounter.Cleanup()
	return provider.hitCounter.Count() < provider.RequestLimit
}

