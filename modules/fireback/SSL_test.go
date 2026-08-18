package fireback

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"io"
	"math/big"
	"net/http"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

func TestAcmeDomainList(t *testing.T) {
	got := AcmeDomainList(" example.com, www.example.com ,, sub.example.com")
	want := []string{"example.com", "www.example.com", "sub.example.com"}
	if len(got) != len(want) {
		t.Fatalf("got %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("got %v, want %v", got, want)
		}
	}
}

func TestInspectCertFileRoundTrip(t *testing.T) {
	certPEM, _, err := GenerateSelfSignedCertPEM([]string{"example.test"})
	if err != nil {
		t.Fatalf("GenerateSelfSignedCertPEM: %v", err)
	}

	dir := t.TempDir()
	certPath := filepath.Join(dir, "cert.pem")
	if err := os.WriteFile(certPath, certPEM, 0644); err != nil {
		t.Fatalf("WriteFile: %v", err)
	}

	details, err := InspectCertFile(certPath)
	if err != nil {
		t.Fatalf("InspectCertFile: %v", err)
	}

	if !details.SelfSigned {
		t.Errorf("expected SelfSigned to be true for a self-signed cert")
	}
	found := false
	for _, d := range details.DNSNames {
		if d == "example.test" {
			found = true
		}
	}
	if !found {
		t.Errorf("expected DNSNames to include example.test, got %v", details.DNSNames)
	}
	if details.DaysRemaining() <= 0 {
		t.Errorf("expected a freshly generated cert to have days remaining, got %d", details.DaysRemaining())
	}
	if details.Expired() {
		t.Errorf("expected a freshly generated cert to not be expired")
	}
}

func TestInspectAcmeCacheMissAndHit(t *testing.T) {
	dir := t.TempDir()

	if _, err := InspectAcmeCache(dir, "example.test"); err != ErrNoCertificateYet {
		t.Fatalf("expected ErrNoCertificateYet for an empty cache dir, got %v", err)
	}

	// Mimic autocert.Manager.cachePut's on-disk format: PEM private key
	// followed by one or more PEM CERTIFICATE blocks, stored under a file
	// named after the domain (see InspectAcmeCache's doc comment).
	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("GenerateKey: %v", err)
	}
	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject:      pkix.Name{CommonName: "example.test"},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(90 * 24 * time.Hour),
		DNSNames:     []string{"example.test"},
	}
	der, err := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)
	if err != nil {
		t.Fatalf("CreateCertificate: %v", err)
	}

	var buf bytes.Buffer
	pem.Encode(&buf, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(priv)})
	pem.Encode(&buf, &pem.Block{Type: "CERTIFICATE", Bytes: der})

	if err := os.WriteFile(filepath.Join(dir, "example.test"), buf.Bytes(), 0600); err != nil {
		t.Fatalf("WriteFile: %v", err)
	}

	details, err := InspectAcmeCache(dir, "example.test")
	if err != nil {
		t.Fatalf("InspectAcmeCache: %v", err)
	}
	if details.Subject != "example.test" {
		t.Errorf("expected Subject example.test, got %q", details.Subject)
	}
	if details.DaysRemaining() <= 0 {
		t.Errorf("expected days remaining > 0, got %d", details.DaysRemaining())
	}
}

