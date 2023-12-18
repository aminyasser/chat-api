package mysql


import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
  )

var db *gorm.DB


func GetDB() (*gorm.DB , error){
	var err error
	if db == nil {
		err = initDB()
	}
	return db , err
}
 
func initDB() error {
	
	dsn := "user:password@tcp(db:3306)/chat_api?charset=utf8&parseTime=True"

	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	return err
}