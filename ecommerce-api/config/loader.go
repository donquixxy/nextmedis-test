package config

import (
	"ecommerce-api/driver"
	"gorm.io/gorm"
	"log"
)

func GetDatabaseInstance(cfg *Config) (*gorm.DB, error) {

	db, err := driver.NewMysqlInstance(driver.DatabaseOpt{
		Name:     cfg.DbCfg.Name,
		Address:  cfg.DbCfg.Host,
		Port:     cfg.DbCfg.Port,
		Username: cfg.DbCfg.Username,
		Password: cfg.DbCfg.Password,
	})

	log.Print("Connected to database")
	return db, err
}
