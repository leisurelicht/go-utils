package utilize

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func HttpCli(method, url string, headers map[string]string, body string, timeout time.Duration, insecureSkipVerify bool, caCertPath string) (response string, err error) {
	// Create the HTTP request
	var (
		request *http.Request
	)

	if len(body) == 0 {
		request, err = http.NewRequest(method, url, nil)
	} else {
		request, err = http.NewRequest(method, url, strings.NewReader(body))
	}
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	for header, value := range headers {
		request.Header.Set(header, value)
	}

	// Configure TLS
	tlsConfig := &tls.Config{InsecureSkipVerify: insecureSkipVerify}
	if !insecureSkipVerify && caCertPath != "" {
		// Load CA certificates
		caCert, err := os.ReadFile(caCertPath)
		if err != nil {
			return "", fmt.Errorf("failed to load CA certificate: %w", err)
		}
		certPool := x509.NewCertPool()
		if !certPool.AppendCertsFromPEM(caCert) {
			return "", fmt.Errorf("failed to append CA certificate")
		}
		tlsConfig.RootCAs = certPool
	}

	// Configure HTTP client
	client := &http.Client{
		Timeout: timeout,
		Transport: &http.Transport{
			DisableKeepAlives: true,
			TLSClientConfig:   tlsConfig,
		},
	}

	// Send the HTTP request
	resp, err := client.Do(request)
	if err != nil {
		return "", fmt.Errorf("failed to execute HTTP request: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	// Remove UTF-8 BOM if present
	respBody = bytes.TrimPrefix(respBody, []byte("\xef\xbb\xbf"))

	return string(respBody), nil
}
