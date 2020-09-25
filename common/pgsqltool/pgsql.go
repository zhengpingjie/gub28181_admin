package pgsqltool

import (
	"admin/common/common"
	"github.com/go-pg/pg"
	_ "github.com/lib/pq"
)

func Pgsql(){
	pgConfig := common.GVA_VP.GetStringMapString("pgsql")
	db :=pg.Connect(&pg.Options{
		Addr:pgConfig["addr"],
		User:pgConfig["user"],
		Password:pgConfig["password"],
		Database:pgConfig["dbname"],
	})
	var n int
	_,err:=db.QueryOne(pg.Scan(&n),"SELECT 1")
	if err != nil{
		panic(err)
	}
	common.GVA_PG = db
}