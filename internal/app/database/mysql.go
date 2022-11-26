package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewConnDatabase() *gorm.DB {
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	// os.Getenv("MYSQL_USER"),
	// os.Getenv("MYSQL_PASSWORD"),
	// os.Getenv("MYSQL_HOST"),
	// os.Getenv("MYSQL_PORT"),
	// os.Getenv("MYSQL_DATABASE"),
	dsn := "sql6580597:JiKgAQK4QF@tcp(sql6.freemysqlhosting.net:3306)/sql6580597?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
