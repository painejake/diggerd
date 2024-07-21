package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/hako/durafmt"
	"github.com/shirou/gopsutil/host"
)

// SystemInfo stats
type SystemInfo struct {
	Platform       string `json:"platform"`
	Family         string `json:"family"`
	Version        string `json:"version"`
	Uptime         uint64 `json:"uptime"`
	UptimeFriendly string `json:"uptime_friendly"`
}

// getHostData returns the data of the host
func getHostData() SystemInfo {

	hostStats := SystemInfo{}

	platform, family, version, err := host.PlatformInformation()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return hostStats
	}

	// Cleanup the version string, remove the wrapping quotes
	cleanVersion := strings.Replace(version, "\"", "", -1)

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

	hostStats = SystemInfo{platform, family, cleanVersion, uptimeInSeconds, uptimeFriendlyString}

	return hostStats
}
