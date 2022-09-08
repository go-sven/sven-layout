package data

import (
	"fmt"
	"github/go-sven/sven-layout/app/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)


func NewDb(conf *conf.MysqlConfig) (*gorm.DB,error)  {
	sqlConn :=fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?%v",conf.Username,conf.Password,conf.Host,conf.Port,conf.Database,conf.Parameter)
	//fmt.Println("sqlConn:",sqlConn)
	db, err := gorm.Open(mysql.Open(sqlConn), &gorm.Config{
		//Logger:newLogger,
		//禁用表名复数
		NamingStrategy:schema.NamingStrategy{
			SingularTable: true,
		},
		//创建并缓存预编译语句
		PrepareStmt:true,
		//跳过默认事务
		SkipDefaultTransaction:true,
	})
	if err != nil {
		return nil,  err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil,err
	}
	sqlDB.SetMaxIdleConns(conf.MaxOpen)
	sqlDB.SetMaxOpenConns(conf.MaxIdle)
	//sqlDB.SetConnMaxLifetime(time.Second * 25)
	return db,err

}