package main

import (
	"IrisDev/router"
	"IrisDev/utils"
	"fmt"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"time"
)

func log(ctx iris.Context) {
	t := time.Now()
	jqs := "?"
	for key, value := range ctx.URLParams() {
		jqs += key + "=" + value
		jqs += "&"
	}
	jqs = jqs[:len(jqs)-1]
	fmt.Printf("%s  %s %s\n", t.Format("2006-01-02 15:04:05"), ctx.Method(), ctx.FullRequestURI()+jqs)
	ctx.Next()
}
func main() {
	fmt.Println("Server init...")
	utils.IndexUrl = ""

	app := iris.New()

	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})
	app.Use(crs)
	app.Use(log)

	app.RegisterView(iris.HTML("views", ".html"))
	app.Get("/", router.GetIndex)

	login := app.Party("/login")
	login.Get("/qr", router.LoginGetQrCode)
	login.Get("/set-status", router.LoginSetStatus)
	login.Get("/get-status", router.LoginGetStatus)

	err := app.Run(iris.Addr(":9102"))
	if err != nil {
		return
	} else {
		fmt.Println("init err" + err.Error())
	}
}
