package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestRootHandler(t *testing.T) {
    req, err := http.NewRequest("GET", "/", nil)
    if err != nil {
        t.Fatal(err)
    }

    // Create a yest ResponseRecorder
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(RootHandler)
    handler.ServeHTTP(rr, req)

    // Check the status code
    if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", 
		status, http.StatusOK)
    }
}