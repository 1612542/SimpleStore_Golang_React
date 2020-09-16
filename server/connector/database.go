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

// package connector

// import (
// 	"log"

// 	"github.com/jinzhu/gorm"

// 	_ "github.com/jinzhu/gorm/dialects/mysql"
// 	_ "github.com/lib/pq"
// )

// func ConnectDB() *gorm.DB {
// 	// dataSource := mysqlConfig.UserName + ":" + mysqlConfig.Password +
// 	// 	"@tcp(" + mysqlConfig.DbHost + ":" + mysqlConfig.DbPort + ")/" + mysqlConfig.DbName + "?parseTime=true"
// 	// db, err := gorm.Open("postgres", "postgres://postgres:postgres@localhost.com:5432/simple_store")
// 	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres password=postgres dbname=simple_store")
// 	if err != nil {
// 		log.Fatal(err)
// 		return nil
// 	}

// 	return db
// }
