package goadmin

import (
	_ "github.com/GoAdminGroup/go-admin/adapter/gin"
	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/examples/datamodel"
	"github.com/GoAdminGroup/go-admin/modules/config"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/mysql"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/plugins/admin"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/chartjs"
	"github.com/GoAdminGroup/themes/adminlte"
	"github.com/gin-gonic/gin"
)

func Goadmin() {
	r := gin.Default()

	eng := engine.Default()

	// global config
	cfg := config.Config{
		Databases: config.DatabaseList{
			"default": config.Database{
				Host:         "127.0.0.1",
				Port:         "3306",
				User:         "root",
				Pwd:          "Mysql112223",
				Name:         "goadmin",
				MaxIdleConns: 10,
				MaxOpenConns: 100,
				Driver:       "mysql",
			},
		},
		UrlPrefix: "admin",
		// STORE 必须设置且保证有写权限，否则增加不了新的管理员用户
		Store: config.Store{
			Path:   "./uploads",
			Prefix: "uploads",
		},
		Language: language.CN,
		// 开发模式
		Debug: true,
		// 日志文件位置，需为绝对路径
		InfoLogPath:   "/var/logs/info.log",
		AccessLogPath: "/var/logs/access.log",
		ErrorLogPath:  "/var/logs/error.log",
		ColorScheme:   adminlte.ColorschemeSkinBlack,
	}

	// Generators： 详见 https://github.com/GoAdminGroup/go-admin/blob/master/examples/datamodel/tables.go
	adminPlugin := admin.NewAdmin(datamodel.Generators)

	// 增加 chartjs 组件
	template.AddComp(chartjs.NewChart())

	// 增加 generator, 第一个参数是对应的访问路由前缀
	// 例子:
	//
	// "user" => http://localhost:9033/admin/info/user
	//
	// adminPlugin.AddGenerator("user", datamodel.GetUserTable)

	// 自定义首页

	// r.GET("/admin", func(ctx *gin.Context) {
	// 	engine.Content(ctx, func(ctx interface{}) (types.Panel, error) {
	// 		return datamodel.GetContent()
	// 	})
	// })

	_ = eng.AddConfig(&cfg).AddPlugins(adminPlugin).Use(r)

	eng.HTML("GET", "/admin", datamodel.GetContent)

	_ = r.Run(":9033")
}
