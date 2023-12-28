// main.go
package main

import (
	"dating-app/configs"
	"dating-app/routes"
)

func main() {
	configs.InitDB()

	e := routes.New()
	// Start server
	e.Start(":3000")
}
