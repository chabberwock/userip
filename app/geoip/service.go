package geoip

import (
	"errors"
	"github.com/chabberwock/userip/app/cache"
)

type geoProvider interface {
	CountryByIp(ip string) string
	Hits() int
	Name() string
	IsAvailable() bool
}

type ProviderData struct {
	Name string
	Hits int
}

type Service struct {
	providers []geoProvider
	cache *cache.Service
}

func Create() *Service {
	item := new(Service)
	item.cache = cache.Create(5)
	return item
}

func (service *Service) CountryByIp(ip string) (string, error) {
	// выбор наименее используемого провайдера
	//var leastUsed int = 0
	//for i := range service.providers {
	//	if service.providers[i].Hits() < service.providers[leastUsed].Hits() {
	//		leastUsed = i
	//	}
	//}
	//return service.providers[leastUsed].CountryByIp(ip)

	country, found := service.cache.Get(ip)
	if found {
		return country, nil
	}
	// просто выбираем первого доступного
	for i := range service.providers {
		if service.providers[i].IsAvailable() {
			country = service.providers[i].CountryByIp(ip)
			service.cache.Set(ip, country)
			return country, nil
		}
	}
	return "", errors.New("No providers available now")
}

func (service *Service) RegisterProvider(provider geoProvider) {
	service.providers = append(service.providers, provider)
}

func (service *Service) GetProviderData() []ProviderData{
	result := []ProviderData{}
	for i := range service.providers {
		result = append(result, ProviderData{Name: service.providers[i].Name(), Hits: service.providers[i].Hits()})
	}
	return result

}

