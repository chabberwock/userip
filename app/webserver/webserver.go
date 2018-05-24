package webserver

import (
	"net/http"
	"log"
	"html/template"
	"github.com/chabberwock/userip/app/geoip"
)

type indexData struct {
	Ip string
	Country string
	Error string
	Providers []geoip.ProviderData

}

type Server struct {
	GeoIp *geoip.Service
	Addr string
}

func (server Server) Start() {
	http.HandleFunc("/", server.indexHandler)
	log.Fatal(http.ListenAndServe(server.Addr, nil))
}

func (server Server) indexHandler(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path == "/" {
		tpl := template.Must(template.ParseFiles("app/webserver/index.html"))
		country, err := server.GeoIp.CountryByIp(request.RemoteAddr)
		//country, err := server.GeoIp.CountryByIp("95.31.50.154")
		data := indexData{Ip: request.RemoteAddr, Country: country, Providers: server.GeoIp.GetProviderData()}
		if err != nil {
			data.Error = err.Error()
		}
		tpl.Execute(writer, data)
	}
}

