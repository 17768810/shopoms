package main

import (
	"github.com/astaxie/beego"
	"github.com/wqdsoft/shopoms/app/jobs"
	"github.com/wqdsoft/shopoms/app/models"
	_ "github.com/wqdsoft/shopoms/app/routers"
	"html/template"
	"net/http"
	"os"
)

const VERSION = "1.0.0"

func initArgs() {
	args := os.Args
	for _, v := range args {
		if v == "-syncdb" {
			models.Syncdb()
			os.Exit(0)
		}
	}
}

func main() {
	initArgs()
	//判断是否安装
	models.Init()
	jobs.InitJobs()
	// 设置默认404页面
	beego.ErrorHandler("404", func(rw http.ResponseWriter, r *http.Request) {
		t, _ := template.New("404.html").ParseFiles(beego.BConfig.WebConfig.ViewsPath + "/error/404.html")
		data := make(map[string]interface{})
		data["content"] = "page not found"
		t.Execute(rw, data)
	})
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.AppConfig.Set("version", VERSION)
	beego.Run()
}
