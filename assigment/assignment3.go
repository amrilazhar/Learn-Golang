package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	filename := os.Args
	fmt.Println("file opened : ", filename[1])
	fl, err := os.Open(filename[1])
	if err != nil {
		fmt.Println("Error :", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, fl)
}
