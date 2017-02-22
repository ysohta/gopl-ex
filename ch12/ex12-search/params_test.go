package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetProjectsHandler(t *testing.T) {
	for _, test := range []struct {
		query  string
		status int
		want   string
	}{
		{
			"x=true&l=golang&l=programming&max=3&m=j@c&cc=1234567890123456",
			200,
			"Search: {Labels:[golang programming] MaxResults:3 Exact:true Emails:[j@c] CreditCard:1234567890123456}\n",
		},
		{
			"m=INVALID",
			400,
			"m: invalid data format:INVALID\n",
		},
		{
			"l=golang&l=programming&cc=INVALID",
			400,
			"cc: invalid data format:INVALID\n",
		},
	} {

		url := "/search"
		if test.query != "" {
			url += "?" + test.query
		}
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(search)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != test.status {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		got := rr.Body.String()
		if got != test.want {
			t.Errorf("got:%q want:%q", got, test.want)
		}
	}
}
