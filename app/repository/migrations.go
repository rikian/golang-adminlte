package repository

import (
	user_entity "golte/app/entity/user"
	"log"

	"gorm.io/gorm"
)

type table struct {
	tableName interface{}
}

func registerTable() []table {
	return []table{
		{tableName: user_entity.StatusUser{}},
		{tableName: user_entity.User{}},
	}
}

func RunMigration(db *gorm.DB) error {
	for _, table := range registerTable() {
		err := db.Migrator().CreateTable(table.tableName)

		if err != nil {
			log.Println("Migrations Failed...")
			log.Print(err)
			break
		}
	}

	result := db.Create(&user_entity.StatusUser{
		Id:     1,
		Status: "admin",
	})

	if result.Error != nil {
		log.Print(result.Error.Error())
	}

	log.Println("Migrations successfully ...")
	return nil
}
