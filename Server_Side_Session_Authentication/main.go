package main

import (
	"sessionauth/app/model"
	"sessionauth/app/routes"
)

func main() {
	model.Setup()
	routes.Setup()
}
