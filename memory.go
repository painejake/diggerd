package main

import (
	"fmt"
	"os"

	"github.com/shirou/gopsutil/mem"
)

// SystemMemory stats
type SystemMemory struct {
	Total  uint64 `json:"total"`
	Used   uint64 `json:"used"`
	Cached uint64 `json:"cached"`
	Free   uint64 `json:"free"`
}

// getMemoryUsage returns the memory usage of the host
func getMemoryUsage() SystemMemory {

	memStats := SystemMemory{}
	v, err := mem.VirtualMemory()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return memStats
	}
	memStats = SystemMemory{v.Total, v.Used, v.Cached, v.Free}

	return memStats
}
