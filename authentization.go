package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strings"
)

func httpRouterAuthentication() {
	user := []byte("gordon")
	pass := []byte("secret!")
	router := httprouter.New()
	router.GET("/", index)
	router.GET("/protected/", basicAuth(protected, user, pass))
	router.GET("/new/", basicAuth(newPage, user, pass))
	_ = http.ListenAndServe(":8080", router)
}

func basicAuth(h httprouter.Handle, user, pass []byte) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		const basicAuthPrefix string = "Basic "
		auth := r.Header.Get("Authorization")
		if strings.HasPrefix(auth, basicAuthPrefix) {
			payload, err := base64.StdEncoding.DecodeString(auth[len(basicAuthPrefix):])
			if err == nil {
				pair := bytes.SplitN(payload, []byte(":"), 2)
				if len(pair) == 2 && bytes.Equal(pair[0], user) && bytes.Equal(pair[1], pass) {
					h(w, r, ps)
					return
				}
			}
		}
		w.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
	}
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_, _ = fmt.Fprint(w, "Not protected!\n")
}

func protected(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_, _ = fmt.Fprint(w, "protected!\n")
}

func newPage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_, _ = fmt.Fprint(w, "new!\n")
}
