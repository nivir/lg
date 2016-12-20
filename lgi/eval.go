package lgi

// symbol table
var st = make(map[LGSymbol]LGExpr)

// define various types
func (x LGBool) Eval() LGExpr {
	return x
}
func (x LGInt) Eval() LGExpr {
	return x
}
func (x LGSymbol) Eval() LGExpr {
	return st[x]
}
func (x LGList) Eval() LGExpr {
	return x.Val[0].Eval().(LGCallable).Call(x.Val[1:])
}
func (x LGLambda) Eval() LGExpr {
	return x
}
func (x LGSpecial) Eval() LGExpr {
	return x
}
