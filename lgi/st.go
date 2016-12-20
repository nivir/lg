package lgi

// symbol table
type SymbolTable map[LGSymbol]LGExpr

func InitSymbolTable() {
	// TODO: no globals!
	st = make(SymbolTable)

	// populate the symbol table
	st[LGSymbol{"+"}] = LGLambda{func(args []LGExpr) LGExpr {
		sum := 0
		for _, arg := range args {
			// TODO: Runnable
			sum += arg.Eval().(LGInt).Val
		}		
		return LGInt{sum}
	}}
	st[LGSymbol{"if"}] = LGSpecial{func(args []LGExpr) LGExpr {
		if args[0].Eval().(LGBool).Val == true {
			return args[1].Eval()
		} else {
			return args[2].Eval()
		}
	}}
}

