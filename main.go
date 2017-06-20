package main

import (
	_ "FriendsManagement/routers"
	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func init() {
	//orm.RegisterDataBase("default", "postgres", "postgres://u76e5497ba90b458586b664afb21ab09f:57c51e87ac4743f3aa871c1c83dfa07f@10.120.8.137:5432/d4280ca62e4be4c2aa5f2f40e03b5c590?sslmode=disable")
	orm.RegisterDataBase("default", "postgres", "postgres://postgres:1qaz2wsx3EDC@192.168.10.199:5432/FMDB?sslmode=disable")
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run(":" + os.Getenv("PORT"))
}
