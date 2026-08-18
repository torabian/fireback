//go:build !wasm

package fireback

import (
	"bytes"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"golang.org/x/crypto/acme/autocert"
)

// ErrNoCertificateYet is returned by InspectAcmeCache when acmeCacheDir has
// no cached certificate for the domain - either `ssl enable` was just run
// (Let's Encrypt only requests a certificate lazily, on the server's first
// HTTPS handshake for that domain - see NewAcmeManager/HttpServer.go) or the
// server has never actually served the domain yet.
var ErrNoCertificateYet = errors.New("no certificate has been issued for this domain yet")

// AcmeDomainList splits config.AcmeDomains ("example.com, www.example.com")
// into a clean slice - shared by NewAcmeManager (HttpServer.go's TLS setup)
// and the `ssl enable`/`ssl status` commands (modules/fireback/clitools/SSLCli.go).
func AcmeDomainList(csv string) []string {
	var out []string
	for _, d := range strings.Split(csv, ",") {
		d = strings.TrimSpace(d)
		if d != "" {
			out = append(out, d)
		}
	}
	return out
}

// NewAcmeManager builds the autocert.Manager backing `sslProvider:
// letsencrypt` (see HttpServer.go's CreateHttpServer): certificates for
// domainsCsv are requested from Let's Encrypt - and kept renewed - lazily,
// on demand, the first time a TLS handshake for that domain comes in over
// the HTTP-01 challenge served on port 80. Nothing is requested eagerly by
// this call itself. Cached under cacheDir (defaulting to
// "./.fireback-acme-cache" like the config field itself) so issuance isn't
// repeated across restarts - Let's Encrypt rate-limits how many certificates
// an account can request for the same domain in a week.
func NewAcmeManager(domainsCsv string, email string, cacheDir string) *autocert.Manager {
	if cacheDir == "" {
		cacheDir = "./.fireback-acme-cache"
	}

	return &autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		Cache:      autocert.DirCache(cacheDir),
		HostPolicy: autocert.HostWhitelist(AcmeDomainList(domainsCsv)...),
		Email:      email,
	}
}

// CertificateDetails is what `fireback ssl status` (and the reuse-or-replace
// prompt inside `fireback ssl enable`) reports about a certificate -
// independent of whether it came from a manual certFile/keyFile pair, a
// persisted self-signed one, or an ACME cert cached by autocert.
type CertificateDetails struct {
	Subject    string
	Issuer     string
	DNSNames   []string
	NotBefore  time.Time
	NotAfter   time.Time
	SelfSigned bool
}

// DaysRemaining returns how many days are left until NotAfter - negative
// once the certificate has expired.
func (d *CertificateDetails) DaysRemaining() int {
	return int(time.Until(d.NotAfter).Hours() / 24)
}

func (d *CertificateDetails) Expired() bool {
	return time.Now().After(d.NotAfter)
}

// describeCertificate turns a parsed x509 leaf certificate into
// CertificateDetails - shared by InspectCertFile and InspectAcmeCache below.
func describeCertificate(cert *x509.Certificate) *CertificateDetails {
	return &CertificateDetails{
		Subject:    cert.Subject.CommonName,
		Issuer:     cert.Issuer.CommonName,
		DNSNames:   cert.DNSNames,
		NotBefore:  cert.NotBefore,
		NotAfter:   cert.NotAfter,
		SelfSigned: bytes.Equal(cert.RawIssuer, cert.RawSubject),
	}
}

// InspectCertFile parses the leaf certificate out of a PEM file at
// certPath - used for sslProvider "manual" and "self-signed", where
// config.CertFile/KeyFile point at real files on disk.
func InspectCertFile(certPath string) (*CertificateDetails, error) {
	raw, err := os.ReadFile(certPath)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(raw)
	if block == nil || block.Type != "CERTIFICATE" {
		return nil, fmt.Errorf("%s does not contain a PEM certificate", certPath)
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, err
	}

	return describeCertificate(cert), nil
}

// InspectAcmeCache reads whatever certificate autocert has already cached
// for domain under cacheDir, without going through a running Manager -
// returns ErrNoCertificateYet if Let's Encrypt hasn't been asked for one
// yet (see NewAcmeManager's doc comment on lazy issuance).
//
// autocert.Manager (default, non-legacy-RSA key type) caches a domain's data
// under a file simply named after the domain itself - see certKey.String()
// and Manager.cachePut in golang.org/x/crypto/acme/autocert - containing the
// PEM-encoded private key followed by one or more PEM CERTIFICATE blocks;
// the first CERTIFICATE block is the leaf.
func InspectAcmeCache(cacheDir string, domain string) (*CertificateDetails, error) {
	if cacheDir == "" {
		cacheDir = "./.fireback-acme-cache"
	}

	raw, err := os.ReadFile(filepath.Join(cacheDir, domain))
	if err != nil {
		if os.IsNotExist(err) {
			return nil, ErrNoCertificateYet
		}
		return nil, err
	}

	rest := raw
	for {
		var block *pem.Block
		block, rest = pem.Decode(rest)
		if block == nil {
			return nil, fmt.Errorf("no certificate block found in cached data for %s", domain)
		}
		if block.Type == "CERTIFICATE" {
			cert, err := x509.ParseCertificate(block.Bytes)
			if err != nil {
				return nil, err
			}
			return describeCertificate(cert), nil
		}
	}
}
