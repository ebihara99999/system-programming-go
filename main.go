package main

import (
	"archive/zip"
	"io"
	"os"
	"strings"
)

func main() {
	file, err := os.Create("new.zip")
	if err != nil {
		panic(err)
	}

	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()

	firstFile, err := zipWriter.Create("first.txt")
	if err != nil {
		panic(err)
	}
	io.Copy(firstFile, strings.NewReader("first contents"))

	secondFile, err := zipWriter.Create("second.txt")
	if err != nil {
		panic(err)
	}

	io.Copy(secondFile, strings.NewReader("second contents"))
}
