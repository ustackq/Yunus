package models

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"path"
	"strings"

	"database/sql"

	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/spf13/pflag"
	// sql engine: mysql、pq、sqlite、mssqldb
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	"github.com/ustack/Yunus/src/app/backend/pkg/setting"
)

// Engine define XORM engine or session
type Engine interface {
	Delete(interface{}) (int64, error)
	Exec(string, ...interface{}) (sql.Result, error)
	Find(interface{}, ...interface{}) error
	Get(interface{}) (bool, error)
	Id(interface{}) *xorm.Session
	Insert(...interface{}) (int64, error)
	InsertOne(interface{}) (int64, error)
	Iterate(interface{}, xorm.IterFunc) error
	Sql(string, ...interface{}) *xorm.Session
	Table(interface{}) *xorm.Session
	Where(interface{}, ...interface{}) *xorm.Session
}

// DbCfg defines dabase config
var (
	engine       *xorm.Engine
	tables       []interface{}
	HasEngine    bool
	ItemsPerPage = 40
	DbCfg        struct {
		Type, Host, Port, Name, User, Passwd, Prefix, SSLMode, Path string
	}

	EnableSQLite3 bool
)

// copy from gogits
func init() {
	tables = append(tables, new(User))

	gonicNames := []string{"SSL"}
	for _, name := range gonicNames {
		core.LintGonicMapper[name] = true
	}
}

// GetConfigs get config from file or post formatted
// 1、get config from flag
// 2、get config from post form
func GetConfigs() {
	setting.Cfg.BindPFlags(pflag.CommandLine)
	path := setting.Cfg.GetString("cfgPath")
	if path != "" {
		setting.Cfg.AddConfigPath(path)
		err := setting.Cfg.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("Fatal error config file: %s ", err))
		}
	}

}

// SetConfigs return load configs from init file or postform
func SetConfigs() {
	setting.Cfg.Set("Verbose", true)
	DbCfg.Type = setting.Cfg.GetString("DB_TYPE")
	DbCfg.Host = setting.Cfg.GetString("DB_HOST")
	DbCfg.User = setting.Cfg.GetString("DB_USER")
	DbCfg.Passwd = setting.Cfg.GetString("DB_PASS")
	DbCfg.Port = setting.Cfg.GetString("DB_PORT")
	DbCfg.SSLMode = setting.Cfg.GetString("DB_SSLMODE")
	DbCfg.Prefix = setting.Cfg.GetString("DB_PREFIX")
	DbCfg.Path = setting.Cfg.GetString("DB_PATH")
}

// NewEngine ...
func NewEngine() (err error) {
	if err = GetEngine(); err != nil {
		return err
	}
	if err = engine.StoreEngine("InnoDB").Sync2(tables...); err != nil {
		return fmt.Errorf("Sync database struct error: %v", err)
	}
	return nil
}

// GetEngine return xorm engine instance
func GetEngine() (err error) {
	engine, err = getEngine()
	if err != nil {
		return fmt.Errorf("Connect database error: %v", err)
	}

	// set struct tag reflect
	engine.SetMapper(core.GonicMapper{})

	// need set logger
	engine.ShowSQL(true)
	return nil

}

func getEngine() (*xorm.Engine, error) {
	connStr, err := GeneratorconnStr()
	if err != nil {
		return nil, err
	}
	return xorm.NewEngine(DbCfg.Type, connStr)
}

// GeneratorconnStr return sql connection string
func GeneratorconnStr() (string, error) {
	connStr := ""
	Param := "?"
	if strings.Contains(DbCfg.Name, Param) {
		Param = "&"
	}
	// Checkout SQL Type
	switch DbCfg.Type {
	case "mysql":
		connStr = fmt.Sprintf("%s:%s@tcp(%s)/%s%scharset=utf8&parseTime=true",
			DbCfg.User, DbCfg.Passwd, DbCfg.Host, DbCfg.Name, Param)
	case "postgres":
		connStr = fmt.Sprintf("postgres://%s:%s@%s:%s/%s%ssslmode=%s",
			url.QueryEscape(DbCfg.User), url.QueryEscape(DbCfg.Passwd), DbCfg.Host, DbCfg.Port, DbCfg.Name, Param, DbCfg.SSLMode)
	case "mssql":
		connStr = fmt.Sprintf("server=%s; port=%s; database=%s; user id=%s; password=%s;", DbCfg.Host, DbCfg.Port, DbCfg.Name, DbCfg.User, DbCfg.Passwd)
	case "sqlite3":
		if err := os.MkdirAll(path.Dir(DbCfg.Path), os.ModePerm); err != nil {
			return "", fmt.Errorf("Fail to create directories: %v", err)
		}
		connStr = "file:" + DbCfg.Path + "?cache=shared&mode=rwc"
	default:
		return "", fmt.Errorf("Unknown database type: %s", DbCfg.Type)
	}
	return connStr, nil
}

// Statistic define resource count
type Statistic struct {
	Counter struct {
		User int64
	}
}

// GetStatistic ...
func GetStatistic() (stats Statistic) {
	stats.Counter.User = CountUsers()
	return
}

// Ping that ping database
func Ping() error {
	return engine.Ping()
}

// Version The version table. Should have only one row with id==1
type Version struct {
	ID      int64
	Version int64
}

// ExportDatabase export all data from database to file system in JSON file
func ExportDatabase(dir string) (err error) {
	os.MkdirAll(dir, os.ModePerm)
	newtables := append(tables, new(Version))
	for _, table := range newtables {
		tableName := strings.TrimPrefix(fmt.Sprintf("%T", table), "*models")
		tableFile := path.Join(dir, tableName+".json")
		f, err := os.Create(tableFile)
		if err != nil {
			return fmt.Errorf("create JSON file %s fail: %v", tableFile, err)
		}
		defer f.Close()
		if err = engine.Asc("id").Iterate(table, func(idx int, bean interface{}) (err error) {
			enc := json.NewEncoder(f)
			return enc.Encode(bean)
		}); err != nil {

			return fmt.Errorf("fail to dump table '%s': %v", tableName, err)
		}

	}
	return nil
}
