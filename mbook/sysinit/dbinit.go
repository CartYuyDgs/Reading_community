package sysinit

import (
	_ "Reading_community/mbook/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//数据库初始化
func dbinit(aliases ...string) {

	isDev := ("dev" == beego.AppConfig.String("runmode"))

	if len(aliases) > 0 {
		for _, alias := range aliases {
			registDatabase(alias)
			if alias == "w" {
				orm.RunSyncdb("default", false, true) //自动建表
			}
		}
	} else {
		registDatabase("w")
		orm.RunSyncdb("default", false, true) //自动建表
	}

	if isDev {
		orm.Debug = isDev
	}
}

func registDatabase(alias string) {
	if len(alias) <= 0 {
		return
	}
	dbAlias := alias //default
	if "w" == alias || "default" == alias {
		dbAlias = "default"
		alias = "w"
	}
	//数据库名
	dbName := beego.AppConfig.String("db_" + alias + "_database")
	//数据库用户名
	dbUser := beego.AppConfig.String("db_" + alias + "_username")
	//mima
	dbPasswd := beego.AppConfig.String("db_" + alias + "_username")
	//IP
	dbHost := beego.AppConfig.String("db_" + alias + "_host")

	dbPort := beego.AppConfig.String("db_" + alias + "_port")
	//root:199411@tcp(192.168.31.205:3306)/test?charset=utf8"
	orm.RegisterDataBase(dbAlias, "mysql", dbUser+":"+dbPasswd+"@tcp("+dbHost+":"+dbPort+")"+"/"+dbName+"?charset=utf8")

}
