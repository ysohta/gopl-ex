package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/ysohta/gopl-ex/ch07/ex13-eval"
)

var validator = template.Must(template.New("validator").
	Parse(`
<h1>Calculator</h1>
<form action="validate" method="post">
  <input type="text" name="expr">
  <input type="submit" value="validate">
</form>
`))

var calcs = template.Must(template.New("calcs").
	Parse(`
<h1>Calculator</h1>
<form action="calculate" method="post">
  <p>{{.Expr}}</p>
  <table>
  {{range $var, $val := .Env}}
    <tr>
      <td>{{$var}}</td>
      <td><input name="{{$var}}" value="{{$val}}"></td>      
    </tr>
  {{end}}
  </table>
  <input type="submit" value="calc">
  <p>Ans={{.Ans}}</p>
</form>
`))

var (
	calc = calculator{nil, map[eval.Var]bool{}, eval.Env{}, 0}
)

func main() {
	http.HandleFunc("/", exec)
	http.HandleFunc("/validate", calc.validate)
	http.HandleFunc("/calculate", calc.calculate)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type calculator struct {
	Expr eval.Expr
	Vars map[eval.Var]bool
	Env  eval.Env
	Ans  float64
}

func exec(w http.ResponseWriter, req *http.Request) {
	if err := validator.Execute(w, calc); err != nil {
		log.Fatal(err)
	}
}

func (c calculator) validate(w http.ResponseWriter, req *http.Request) {
	var err error
	calc.Expr, err = eval.Parse(req.FormValue("expr"))
	if err != nil {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "failed to parse expression: %q\n", err)
		return
	}

	// clear var
	calc.Vars = map[eval.Var]bool{}

	if err := calc.Expr.Check(calc.Vars); err != nil {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "expression check error: %q\n", err)
		return
	}

	calc.Env = eval.Env{}
	for key := range calc.Vars {
		calc.Env[key] = 0
	}

	if err := calcs.Execute(w, calc); err != nil {
		log.Fatal(err)
	}
}

func (c calculator) calculate(w http.ResponseWriter, req *http.Request) {
	for key := range calc.Vars {
		v := req.FormValue(string(key))
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			w.WriteHeader(http.StatusNotFound) // 404
			fmt.Fprintf(w, "invalid number format: %q\n", err)
			return
		}
		calc.Env[key] = f
	}

	calc.Ans = calc.Expr.Eval(calc.Env)
	if err := calcs.Execute(w, calc); err != nil {
		log.Fatal(err)
	}
}
