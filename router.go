package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// logMsg will log any incoming API requests to the console
func logMsg(endpoint string, r *http.Request) {
	ipAdd := getIPAddress(r)
	ua := r.UserAgent()
	fmt.Println(fmt.Sprintf("Served: %s %s %s", endpoint, ipAdd, ua))
}

// getIPAddress gets a requests IP address by reading off the forwarded-for
// header (for proxies) and falls back to use the remote address.
func getIPAddress(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}

// getStats handles the JSON response containing the SystemStats
func getStats(w http.ResponseWriter, r *http.Request) {
	logMsg("/stats", r)
	disableCors(&w)
	json.NewEncoder(w).Encode(getSystemStats())
}

// getStats handles the JSON response containing the SystemMemory only
func getMemoryStats(w http.ResponseWriter, r *http.Request) {
	logMsg("/stats/memory", r)
	disableCors(&w)
	json.NewEncoder(w).Encode(getMemoryUsage())
}

// getCPUStats handles the JSON response containing the SystemCPU only
func getCPUStats(w http.ResponseWriter, r *http.Request) {
	logMsg("/stats/cpu", r)
	disableCors(&w)
	json.NewEncoder(w).Encode(getCPUUsage())
}

// getNetStats handles the JSON response containing the SystemNet only
func getNetStats(w http.ResponseWriter, r *http.Request) {
	logMsg("/stats/net", r)
	disableCors(&w)
	json.NewEncoder(w).Encode(getNetUsage())
}

// disableCors will set the access allowed origin to *
// as we're not serving anything important or accepting any
// data doesn't really matter
func disableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

// handleRequests will route our incoming API calls
func handleRequests() {
	http.HandleFunc("/stats", getStats)
	http.HandleFunc("/stats/cpu", getCPUStats)
	http.HandleFunc("/stats/memory", getMemoryStats)
	http.HandleFunc("/stats/net", getNetStats)

	fmt.Println("Listening on *:10000")

	log.Fatal(http.ListenAndServe(":10000", nil))
}
