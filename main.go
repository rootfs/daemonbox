package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

func InfoHandler(res http.ResponseWriter, req *http.Request) {
	info := "This is a daemonbox"

	r := render.New(render.Options{})
	r.JSON(res, 200, info)

}

func CmdHandler(res http.ResponseWriter, req *http.Request) {
	st := 200
	vars := mux.Vars(req)
	c := vars["cmd"]
	println(c)
	r := render.New(render.Options{})
	cmd := exec.Command(c)
	output, err := cmd.CombinedOutput()
	if err != nil {
		st = 400
		fmt.Printf("%v", string(output))
	}

	r.JSON(res, st, string(output))

}

func CmdArgHandler(res http.ResponseWriter, req *http.Request) {
	st := 200
	vars := mux.Vars(req)
	r := render.New(render.Options{})
	c := vars["cmd"]
	arg := vars["arg"]
	argStr, err := base64.URLEncoding.DecodeString(arg)
	if err != nil {
		st = 400
		fmt.Printf("bad arg %v", err)
		r.JSON(res, st, err.Error())
		return
	}
	argArr := strings.Fields(string(argStr))
	println(c, string(argStr))
	cmd := exec.Command(c, argArr...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		st = 400
	}
	fmt.Printf("%v", string(output))

	r.JSON(res, st, string(output))

}

func main() {
	var port = ":" + os.Getenv("SERVICE_PORT")
	if port == ":" {
		port = ":3000"
	}

	r := mux.NewRouter()

	// define RESTful handlers
	r.Path("/info").Methods("GET").HandlerFunc(InfoHandler)
	r.Path("/cmd/{cmd}").Methods("GET").HandlerFunc(CmdHandler)
	r.Path("/cmd/{cmd}/{arg}").Methods("GET").HandlerFunc(CmdArgHandler)

	n := negroni.New(negroni.NewLogger())

	n.UseHandler(r)

	n.Run(port)
}
