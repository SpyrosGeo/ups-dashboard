package main

import (
    "encoding/json"
    "log"
    "net/http"
    "os/exec"
    "strings"
)

// Function to fetch UPS status using apcaccess
func getUPSStatus() (map[string]string, error) {
    // Run apcaccess command
    output, err := exec.Command("sudo", "apcaccess", "status").Output()
    if err != nil {
        return nil, err
    }

    lines := strings.Split(string(output), "\n")
    status := make(map[string]string)

    // Parse the output and store key-value pairs in a map
    for _, line := range lines {
        if strings.Contains(line, ":") {
            parts := strings.SplitN(line, ":", 2)
            key := strings.TrimSpace(parts[0])
            value := strings.TrimSpace(parts[1])
            status[key] = value
        }
    }

    return status, nil
}

// HTTP handler to serve the UPS status as JSON
func upsHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*") 
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS") 
    status, err := getUPSStatus()
    if err != nil {
        http.Error(w, "Error retrieving UPS status", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(status)
}

func main() {
    http.HandleFunc("/api/ups", upsHandler)

    port := ":8080"
    log.Printf("UPS API server running at http://localhost%s/api/ups\n", port)
    if err := http.ListenAndServe(port, nil); err != nil {
        log.Fatalf("Server failed: %s", err)
    }
}
