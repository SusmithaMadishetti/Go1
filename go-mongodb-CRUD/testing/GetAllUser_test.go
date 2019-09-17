package main

import (
	"../controllers"

	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllUser(t *testing.T) {
	req, err := http.NewRequest("GET", "/user", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	uc := controllers.NewUserController(getSession())
	handler := http.HandlerFunc(uc.GetAllUser)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `[{"name":"sush","gender":"f","age":24},{"name":"Karth","gender":"m","age":24}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
