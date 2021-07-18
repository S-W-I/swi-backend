package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"runtime/debug"
	"swi/server/router"

	fasthttp "github.com/valyala/fasthttp"
)


func main() {

	swiRouter := router.NewRouter()
	fastHTTPHandler := func (ctx *fasthttp.RequestCtx) {
		swiRouter.Route(ctx)
	}

	fasthttp.ListenAndServe(":8081", fastHTTPHandler)
}

