package storage

import (
	"fmt"
	"strconv"
)

type Guage float64
type Counter int64

type Repositories interface {
	SaveGuage(string, Guage)
	SaveCounter(string, Counter)
	// ReadAll() []byte
	GetGauge(string) (Guage, error)
	GetCounter(string) (Counter, error)
	GetAll() (map[string]Guage, map[string]Counter, error)
}

type MemStorage struct {
	Gauges   map[string]Guage
	Counters map[string]Counter
}

func NewMemStorage() *MemStorage {
	return &MemStorage{
		make(map[string]Guage),
		make(map[string]Counter)}
}

func (m *MemStorage) SaveGuage(key string, val Guage) {
	m.Gauges[key] = val
}

func (m *MemStorage) SaveCounter(key string, val Counter) {

	_, ok := m.Counters[key]
	if !ok {
		m.Counters[key] = val
		return
	}

	m.Counters[key] += val
}

func (g *Guage) FromString(str string) error {
	val, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return err
	}
	*g = Guage(val)
	return nil
}

func (c *Counter) FromString(str string) error {
	val, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return err
	}
	*c = Counter(val)
	return nil
}

func (m *MemStorage) GetGauge(key string) (Guage, error) {
	gauge, ok := m.Gauges[key]
	if !ok {
		return Guage(0), fmt.Errorf("gauge metric with key '%s' not found", key)
	}

	return gauge, nil
}

func (m *MemStorage) GetCounter(key string) (Counter, error) {
	counter, ok := m.Counters[key]
	if !ok {
		return Counter(0), fmt.Errorf("counter metric with key '%s' not found", key)
	}

	return counter, nil
}

func (m *MemStorage) GetAll() (map[string]Guage, map[string]Counter, error) {
	return m.Gauges, m.Counters, nil
}

// func (m *MemStorage) ReadAll() []byte {
// 	l := "ALL METRICS ===============\r\n"
// 	l += "GUAGE ===============\r\n"
// 	for k, v := range m.Gauges {
// 		l += fmt.Sprintf("%s: %v\r\n", k, v)
// 	}
// 	l += "COUNTER ===============\r\n"
// 	for k, v := range m.Counters {
// 		l += fmt.Sprintf("%s: %v\r\n", k, v)
// 	}
// 	return []byte(l)
// }
