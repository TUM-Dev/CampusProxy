package main

import "github.com/TUM-Dev/CampusProxy/proxy"

//go:generate swag init --parseDependency proxy/proxy.go -o docs

// @title CAMPUSOnline Webservice proxy
// @version 1.0
// @description This is the proxy server for CAMPUSOnline Webservices, enabling a nicely documented and uniform rest access to CAMPUSOnline resources.

// @contact.name   Joscha Henningsen
// @contact.url    https://joschas.page

// @securityDefinitions.apikey   ApiKeyAuth
// @in                           header
// @name                         x-api-key
// @description					 A valid token to authenticate with CampusOnline.

// @BasePath /api/v1
func main() {
	p := proxy.Proxy{}
	p.Start()
}
