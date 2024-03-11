package main

import (
	"fmt"
	"log"
	"os"

	onu "github.com/rivaldito/sopa/ONU"
	sopa "github.com/rivaldito/sopa/Sopa"
	variables "github.com/rivaldito/sopa/Variables"
)

func main() {

	// sopa, err := sopa.Constructor()
	// if err != nil {
	// 	log.Println(err)
	// }

	// sopa.PostAndSetCookie(variables.URL_POST_LOGIN, variables.Payload)

	// err = sopa.Get(variables.URL_DEVICE_INFO)
	// if err != nil {
	// 	log.Println(err)
	// }

	// sopa.HTMLParse()

	ONU()

}

func ONU() {
	sopa, err := sopa.Constructor()
	if err != nil {
		log.Println(err)
	}

	sopa.PostAndSetCookie(variables.URL_POST_LOGIN, variables.Payload)

	sopa.Get("http://192.168.101.1/cgi-bin/fwupgrade.cgi")
	key, err := onu.ExtractValue(*sopa.HTMLResponse)
	if err != nil {
		log.Println(err)
	}

	b, err := os.ReadFile("up2.img")
	if err != nil {
		log.Println(err)
	}
	variables.PayloadFW.Set("sessionkey", key)
	fmt.Println(variables.PayloadFW)
	sopa.PostUploadBinaryFile("http://192.168.101.1/cgi-bin/fwupgrade.cgi",
		variables.PayloadFR,
		b)
	sopa.Get("http://192.168.101.1/cgi-bin/loading.cgi?url=fwupgrade.cgi&waittime=90&operation=docmd_fwupgrade")
	key, err = onu.ExtractValue(*sopa.HTMLResponse)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(key)
	variables.PayloadFWDo.Set("sessionkey", key)
	fmt.Println(variables.PayloadFWDo)

	sopa.Post("http://192.168.101.1/cgi-bin/fwupgrade.cgi", variables.PayloadFWDo)

	// // Factory Reset
	// sopa.Get(variables.URL_DEVICE_FR)
	// key, err := onu.ExtractValue(*sopa.HTMLResponse)
	// if err != nil {
	// 	log.Println(err)
	// }
	// variables.PayloadFR.Set("sessionkey", key)
	// sopa.Post(variables.URL_DEVICE_DoFR, variables.PayloadFR)

}
