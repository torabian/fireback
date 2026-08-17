//go:build !wasm

package clitools

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/torabian/fireback/modules/fireback"
)

// withSSLConfig temporarily overrides every SSL-related field on the shared
// config - same reasoning as withConfig in DatabaseWizard_test.go: config is
// the same package-level pointer production code (currentSSLStatus,
// sslEnableWizard, ...) reads, so tests must not leak changes into each
// other.
func withSSLConfig(t *testing.T, mutate func(), fn func()) {
	t.Helper()
	prev := *config
	mutate()
	defer func() { *config = prev }()
	fn()
}

func TestCurrentSSLStatus_NothingConfigured(t *testing.T) {
	withSSLConfig(t, func() {
		config.SslProvider = ""
		config.AcmeDomains = ""
		config.CertFile = ""
		config.KeyFile = ""
	}, func() {
		status := currentSSLStatus()
		if status.source != "" {
			t.Errorf("source = %q, want empty when nothing is configured", status.source)
		}
	})
}

func TestCurrentSSLStatus_LetsEncryptPending(t *testing.T) {
	dir := t.TempDir()

	withSSLConfig(t, func() {
		config.SslProvider = "letsencrypt"
		config.AcmeDomains = "example.com"
		config.AcmeCacheDir = dir // empty - no certificate cached yet
		config.CertFile = ""
		config.KeyFile = ""
	}, func() {
		status := currentSSLStatus()
		if status.source == "" {
			t.Fatalf("expected a non-empty source once sslProvider/acmeDomains are set")
		}
		if status.pendingReason == "" {
			t.Errorf("expected pendingReason to explain the certificate hasn't been issued yet, got status=%+v", status)
		}
		if status.details != nil {
			t.Errorf("expected no details for a certificate that hasn't been issued, got %+v", status.details)
		}
	})
}

func TestCurrentSSLStatus_LetsEncryptIssued(t *testing.T) {
	dir := t.TempDir()

	// Mimic autocert's own on-disk cache format (private key PEM followed by
	// a CERTIFICATE PEM block, stored under a file named after the domain -
	// see fireback.InspectAcmeCache's doc comment) using the same
	// certificate-generation helper `ssl enable`'s self-signed option uses,
	// rather than hand-rolling x509 here too.
	certPEM, keyPEM, err := fireback.GenerateSelfSignedCertPEM([]string{"example.com"})
	if err != nil {
		t.Fatalf("GenerateSelfSignedCertPEM: %v", err)
	}
	if err := os.WriteFile(filepath.Join(dir, "example.com"), append(keyPEM, certPEM...), 0600); err != nil {
		t.Fatalf("WriteFile: %v", err)
	}

	withSSLConfig(t, func() {
		config.SslProvider = "letsencrypt"
		config.AcmeDomains = "example.com"
		config.AcmeCacheDir = dir
		config.CertFile = ""
		config.KeyFile = ""
	}, func() {
		status := currentSSLStatus()
		if status.pendingReason != "" {
			t.Errorf("expected no pendingReason once a certificate is cached, got %q", status.pendingReason)
		}
		if status.inspectionErr != nil {
			t.Fatalf("unexpected inspection error: %v", status.inspectionErr)
		}
		if status.details == nil {
			t.Fatalf("expected details for an issued certificate")
		}
		if status.details.DaysRemaining() <= 0 {
			t.Errorf("expected a freshly generated certificate to have days remaining, got %d", status.details.DaysRemaining())
		}
	})
}

func TestCurrentSSLStatus_ManualAndSelfSignedLabels(t *testing.T) {
	dir := t.TempDir()
	certPEM, keyPEM, err := fireback.GenerateSelfSignedCertPEM([]string{"internal.test"})
	if err != nil {
		t.Fatalf("GenerateSelfSignedCertPEM: %v", err)
	}
	certPath := filepath.Join(dir, "cert.pem")
	keyPath := filepath.Join(dir, "key.pem")
	if err := os.WriteFile(certPath, certPEM, 0644); err != nil {
		t.Fatalf("WriteFile cert: %v", err)
	}
	if err := os.WriteFile(keyPath, keyPEM, 0600); err != nil {
		t.Fatalf("WriteFile key: %v", err)
	}

	cases := []struct {
		provider   string
		wantSubstr string
	}{
		{"manual", "Manual certificate"},
		{"self-signed", "Self-signed certificate"},
	}

	for _, c := range cases {
		withSSLConfig(t, func() {
			config.SslProvider = c.provider
			config.CertFile = certPath
			config.KeyFile = keyPath
			config.AcmeDomains = ""
		}, func() {
			status := currentSSLStatus()
			if status.inspectionErr != nil {
				t.Fatalf("provider %q: unexpected inspection error: %v", c.provider, status.inspectionErr)
			}
			if status.details == nil {
				t.Fatalf("provider %q: expected details", c.provider)
			}
			if !status.details.SelfSigned {
				t.Errorf("provider %q: expected the generated certificate to report SelfSigned=true", c.provider)
			}
			found := false
			for _, d := range status.details.DNSNames {
				if d == "internal.test" {
					found = true
				}
			}
			if !found {
				t.Errorf("provider %q: expected DNSNames to include internal.test, got %v", c.provider, status.details.DNSNames)
			}
			if !strings.Contains(status.source, c.wantSubstr) {
				t.Errorf("provider %q: source = %q, want it to contain %q", c.provider, status.source, c.wantSubstr)
			}
		})
	}
}

func TestSSLCliCommandStructure(t *testing.T) {
	cmd := sslCliCommand()
	if cmd.Name != "ssl" {
		t.Fatalf("Name = %q, want %q", cmd.Name, "ssl")
	}

	names := map[string]bool{}
	for _, sub := range cmd.Commands {
		names[sub.Name] = true
	}
	if !names["enable"] {
		t.Errorf("expected an %q subcommand, got %v", "enable", names)
	}
	if !names["status"] {
		t.Errorf("expected a %q subcommand, got %v", "status", names)
	}
}

// TestSSLEnableLetsEncrypt_NeedsRealDomain documents (rather than tests) why
// there is no automated end-to-end test that actually requests a real
// certificate through enableLetsEncrypt/NewAcmeManager: ACME's HTTP-01
// challenge requires the domain to already resolve to this machine with
// port 80 reachable from the public internet, which no CI/sandbox
// environment here provides. What's automatable without a real domain is
// already covered elsewhere:
//   - the domain allow-list/HostPolicy wiring: fireback.TestNewAcmeManagerHostPolicy
//   - reading back whatever autocert has cached: fireback.TestInspectAcmeCacheMissAndHit
//   - the "not issued yet" status this package reports before a domain
//     resolves: TestCurrentSSLStatus_LetsEncryptPending above
//
// (both fireback.Test* live in modules/fireback/SSL_test.go)
func TestSSLEnableLetsEncrypt_NeedsRealDomain(t *testing.T) {
	t.Skip("issuing a real Let's Encrypt certificate needs a domain that already resolves to this machine with port 80 reachable from the public internet - not available here. See this test's doc comment for what is covered instead.")
}
