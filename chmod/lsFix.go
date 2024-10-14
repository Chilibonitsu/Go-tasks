package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("/usr/bin/ls")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fi, err := f.Stat()

	fmt.Println("Текущие права ls\n", fi.Mode())

	errChmod := f.Chmod(0755)

	if errChmod != nil {
		panic(errChmod)
	}

	fi, err = f.Stat()
	fmt.Println("Измененные права ls\n", fi.Mode())

}
