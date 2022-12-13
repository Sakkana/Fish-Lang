package main

import (
	"Fish-Lang/repl"
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Printf("Fish programming language 1.0.0 (%s %d %d, %d:%d:%d)\n[GCC 9.4.0] on linux\n",
		time.Now().Month(), time.Now().Day(), time.Now().Year(),
		time.Now().Hour(), time.Now().Minute(), time.Now().Second())

	fmt.Printf("This language support basic type system and C-like grammar.\n")
	repl.Start(os.Stdin, os.Stdout)
}
