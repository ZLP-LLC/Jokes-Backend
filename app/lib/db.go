package lib

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"jokes/models"
)

type Database struct {
	*gorm.DB
}

func NewDatabase(env Env, logger Logger) Database {
	dsn := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		env.DBUser,
		env.DBPassword,
		env.DBName,
		env.DBHost,
		env.DBPort,
		env.DBSSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.GetGormLogger(),
	})

	if err != nil {
		logger.Info("Connect to sqlite...")
		dsn = "test.db"
		db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{
			Logger: logger.GetGormLogger(),
		})
		if err != nil {
			logger.Panic("Can't open DB: ", err.Error())
		}
	}
	logger.Info("Connected to database")

	if err := db.AutoMigrate(&models.User{}, &models.Joke{}, &models.Rating{}, &models.Annotation{}); err != nil {
		logger.Panic("Can't migrate database: ", err.Error())
	}
	logger.Info("Migrated database")

	// TEST: test data
	//err = db.Create(&models.Joke{Text: "test"}).Error
	//if err != nil {
	//	logger.Errorln(err.Error())
	//}
	//err = db.Create(&models.Annotation{Text: "test", JokeID: 1, Approved: true}).Error
	//if err != nil {
	//	logger.Errorln(err.Error())
	//}

	return Database{
		db,
	}
}
