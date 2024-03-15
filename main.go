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

	UPG()

}

func ONU_FR() {
	sopa, err := sopa.Constructor()
	if err != nil {
		log.Println(err)
	}

	sopa.PostAndSetCookie(variables.URL_POST_LOGIN, variables.Payload)

	// Factory Reset
	sopa.Get(variables.URL_DEVICE_FR)
	key, err := onu.ExtractValue(*sopa.HTMLResponse)
	if err != nil {
		log.Println(err)
	}
	variables.PayloadFR.Set("sessionkey", key)
	sopa.Post(variables.URL_DEVICE_DoFR, variables.PayloadFR)

}

func ONU_FW() {
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

	b, err := os.ReadFile("up3.img")
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

}

func ONU_FW2() {

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
	variables.PayloadFW.Set("sessionkey", key)

	// values := map[string]io.Reader{
	// 	"file":       sopa.MustOpen("main.go"), // lets assume its this file
	// 	"other":      strings.NewReader("hello world!"),
	// 	"onSubmit":   strings.NewReader("loading_fwupgrade"),
	// 	"FW_File":    strings.NewReader("up3.img"),
	// 	"sessionkey": strings.NewReader(key),
	// }

	sopa.Up("http://192.168.101.1/cgi-bin/fwupgrade.cgi")
	sopa.Get("http://192.168.101.1/cgi-bin/loading.cgi?url=fwupgrade.cgi&waittime=90&operation=docmd_fwupgrade")
	key, err = onu.ExtractValue(*sopa.HTMLResponse)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(key)
	variables.PayloadFWDo.Set("sessionkey", key)
	fmt.Println(variables.PayloadFWDo)

	sopa.Post("http://192.168.101.1/cgi-bin/fwupgrade.cgi", variables.PayloadFWDo)

}

func UPG() {

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

	data := map[string]string{
		"onSubmit":   "loading_fwupgrade",
		"sessionkey": key,
	}

	b, err := os.ReadFile("up.img")
	if err != nil {
		fmt.Println("ssdfasdfasdfasdfg")
		log.Println(err)
	}

	sopa.Up2(
		"http://192.168.101.1/cgi-bin/fwupgrade.cgi",
		data,
		b,
	)

	sopa.Get("http://192.168.101.1/cgi-bin/fwupgrade.cgi")
	key, err = onu.ExtractValue(*sopa.HTMLResponse)
	if err != nil {
		log.Println(err)
	}

	data = map[string]string{
		"onSubmit":   "loading_fwupgrade",
		"sessionkey": key,
	}

	b, err = os.ReadFile("up3.img")
	if err != nil {
		fmt.Println("ssdfasdfasdfasdfg")
		log.Println(err)
	}

	sopa.Up2(
		"http://192.168.101.1/cgi-bin/fwupgrade.cgi",
		data,
		b,
	)

	///

	sopa.Get("http://192.168.101.1/cgi-bin/loading.cgi?url=fwupgrade.cgi&waittime=90&operation=docmd_fwupgrade")
	key, err = onu.ExtractValue(*sopa.HTMLResponse)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(key)
	variables.PayloadFWDo.Set("sessionkey", key)
	fmt.Println(variables.PayloadFWDo)

	sopa.Post("http://192.168.101.1/cgi-bin/fwupgrade.cgi", variables.PayloadFWDo)

}
