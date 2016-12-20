package lgi

import "io"
import "bufio"

func Parse(reader io.Reader) LGExpr {
	_reader := bufio.NewReader(reader)
	
	return LGInt{5}
}
