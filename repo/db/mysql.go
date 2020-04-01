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

// InitDB initialization mysql db
func InitDB() {
	var err error
	db, err = gorm.Open(driver, fmt.Sprintf("%s:%s@(%s:%d)/%s%s", username, password, host, port, dbName, extras))
	if err != nil {
		log.Fatalf("数据库连接失败：%v", err)
		return
	}
	if err = db.DB().Ping(); err != nil {
		log.Fatalf("数据库连接失败：%v", err)
		return
	}

	db.LogMode(true)
	db.SingularTable(true)

	db.AutoMigrate(
		model.Attachment{},
		model.Category{},
		model.CommentBlackList{},
		model.Comment{},
		model.FlywaySchemaHistory{},
		model.Journal{},
		model.Link{},
		model.Log{},
		model.Menu{},
		model.Meta{},
		model.Option{},
		model.Photo{},
		model.PostCategory{},
		model.PostTag{},
		model.Post{},
		model.Tag{},
		model.ThemeSetting{},
		model.User{},
	)
}

// GetMDB get global gorm.DB
func GetMDB() *gorm.DB {
	return db
}
