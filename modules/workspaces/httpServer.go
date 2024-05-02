package workspaces

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
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
		log.Fatal(err)
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
	config := GetAppConfig()

	if config.Gin.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	if config.Gin.Mode == "test" {
		gin.SetMode(gin.TestMode)
	}

	if config.Gin.Mode == "debug" {
		gin.SetMode(gin.DebugMode)
	}

	if config.PublicServer.Enabled {

		port := config.PublicServer.Port

		if os.Getenv("PORT") != "" {
			port = os.Getenv("PORT")
		}

		server01 := &http.Server{
			Addr:         ":" + port,
			Handler:      handler,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
		}

		fmt.Println("Http server is listening on ")
		fmt.Println("http://localhost" + server01.Addr + "/ping")
		fmt.Println("")
		fmt.Println("Internal server ip: ** / in the end is important in some sdks we generate **")
		fmt.Println("http://" + GetOutboundIP().String() + server01.Addr + "/")
		fmt.Println(GetOutboundIP().String() + server01.Addr + "/")

		g.Go(func() error {
			return server01.ListenAndServe()
		})

	}

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Fireback exited. You need to either set publicServer.enabled to true or backOfficeServer.enabled to true in order to get web response.")
}
