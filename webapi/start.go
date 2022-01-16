package webapi

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

var httpServ *http.Server

func Launch(addr string, port int16) {
	route := getRoute()

	if httpServ == nil {
		httpServ = &http.Server{
			Addr:    fmt.Sprintf("%s:%d", addr, port),
			Handler: route,
		}
	}

	if err := httpServ.ListenAndServe(); err != nil {
		log.Fatalf("listen error:%+v\n", err)
	}
}
func Stop() {
	if httpServ == nil {
		return
	}
	httpServ.Shutdown(context.Background())
	httpServ = nil
}
