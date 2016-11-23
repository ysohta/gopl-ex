package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/ysohta/gopl-ex/ch07/ex13-eval"
)

func TestHandlerExec(t *testing.T) {
	tests := []struct {
		handler    func(http.ResponseWriter, *http.Request)
		statusCode int
		out        []string
	}{
		{
			exec,
			200,
			[]string{
				"Calculator",
				"validate",
			},
		},
	}

	for _, test := range tests {
		calc = calculator{nil, map[eval.Var]bool{}, eval.Env{}, 0}
		ts := httptest.NewServer(http.HandlerFunc(test.handler))
		defer ts.Close()

		resp, err := http.Get(ts.URL)
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

func TestHandlerValidate(t *testing.T) {
	tests := []struct {
		handler    func(http.ResponseWriter, *http.Request)
		expr       string
		statusCode int
		out        []string
	}{
		{
			calc.validate,
			"x * y",
			200,
			[]string{
				"x * y",
			},
		},
		{
			calc.validate,
			"3 * x++",
			404,
			[]string{
				"failed to parse expression: ",
			},
		},
		{
			calc.validate,
			"pow(1, 2, 3)",
			404,
			[]string{
				"expression check error: ",
			},
		},
	}

	for _, test := range tests {
		calc = calculator{nil, map[eval.Var]bool{}, eval.Env{}, 0}
		ts := httptest.NewServer(http.HandlerFunc(test.handler))
		defer ts.Close()

		resp, err := http.PostForm(ts.URL,
			url.Values{"expr": {test.expr}})
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

func TestHandlerCalculate(t *testing.T) {
	tests := []struct {
		handler    func(http.ResponseWriter, *http.Request)
		expr       string
		vars       map[string][]string
		statusCode int
		out        []string
	}{
		{
			calc.calculate,
			"x * y",
			map[string][]string{"x": {"3"}, "y": {"2"}},
			200,
			[]string{
				"x * y",
				"Ans=6",
			},
		},
		{
			calc.calculate,
			"x * y",
			map[string][]string{"x": {""}, "y": {"2"}},
			404,
			[]string{
				"invalid number format: ",
			},
		},
	}

	for _, test := range tests {
		var expr eval.Expr
		var err error
		expr, err = eval.Parse(test.expr)
		if err != nil {
			t.Errorf("parse error: %s", err)
			return
		}

		calc = calculator{
			Expr: expr,
			Vars: map[eval.Var]bool{},
			Env:  eval.Env{},
		}

		for k := range test.vars {
			calc.Vars[eval.Var(k)] = true
		}

		ts := httptest.NewServer(http.HandlerFunc(test.handler))
		defer ts.Close()

		resp, err := http.PostForm(ts.URL,
			url.Values(test.vars))
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
