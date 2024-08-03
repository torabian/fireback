package workspaces

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)

func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return nil
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

// We lift two instances of webserver per application.
// One is for manager of the server, to let them have control on their
// users, workspace, support them, make changes to their credentials.

// Other one is used for public, anyone who wants to use their software,
// create account, etc.
func CreateHttpServer(handler *gin.Engine) {

	if config.GinMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	if config.GinMode == "test" {
		gin.SetMode(gin.TestMode)
	}

	if config.GinMode == "debug" {
		gin.SetMode(gin.DebugMode)
	}

	server01 := &http.Server{
		Addr:         ":" + fmt.Sprintf("%v", config.Port),
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("Http server is listening on ")
	fmt.Println("http://localhost" + server01.Addr + "/ping")
	fmt.Println("")
	fmt.Println("Internal server ip: ** slash char \"/\" in the end is important in some sdks we generate depend on it **")

	ipData := GetOutboundIP()

	if ipData != nil {

		fmt.Println("http://" + ipData.String() + server01.Addr + "/")
		fmt.Println(ipData.String() + server01.Addr + "/")
	}

	g.Go(func() error {
		return server01.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Fireback exited. You need to either set publicServer.enabled to true or backOfficeServer.enabled to true in order to get web response.")
}
