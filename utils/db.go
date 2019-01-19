package utils

import (
	"fmt"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nullscc/jenkins_requirements/models"
	"os"
)

var (
	Db *xorm.Engine
)

func init() {
	if len(os.Args) > 1 {
		return
	}

	engine, _ := xorm.NewEngine("sqlite3", "./sqlite3.db")
	var err error

	err = engine.Sync2(new(models.RepoRequirement))
	if err != nil {
		panic(fmt.Sprintf("sqlite3 sync2 error: %s", err.Error()))
	}

	Db = engine
}
