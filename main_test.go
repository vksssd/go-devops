package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestIndex tests serving the index.html file
func TestIndex(t *testing.T) {
	req, err := http.NewRequest("GET", "/vnayak", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlerName)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// expected, err := ioutil.ReadFile("./public/index.html")
	// if err != nil {
	// 	t.Fatalf("could not read index.html: %v", err)
	// }
	
	expected := webPage

	if rr.Body.String() != string(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), string(expected))
	}
}

// // TestPage tests serving a specific HTML file
// func TestPage(t *testing.T) {
// 	tests := []struct {
// 		page     string
// 		expected string
// 	}{
// 		{"/page1.html", "./public/page1.html"},
// 		{"/page2.html", "./public/page2.html"},
// 		// {"/page3.html", "./public/page3.html"},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.page, func(t *testing.T) {
// 			req, err := http.NewRequest("GET", tt.page, nil)
// 			if err != nil {
// 				t.Fatal(err)
// 			}

// 			rr := httptest.NewRecorder()
// 			handler := http.HandlerFunc(handler)

// 			handler.ServeHTTP(rr, req)

// 			if status := rr.Code; status != http.StatusOK {
// 				t.Errorf("handler returned wrong status code: got %v want %v",
// 					status, http.StatusOK)
// 			}

// 			expected, err := ioutil.ReadFile(tt.expected)
// 			if err != nil {
// 				t.Fatalf("could not read %s: %v", tt.expected, err)
// 			}

// 			if rr.Body.String() != string(expected) {
// 				t.Errorf("handler returned unexpected body: got %v want %v",
// 					rr.Body.String(), string(expected))
// 			}
// 		})
// 	}
// }
