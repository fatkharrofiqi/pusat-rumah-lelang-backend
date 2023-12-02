package main

import (
	"log"
	"pusat-rumah-lelang-backend/internal"
	"pusat-rumah-lelang-backend/internal/seeds"
	"pusat-rumah-lelang-backend/migrations"
)

func main() {
	db, err := internal.InitializeDatabase()
	if err != nil {
		log.Fatalln(err)
	}

	migrations.RunMigrations(db)

	seeds.SeedData(db)
}
