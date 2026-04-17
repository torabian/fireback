package fireback

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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

	for _, vd := range config2.VirtualDomains {
		if vd == "" {
			continue
		}

		fmt.Println("Starting virtual domain: ", vd, EnableDomain(vd))
	}

	forceSSL := config2.SSL || config.UseSSL

	mainServer := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	var redirectServer *http.Server

	if forceSSL {
		mainServer.Addr = ":443"

		redirectServer = &http.Server{
			Addr: ":80",
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				target := "https://" + r.Host + r.URL.RequestURI()
				http.Redirect(w, r, target, http.StatusMovedPermanently)
			}),
		}

		go func() {
			if err := redirectServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatal(err)
			}
		}()
	}

	// Run main server
	go func() {
		var err error
		if forceSSL {
			err = mainServer.ListenAndServeTLS(config.CertFile, config.KeyFile)
		} else {
			err = mainServer.ListenAndServe()
		}
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	// --- Graceful shutdown ---
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop
	LOG.Info("Shutting down...")

	for _, vd := range config2.VirtualDomains {
		if vd == "" {
			continue
		}
		fmt.Println("Stopping virtual domain: ", vd, DisableDomain(vd))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := mainServer.Shutdown(ctx); err != nil {
		LOG.Error("Main server shutdown failed", zap.Error(err))
	}

	if redirectServer != nil {
		_ = redirectServer.Shutdown(ctx)
	}

	LOG.Info("Server exited properly")
}
