# GeoIP Cheking tool

## About

This application uses open source geoip APIs to detect user's country. It supports multiple services, result caching and load balancing

## Configuration

Configuration is done using `config.json` file that should be located in the same directory as executable. Please note `TemplatePath`
parameter, that should point to a valid directory containing website templates. by default they are located in 
`app/webserver` directory. By default app is running on port `8080`

## Load balancing

Application uses simple load balancer that basically stores all requests i performed on provider along with their dates
Once allowed limit is exceeded, requests are forwarded to next provider. You can change limits in `RequestLimit` parameter

## Overriding IP

For convinient manual testing you can specify GET parameter `ip` in request url to override default IP detection
for example `http://localhost:8080/?ip=95.31.150.150`



