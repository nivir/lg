package main

import "fmt"
import "lg/lgi"


// main program
func main() {

	lgi.InitSymbolTable()
	fmt.Printf("%s\n", lgi.Parse())
}
