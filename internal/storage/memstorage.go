package storage

import "strconv"

type Guage float64
type Counter int64

type Repositories interface {
	SaveGuage(string, Guage)
	SaveCounter(string, Counter)
}

type MemStorage struct {
	gauges   map[string]Guage
	counters map[string]Counter
}

func NewMemStorage() *MemStorage {
	return &MemStorage{
		make(map[string]Guage),
		make(map[string]Counter)}
}

func (m *MemStorage) SaveGuage(key string, val Guage) {
	m.gauges[key] = val
}

func (m *MemStorage) SaveCounter(key string, val Counter) {

	_, ok := m.counters[key]
	if !ok {
		m.counters[key] = val
		return
	}

	m.counters[key] += val
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
