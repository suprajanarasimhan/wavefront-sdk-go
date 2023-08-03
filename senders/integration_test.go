package senders

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestEndToEnd(t *testing.T) {
	testServer := startTestServer()
	defer testServer.Close()
	sender, err := NewSender(testServer.URL)
	require.NoError(t, err)
	require.NoError(t, sender.SendMetric("my metric", 20, 0, "localhost", nil))
	require.NoError(t, sender.Flush())

	assert.Equal(t, 1, len(testServer.MetricLines))
	assert.Equal(t, "\"my-metric\" 20 source=\"localhost\"", testServer.MetricLines[0])
	assert.Equal(t, "/report?f=wavefront", testServer.LastRequestURL)
}

func TestEndToEndWithPath(t *testing.T) {
	testServer := startTestServer()
	defer testServer.Close()
	sender, err := NewSender(testServer.URL + "/test-path")
	require.NoError(t, err)
	require.NoError(t, sender.SendMetric("my metric", 20, 0, "localhost", nil))
	require.NoError(t, sender.Flush())

	assert.Equal(t, 1, len(testServer.MetricLines))
	assert.Equal(t, "\"my-metric\" 20 source=\"localhost\"", testServer.MetricLines[0])
	assert.Equal(t, "/test-path/report?f=wavefront", testServer.LastRequestURL)
}

func TestTLSEndToEnd(t *testing.T) {
	testServer := startTLSTestServer()
	defer testServer.Close()
	testServer.httpServer.Client()
	tlsConfig := testServer.TLSConfig()

	sender, err := NewSender(testServer.URL, TLSConfigOptions(tlsConfig))
	require.NoError(t, err)
	require.NoError(t, sender.SendMetric("my metric", 20, 0, "localhost", nil))
	require.NoError(t, sender.Flush())

	assert.Equal(t, 1, len(testServer.MetricLines))
	assert.Equal(t, "\"my-metric\" 20 source=\"localhost\"", testServer.MetricLines[0])
}

func TestEndToEndWithInternalMetrics(t *testing.T) {
	testServer := startTestServer()
	defer testServer.Close()
	sender, err := NewSender(testServer.URL)
	require.NoError(t, err)
	require.NoError(t, sender.SendMetric("my metric", 20, 0, "localhost", nil))
	require.NoError(t, sender.Flush())
	time.Sleep(2 * time.Second)

	assert.Equal(t, 12, len(testServer.MetricLines))
	// Assert that testServer.MetricLines contains a string with a pattern that matches ~sdk.go.core.sender.proxy.points.valid
	// Options: Use a for loop, or create a matcher
	// assert.(t, "\"~sdk.go.core.sender.proxy.points.valid\" 3 source=\"localhost\"", testServer.MetricLines)
	assert.Equal(t, "\"my-metric\" 20 source=\"localhost\"", testServer.MetricLines[0])
	assert.Equal(t, "/report?f=wavefront", testServer.LastRequestURL)
}

func TestEndToEndWithInternalMetricsDisabled(t *testing.T) {
	testServer := startTestServer()
	defer testServer.Close()
	sender, err := NewSender(testServer.URL, InternalMetricsEnabled(false))
	require.NoError(t, err)
	require.NoError(t, sender.SendMetric("my metric", 20, 0, "localhost", nil))
	require.NoError(t, sender.Flush())

	assert.Equal(t, 1, len(testServer.MetricLines))
	assert.Equal(t, "\"my-metric\" 20 source=\"localhost\"", testServer.MetricLines[0])
	assert.Equal(t, "/report?f=wavefront", testServer.LastRequestURL)
}
