package main

import "fmt"

func main() {

	var myName = "jake-flip flop"

	fmt.Println("my name is", myName)

	var flexName string = "flex name is flex name"
	fmt.Println("excuse me :", flexName)

	usingCreateAssigned := "holla funmi"

	fmt.Println("what do you say", usingCreateAssigned)

	var sum int
	fmt.Println("the total sumis : ", sum)

	part1, other := 1, 2

	fmt.Println("this is a part one var", part1, "\n this is other var : ", other)

	part2, other := 4, 5

	fmt.Println("part two:", part2, "this is reassigned other val : ", other)

	sum = part1 + part2 + other

	fmt.Println("the total sum: ", sum)

	var (
		testingBlock1 = 3
		testingBlock2 = 9
	)

	fmt.Println("logging testingBlock1: ", testingBlock1, " AND Logging TestingBlock2::", testingBlock2)

	word1, word2, _ := "hello", "world", "!"

	fmt.Println(word1, word2)
}
