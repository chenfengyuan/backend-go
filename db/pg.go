package db

import (
	"github.com/graphql-stack/backend-go/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"os"
)

var ORM *gorm.DB

func InitDB(fns ...func(db *gorm.DB)) {
	db, err := gorm.Open("postgres", os.Getenv("PG_URL"))
	if err != nil {
		log.Fatal(err)
	}

	for _, fn := range fns {
		fn(db)
	}

	ORM = db.Debug()
}

func init() {
	InitDB(func(db *gorm.DB) {
		db.AutoMigrate(new(model.User))
		db.AutoMigrate(new(model.Token))
		db.AutoMigrate(new(model.Post))
		db.AutoMigrate(new(model.Comment))
	})
}
