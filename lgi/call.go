package lgi

type LGCallable interface {
	Call(args []LGExpr) LGExpr
}

func (x LGLambda) Call(args []LGExpr) LGExpr {
	var _args []LGExpr;
	for _, arg := range args {
		_args = append(_args, arg.Eval())
	}
	return x.Val(_args)
}
func (x LGSpecial) Call(args []LGExpr) LGExpr {
	return x.Val(args)
}

