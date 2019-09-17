package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllExpanses(t *testing.T) {
	req, err := http.NewRequest("GET", "/expanses", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	//	uc := controllers.NewUserController(getSession())
	handler := http.HandlerFunc(GetAllExpanses)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `[{"id":"1","name":"Rent","description":"For the mont Aug","amount":36000},{"id":"2","name":"Power Bill","description":"For the mont Aug","amount":3000},{"id":"1","name":"Rent","description":"For the mont Aug","amount":0},{"id":"7","name":"elec","description":"Aug","amount":0}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
