package db

import (
	"errors"

	"github.com/go-pg/pg/v10"
	"github.com/loco/conf"
)

//set postgres connection options from conf
func getPostgresOptions() (pgOption pg.Options) {
	container := "database.master"
	host := conf.String(container+".host", "")
	port := conf.String(container+".port", "")
	addr := ""
	if host != "" && port != "" {
		addr = host + ":" + port
	}
	pgOption.Addr = addr
	pgOption.User = conf.String(container+".username", "")
	pgOption.Password = conf.String(container+".password", "")
	pgOption.Database = conf.String(container+".db", "")
	return
}

func Conn() (conn *pg.DB, err error) {
	pgOption := getPostgresOptions()

	conn = pg.Connect(&pgOption)

	if conn == nil {
		err = errors.New("could not connect DB")
		return
	}

	return
}