// TestSSLEnableSelfSignedEndToEnd exercises the whole `ssl enable` ->
// `.env` -> real server flow for the one certificate source that doesn't
// need a real domain or internet-reachable port 80 (see
// modules/fireback/clitools/SSLCli_test.go's
// TestSSLEnableLetsEncrypt_NeedsRealDomain for why Let's Encrypt itself
// isn't covered end-to-end here):
//
//  1. generates + persists a self-signed cert, exactly like `ssl enable`'s
//     self-signed option (clitools/SSLCli.go's enableSelfSignedCert)
//  2. saves it into a real .env file via Config.Save, exactly like the
//     wizard does
//  3. reloads that .env with godotenv.Read - a plain file parse, independent
//     of the real process environment, so this test can't be made to pass
//     or fail by whatever's already exported in the shell it happens to run
//     in - proving the values actually round-trip through disk
//  4. starts a real HTTPS server using exactly the cert/key files the
//     reloaded config points at, the same pair CreateHttpServer loads via
//     ListenAndServeTLS(config.CertFile, config.KeyFile) for sslProvider
//     "manual"/"self-signed" (see HttpServer.go) - bound to an OS-assigned
//     port rather than CreateHttpServer's hardcoded :443, so this test
//     doesn't need root or a free privileged port - and has a real client
//     complete a TLS handshake and fetch a response over it.
func TestSSLEnableSelfSignedEndToEnd(t *testing.T) {
	dir := t.TempDir()
	envPath := filepath.Join(dir, ".env")
	certPath := filepath.Join(dir, "self-signed.crt")
	keyPath := filepath.Join(dir, "self-signed.key")

	certPEM, keyPEM, err := GenerateSelfSignedCertPEM([]string{"selfsigned.test"})
	if err != nil {
		t.Fatalf("GenerateSelfSignedCertPEM: %v", err)
	}
	if err := os.WriteFile(certPath, certPEM, 0644); err != nil {
		t.Fatalf("WriteFile cert: %v", err)
	}
	if err := os.WriteFile(keyPath, keyPEM, 0600); err != nil {
		t.Fatalf("WriteFile key: %v", err)
	}

	saved := Config{
		SslProvider: "self-signed",
		CertFile:    certPath,
		KeyFile:     keyPath,
		UseSSL:      true,
	}
	if err := saved.Save(envPath); err != nil {
		t.Fatalf("Save: %v", err)
	}

	reloaded, err := godotenv.Read(envPath)
	if err != nil {
		t.Fatalf("reading generated .env: %v", err)
	}
	if reloaded["SSL_PROVIDER"] != "self-signed" {
		t.Errorf("SSL_PROVIDER = %q, want %q", reloaded["SSL_PROVIDER"], "self-signed")
	}
	if reloaded["CERT_FILE"] != certPath {
		t.Errorf("CERT_FILE = %q, want %q", reloaded["CERT_FILE"], certPath)
	}
	if reloaded["KEY_FILE"] != keyPath {
		t.Errorf("KEY_FILE = %q, want %q", reloaded["KEY_FILE"], keyPath)
	}
	if reloaded["USE_SSL"] != "true" {
		t.Errorf("USE_SSL = %q, want %q", reloaded["USE_SSL"], "true")
	}

	cert, err := tls.LoadX509KeyPair(reloaded["CERT_FILE"], reloaded["KEY_FILE"])
	if err != nil {
		t.Fatalf("tls.LoadX509KeyPair on the reloaded paths: %v", err)
	}

	listener, err := tls.Listen("tcp", "127.0.0.1:0", &tls.Config{Certificates: []tls.Certificate{cert}})
	if err != nil {
		t.Fatalf("tls.Listen: %v", err)
	}
	defer listener.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("pong"))
	})
	srv := &http.Server{Handler: mux}
	go srv.Serve(listener)
	defer srv.Close()

	client := &http.Client{
		Transport: &http.Transport{
			// The certificate is self-signed by design (see
			// enableSelfSignedCert's own doc comment) - no CA in any real
			// trust store issued it, so verification has to be skipped the
			// same way a browser's "proceed anyway" click would.
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	resp, err := client.Get("https://" + listener.Addr().String() + "/ping")
	if err != nil {
		t.Fatalf("https GET failed: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("reading response body: %v", err)
	}
	if string(body) != "pong" {
		t.Errorf("response body = %q, want %q", body, "pong")
	}

	if resp.TLS == nil || len(resp.TLS.PeerCertificates) == 0 {
		t.Fatalf("expected the response to report the server's TLS certificate")
	}
	served := resp.TLS.PeerCertificates[0]
	if len(served.DNSNames) == 0 || served.DNSNames[0] != "selfsigned.test" {
		t.Errorf("served certificate DNSNames = %v, want it to include %q - the http server is not using the certificate ssl enable generated", served.DNSNames, "selfsigned.test")
	}
}

func TestNewAcmeManagerHostPolicy(t *testing.T) {
	mgr := NewAcmeManager("example.com, other.example.com", "ops@example.com", "")
	if err := mgr.HostPolicy(nil, "example.com"); err != nil {
		t.Errorf("expected example.com to be allowed: %v", err)
	}
	if err := mgr.HostPolicy(nil, "not-allowed.example.com"); err == nil {
		t.Errorf("expected not-allowed.example.com to be rejected")
	}
}
