package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"sync"
	"wscmakebygo.com/global/envConfig"
	"wscmakebygo.com/tools/logUtil"
)

var (
	db   *gorm.DB
	once sync.Once
)

func GetDatabase() *gorm.DB {
	if db == nil {
		panic("database not initialized")
	}
	return db
}

func CloseDatabase() {
	sqldb, err := db.DB()
	if err != nil {
		panic(err)
	}
	err = sqldb.Close()
	if err != nil {
		logUtil.Log.Fatal(err.Error())
		return
	}
}

func crateDbAddr() string {
	Addr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&loc=%s&parseTime=true",
		envConfig.GetConfig().Db.User,
		envConfig.GetConfig().Db.Password,
		envConfig.GetConfig().Db.Host,
		envConfig.GetConfig().Db.Port,
		envConfig.GetConfig().Db.DbName,
		envConfig.GetConfig().Db.Charset,
		envConfig.GetConfig().Db.Loc,
	)
	return Addr
}

func InitVal() {
	once.Do(func() {
		logStr := fmt.Sprintf("%s:%d", envConfig.GetConfig().Db.Host, envConfig.GetConfig().Db.Port)
		addr := crateDbAddr()
		init, err := gorm.Open(mysql.Open(addr), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		logUtil.Log.Println("created Db connection:" + logStr)
		init.Logger = NewLogger(
			logUtil.DBLog,
			logger.Info)
		db = init
	})
}
