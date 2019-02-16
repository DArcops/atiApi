package main

import (
	"github.com/darcops/atiApi/models"
	"github.com/darcops/atiApi/routes"
)

func main() {
	models.Connect()
	models.Migrate()
	routes.Init()
}
