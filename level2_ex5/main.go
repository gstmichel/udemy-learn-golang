package main

import "fmt"

func main() {
	x := `This ?
is what raw mean...
Can go multiple line with backtick, not with
"" double quotes...
`
	fmt.Println(x)
}
