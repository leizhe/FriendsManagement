package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["FriendsManagement/controllers:FirendController"] = append(beego.GlobalControllerRouter["FriendsManagement/controllers:FirendController"],
		beego.ControllerComments{
			Method: "AddFriend",
			Router: `/AddFriends`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["FriendsManagement/controllers:FirendController"] = append(beego.GlobalControllerRouter["FriendsManagement/controllers:FirendController"],
		beego.ControllerComments{
			Method: "GetAllFriends",
			Router: `/GetFriends`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["FriendsManagement/controllers:FirendController"] = append(beego.GlobalControllerRouter["FriendsManagement/controllers:FirendController"],
		beego.ControllerComments{
			Method: "GetCommonFriends",
			Router: `/GetCommonFriends`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["FriendsManagement/controllers:SubscribeController"] = append(beego.GlobalControllerRouter["FriendsManagement/controllers:SubscribeController"],
		beego.ControllerComments{
			Method: "AddSubscribe",
			Router: `/AddSubscribe`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["FriendsManagement/controllers:SubscribeController"] = append(beego.GlobalControllerRouter["FriendsManagement/controllers:SubscribeController"],
		beego.ControllerComments{
			Method: "BlockSubscribe",
			Router: `/BlockSubscribe`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["FriendsManagement/controllers:SubscribeController"] = append(beego.GlobalControllerRouter["FriendsManagement/controllers:SubscribeController"],
		beego.ControllerComments{
			Method: "RetrieveSubscribe",
			Router: `/RetrieveSubscribe`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

}
