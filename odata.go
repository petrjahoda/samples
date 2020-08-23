package main

import (
	"github.com/Azure/go-ntlmssp"
	"io/ioutil"
	"log"
	"net/http"
)

func Odata() {
	url, user, password := "http://AppNav.Likov.local:7348/LikovNAV90BetaTest/OData/Company('LIKOV%20s.r.o.')/ZAPSI_Order", "LIKOV\\SQL_NAV_SV", "SV-Likov01"
	client := &http.Client{
		Transport: ntlmssp.Negotiator{
			RoundTripper: &http.Transport{},
		},
	}
	req, _ := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(user, password)
	res, _ := client.Do(req)
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err.Error())
	}
	bodyString := string(bodyBytes)
	println(bodyString)

}
