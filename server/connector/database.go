// package connector

// import (
// 	"database/sql"
// 	"log"
// 	"server/config"

// 	_ "github.com/go-sql-driver/mysql"
// )

// func ConnectDB(mysqlConfig config.MysqlConfig) *sql.DB {
// 	dataSource := mysqlConfig.UserName + ":" + mysqlConfig.Password +
// 		"@tcp(" + mysqlConfig.DbHost + ":" + mysqlConfig.DbPort + ")/" + mysqlConfig.DbName
// 	db, err := sql.Open(mysqlConfig.DbDriver, dataSource)
// 	if err != nil {
// 		log.Fatal(err)
// 		return nil
// 	}
// 	return db
// }

package connector

import (
	"log"
	"server/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func ConnectDB(mysqlConfig config.MysqlConfig) *gorm.DB {
	dataSource := mysqlConfig.UserName + ":" + mysqlConfig.Password +
		"@tcp(" + mysqlConfig.DbHost + ":" + mysqlConfig.DbPort + ")/" + mysqlConfig.DbName + "?parseTime=true"
	db, err := gorm.Open(mysqlConfig.DbDriver, dataSource)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return db
}
