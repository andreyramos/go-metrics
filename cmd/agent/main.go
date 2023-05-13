package main

import (
	"flag"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/andreyramos/go-metrics/internal/agent"
)

func main() {

	var flagAddr string
	var flagPullInt int
	var flagRepInt int

	flag.StringVar(&flagAddr, "a", "localhost:8080", "address and port to run agent")
	flag.IntVar(&flagPullInt, "p", 2, "frequency of polling metrics from the runtime package")
	flag.IntVar(&flagRepInt, "r", 10, "frequency of sending metrics to the server")

	flag.Parse()

	if envAddr := os.Getenv("V"); envAddr != "" {
		flagAddr = envAddr
	}
	if envPullInt := os.Getenv("V"); envPullInt != "" {
		val, err := strconv.ParseInt(envPullInt, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		flagPullInt = int(val)
	}
	if envRepInt := os.Getenv("V"); envRepInt != "" {
		val, err := strconv.ParseInt(envRepInt, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		flagRepInt = int(val)
	}

	cfg := agent.Config{
		PullInterval:   time.Duration(flagPullInt) * time.Second,
		ReportInterval: time.Duration(flagRepInt) * time.Second,
		Address:        flagAddr,
	}

	agent := agent.New(cfg)
	agent.Run()
}
