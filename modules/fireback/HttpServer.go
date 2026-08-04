package fireback

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"math/big"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/torabian/fireback/modules/fireback/vdomain"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)

// CreateHttpServer lifts the gin engine into a real HTTP(S) server with
// graceful shutdown on OS signals - unavailable under wasm. Real
// implementation lives in modules/fireback/clitools (tagged !wasm) and
// registers itself here via init().

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

		fmt.Println("Starting virtual domain: ", vd, vdomain.EnableDomain(vd))
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
			cert, err := GenerateSelfSignedCert(config2.VirtualDomains)
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
	LOG.Info("Shutting down...")

	for _, vd := range config2.VirtualDomains {
		if vd == "" {
			continue
		}
		fmt.Println("Stopping virtual domain: ", vd, vdomain.DisableDomain(vd))
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

func GenerateSelfSignedCert(domains []string) (tls.Certificate, error) {
	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return tls.Certificate{}, err
	}

	if len(domains) == 0 {
		domains = []string{"localhost"}
	}

	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(365 * 24 * time.Hour),

		KeyUsage:    x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},

		BasicConstraintsValid: true,
		DNSNames:              domains,
	}

	// add localhost fallback
	template.DNSNames = append(template.DNSNames, "localhost")

	// add IP support
	template.IPAddresses = []net.IP{net.ParseIP("127.0.0.1")}

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)
	if err != nil {
		return tls.Certificate{}, err
	}

	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(priv)})

	return tls.X509KeyPair(certPEM, keyPEM)
}
