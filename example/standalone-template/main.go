package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/bgadrian/fastfaker/faker"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Printf("usage: <template> [seed]\n")
		fmt.Printf("Example: 'Names: {name}, {name}, {name}.'\n")
		os.Exit(1)
	}

	template := os.Args[1]

	if template == "" {
		fmt.Println("the template cannot be empty")
		os.Exit(1)
	}

	fastFaker := faker.NewFastFaker()
	if len(os.Args) > 2 {
		seed, err := strconv.ParseInt(os.Args[2], 10, 64)
		if err != nil {
			fmt.Printf("seed error: %s\n", err)
			os.Exit(1)
		}
		fastFaker.Seed(seed)
	}

	fmt.Println(fastFaker.Template(template))
}
