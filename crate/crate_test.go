package tests

import (
	"database/sql"
	"testing"

	. "github.com/go-xorm/tests"
	"github.com/go-xorm/xorm"
	_ "github.com/herenow/go-crate"
)

func connStr() string {
	return "http://localhost:4200"
}

func newCrateEngine() (*xorm.Engine, error) {
	orm, err := xorm.NewEngine("crate", connStr())
	if err != nil {
		return nil, err
	}
	orm.ShowSQL(ShowTestSql)

	tables, err := orm.DBMetas()
	if err != nil {
		return nil, err
	}
	for _, table := range tables {
		_, err = orm.Exec("drop table \"" + table.Name + "\"")
		if err != nil {
			return nil, err
		}
	}

	return orm, err
}

func newCrateDriverDB() (*sql.DB, error) {
	return sql.Open("crate", connStr())
}

func TestCrate(t *testing.T) {
	engine, err := newCrateEngine()
	if err != nil {
		t.Error(err)
		return
	}
	defer engine.Close()

	BaseTestAll(engine, t)
	// UserTest1(engine, t)
	// BaseTestAllSnakeMapper(engine, t)
	// BaseTestAll2(engine, t)
	// BaseTestAll3(engine, t)
}
