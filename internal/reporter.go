package internal

import (
	"bytes"
	"compress/gzip"
	"crypto/tls"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// The implementation of a Reporter that reports points directly to a Wavefront server.
type reporter struct {
	serverURL string
	token     string
	client    *http.Client
}

// NewReporter create a metrics Reporter
func NewReporter(server string, token string, timeout time.Duration, tlsConfigOptions *tls.Config) Reporter {

	var client *http.Client
	if tlsConfigOptions == nil {
		client = &http.Client{Timeout: timeout}
	} else {
		transport := &http.Transport{TLSClientConfig: tlsConfigOptions}
		client = &http.Client{Timeout: timeout, Transport: transport}
	}

	return &reporter{
		serverURL: server,
		token:     token,
		client:    client,
	}
}

// Report creates and sends a POST to the reportEndpoint with the given pointLines
func (reporter reporter) Report(format string, pointLines string) (*http.Response, error) {
	if format == "" || pointLines == "" {
		return nil, formatError
	}

	requestBody, err := linesToGzippedBytes(pointLines)
	if err != nil {
		return &http.Response{}, err
	}

	req, err := reporter.buildRequest(format, requestBody)
	if err != nil {
		return &http.Response{}, err
	}

	return reporter.execute(req)
}

func linesToGzippedBytes(pointLines string) ([]byte, error) {
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)
	_, err := zw.Write([]byte(pointLines))
	if err != nil {
		zw.Close()
		return nil, err
	}
	if err = zw.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), err
}

func (reporter reporter) buildRequest(format string, body []byte) (*http.Request, error) {
	apiURL := reporter.serverURL + reportEndpoint
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set(contentType, octetStream)
	req.Header.Set(contentEncoding, gzipFormat)
	if len(reporter.token) > 0 {
		req.Header.Set(authzHeader, bearer+reporter.token)
	}

	q := req.URL.Query()
	q.Add(formatKey, format)
	req.URL.RawQuery = q.Encode()
	return req, nil
}

func (reporter reporter) ReportEvent(event string) (*http.Response, error) {
	if event == "" {
		return nil, formatError
	}

	apiURL := reporter.serverURL + eventEndpoint
	req, err := http.NewRequest("POST", apiURL, strings.NewReader(event))
	if err != nil {
		return &http.Response{}, err
	}

	req.Header.Set(contentType, applicationJSON)
	if len(reporter.token) > 0 {
		req.Header.Set(contentEncoding, gzipFormat)
		req.Header.Set(authzHeader, bearer+reporter.token)
	}

	return reporter.execute(req)
}

func (reporter reporter) execute(req *http.Request) (*http.Response, error) {
	resp, err := reporter.client.Do(req)
	if err != nil {
		return resp, err
	}
	io.Copy(ioutil.Discard, resp.Body)
	defer resp.Body.Close()
	return resp, nil
}
