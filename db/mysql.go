package db

import (
	"blog/model"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db       *gorm.DB
	driver   = "mysql"
	host     = "localhost"
	port     = 3306
	dbName   = "blog"
	username = "root"
	password = "root"
	extras   = "?charset=UTF8&parseTime=true"
)

func InitDB() {
	db = connectMDB()
	db.LogMode(true)
	db.SingularTable(true)

	db.AutoMigrate(
		model.Attachment{},
		model.Category{},
		model.CommentBlackList{},
		model.FlywaySchemaHistory{},
		model.Journal{},
		model.Link{},
		model.Log{},
		model.Menu{},
		model.Meta{},
		model.Option{},
		model.Post{},
		model.PostCategory{},
		model.PostTag{},
		model.Tag{},
		model.ThemeSetting{},
		model.User{},
	)
}

func connectMDB() *gorm.DB {
	db, err := gorm.Open(driver, fmt.Sprintf("%s:%s@(%s:%d)/%s%s", username, password, host, port, dbName, extras))
	if err != nil {
		log.Fatalf("数据库连接失败：%v", err)
		return nil
	}
	return db
}

// GetDB获取全局 DB
func GetMDB() *gorm.DB {
	return db
}
