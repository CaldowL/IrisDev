package router

import (
	"github.com/kataras/iris/v12"
)

func GetIndex(ctx iris.Context) {
	err := ctx.View("index.html")
	if err != nil {
		return
	}
}
