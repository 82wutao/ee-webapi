package web

import (
	"fmt"
	"net/http"
)

func Launch(hostIp string, port int16, frontEnd http.Handler) {

	httpServ := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", hostIp, port),
		Handler: frontEnd,
	}

	if err := httpServ.ListenAndServe(); err != nil {
		panic("launch http server fail.")
	}

}
