package main

import (
	"github.com/jeanmorais/beatport/domain/beatport"
	"github.com/jeanmorais/beatport/infrastructure/client"
	"github.com/jeanmorais/beatport/infrastructure/server/http"
	"github.com/jeanmorais/beatport/pkg/http_client"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	beatPortURL = "https://www.beatport.com"
)

func main() {

	beatPortClient := client.NewBeatPortClient(httpclient.NewHTTPClient(60*time.Second), beatPortURL)

	beatPortService := beatport.NewService(beatPortClient)

	handler := http.NewHandler(beatPortService)

	server := http.New("8080", handler)
	server.ListenAndServe()

	// Graceful shutdown
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGINT)
	<-stopChan
	server.Shutdown()
}
