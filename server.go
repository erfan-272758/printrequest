package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	err := bootstrap()
	if err != nil {
		log.Fatalf("%v", err)
	}
}

func initPort() (port int64) {
	if p, ok := os.LookupEnv("PORT"); ok {
		pp, err := strconv.ParseInt(p, 10, 64)
		if err != nil {
			port = 8080
		} else {
			port = pp
		}
	} else {
		port = 8080
	}
	return
}

func onReq(ctx *gin.Context) {
	resp, err := httputil.DumpRequest(ctx.Request, true)
	if err != nil {
		ctx.AbortWithError(500, err)
	} else {
		ctx.Status(200)
		ctx.String(200, "%v", string(resp))
	}

}

func bootstrap() error {
	gin.SetMode(gin.DebugMode)
	server := gin.Default()
	port := initPort()
	log.Printf("Server listen at http://127.0.0.1:%v", port)

	server.NoRoute(onReq)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%v", port),
		Handler: server.Handler(),
	}
	err := srv.ListenAndServe()
	return err
}
