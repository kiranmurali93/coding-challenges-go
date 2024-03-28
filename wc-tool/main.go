package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 3 {
		return
	}
	flag := args[1]
	filepath := args[2]

	if flag == "-c" {
		file, err := os.Stat(filepath)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(file.Size(), filepath)

	}

}
