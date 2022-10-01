package main

import "GoHTML2PDF/pkg/endpoint"

func main() {
	endpoint.RegisterRoutes()
	endpoint.Listen()

}
