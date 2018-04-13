package filter

import (
	"net/http"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	beego.InsertFilter("/*", beego.BeforeRouter, CrossDomainFilter)
}

func CrossDomainFilter(ctx *context.Context) {
	if ctx.Request.Method == "OPTIONS" {
		ctx.Output.Header("Access-Control-Allow-Origin", "*")
		ctx.Output.Header("Access-Control-Allow-Headers", "authorization")
		ctx.Output.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		ctx.ResponseWriter.WriteHeader(http.StatusOK)
	}
}
