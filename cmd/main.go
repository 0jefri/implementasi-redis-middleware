package main

import (
	"github.com/project-sistem-voucher/api/seeders"
	"github.com/project-sistem-voucher/config"
	_ "github.com/project-sistem-voucher/docs"
	routes "github.com/project-sistem-voucher/router"
)

// http://localhost:8080/swagger/index.html
func init() {
	config.InitiliazeConfig()
	config.InitDB()
	config.SyncDB()
	seeders.SeedVouchers(config.DB)
	seeders.SeedRedeem(config.DB)
}

// @title Example API
// @version 1.0
// @description This is a sample server for a Swagger API.
// @termsOfService http://example.com/terms/
// @contact.name API Support
// @contact.url https://academy.lumoshive.com/contact-us
// @contact.email lumoshive.academy@gmail.com
// @license.name Lumoshive Academy
// @license.url https://academy.lumoshive.com
// @host localhost:8080
// @schemes http
// @BasePath /
// @securityDefinitions.apikey BasixAuth
// @in header
// @name Authorization

func main() {
	routes.Server().Run()
}
