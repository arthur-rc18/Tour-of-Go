package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Teg())
	v := (Teg())
	fmt.Println(v)
	fmt.Println(len(v))
}

func Teg() []string {
	return strings.Fields("teste te s")
}
