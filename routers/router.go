// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"github.com/document-server/controllers"
	_ "github.com/document-server/filter"
)

func init() {
	ns := beego.NewNamespace("/api/v1",
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		//beego.NSInclude(
		//	&controllers.UploadFile{},
		//),
		beego.NSInclude(
			&controllers.TempFile{},
		),
	)
	beego.AddNamespace(ns)
}
