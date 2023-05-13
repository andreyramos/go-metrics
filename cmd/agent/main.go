package main

import (
	"flag"
	"time"

	"github.com/andreyramos/go-metrics/internal/agent"
)

func main() {

	var flagAddr string
	var flagPullInt int
	var flagRepInt int

	flag.StringVar(&flagAddr, "a", "127.0.0.1:8080", "address and port to run agent")
	flag.IntVar(&flagPullInt, "p", 2, "frequency of polling metrics from the runtime package")
	flag.IntVar(&flagRepInt, "r", 10, "frequency of sending metrics to the server")

	cfg := agent.Config{
		PullInterval:   time.Duration(flagPullInt) * time.Second,
		ReportInterval: time.Duration(flagRepInt) * time.Second,
		Address:        flagAddr,
	}

	agent := agent.New(cfg)
	agent.Run()
}
