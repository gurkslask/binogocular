package main

import (
	"os"
	"strings"
)
func readFile(path string) (string, string){
    f, err := os.ReadFile(path)
    check(err)
	ss := strings.Split(string(f), "\n")
	startvariables, endvariables, startprogram, endprogram := 0, 0, 0, 0
	for k, v := range ss {
		if strings.Contains(v, "VARIABLES") {
			startvariables = k + 1
		} else if strings.Contains(v, "ENDVARS") {
			endvariables = k
		} else if strings.Contains(v, "PROGRAM") {
			startprogram = k + 1
		} else if strings.Contains(v, "ENDPROG") {
			endprogram = k
		}
	}
	variables := strings.Join(ss[startvariables:endvariables],"\n")
	program := strings.Join(ss[startprogram:endprogram], "\n")
	return variables, program
}
func check(e error) {
    if e != nil {
        panic(e)
    }
}
