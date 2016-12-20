package lgi

import "fmt"

func (x LGBool) String() string {
	return fmt.Sprintf("%t", x.Val)
}
func (x LGInt) String() string {
	return fmt.Sprintf("%d", x.Val)
}
func (x LGSymbol) String() string {
	return fmt.Sprintf("%s", x.Val)
}
func (x LGList) String() string {
	res := "("
	if(len(x.Val) > 0) {
		res += fmt.Sprintf("%s", x.Val[0])
	}
	if(len(x.Val) > 1) {
		for _, elem := range x.Val[1:] {
			res += " "
			res += fmt.Sprintf("%s", elem)
		}
	}
	res += ")"
	return res
}
func (x LGLambda) String() string {
	return fmt.Sprintf("<lambda %p>", &x.Val)
}
func (x LGSpecial) String() string {
	return fmt.Sprintf("<special %p>", &x.Val)
}
