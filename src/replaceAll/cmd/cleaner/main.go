package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	flag.Parse()
	name := flag.Arg(0)
	message := strings.ReplaceAll(name, "(", "")
	message = strings.ReplaceAll(message, ")", "")
	message = strings.ReplaceAll(message, "-", "")
	message = strings.ReplaceAll(message, " ", "")
	fmt.Println(message)

}
