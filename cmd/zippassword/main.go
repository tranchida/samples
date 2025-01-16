package main

import (
	"fmt"
	"io"

	"github.com/yeka/zip"
)

func main() {

	zipfile, err := zip.OpenReader("test.zip")
	if err != nil {
		panic(err)
	}

	defer zipfile.Close()

	for _, file := range zipfile.File {

		if file.IsEncrypted() {
			file.SetPassword("blabla")
		}

		f, err := file.Open()
		if err != nil {
			panic(err)
		}

		buff, err := io.ReadAll(f)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(buff))
	}

}
