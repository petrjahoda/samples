package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func faceDetectionHtmlWithStaticFilesAndHttpRouter() {
	router := httprouter.New()
	router.ServeFiles("/face_detection/*filepath", http.Dir("face_detection"))
	router.ServeFiles("/models/*filepath", http.Dir("face_detection/models"))
	router.GET("/", home)
	log.Println("Listening on :3000...")
	_ = http.ListenAndServe(":3000", router)
}

func home(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	http.ServeFile(writer, request, "home.html")
}
