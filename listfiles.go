package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func ListFilesInDirectory() {
	var files []string

	root := "./ipg"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		dat, err := ioutil.ReadFile(file)
		if err != nil {
			println(err.Error())
		}
		fmt.Print(string(dat))
	}
}
