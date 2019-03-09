package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestParseOutputPluginIfttt(t *testing.T) {
	input := `{"type": "ifttt", "value": {"event_name": "test_event", "api_key": "test_key"}}`
	outputPlugin, err := parseOutputPlugin(input)

	assert.Nil(t, err)
	ifttt, ok := outputPlugin.(Ifttt)
	assert.Equal(t, ok, true)
	assert.Equal(t, ifttt.eventName, "test_event")
	assert.Equal(t, ifttt.apiKey, "test_key")
}

func TestParseOutputPluginHttp(t *testing.T) {
	input := `{"type": "http", "value": {"endpoint": "http://example.com"}}`
	outputPlugin, err := parseOutputPlugin(input)

	assert.Nil(t, err)
	http, ok := outputPlugin.(Http)
	assert.Equal(t, ok, true)
	assert.Equal(t, http.endpoint, "http://example.com")
}

func TestParseOutputPluginUnknowType(t *testing.T) {
	input := `{"type": "foobar", "value": {"endpoint": "http://example.com"}}`
	outputPlugin, err := parseOutputPlugin(input)
	assert.NotNil(t, err)
	assert.Nil(t, outputPlugin)
}
