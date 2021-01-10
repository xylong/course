package lib

import (
	"course/src/domain/model/course"
	"course/src/domain/model/price"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var (
	DB *gorm.DB
)

func init() {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=True&loc=Local",
		"root", "root", "course", "utf8mb4"))
	if err != nil {
		log.Fatal(err)
	}

	db.DB().SetMaxIdleConns(5)
	db.DB().SetMaxOpenConns(10)

	db.AutoMigrate(
		&course.Course{},
		&price.CoursePrice{},
	)

	DB = db
}
