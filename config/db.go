package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
	"gopkg.in/gorp.v1"
)

var DbConnect *gorp.DbMap

// InitDb initialize la database
func InitDb() {
	db, err := sql.Open("sqlite3", ConfigLoaded.Datasource)

	if err != nil {
		logrus.Error("Failed to connect to database:") //NOSONAR
		logrus.Panic(err)                              //NOSONAR
	}
	if err := db.Ping(); err != nil {
		logrus.Error("Failed to ping database")
		logrus.Panic(err)
	}
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}            //NOSONAR
	dbmap.TraceOn("[gorp]", log.New(os.Stdout, "myapp:", log.Lmicroseconds)) //NOSONAR
	logrus.Debug("Connected to db")                                          //NOSONAR
	DbConnect = dbmap                                                        //NOSONAR
}
