package main

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

func updateAndCreate() {
	UpdateMain()
	UpdateCreate()
}

func UpdateCreate() {
	readFile, err := ioutil.ReadFile("create.sh")
	if err != nil {
		println("Error reading create.sh")
		os.Exit(1)
	}
	println("create.sh: file opened")

	result := ""
	date := time.Now()
	year := date.Year()
	month := date.Month()
	quarter := (month-1)/3 + 1
	monthInQuarter := month - 3*(quarter-1)
	for _, line := range strings.Split(strings.TrimSuffix(string(readFile), "\n"), "\n") {
		if strings.Contains(line, ":2") && strings.Contains(line, "docker rmi -f") {
			i := strings.Index(line, ":")
			result += line[:i+1] + strconv.Itoa(year) + "." + strconv.Itoa(int(quarter)) + "." + strconv.Itoa(int(monthInQuarter)) + "\n"
		} else if strings.Contains(line, ":2") && strings.Contains(line, "docker build -t") {
			i := strings.Index(line, ":")
			result += line[:i+1] + strconv.Itoa(year) + "." + strconv.Itoa(int(quarter)) + "." + strconv.Itoa(int(monthInQuarter)) + " .\n"
		} else if strings.Contains(line, ":2") && strings.Contains(line, "docker push") {
			i := strings.Index(line, ":")
			result += line[:i+1] + strconv.Itoa(year) + "." + strconv.Itoa(int(quarter)) + "." + strconv.Itoa(int(monthInQuarter)) + "\n"
		} else {
			result += line + "\n"
		}
	}
	err = ioutil.WriteFile("create.sh", []byte(result), 0644)
	if err != nil {
		println("Error writing create.sh")
		os.Exit(1)
	}
	println("create.sh: file updated to version " + strconv.Itoa(year) + "." + strconv.Itoa(int(quarter)) + "." + strconv.Itoa(int(monthInQuarter)))
}

func UpdateMain() {
	readFile, err := ioutil.ReadFile("main.go")
	if err != nil {
		println("Error reading main.go")
		os.Exit(1)
	}
	println("main.go: file opened")

	result := ""
	date := time.Now()
	day := date.Day()
	year := date.Year()
	month := date.Month()
	quarter := (month-1)/3 + 1
	monthInQuarter := month - 3*(quarter-1)
	for _, line := range strings.Split(strings.TrimSuffix(string(readFile), "\n"), "\n") {
		if strings.Contains(line, "const version =") {
			result += "const version = \"" + strconv.Itoa(year) + "." + strconv.Itoa(int(quarter)) + "." + strconv.Itoa(int(monthInQuarter)) + "." + strconv.Itoa(day) + "\"\n"
		} else {
			result += line + "\n"
		}
	}
	err = ioutil.WriteFile("main.go", []byte(result), 0644)
	if err != nil {
		println("Error writing main.go")
		os.Exit(1)
	}
	println("main.go: file updated to version " + strconv.Itoa(year) + "." + strconv.Itoa(int(quarter)) + "." + strconv.Itoa(int(monthInQuarter)) + "." + strconv.Itoa(day))
}
