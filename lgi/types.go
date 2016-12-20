package lgi

// parser types
type LGBool struct {
	Val bool
}
type LGInt struct {
	Val int
}
type LGSymbol struct {
	Val string
}
type LGList struct {
	Val []LGExpr
}
type LGLambda struct {
	Val func([]LGExpr) LGExpr
}
type LGSpecial struct {
	Val func([]LGExpr) LGExpr
}

type LGExpr interface {
	Eval() LGExpr
	String() string
}

