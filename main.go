package main

import (
	"encoding/json"
	"time"

	"net/http"
)

type Entry struct {
	  Value string
		ExpiresAt time.Time

}
type Store struct {
	data map[string]Entry
}
type SetRequest struct {
	Key string `json:"key"`
	Value string `json:"value"`
}





func NewStore() *Store {
	return &Store{
		data: make(map[string]Entry),
	}
}
func (s *Store) Set(key string, value string, ttl time.Duration) {
	expiry := time.Now().Add(ttl)
	s.data[key] = Entry{Value: value, ExpiresAt: expiry}
}
func (s *Store) Get(key string) (string, bool) {
	entry, ok := s.data[key]
	if !ok {
		return "", false
	}
	if time.Now().After(entry.ExpiresAt){
		delete(s.data, key)
		return "", false
	}
	return entry.Value, true
}






func main() {
  store := NewStore()
	


	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response := map[string]string{"status": "ok"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})
	http.HandleFunc("/keys", func(w http.ResponseWriter, r *http.Request) {
		var req SetRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w,"invaid request", http.StatusBadRequest)
			return
		}
      store.Set(req.Key, req.Value, 24*time.Hour)

			response := map[string]string{"status":"ok"}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
	})
	http.HandleFunc("/keys/", func(w http.ResponseWriter, r *http.Request){
		key := r.URL.Path[len("/keys/"):]
		value, ok := store.Get(key)

		if !ok {
			http.Error(w, "key not found", http.StatusNotFound)
			return
		}
		response := map[string]string{"value": value}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})







	http.ListenAndServe(":8080", nil)

	
}