package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["im/client/controllers/user:UserController"] = append(beego.GlobalControllerRouter["im/client/controllers/user:UserController"],
		beego.ControllerComments{
			Method: "CheckAuth",
			Router: `/check_auth`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["im/client/controllers/user:UserController"] = append(beego.GlobalControllerRouter["im/client/controllers/user:UserController"],
		beego.ControllerComments{
			Method: "ExistEmail",
			Router: `/existEmail`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["im/client/controllers/user:UserController"] = append(beego.GlobalControllerRouter["im/client/controllers/user:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["im/client/controllers/user:UserController"] = append(beego.GlobalControllerRouter["im/client/controllers/user:UserController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["im/client/controllers/user:UserController"] = append(beego.GlobalControllerRouter["im/client/controllers/user:UserController"],
		beego.ControllerComments{
			Method: "Register",
			Router: `/register`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})
	beego.GlobalControllerRouter["im/client/controllers/user:UserController"] = append(beego.GlobalControllerRouter["im/client/controllers/user:UserController"],
		beego.ControllerComments{
			Method: "GGim",
			Router: `/ggim`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
