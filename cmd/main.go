package main

import "github.com/kevinharv/idp/internal/server"

/*
	- read in configuration
		- web
		- ad
		- idp
	- create OAuth and OIDC endpoints
	- create SAML endpoints
	- register middleware
	- 
*/

func main() {
	server.Serve()
}
