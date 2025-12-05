package main

import "subscription/internal/app"

// @title           Subscription Rest API
// @version         1.0
// @description     This is a server which can add subscription and get summary cost of user subscriptions.

// @host      localhost:8080
// @BasePath  /api
func main() {
	app := app.NewApp()
	app.Init()
}