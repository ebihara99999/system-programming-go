package main

import (
	"archive/zip"
	"io"
	"net/http"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment; filename=ascii_sample.zip")

	zipWriter := zip.NewWriter(w)
	defer zipWriter.Close()
	firstFile, err := zipWriter.Create("a.txt")
	if err != nil {
		panic(err)
	}
	io.Copy(firstFile, strings.NewReader("hogehoge"))

	secondFile, err := zipWriter.Create("b.txt")
	if err != nil {
		panic(err)
	}
	io.Copy(secondFile, strings.NewReader("fugaguga"))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
