package	main

//import "fmt"

// VARIABLES
var kontaktor1 bool
var startknapp1 bool
var stoppknapp1 bool
var kontaktor2 bool
var startknapp2 bool
var stoppknapp2 bool
// ENDVARS

func program() {
	// PROGRAM
	//fmt.Println(kontaktor)
	kontaktor1 = !stoppknapp1 && (startknapp1 || kontaktor1)
	_ = kontaktor1

	// ENDPROG
}
