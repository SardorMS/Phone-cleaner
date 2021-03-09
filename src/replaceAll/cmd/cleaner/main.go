package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	flag.Parse()
	name := flag.Arg(0)
	message1 := strings.ReplaceAll(name, "(", "")
	message1 = strings.ReplaceAll(message1, ")", "")
	message1 = strings.ReplaceAll(message1, "-", "")
	message1 = strings.ReplaceAll(message1, " ", "")
	fmt.Println(message1)

}
