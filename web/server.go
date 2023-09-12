package web

import (
	"context"
	"fmt"
	"net/http"
)

var httpServ *http.Server

func Launch(hostIp string, port int16, frontEnd http.Handler) {

	httpServ = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", hostIp, port),
		Handler: frontEnd,
	}

	if err := httpServ.ListenAndServe(); err != nil {
		panic("launch http server fail.")
	}

}

func Stop() {
	if httpServ == nil {
		return
	}
	httpServ.Shutdown(context.Background())
	httpServ = nil
}
