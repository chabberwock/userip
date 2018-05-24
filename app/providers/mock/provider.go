package mock

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

func (provider *Provider) CountryByIp(ip string) string {
	provider.hitCounter.RecordHit()
	return "Narnia"
}

func (provider *Provider) Hits() int {
	return provider.hitCounter.Count()
}

func (provider *Provider) Name() string {
	return "Mock Provider"
}

func (provider *Provider) IsAvailable() bool {
	provider.hitCounter.Cleanup()
	return provider.hitCounter.Count() < provider.RequestLimit
}



