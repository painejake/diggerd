package main

import (
	"fmt"
	"os"

	"github.com/shirou/gopsutil/net"
)

// InterfaceStats stats hold the stats for each interface
type InterfaceStats struct {
	Interface       string `json:"interface"`
	BytesSent       uint64 `json:"bytes_sent"`
	BytesRecieved   uint64 `json:"bytes_recieved"`
	PacketsSent     uint64 `json:"packets_sent"`
	PacketsRecieved uint64 `json:"packets_recieved"`
}

// SystemNet holds the stats for each interface
type SystemNet struct {
	Stats []InterfaceStats `json:"stats"`
}

// appendStat will append an InterfaceStats stuct to the SystemNet struct
func (netStat *SystemNet) appendStat(infStat InterfaceStats) []InterfaceStats {
	netStat.Stats = append(netStat.Stats, infStat)
	return netStat.Stats
}

// getNetUsage will return the network stats for each
// interface found on the system
func getNetUsage() SystemNet {

	netStats := SystemNet{}
	n, err := net.IOCounters(true)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return netStats
	}

	// Loop through the interfaces and build up
	// the results, append
	for i := 0; i < len(n); i++ {
		s := n[i]
		intf := InterfaceStats{s.Name, s.BytesSent, s.BytesRecv, s.PacketsSent, s.PacketsRecv}
		netStats.appendStat(intf)
	}

	return netStats
}
