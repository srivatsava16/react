//The database connection is written in this file

package dao

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Databaseconn() {
	connect := "host=psql-mock-database-cloud.postgres.database.azure.com user=dxozbydqjounicdswzhjgemx@psql-mock-database-cloud password=hmbsrbicqemslkzewfjszgvd dbname=tasks1666187522241pjkcvkljmkrjnxrt port=5432"
	db, err1 := gorm.Open(postgres.Open(connect), &gorm.Config{})
	if err1 != nil {
		panic(err1)
	}
	fmt.Print("success", db)
	DB = db
}
