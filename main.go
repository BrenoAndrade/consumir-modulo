package main

import (
	"fmt"

	"github.com/brenoandrade/modulo/hello"
	v2 "github.com/brenoandrade/modulo/v2/hello"
)

func main() {
	fmt.Println(hello.Hello())
	fmt.Println(v2.HelloV2())
}
