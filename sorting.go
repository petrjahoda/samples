package main

import (
	"math/rand"
	"sort"
	"strconv"
	"time"
)

func sorting() {
	max := 1000000
	secondMax := 10000
	var randomIntegers []string
	var randomIntegersTwo []string
	var secondRandomIntegersTwo []string
	var thirdRandomIntegersTwo []string
	var fourthRandomIntegersTwo []string

	for j := 0; j < max; j++ {
		number := rand.Intn(max)
		randomIntegers = append(randomIntegers, strconv.Itoa(number))
	}
	for j := 0; j < secondMax; j++ {
		number := rand.Intn(max)
		randomIntegersTwo = append(randomIntegersTwo, strconv.Itoa(number))
		secondRandomIntegersTwo = append(secondRandomIntegersTwo, strconv.Itoa(number))
		thirdRandomIntegersTwo = append(thirdRandomIntegersTwo, strconv.Itoa(number))
		fourthRandomIntegersTwo = append(thirdRandomIntegersTwo, strconv.Itoa(number))
	}

	same := 0
	simple := time.Now()
	for _, integerFirst := range randomIntegers {
		for _, integerSecond := range randomIntegersTwo {
			if integerFirst == integerSecond {
				same++
				break
			}
		}
	}
	println(strconv.Itoa(same) + " as linear in " + time.Since(simple).String())

	sort.Slice(fourthRandomIntegersTwo, func(i, j int) bool {
		return fourthRandomIntegersTwo[i] <= fourthRandomIntegersTwo[j]
	})
	same = 0
	finaly := time.Now()
	for _, integerFirst := range randomIntegers {
		i := sort.SearchStrings(fourthRandomIntegersTwo, integerFirst)
		if i < len(fourthRandomIntegersTwo) && fourthRandomIntegersTwo[i] == integerFirst {
			same++
		}
	}
	println(strconv.Itoa(same) + " as binary with searchInts in " + time.Since(finaly).String())

	sort.Slice(fourthRandomIntegersTwo, func(i, j int) bool {
		return fourthRandomIntegersTwo[i] <= fourthRandomIntegersTwo[j]
	})
	same = 0
	binary := time.Now()
	for _, integerFirst := range randomIntegers {
		i := sort.Search(len(fourthRandomIntegersTwo), func(i int) bool { return fourthRandomIntegersTwo[i] >= integerFirst })
		if i < len(fourthRandomIntegersTwo) && fourthRandomIntegersTwo[i] == integerFirst {
			same++
		}
	}
	println(strconv.Itoa(same) + " as binary in " + time.Since(binary).String())
}
