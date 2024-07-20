package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/hako/durafmt"
	"github.com/shirou/gopsutil/host"
)

// SystemInfo stats
type SystemInfo struct {
	// Hostname  uint64 `json:"hostname"`
	Uptime  		uint64 `json:"uptime"`
	UptimeFriendly  string `json:"uptime_friendly"`
	// Procs  uint64 `json:"Procs"`
}

// getHostData returns the data of the host
func getHostData() SystemInfo {

	hostStats := SystemInfo{}

	uptimeInSeconds, err := host.Uptime()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return hostStats
	}
	
	// Convert uint64 to string
	uptimeInSecondsString := strconv.FormatUint(uptimeInSeconds, 10)
	// Append 's' to the string
	uptimeString := uptimeInSecondsString + "s"
	
	parsedUptimeString, err := durafmt.ParseString(uptimeString)
	if err != nil {
		fmt.Println(err)
		return hostStats
	}
	
	// Convert *durafmt.Durafmt to string
	uptimeFriendlyString := parsedUptimeString.String()
	
	hostStats = SystemInfo{uptimeInSeconds, uptimeFriendlyString}

	return hostStats
}
