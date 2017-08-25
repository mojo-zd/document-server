package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/document-server/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/document-server/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/document-server/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/document-server/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/document-server/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/document-server/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/document-server/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/document-server/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/document-server/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/document-server/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/document-server/controllers:UploadFile"] = append(beego.GlobalControllerRouter["github.com/document-server/controllers:UploadFile"],
		beego.ControllerComments{
			Method: "UploadFile",
			Router: `/file/upload`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/document-server/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/document-server/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
