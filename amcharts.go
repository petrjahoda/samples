package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
)

func amchartsSample() {

	println(time.Now().Format("2006-01-02T15:04:05Z"))
	println(time.Now().UTC().Format("2006-01-02T15:04:05Z"))
	//println(time.Now().Add(-11*time.Hour).Format("2006-01-02T15:04:05Z"))

	router := httprouter.New()
	router.GET("/", amcharts)
	router.GET("/data", jsonData)
	_ = http.ListenAndServe(":82", router)
}

func amcharts(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	println("called")
	http.ServeFile(writer, request, "main.html")
}
func jsonData(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	println("called 2")
	http.ServeFile(writer, request, "data.json")
}
