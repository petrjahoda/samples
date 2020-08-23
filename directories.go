package main

import (
	"github.com/fatih/color"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func Directories() {
	// create directory even if exists
	NewPath := filepath.Join(".", "public")
	mkdirError := os.MkdirAll(NewPath, 0777)
	if mkdirError != nil {
		panic("Unable to create directory")
	} else {
		color.Green("Directory created")
	}

	//create file even if exists
	fileToMake := "test.txt"
	fileOut, err := os.Create(strings.Join([]string{NewPath, fileToMake}, "/"))
	if err != nil {
		panic("Unable to create file")
	} else {
		println("File created")
	}
	defer fileOut.Close()

	// write data to file
	data := "Kobyla\nma\nmaly bok"
	writingError := ioutil.WriteFile(NewPath+"/"+fileToMake, []byte(data), 0777)
	if writingError != nil {
		panic("Unable write to file")
	} else {
		println("Data writen to file")
	}
}
