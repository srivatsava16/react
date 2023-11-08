package mini

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func databaseconn() {
	connect := "host=psql-mock-database-cloud.postgres.database.azure.com user=dxozbydqjounicdswzhjgemx@psql-mock-database-cloud password=hmbsrbicqemslkzewfjszgvd dbname=tasks1666187522241pjkcvkljmkrjnxrt port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(connect), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Print("success", db)
	//db.AutoMigrate(&users{})
	DB = db

}
