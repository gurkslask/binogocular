# binogocular

This project aims to show my students how variables and conditions interact,
primarily focusing on how booleans interact with boolean keywords like:

 - AND (&&)
 - OR (||)
 - NOT (!)

## How to use
The entry point for the studen is the file *program.go*. Here some variables
are configured. At first its only 
 - kontaktor1 (output)
 - startknapp1 (input)
 - stoppknapp1 (input)

The students can and should manipulate the code in *program.go* so that it
behaves in a manner that is described by combining boolean operators.

For example:

    kontaktor1 = startknapp1 && !stoppknapp1

This sets the kontaktor1 if startknapp and not stoppknapp1. To run this program
and visualize it, run the command 

    go run *.go

in the root directory. Now you should get a TUI where we can manipulate
the inputs in the runtime.
