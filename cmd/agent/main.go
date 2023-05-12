package main

import (
	"time"

	"github.com/andreyramos/go-metrics/internal/agent"
)

func main() {

	cfg := agent.Config{
		PullInterval:   2 * time.Second,
		ReportInterval: 10 * time.Second,
		Address:        "127.0.0.1",
		Port:           "8080",
	}

	agent := agent.New(cfg)
	agent.Run()
}
