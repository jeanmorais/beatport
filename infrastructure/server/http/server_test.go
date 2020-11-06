package http_test

import (
	"github.com/jeanmorais/beatport/infrastructure/server/http"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestServer_ListenAndServe(t *testing.T) {
	server := http.New("9000", nil)
	server.ListenAndServe()

	stopChan := make(chan bool)

	go func() {
		time.Sleep(1 * time.Second)
		stopChan <- true
	}()

	var result bool
	result = <-stopChan
	server.Shutdown()

	assert.True(t, result)
}
