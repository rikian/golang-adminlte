package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func ConnectDB() (string, bool) {
	dsn := "host=" + Host + " user=" + Username + " password=" + Password + " dbname=" + Database + " port=" + Port + " sslmode=disable TimeZone=Asia/Jakarta"

	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "public.",
		},
	})

	if err != nil {
		return err.Error(), true
	}

	// for _, model := range models.RegisterModels() {
	// 	err = DB.Debug().AutoMigrate(model.Model)

	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }

	// for i := 0; i < 1000; i++ {
	// 	seeders.DBSeed(DB)
	// }

	// log.Println("Migrations successfully ...")

	return "Postgres server listening on port " + Port + "...\n", false
}
