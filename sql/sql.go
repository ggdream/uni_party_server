package sql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

const (
	dsn = "root:hpxhJvhnrIAbX6AWd9DVLhcaZAoaGD8m@tcp(101.34.2.166:3306)/test?charset=utf8mb4&parseTime=True"
)

var db *gorm.DB
//db.Set("gorm:table_options", "AUTO_INCREMENT=100")

// Init 初始化
func Init() (err error) {
	dialector := mysql.Open(dsn)
	config := gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		//Logger: logger.Default.LogMode(logger.Info),
	}
	db, err = gorm.Open(dialector, &config)
	if err != nil {
		return
	}

	return db.Exec(fmt.Sprintf("ALTER TABLE %s AUTO_INCREMENT=1000000", UserInfoTable{}.TableName())).Error
}















