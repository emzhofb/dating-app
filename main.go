// main.go
package main

import (
	"dating-app/configs"
	"dating-app/controllers"
	"dating-app/routes"
	"dating-app/seeds"
	"fmt"

	cron "gopkg.in/robfig/cron.v2"
)

func main() {
	configs.InitDB()
	seeds.CreateSeedUser()

	c := cron.New()
	// Schedule the job to run every day at a midnight
	_, err := c.AddFunc("0 0 * * *", func() {
		// Your logic to reset the user's limit goes here
		controllers.ResetUserLimits()
	})

	if err != nil {
		fmt.Println("Error scheduling cron job:", err)
	}

	c.Start()

	e := routes.New()
	// Start server
	e.Start(":3000")
}
