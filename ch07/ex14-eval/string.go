package eval

import "fmt"

func (v Var) String() string {
	return string(v)
}

func (l literal) String() string {
	return fmt.Sprintf("%.6g", l)
}

func (u unary) String() string {
	return fmt.Sprintf("%c%s", u.op, u.x)
}

func (b binary) String() string {
	var str string
	// lower precedence
	if b.x.Prec() < b.Prec() {
		str = fmt.Sprintf("(%s)", b.x)
	} else {
		str = b.x.String()
	}

	// lower precedence
	if b.y.Prec() < b.Prec() {
		str += fmt.Sprintf(" %c (%s)", b.op, b.y)
	} else {
		str += fmt.Sprintf(" %c %s", b.op, b.y)
	}

	return str
}

func (c call) String() string {
	args := fmt.Sprintf("%s", c.args[0])
	for i := 1; i < len(c.args); i++ {
		args += fmt.Sprintf(", %s", c.args[i])
	}

	return fmt.Sprintf("%s(%s)", c.fn, args)
}

func (c call2) String() string {
	return fmt.Sprintf("%s(%s, %s)", c.fn, c.x, c.y)
}
