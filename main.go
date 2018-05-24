package main

import (
	"github.com/chabberwock/userip/app/webserver"
	"github.com/chabberwock/userip/app/geoip"
	"github.com/chabberwock/userip/app/providers/mock"
	"github.com/chabberwock/userip/app/providers/ipstack"
	"github.com/chabberwock/userip/app/providers/nekudo"
)

func main() {
	geoipService := geoip.Create()
	geoipService.RegisterProvider(mock.Create(5))
	geoipService.RegisterProvider(ipstack.Create(5, "b2a856f222be417cb767f831cbca347d"))
	geoipService.RegisterProvider(nekudo.Create(5))

	server := webserver.Server{Addr: ":8080", GeoIp: geoipService}
	server.Start()
}



