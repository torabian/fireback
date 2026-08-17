//go:build !wasm

// SSLCli.go implements `fireback ssl enable` and `fireback ssl status` - see
// fireback.SSLCliCommand's doc comment (CliActions.go) for how this wires
// into the core package.
//
// `ssl enable` first checks whether a certificate is already configured
// (any of config.SslProvider's three flavors - manual, self-signed,
// letsencrypt - see Configuration's doc comment on that field) and, if one
// is found, asks whether to keep it or replace it, rather than silently
// clobbering a working setup. Replacing (or starting from nothing) then
// offers three ways to obtain a certificate:
//
//   - Let's Encrypt: no external certbot/acme.sh install needed - fireback
//     itself speaks ACME via golang.org/x/crypto/acme/autocert
//     (see fireback.NewAcmeManager). The certificate isn't requested here;
//     `ssl enable` only saves config, and the running server requests (and
//     transparently renews) it lazily on the first real HTTPS handshake for
//     the domain - see HttpServer.go's CreateHttpServer.
//   - A certificate/key pair the operator already has.
//   - A self-signed certificate, generated here and persisted to disk (unlike
//     the ephemeral one HttpServer.go falls back to when useSSL is on with no
//     cert files at all - that one is regenerated, and thus changes, on every
//     restart).
package clitools

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/torabian/fireback/modules/fireback"
	"github.com/urfave/cli/v3"
)

func init() {
	fireback.SSLCliCommand = sslCliCommand
}

const (
	sslPickLetsEncrypt = "Let's Encrypt - automatic and trusted by browsers (recommended)"
	sslPickManual      = "I already have a certificate + key file"
	sslPickSelfSigned  = "Self-signed - for local/dev use, not trusted by browsers"

	sslPickKeep    = "Keep the existing certificate"
	sslPickReplace = "Replace it with a new certificate"
)

func sslCliCommand() *cli.Command {
	return &cli.Command{
		Name:  "ssl",
		Usage: "Enable, inspect, or replace the SSL certificate the http server uses",
		Commands: []*cli.Command{
			{
				Name:  "enable",
				Usage: "Wizard to configure SSL: detects any certificate already set up and offers Let's Encrypt, a manual cert+key pair, or a self-signed certificate",
				Action: func(ctx context.Context, c *cli.Command) error {
					sslEnableWizard()
					return nil
				},
			},
			{
				Name:  "status",
				Usage: "Shows details (domains, issuer, expiry) about the certificate currently configured",
				Action: func(ctx context.Context, c *cli.Command) error {
					sslStatus()
					return nil
				},
			},
		},
	}
}

// sslCertificateStatus is the outcome of looking at whatever config already
// says about SSL - shared by sslEnableWizard (to decide whether to offer
// keep-or-replace) and sslStatus (to print it).
type sslCertificateStatus struct {
	// Human label of where the certificate comes from, e.g. "Let's Encrypt
	// (domains: example.com)" - empty if SSL isn't configured at all.
	source string
	// nil if configured but not yet resolvable (e.g. a Let's Encrypt
	// certificate that hasn't been issued yet - see pendingReason).
	details       *fireback.CertificateDetails
	pendingReason string
	inspectionErr error
}

func currentSSLStatus() sslCertificateStatus {
	switch {
	case config.SslProvider == "letsencrypt" && strings.TrimSpace(config.AcmeDomains) != "":
		domains := fireback.AcmeDomainList(config.AcmeDomains)
		status := sslCertificateStatus{source: fmt.Sprintf("Let's Encrypt (domains: %s)", strings.Join(domains, ", "))}
		if len(domains) == 0 {
			return status
		}
		details, err := fireback.InspectAcmeCache(config.AcmeCacheDir, domains[0])
		if err == fireback.ErrNoCertificateYet {
			status.pendingReason = "configured, but not requested yet - Let's Encrypt issues it automatically on the server's first HTTPS handshake for this domain"
		} else if err != nil {
			status.inspectionErr = err
		} else {
			status.details = details
		}
		return status

	case config.CertFile != "" && config.KeyFile != "":
		label := "Manual certificate"
		if config.SslProvider == "self-signed" {
			label = "Self-signed certificate"
		}
		status := sslCertificateStatus{source: fmt.Sprintf("%s (%s)", label, config.CertFile)}
		details, err := fireback.InspectCertFile(config.CertFile)
		if err != nil {
			status.inspectionErr = err
		} else {
			status.details = details
		}
		return status

	default:
		return sslCertificateStatus{}
	}
}

func printCertificateDetails(d *fireback.CertificateDetails) {
	fmt.Println("  Subject:      ", d.Subject)
	fmt.Println("  Issuer:       ", d.Issuer)
	fmt.Println("  Domains:      ", strings.Join(d.DNSNames, ", "))
	fmt.Println("  Valid from:   ", d.NotBefore.Format("2006-01-02"))
	fmt.Println("  Valid until:  ", d.NotAfter.Format("2006-01-02"))
	if d.SelfSigned {
		fmt.Println("  Self-signed:   yes (not trusted by browsers)")
	}
	switch {
	case d.Expired():
		fmt.Printf("  Status:        EXPIRED %d days ago\n", -d.DaysRemaining())
	case d.DaysRemaining() <= 14:
		fmt.Printf("  Status:        expires soon, in %d days\n", d.DaysRemaining())
	default:
		fmt.Printf("  Status:        valid, %d days remaining\n", d.DaysRemaining())
	}
}

