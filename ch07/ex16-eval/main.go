package main

import (
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
  {{range $var, $val := .Vars}}
    <tr>
      <td>{{$var}}</td>
      <td><input name="{{$var}}" value="0"></td>      
    </tr>
  {{end}}
  </table>
  <input type="submit" value="calc">
  <p>Ans={{.Ans}}</p>
</form>
`))

var (
	calc = calculator{nil, map[eval.Var]bool{}, 0}
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
		log.Fatal(err)
	}

	if err := calc.Expr.Check(calc.Vars); err != nil {
		log.Fatal(err)
	}

	if err := calcs.Execute(w, calc); err != nil {
		log.Fatal(err)
	}
}

func (c calculator) calculate(w http.ResponseWriter, req *http.Request) {
	env := eval.Env{}
	for key := range calc.Vars {
		v := req.FormValue(string(key))
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			log.Fatal(err)
			return
		}
		env[key] = f
	}

	calc.Ans = calc.Expr.Eval(env)
	if err := calcs.Execute(w, calc); err != nil {
		log.Fatal(err)
	}
}
