package models

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"database/sql"
	"database/sql/driver"
	// sql engine: mysql、pq、sqlite
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	"github.com/golang/glog"
)

var sqlFile = []string{"mysql.sql", "systemSetting.sql"}

// DbCfg defines dabase config
var (
	x     *sql.DB
	DbCfg struct {
		Type, Host, Port, Name, User, Passwd, Prefix, SSLMode string
	}
)

// GeneratorconnStr return sql connection string
func GeneratorconnStr() string {
	connStr := ""
	// Checkout SQL Type
	switch DbCfg.Type {
	case "postgres":
		connStr = fmt.Sprintf("postgres://%s:%s@%s/%s%ssslmode=%s", DbCfg.User, DbCfg.Passwd, DbCfg.Host, DbCfg.Name, "", DbCfg.SSLMode)
	case "mysql":
		connStr = fmt.Sprintf("%s:%s@tcp(%s)/%s%scharset=utf8&parseTime=true", DbCfg.User, DbCfg.Passwd, DbCfg.Host, DbCfg.Name, "")
	case "sqlite3":
		pwd, err := os.Getwd()
		if err != nil {
			glog.Errorf("Getwd error which is %s", err)
		}
		connStr = "file: " + pwd + "?cache=shared&mode=rwc"
	default:
		glog.Errorf("Unknown database type: %s", DbCfg.Type)
	}
	return connStr
}

// CreateSQLConn create sql connection
func CreateSQLConn(sqlType string, conStr string) (*sql.DB, error) {
	if x == nil {
		db, err := sql.Open(sqlType, conStr)
		if err != nil {
			glog.Fatalf("open %s found err: %v", sqlType, err)
			return nil, err
		}
		err = db.Ping()
		if err != nil {
			glog.Fatalf("db.Ping found err:%v", err)
		}
		x = db

	}
	return x, nil
}

//
// EnSureTableExist ensure table exist
func EnSureTableExist(sqlType string, conStr string) error {
	db, err := CreateSQLConn(sqlType, conStr)
	if err != nil {
		glog.Fatal(err)
	}
	defer db.Close()
	if db.Ping() == driver.ErrBadConn {
		return fmt.Errorf("mysql: could not connect to the database. " +
			"could be bad address, or this address is not whitelisted for access.")
	}

	if _, err = db.Exec("USE report"); err != nil {
		fmt.Print("mysql: creating database report")
		return createTable(db)
	}
	return err
}

// createTable create table
func createTable(connection *sql.DB) error {

	for _, sf := range sqlFile {
		table, err := os.Open(sf)
		if err != nil {
			glog.Fatal(err)
		}
		defer table.Close()
		fd, _ := ioutil.ReadAll(table)
		createTableStatements := strings.Split(string(fd), ";")
		for _, stmt := range createTableStatements {
			if _, err := connection.Exec(stmt); err != nil {
				return err
			}
		}
	}
	return nil
}

// GetSQLEngine return sql.DB handler
func GetSQLEngine() (*sql.DB, error) {
	connStr := GeneratorconnStr()
	return CreateSQLConn(DbCfg.Type, connStr)
}
