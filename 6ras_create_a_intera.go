package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Dashboard represents the security tool dashboard
type Dashboard struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	SystemStats struct {
		CPUUsage float64 `json:"cpu_usage"`
		MemUsage float64 `json:"mem_usage"`
		DiskUsage float64 `json:"disk_usage"`
	} `json:"system_stats"`
	Threats []Threat `json:"threats"`
}

// Threat represents a security threat
type Threat struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Level string `json:"level"`
}

var dashboards = map[string]Dashboard{}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/dashboards", getDashboards).Methods("GET")
	router.HandleFunc("/dashboards/{username}", getDashboard).Methods("GET")
	router.HandleFunc("/dashboards", createDashboard).Methods("POST")
	router.HandleFunc("/dashboards/{username}", updateDashboard).Methods("PUT")
	router.HandleFunc("/dashboards/{username}", deleteDashboard).Methods("DELETE")

	fmt.Println("Server listening on port 8000")
	http.ListenAndServe(":8000", router)
}

func getDashboards(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(dashboards)
}

func getDashboard(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	username := params["username"]
	dashboard, ok := dashboards[username]
	if !ok {
		http.Error(w, "Dashboard not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(dashboard)
}

func createDashboard(w http.ResponseWriter, r *http.Request) {
	var dashboard Dashboard
	_ = json.NewDecoder(r.Body).Decode(&dashboard)
	dashboards[dashboard.Username] = dashboard
	w.WriteHeader(http.StatusCreated)
}

func updateDashboard(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	username := params["username"]
	var dashboard Dashboard
	_ = json.NewDecoder(r.Body).Decode(&dashboard)
	dashboards[username] = dashboard
}

func deleteDashboard(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	username := params["username"]
	delete(dashboards, username)
}