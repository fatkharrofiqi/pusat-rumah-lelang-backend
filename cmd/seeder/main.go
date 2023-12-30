package main

import (
	"pusat-rumah-lelang-backend/db/seeder"
	"pusat-rumah-lelang-backend/internal/config"
	"pusat-rumah-lelang-backend/internal/migrations"
)

func main() {
	viperConfig := config.NewViper()
	log := config.NewLogger(viperConfig)
	db := config.NewDatabase(viperConfig, log)

	migrations.DropTable(db)
	migrations.RunMigrations(db)

	seeder.Seeds(db)
}
