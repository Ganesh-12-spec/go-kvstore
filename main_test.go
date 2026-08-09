package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestSetHandler(t *testing.T) {
	store := NewStore()
	handler := func(w http.ResponseWriter, r *http.Request) {
		var p struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		}
		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		store.Set(p.Key, p.Value, 24*time.Hour)
		w.WriteHeader(http.StatusOK)
	}

	req := httptest.NewRequest("POST", "/keys", strings.NewReader(`{"key":"testKey","value":"testValue"}`))
	w := httptest.NewRecorder()

	handler(w, req)

	if w.Code != 200 {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	val, ok := store.Get("testKey")
	if !ok || val != "testValue" {
		t.Errorf("expected testValue, got %s", val)
	}
}