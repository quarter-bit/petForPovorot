package main

import (
	// _ "proxy/docs"

	_ "Pet_store/docs"
	"Pet_store/internal/controller/app"
	"Pet_store/internal/controller/routs"
	"Pet_store/internal/controller/server"
)

// @title Pet Store
// @version 1.0
// @description Pet Store from KataAcademy
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// for {
	// }
	app := app.Inject()
	r := routs.InitAllRouts(app)
	server := server.NewServer(":8080", r)
	go server.Serve()
	server.Shutdown()

}
