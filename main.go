package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Entry struct {
	  Value string

}
type Store struct {
	data map[string]Entry
}
func NewStore() *Store {
	return &Store{
		data: make(map[string]Entry),
	}
}
func (s *Store) Set(key string, value string) {
	s.data[key] = Entry{Value: value}
}
func (s *Store) Get(key string) (string, bool) {
	entry, ok := s.data[key]
	if !ok {
		return "", false
	}
	return entry.Value, true
}






func main() {
  store := NewStore()
	store.Set("name", "Ganesh")
  val, ok := store.Get("name")
	fmt.Println(val, ok)


	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response := map[string]string{"status": "ok"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	http.ListenAndServe(":8080", nil)

	
}