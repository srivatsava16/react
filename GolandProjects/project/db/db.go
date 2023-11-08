//The database connection is written in this file

package dao

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Databaseconn() *gorm.DB {
	connect := "host=psql-mock-database-cloud.postgres.database.azure.com user=dxozbydqjounicdswzhjgemx@psql-mock-database-cloud password=hmbsrbicqemslkzewfjszgvd dbname=tasks1666187522241pjkcvkljmkrjnxrt port=5432"
	db, err1 := gorm.Open(postgres.Open(connect), &gorm.Config{
		SkipDefaultTransaction: false,
		DryRun:                 false,
	})
	if err1 != nil {
		panic(err1)
	}
	fmt.Print("success", db)
	return db
}
