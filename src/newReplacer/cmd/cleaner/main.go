package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	flag.Parse()
	name := flag.Arg(0)
	replacer := strings.NewReplacer("(", "", ")", "", "-", "", " ", "")
	message := replacer.Replace(name)
	fmt.Println(message)

}
