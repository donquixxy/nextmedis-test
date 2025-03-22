package driver

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

type DatabaseOpt struct {
	Name     string
	Address  string
	Port     int
	Username string
	Password string
}

func NewMysqlInstance(opt DatabaseOpt) (*gorm.DB, error) {
	addr := fmt.Sprintf("%s:%d", opt.Address, opt.Port)
	dsn := fmt.Sprintf("%s:%v@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		opt.Username,
		opt.Password,
		addr,
		opt.Name,
	)

	cfg := &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Info),
		DisableForeignKeyConstraintWhenMigrating: true,
	}

	db, err := gorm.Open(mysql.Open(dsn), cfg)

	if err != nil {
		log.Fatalf("[NewMysqlInstance] - Failed on connect database %v", err)
		return nil, err
	}

	return db, nil
}
