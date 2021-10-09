package main

import (
	"fmt"
	"net/http"
	"testing"
	"net/http/httptest"
)

func TestGetEntries(t *testing.T) { 

req, err := http.NewRequest("GET", "/users/0", nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(req)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetUser)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf(" returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
   expected := `{"UserID": 0,"Name": "NAVEEN","Email": "mknaveen837@gmail.com","Password": "14785269","Time": "2021-10-09T11:41:01.895Z"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
  //not fixed
}