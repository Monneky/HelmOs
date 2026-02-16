package handler

import (
	"encoding/json"
	"net/http"
)

// Health returns service health (GET /health).
func Health(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]string{
		"status":  "ok",
		"service": "@helmos/api",
	})
}

// DBCheck checks SQLite connectivity (GET /db-check).
func DBCheck(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	type result struct {
		DB  string `json:"db"`
		OK  int    `json:"ok,omitempty"`
		Err string `json:"message,omitempty"`
	}
	var res result
	if db == nil {
		res.DB = "error"
		res.Err = "database not initialized"
		w.WriteHeader(http.StatusServiceUnavailable)
		_ = json.NewEncoder(w).Encode(res)
		return
	}
	var ok int
	if err := db.QueryRow("SELECT 1").Scan(&ok); err != nil {
		res.DB = "error"
		res.Err = err.Error()
		w.WriteHeader(http.StatusServiceUnavailable)
		_ = json.NewEncoder(w).Encode(res)
		return
	}
	res.DB = "ok"
	res.OK = 1
	_ = json.NewEncoder(w).Encode(res)
}
