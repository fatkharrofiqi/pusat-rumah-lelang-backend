package main

import (
	"pusat-rumah-lelang-backend/config"
	"pusat-rumah-lelang-backend/migrations"
	"pusat-rumah-lelang-backend/seeder"
)

func main() {
	db := config.OpenDB()

	migrations.RunMigrations(db)

	seeder.Seeds(db)
}
