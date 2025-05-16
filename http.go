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

const (
	HeaderKeyContentType    = "Content-type"
	HeaderKeyAuthorization  = "Authorization"
	HeaderKeyConnect        = "Connection"
	HeaderKeyUserAgent      = "User-Agent"
	HeaderKeyAccept         = "Accept"
	HeaderKeyAcceptCharset  = "Accept-Charset"
	HeaderKeyAcceptEncoding = "Accept-Encoding"
	HeaderKeyAcceptLanguage = "Accept-Language"
	HeaderKeyAcceptRanges   = "Accept-Ranges"
	HeaderKeyCacheControl   = "Cache-Control"
	HeaderKeyCookie         = "Cookie"
	HeaderKeyExpect         = "Expect"
	HeaderKeyForwarded      = "Forwarded"
	HeaderKeyFrom           = "From"

	HeaderValueApplicationJson        = "application/json"
	HeaderValueApplicationXml         = "application/xml"
	HeaderValueApplicationForm        = "application/x-www-form-urlencoded"
	HeaderValueApplicationOctetStream = "application/octet-stream"
	HeaderValueApplicationPdf         = "application/pdf"
	HeaderValueApplicationZip         = "application/zip"
	HeaderValueApplicationGzip        = "application/gzip"
	HeaderValueApplicationXhtml       = "application/xhtml+xml"
	HeaderValueApplicationRss         = "application/rss+xml"
	HeaderValueConnectionKeepAlive    = "keep-alive"
	HeaderValueConnectionClose        = "close"
	HeaderValueConnectionUpgrade      = "upgrade"
	HeaderValueConnectionHttp2        = "HTTP/2"
	HeaderValueConnectionWebsocket    = "websocket"
	HeaderValueConnectionSpdy         = "spdy"
	HeaderValueAcceptJson             = "application/json"
)

func HttpCli(method, url string, headers http.Header, body string, timeout time.Duration, insecureSkipVerify bool, caCertPath string) (response []byte, err error) {
	// Create the HTTP request
	var request *http.Request

	if len(body) == 0 {
		request, err = http.NewRequest(method, url, nil)
	} else {
		request, err = http.NewRequest(method, url, strings.NewReader(body))
	}
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	request.Header = headers

	// Configure TLS
	tlsConfig := &tls.Config{InsecureSkipVerify: insecureSkipVerify}
	if !insecureSkipVerify && caCertPath != "" {
		// Load CA certificates
		caCert, err := os.ReadFile(caCertPath)
		if err != nil {
			return nil, fmt.Errorf("failed to load CA certificate: %w", err)
		}
		certPool := x509.NewCertPool()
		if !certPool.AppendCertsFromPEM(caCert) {
			return nil, fmt.Errorf("failed to append CA certificate")
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
		return nil, fmt.Errorf("failed to execute HTTP request: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	// Read the response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Remove UTF-8 BOM if present
	respBody = bytes.TrimPrefix(respBody, []byte("\xef\xbb\xbf"))

	return respBody, nil
}

func HttpGet(url string, headers http.Header, timeout time.Duration) (response []byte, err error) {
	return HttpCli(http.MethodGet, url, headers, "", timeout, true, "")
}

func HttpPost(url string, headers http.Header, body string, timeout time.Duration) (response []byte, err error) {
	return HttpCli(http.MethodPost, url, headers, body, timeout, true, "")
}

func HttpPut(url string, headers http.Header, body string, timeout time.Duration) (response []byte, err error) {
	return HttpCli(http.MethodPut, url, headers, body, timeout, true, "")
}

func HttpDelete(url string, headers http.Header, body string, timeout time.Duration) (response []byte, err error) {
	return HttpCli(http.MethodDelete, url, headers, body, timeout, true, "")
}

func HttpPatch(url string, headers http.Header, body string, timeout time.Duration) (response []byte, err error) {
	return HttpCli(http.MethodPatch, url, headers, body, timeout, true, "")
}

func HttpHead(url string, headers http.Header, timeout time.Duration) (response []byte, err error) {
	return HttpCli(http.MethodHead, url, headers, "", timeout, true, "")
}

func HttpOptions(url string, headers http.Header, timeout time.Duration) (response []byte, err error) {
	return HttpCli(http.MethodOptions, url, headers, "", timeout, true, "")
}

func HttpsGet(url string, headers http.Header, timeout time.Duration, caCertPath string) (response []byte, err error) {
	return HttpCli(http.MethodGet, url, headers, "", timeout, false, caCertPath)
}

func HttpsPost(url string, headers http.Header, body string, timeout time.Duration, caCertPath string) (response []byte, err error) {
	return HttpCli(http.MethodPost, url, headers, body, timeout, false, caCertPath)
}

func HttpsPut(url string, headers http.Header, body string, timeout time.Duration, caCertPath string) (response []byte, err error) {
	return HttpCli(http.MethodPut, url, headers, body, timeout, false, caCertPath)
}

func HttpsDelete(url string, headers http.Header, body string, timeout time.Duration, caCertPath string) (response []byte, err error) {
	return HttpCli(http.MethodDelete, url, headers, body, timeout, false, caCertPath)
}

func HttpsPatch(url string, headers http.Header, body string, timeout time.Duration, caCertPath string) (response []byte, err error) {
	return HttpCli(http.MethodPatch, url, headers, body, timeout, false, caCertPath)
}

func HttpsHead(url string, headers http.Header, timeout time.Duration, caCertPath string) (response []byte, err error) {
	return HttpCli(http.MethodHead, url, headers, "", timeout, false, caCertPath)
}

func HttpsOptions(url string, headers http.Header, timeout time.Duration, caCertPath string) (response []byte, err error) {
	return HttpCli(http.MethodOptions, url, headers, "", timeout, false, caCertPath)
}
