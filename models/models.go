package models

import (
	"github.com/exlskills/demo-go-webservice/config"
	// Our lightweight ORM
	"github.com/jinzhu/gorm"
	// Use the _ import syntax to ensure that the mysql init()
	// gets run and so that Go doesn't complain about an unused import
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Log = config.Cfg().GetLogger()
var db *gorm.DB

func Setup() (err error) {
	db, err = gorm.Open("mysql", config.Cfg().DBPath)
	if err != nil {
		return err
	}

	// Some reasonable pool settings
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	// Log all our queries
	// NB: Watch out for this in production. Are you inserting private user data?
	//     That would end up in your logs. Consider your privacy mandate, adjust accordingly.
	db.LogMode(true)

	// Usually simplest to stick to singular table names
	// (less potential for ambiguity and works neatly with struct/model names)
	db.SingularTable(true)

	return nil
}
