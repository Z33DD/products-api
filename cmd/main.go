package main

import (
	ctrl "eulabs/internal/controller"
	"eulabs/internal/service"
)

func main() {
	service.GetDatabase()
	server := ctrl.RouterFactory()

	server.Logger.Fatal(server.Start(":1323"))
}
