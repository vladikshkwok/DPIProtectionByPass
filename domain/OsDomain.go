package domain

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
