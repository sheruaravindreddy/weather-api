package httpclient_test

import (
	"testing"
	"time"
	"weather-api/pkg/httpclient"

	"github.com/stretchr/testify/assert"
)

func TestNewHTTPClient(t *testing.T) {
	client := httpclient.NewHTTPClient()
	assert.NotNil(t, client)
	assert.Equal(t, client.Timeout, time.Second*10)
}