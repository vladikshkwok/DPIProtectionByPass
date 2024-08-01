package domain

import (
	"fmt"
)

type Memory struct {
	MemTotal     int
	MemFree      int
	MemAvailable int
}

type LoadAvg struct {
	Load1          float64
	Load5          float64
	Load15         float64
	LastCreatedPid int
}

func (l LoadAvg) String() string {
	return fmt.Sprintf("Load average: %.2f %.2f %.2f. Last created PID: %d", l.Load1, l.Load5, l.Load15, l.LastCreatedPid)
}

func (m Memory) String() string {
	return fmt.Sprintf("Memory total %d, Memory free: %d, Used memory: %.2f%%", m.MemTotal, m.MemFree, float64(m.MemTotal-m.MemFree)/float64(m.MemTotal)*100)
}
