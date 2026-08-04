//go:build !wasm

package clitools

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/torabian/fireback/modules/fireback"
	"go.uber.org/zap"
)

func init() {
	fireback.CreateHttpServer = createHttpServer
}

// We lift two instances of webserver per application.
// One is for manager of the server, to let them have control on their
// users, workspace, support them, make changes to their credentials.

// Other one is used for public, anyone who wants to use their software,
// create account, etc.
func createHttpServer(handler *gin.Engine, config2 fireback.HttpServerInstanceConfig) {
	port := config.Port
	if config2.Port != 0 {
		port = config2.Port
	}

	for _, vd := range config2.VirtualDomains {
		if vd == "" {
			continue
		}

		fmt.Println("Starting virtual domain: ", vd, fireback.EnableDomain(vd))
	}

	mainServer := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	forceSSL := config2.SSL || config.UseSSL

	var useVirtualCert bool

	if forceSSL {
		useVirtualCert = config.CertFile == "" || config.KeyFile == ""

		if useVirtualCert {
			cert, err := fireback.GenerateSelfSignedCert(config2.VirtualDomains)
			if err != nil {
				log.Fatal("failed to generate self-signed cert:", err)
			}

			mainServer.TLSConfig = &tls.Config{
				Certificates: []tls.Certificate{cert},
			}

			fmt.Println("Using self-signed certificate (no cert files provided)")
		}
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
			if useVirtualCert {
				err = mainServer.ListenAndServeTLS("", "")
			} else {
				err = mainServer.ListenAndServeTLS(config.CertFile, config.KeyFile)
			}
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
	fireback.LOG.Info("Shutting down...")

	for _, vd := range config2.VirtualDomains {
		if vd == "" {
			continue
		}
		fmt.Println("Stopping virtual domain: ", vd, fireback.DisableDomain(vd))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := mainServer.Shutdown(ctx); err != nil {
		fireback.LOG.Error("Main server shutdown failed", zap.Error(err))
	}

	if redirectServer != nil {
		_ = redirectServer.Shutdown(ctx)
	}

	fireback.LOG.Info("Server exited properly")
}
