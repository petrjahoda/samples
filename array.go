package main

func arrayOverflow() {
	primes := []int{2, 3, 5, 7, 11, 13}
	printOverflow(primes, 17)
	println("Survived index out of range error")
}

func printOverflow(primes []int, i int) {
	defer func() {
		if err := recover(); err != nil {
			println("Index out of range while printing")
		}
	}()
	println(primes[i])
}
