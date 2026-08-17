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
	"strings"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/torabian/fireback/modules/fireback/vdomain"
	"go.uber.org/zap"
	"golang.org/x/crypto/acme/autocert"
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
	var acmeManager *autocert.Manager

	if forceSSL {
		switch {
		// config.SslProvider "manual" and "self-signed" (persisted - see
		// modules/fireback/clitools/SSLCli.go's `ssl enable`) both just point
		// CertFile/KeyFile at real files on disk and fall through to the
		// plain ListenAndServeTLS(config.CertFile, config.KeyFile) branch
		// below like always - only "letsencrypt" needs a TLSConfig of its
		// own, since its certificate isn't a static file but requested (and
		// kept renewed) on demand by autocert.Manager.
		case config.SslProvider == "letsencrypt" && strings.TrimSpace(config.AcmeDomains) != "":
			acmeManager = NewAcmeManager(config.AcmeDomains, config.AcmeEmail, config.AcmeCacheDir)
			mainServer.TLSConfig = acmeManager.TLSConfig()
			fmt.Printf("Using Let's Encrypt certificate for %s (requested on first handshake, cached in %s)\n", config.AcmeDomains, config.AcmeCacheDir)

		case config.CertFile == "" || config.KeyFile == "":
			useVirtualCert = true
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

		var portEightyHandler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			target := "https://" + r.Host + r.URL.RequestURI()
			http.Redirect(w, r, target, http.StatusMovedPermanently)
		})

		if acmeManager != nil {
			// Serves the ACME HTTP-01 challenge under /.well-known/acme-challenge/
			// (required to be plain, unredirected HTTP on port 80), falling back
			// to the https redirect above for every other request.
			portEightyHandler = acmeManager.HTTPHandler(portEightyHandler)
		}

		redirectServer = &http.Server{
			Addr:    ":80",
			Handler: portEightyHandler,
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
			if useVirtualCert || acmeManager != nil {
				// Both branches supply their certificate via mainServer.TLSConfig
				// (a static self-signed cert, or acmeManager.GetCertificate)
				// rather than a certFile/keyFile pair on disk.
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

// GenerateSelfSignedCertPEM does the actual cert generation for
// GenerateSelfSignedCert, returning the raw PEM bytes rather than a parsed
// tls.Certificate - split out so `ssl enable`'s self-signed option
// (modules/fireback/clitools/SSLCli.go) can persist them to certFile/keyFile
// on disk instead of the ephemeral, regenerated-every-start certificate this
// package falls back to when forceSSL is on but no cert files are configured.
func GenerateSelfSignedCertPEM(domains []string) (certPEM []byte, keyPEM []byte, err error) {
	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
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
		return nil, nil, err
	}

	certPEM = pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	keyPEM = pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(priv)})

	return certPEM, keyPEM, nil
}

func GenerateSelfSignedCert(domains []string) (tls.Certificate, error) {
	certPEM, keyPEM, err := GenerateSelfSignedCertPEM(domains)
	if err != nil {
		return tls.Certificate{}, err
	}

	return tls.X509KeyPair(certPEM, keyPEM)
}
