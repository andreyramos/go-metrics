package agent

import (
	"fmt"
	"math/rand"
	"net/http"
	"runtime"
	"time"

	"github.com/andreyramos/go-metrics/internal/storage"
)

type Config struct {
	PullInterval   time.Duration
	ReportInterval time.Duration
	Address        string
}

type Agent struct {
	metrics *storage.MemStorage
	cfg     Config
	client  http.Client
}

func New(cfg Config) *Agent {
	a := &Agent{
		metrics: storage.NewMemStorage(),
		cfg:     cfg,
		client: http.Client{
			Timeout: 1 * time.Second,
		},
	}
	return a
}

func (a *Agent) Update() {
	m := &runtime.MemStats{}
	runtime.ReadMemStats(m)

	a.metrics.SaveGuage("Alloc", storage.Guage(m.Alloc))
	a.metrics.SaveGuage("BuckHashSys", storage.Guage(m.BuckHashSys))
	a.metrics.SaveGuage("Frees", storage.Guage(m.Frees))
	a.metrics.SaveGuage("GCCPUFraction", storage.Guage(m.GCCPUFraction))
	a.metrics.SaveGuage("GCSys", storage.Guage(m.GCSys))
	a.metrics.SaveGuage("HeapAlloc", storage.Guage(m.HeapAlloc))
	a.metrics.SaveGuage("HeapIdle", storage.Guage(m.HeapIdle))
	a.metrics.SaveGuage("HeapInuse", storage.Guage(m.HeapInuse))
	a.metrics.SaveGuage("HeapObjects", storage.Guage(m.HeapObjects))
	a.metrics.SaveGuage("HeapReleased", storage.Guage(m.HeapReleased))
	a.metrics.SaveGuage("HeapSys", storage.Guage(m.HeapSys))
	a.metrics.SaveGuage("LastGC", storage.Guage(m.LastGC))
	a.metrics.SaveGuage("Lookups", storage.Guage(m.Lookups))
	a.metrics.SaveGuage("MCacheInuse", storage.Guage(m.MCacheInuse))
	a.metrics.SaveGuage("MCacheSys", storage.Guage(m.MCacheSys))
	a.metrics.SaveGuage("MSpanInuse", storage.Guage(m.MSpanInuse))
	a.metrics.SaveGuage("MSpanSys", storage.Guage(m.MSpanSys))
	a.metrics.SaveGuage("Mallocs", storage.Guage(m.Mallocs))
	a.metrics.SaveGuage("NextGC", storage.Guage(m.NextGC))
	a.metrics.SaveGuage("NumForcedGC", storage.Guage(m.NumForcedGC))
	a.metrics.SaveGuage("NumGC", storage.Guage(m.NumGC))
	a.metrics.SaveGuage("OtherSys", storage.Guage(m.OtherSys))
	a.metrics.SaveGuage("PauseTotalNs", storage.Guage(m.PauseTotalNs))
	a.metrics.SaveGuage("StackInuse", storage.Guage(m.StackInuse))
	a.metrics.SaveGuage("StackSys", storage.Guage(m.StackSys))
	a.metrics.SaveGuage("Sys", storage.Guage(m.Sys))
	a.metrics.SaveGuage("TotalAlloc", storage.Guage(m.TotalAlloc))
	a.metrics.SaveGuage("RandomValue", storage.Guage(rand.Float64()))
	a.metrics.SaveCounter("PollCount", storage.Counter(1))

}

func (a *Agent) SendMetric(mtype, name, value string) error {
	url := fmt.Sprintf("http://%s/update/%s/%s/%s", a.cfg.Address, mtype, name, value)
	request, err := http.NewRequest(http.MethodPost, url, nil)

	if err != nil {
		return err
	}

	request.Header.Set("Conten-Type", "text/plain")
	res, err := a.client.Do(request)

	if err != nil {
		return err
	}

	res.Body.Close()
	return nil
}

func (a *Agent) SendReport() error {

	for k, v := range a.metrics.Gauges {
		err := a.SendMetric("gauge", k, fmt.Sprintf("%f", v))
		if err != nil {
			return err
		}
	}
	for k, v := range a.metrics.Counters {
		err := a.SendMetric("counter", k, fmt.Sprintf("%d", v))
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *Agent) Run() {
	go func() {
		for {
			time.Sleep(a.cfg.PullInterval)
			a.Update()
		}
	}()

	for {
		time.Sleep(a.cfg.ReportInterval)
		err := a.SendReport()
		if err != nil {
			panic(err)
		}
	}
}
