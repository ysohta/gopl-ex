package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		handler    func(http.ResponseWriter, *http.Request)
		q          string
		statusCode int
		out        []string
	}{
		{
			db.list,
			"",
			200,
			[]string{
				"<td>socks</td>", "<td>$5.00</td>",
				"<td>shoes</td>", "<td>$50.00</td>"},
		},
		{
			db.price,
			"?item=socks",
			200,
			[]string{"$5.00"},
		},
		{
			db.price,
			"?item=invalid",
			404,
			[]string{"no such item"},
		},
		{
			db.create,
			"?item=hat&price=100",
			200,
			[]string{"created hat: $100.00"},
		},
		{
			db.create,
			"",
			404,
			[]string{"invalid item name: \"\""},
		},
		{
			db.create,
			"?item=shoes&price=100",
			404,
			[]string{"item already exists: \"shoes\""},
		},
		{
			db.create,
			"?item=glasses&price=invalid",
			404,
			[]string{"invalid price: "},
		},
		{
			db.update,
			"?item=shoes&price=30",
			200,
			[]string{"updated shoes: $30.00"},
		},
		{
			db.update,
			"",
			404,
			[]string{"item not found: \"\""},
		},
		{
			db.update,
			"?item=shoe&price=30",
			404,
			[]string{"item not found: \"shoe\""},
		},
		{
			db.update,
			"?item=shoes&price=invalid",
			404,
			[]string{"invalid price: "},
		},
		{
			db.del,
			"?item=shoes",
			200,
			[]string{"item deleted: "},
		},
		{
			db.del,
			"?item=boots",
			404,
			[]string{"item not found: "},
		},
		{
			db.del,
			"",
			404,
			[]string{"item not found: "},
		},
	}

	for _, test := range tests {
		db = database{"shoes": 50, "socks": 5}
		ts := httptest.NewServer(http.HandlerFunc(test.handler))
		defer ts.Close()

		resp, err := http.Get(ts.URL + test.q)
		if err != nil {
			t.Errorf("unexpected: %v", err)
			return
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("Status code error: %d", resp.StatusCode)
			return
		}

		w := bytes.NewBufferString("")

		_, err = io.Copy(w, resp.Body)
		defer resp.Body.Close()
		if err != nil {
			t.Error("failed to copy")
			return
		}

		got := w.String()
		for _, w := range test.out {
			if strings.Index(got, w) == -1 {
				t.Errorf("%q should contain %q", got, w)
			}
		}
	}
}
