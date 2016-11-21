package eval

func (v Var) Prec() prec {
	return operatorPrior
}

func (l literal) Prec() prec {
	return operatorPrior
}

func (u unary) Prec() prec {
	return operatorPrior
}

func (b binary) Prec() prec {
	return prec(precedence(b.op))
}

func (c call) Prec() prec {
	return operatorPrior
}
