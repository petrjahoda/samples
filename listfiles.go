package main

import (
	"encoding/csv"
	"io"
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
		println(err.Error())
	}
	for _, file := range files {
		csvFile, err := os.Open(file)
		if err != nil {
			println("Cannot open file: " + err.Error())
		} else {
			r := csv.NewReader(csvFile)
			r.Comma = ';'
			for {
				record, err := r.Read()
				if err == io.EOF {
					break
				}
				if err != nil {
					println("Cannot read file: " + err.Error())
					break
				}
				for _, data := range record {
					print(data + " ")
				}
				println("")
			}
		}
	}
}
