package main

import ctrl "eulabs/internal/controller"

func main() {
	server := ctrl.RouterFactory()

	server.Logger.Fatal(server.Start(":1323"))
}
