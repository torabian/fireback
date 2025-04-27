package fireback

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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

var SERVER_INSTANCE string = UUID_Long()

var LOG *zap.Logger

// We lift two instances of webserver per application.
// One is for manager of the server, to let them have control on their
// users, workspace, support them, make changes to their credentials.

// Other one is used for public, anyone who wants to use their software,
// create account, etc.
func CreateHttpServer(handler *gin.Engine, config2 HttpServerInstanceConfig) {

	port := config.Port

	if config2.Port != 0 {
		port = config2.Port
	}

	forceSSL := config2.SSL || config.UseSSL

	if forceSSL {
		port = 443

		go func() {
			redirectServer := &http.Server{
				Addr: ":80",
				Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					target := "https://" + r.Host + r.URL.Path
					if r.URL.RawQuery != "" {
						target += "?" + r.URL.RawQuery
					}
					http.Redirect(w, r, target, http.StatusMovedPermanently)
				}),
			}
			log.Fatal(redirectServer.ListenAndServe())
		}()
	}

	server01 := &http.Server{
		Addr:         ":" + fmt.Sprintf("%v", port),
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	LOG.Info("Http server running on:", zap.String("url", "http://localhost"+server01.Addr+"/"))
	LOG.Info("Ping available on:", zap.String("url", "http://localhost"+server01.Addr+"/ping"))
	LOG.Info("Internal server ip: ** slash char \"/\" in the end is important in some sdks we generate depend on it **")

	// Get's the local IP.
	ipData := GetOutboundIP()
	if ipData != nil {
		url := "http://" + ipData.String() + server01.Addr + "/"
		LOG.Info("Local network address:", zap.String("url", url))
	}

	g.Go(func() error {
		if forceSSL {
			return server01.ListenAndServeTLS(config.CertFile, config.KeyFile)
		}
		return server01.ListenAndServe()
	})

	if config2.Monitor {
		go monitor()
	}

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
