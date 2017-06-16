// @APIVersion 1.0.0
// @FriendsManagement API
// @Description FriendsManagement API
// @Contact leizhe@chinasofti.com
package routers

import (
	"FriendsManagement/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/Subscribe",
			beego.NSInclude(
				&controllers.SubscribeController{},
			),
		),

		beego.NSNamespace("/Firend",
			beego.NSInclude(
				&controllers.FirendController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
