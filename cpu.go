package main

import (
	"fmt"
	"os"

	"github.com/shirou/gopsutil/cpu"
)

// SystemCPU stats
type SystemCPU struct {
	UsedPercent []float64 `json:"used_percent"`
}

// getCPUUsage will return the CPU usage as
// a percentage for the host
func getCPUUsage() SystemCPU {

	cpuStats := SystemCPU{}
	procPerc, err := cpu.Percent(0, false)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return cpuStats
	}
	cpuStats = SystemCPU{procPerc}

	return cpuStats
}