func sslStatus() {
	if !config.UseSSL {
		fmt.Println("SSL is not enabled. Run `fireback ssl enable` to configure it.")
		return
	}

	status := currentSSLStatus()

	if status.source == "" {
		fmt.Println("useSSL is on, but no certificate is configured - the http server falls back to an ephemeral self-signed certificate, regenerated on every restart. Run `fireback ssl enable` to fix this.")
		return
	}

	fmt.Println(status.source)
	switch {
	case status.pendingReason != "":
		fmt.Println("  Status:       ", status.pendingReason)
	case status.inspectionErr != nil:
		fmt.Println("  Could not read this certificate:", status.inspectionErr)
	default:
		printCertificateDetails(status.details)
	}
}

func sslEnableWizard() {
	existing := currentSSLStatus()

	if existing.source != "" {
		fmt.Println("An SSL certificate is already configured:")
		fmt.Println(existing.source)
		switch {
		case existing.pendingReason != "":
			fmt.Println("  Status:       ", existing.pendingReason)
		case existing.inspectionErr != nil:
			fmt.Println("  Could not read this certificate:", existing.inspectionErr)
		case existing.details != nil:
			printCertificateDetails(existing.details)
		}

		if AskForSelect("What do you want to do?", []string{sslPickKeep, sslPickReplace}) == sslPickKeep {
			config.UseSSL = true
			config.Save(".env")
			fmt.Println("Kept the existing configuration. useSSL is on.")
			return
		}
	}

	pick := AskForSelect("How do you want to obtain the certificate?", []string{sslPickLetsEncrypt, sslPickManual, sslPickSelfSigned})

	switch pick {
	case sslPickLetsEncrypt:
		enableLetsEncrypt()
	case sslPickManual:
		enableManualCert()
	case sslPickSelfSigned:
		enableSelfSignedCert()
	}

	config.UseSSL = true
	config.Save(".env")
}

func enableLetsEncrypt() {
	domains := AskForInput("Domain names for the certificate (comma separated - each must already resolve to this server)", config.AcmeDomains)
	email, _, _ := askForInputOptional("Contact email for Let's Encrypt expiry notices (optional)", config.AcmeEmail)
	cacheDir := AskForInput("Directory to cache the ACME account key and certificates in (must persist across restarts)", orDefault(config.AcmeCacheDir, "./.fireback-acme-cache"))

	config.SslProvider = "letsencrypt"
	config.AcmeDomains = domains
	config.AcmeEmail = email
	config.AcmeCacheDir = cacheDir
	config.CertFile = ""
	config.KeyFile = ""

	fmt.Println()
	fmt.Println("Saved. No certificate has been requested yet - Let's Encrypt issues (and later renews) it")
	fmt.Println("automatically the first time the running server gets a real HTTPS request for one of these")
	fmt.Println("domains, so make sure both port 80 (HTTP-01 challenge) and port 443 are reachable from the")
	fmt.Println("internet for them before starting `fireback http start`. Check progress any time with `fireback ssl status`.")
}

func enableManualCert() {
	certPath := AskForInput("Path to the certificate file (PEM)", config.CertFile)
	keyPath := AskForInput("Path to the private key file (PEM)", config.KeyFile)

	if details, err := fireback.InspectCertFile(certPath); err != nil {
		fmt.Println("Warning: could not read/parse that certificate yet:", err)
	} else {
		fmt.Println()
		printCertificateDetails(details)
	}

	config.SslProvider = "manual"
	config.CertFile = certPath
	config.KeyFile = keyPath
}

func enableSelfSignedCert() {
	domainsInput, _, _ := askForInputOptional("Domain names for the certificate (comma separated, optional - defaults to localhost)", "")
	outDir := AskForInput("Directory to save the generated certificate + key", "./.fireback-ssl")

	certPEM, keyPEM, err := fireback.GenerateSelfSignedCertPEM(fireback.AcmeDomainList(domainsInput))
	if err != nil {
		fmt.Println("Failed to generate self-signed certificate:", err)
		return
	}

	if err := os.MkdirAll(outDir, 0700); err != nil {
		fmt.Println("Failed to create", outDir, ":", err)
		return
	}

	certPath := filepath.Join(outDir, "self-signed.crt")
	keyPath := filepath.Join(outDir, "self-signed.key")

	if err := os.WriteFile(certPath, certPEM, 0644); err != nil {
		fmt.Println("Failed to write", certPath, ":", err)
		return
	}
	if err := os.WriteFile(keyPath, keyPEM, 0600); err != nil {
		fmt.Println("Failed to write", keyPath, ":", err)
		return
	}

	config.SslProvider = "self-signed"
	config.CertFile = certPath
	config.KeyFile = keyPath

	fmt.Println()
	fmt.Println("Generated a self-signed certificate, valid for 365 days, at", certPath)
	fmt.Println("Browsers will show a trust warning for it - it's meant for local/dev use only.")
}
