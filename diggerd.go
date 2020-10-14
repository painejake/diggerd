package diggerd

import "fmt"

// SystemStats stats
type SystemStats struct {
	CPU    SystemCPU    `json:"cpu"`
	Memory SystemMemory `json:"memory"`
	Net    SystemNet    `json:"net"`
}

// getSystemStats populates the SystemStats struct
// with all the required data and returns
func getSystemStats() SystemStats {
	stats := SystemStats{
		getCPUUsage(),
		getMemoryUsage(),
		getNetUsage(),
	}
	return stats
}

func main() {
	fmt.Println("Starting diggerd...")
	handleRequests()
}
