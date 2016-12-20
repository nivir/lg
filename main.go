package main

import "os"
import "fmt"
import "lg/lgi"


// main program
func main() {

	lgi.InitSymbolTable()
	fmt.Printf("%s\n", lgi.Parse(os.Stdin))
}
