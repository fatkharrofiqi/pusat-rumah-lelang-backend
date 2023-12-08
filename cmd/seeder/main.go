package main

import (
	"pusat-rumah-lelang-backend/config"
	"pusat-rumah-lelang-backend/migrations"
	"pusat-rumah-lelang-backend/seeder"
)

func main() {
	config.LoadEnv()
	config.LoadConstant()
	db := config.OpenDB()

	migrations.RunMigrations(db)

	seeder.Seeds(db)
}
