package webserver

import (
	"net/http"
	"log"
	"html/template"
	"github.com/chabberwock/userip/app/geoip"
	"strings"
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
	TemplatePath string
}

func (server Server) Start() {
	http.HandleFunc("/", server.indexHandler)
	log.Fatal(http.ListenAndServe(server.Addr, nil))
}

func (server Server) indexHandler(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path == "/" {
		tpl := template.Must(template.ParseFiles(server.TemplatePath + "/index.html"))
		addr := strings.Split(request.RemoteAddr, ":")
		country, err := server.GeoIp.CountryByIp(addr[0])
		//country, err := server.GeoIp.CountryByIp("95.31.50.154")
		data := indexData{Ip: addr[0], Country: country, Providers: server.GeoIp.GetProviderData()}
		if err != nil {
			data.Error = err.Error()
		}
		tpl.Execute(writer, data)
	}
}

