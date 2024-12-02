package main

import (
	"github.com/project-sistem-voucher/api/seeders"
	"github.com/project-sistem-voucher/config"
	routes "github.com/project-sistem-voucher/router"
)

func init() {
	config.InitiliazeConfig()
	config.InitDB()
	config.SyncDB()
	seeders.SeedVouchers(config.DB)
	seeders.SeedRedeem(config.DB)
}

func main() {
	routes.Server().Run()
}
