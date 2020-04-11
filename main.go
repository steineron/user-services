package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/steineron/users-service/morestrings"
	"log"
)

const (
	dbDns      = "root:Ilts0nitm@tcp(:3306)"
	tableName  = "domain_users"
	driverName = "mysql"
)

func main() {

	fmt.Println(morestrings.ReverseRunes("!oG, olleH"))
	fmt.Println(morestrings.ReverseRunes("Hello Go"))
	fmt.Println(morestrings.Runes("JUST THIS"))

	fmt.Println(morestrings.Runes("JUST THIS"))
	// MySql setup using https://github.com/datacharmer/dbdeployer
	db, err := sql.Open(driverName, dbDns+"/"+tableName)

	defer db.Close()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("DB Opened")
		// log.Printf("DB name:%q", db)
	}

	_, err = db.Exec(
		"CREATE TABLE IF NOT EXISTS " + tableName + ".hello(world varchar(50))")
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Table Created!")
	}
}
