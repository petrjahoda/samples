package main

import (
	"strconv"
	"strings"
)

func RompaRfidConverter() {
	originalRFIDString := "ed0004b1ae"
	var convertedString = MakeConversion(originalRFIDString)
	println(originalRFIDString + " - " + strconv.Itoa(int(convertedString)))

}

func MakeConversion(originalData string) int64 {
	shortenedData := strings.ToUpper(originalData[len(originalData)-6:])
	convertedData := ConvertHexData(shortenedData)
	finalData, _ := strconv.ParseInt(convertedData, 16, 64)
	return finalData
}

func ConvertHexData(hexData string) string {
	originalData := "0123456789ABCDEF"
	replacementData := "084C2A6E195D3B7F"
	var finalString string
	for position, _ := range hexData {
		for index, _ := range originalData {
			if hexData[position] == originalData[index] {
				finalString += replacementData[index : index+1]
			}
		}
	}
	return finalString
}
