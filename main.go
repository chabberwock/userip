package main

import (
	"github.com/chabberwock/userip/app/webserver"
	"github.com/chabberwock/userip/app/geoip"
	"github.com/chabberwock/userip/app/config"
	"github.com/chabberwock/userip/app/providers/ipstack"
	"github.com/chabberwock/userip/app/providers/nekudo"
)

func main() {
	configData := config.Read("config.json")
	geoipService := geoip.Create(configData.CacheTTL)
	//geoipService.RegisterProvider(mock.Create(1))
	if configData.IpStack.Enabled {
		geoipService.RegisterProvider(ipstack.Create(configData.IpStack.RequestLimit, configData.IpStack.AccessKey))
	}
	if configData.Nekudo.Enabled {
		geoipService.RegisterProvider(nekudo.Create(configData.Nekudo.RequestLimit))
	}


	server := webserver.Server{Addr: configData.Bind, GeoIp: geoipService, TemplatePath:configData.TemplatePath}
	server.Start()
}





