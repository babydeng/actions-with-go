package main

import (
	"fmt"
	"github.com/dengbei-victor/post2"
)

func main() {
	post2.Hostname = "10.177.21.113"
	fmt.Println(post2.Port)
	fmt.Println(post2.Hostname)
}
