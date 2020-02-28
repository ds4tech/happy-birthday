package birthday

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestIndex(t *testing.T) {
    req, err := http.NewRequest("GET", "/", nil)
    if err != nil {
        t.Fatal(err)
    }
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(Index)
    handler.ServeHTTP(rr, req)
    // Check the status code is what we expect.
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }
    // Check the response body is what we expect.
    expected := string(`{"msg": Happy Birthday Service!!!}`)
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

func TestHello(t *testing.T) {
    req, err := http.NewRequest("GET", "/hello", nil)
    if err != nil {
        t.Fatal(err)
    }
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(Hello)
    handler.ServeHTTP(rr, req)
    // Check the status code is what we expect.
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }
    // Check the response body is what we expect.
    expected := string(`{"msg": Hello People}`)
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}

func TestHelloSomebodyOK(t *testing.T) {
    req, err := http.NewRequest("GET", "/hello/John", nil)
    if err != nil {
        t.Fatal(err)
    }
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(HelloSomebody)
    handler.ServeHTTP(rr, req)
    // Check the status code is what we expect.
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }
    // Check the response body is what we expect.
    expected := `{"msg": Hello, 'John'! Your birthday is in 5 days.}`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}
func TestHelloSomebodyNotOK(t *testing.T) {
    req, err := http.NewRequest("GET", "/hello/John", nil)
    if err != nil {
        t.Fatal(err)
    }
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(HelloSomebody)
    handler.ServeHTTP(rr, req)
    // Check the status code is what we expect.
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }
    // Check the response body is what we expect.
    expected := `{"msg": Unfortunatelly, name 'John' is not in the database.}`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}
