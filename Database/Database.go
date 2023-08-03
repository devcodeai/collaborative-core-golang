package Database

import (
	"fmt"

	"github.com/devcodeai/collaborative-core-golang/Models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GetDatabaseURL(
	user string,
	password string,
	host string,
	port string,
	db_name string,
) string {
	db_url := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		port,
		db_name,
	)
	return db_url
}

func ConnectDatabase(db_url string) {
	database, err := gorm.Open(mysql.Open(db_url))
	if err != nil {
		fmt.Println("Error: ", err)
		panic(err)
	}

	// database tables migration 
	database.AutoMigrate(&Models.Company{})
	database.AutoMigrate(&Models.Product{})
	database.AutoMigrate(&Models.Campus{})
	database.AutoMigrate(&Models.Major{})
	database.AutoMigrate(&Models.Talent{})
	database.AutoMigrate(&Models.Community{})

	DB = database

	println("Database Connected Successfully!")
}
