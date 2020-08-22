package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func FaceDetectionHtmlWithStaticFilesAndHttpRouter() {
	router := httprouter.New()
	router.ServeFiles("/face_detection/*filepath",http.Dir("face_detection"))
	router.ServeFiles("/models/*filepath",http.Dir("face_detection/models"))
	router.GET("/", Authorization(Home()))
	log.Println("Listening on :3000...")
	_ = http.ListenAndServe(":3000", router)
}


func Authorization(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		user, pass, ok := r.BasicAuth()
		if ok {
			println(user, pass)
		} else {
			println("no auth")
		}
		next(w, r, ps)
	}
}

func Home() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		http.ServeFile(w, r, "home.html")
	}
}
