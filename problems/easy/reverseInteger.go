package problems

import (
	"fmt"
	"strconv"
)

func reverse(x int) int {
	isMinus := false
	if x > 2147483647 {
		return 0
	}
	if x < 0 {
		x = x * -1
		isMinus = true
	}
	strNumber := strconv.Itoa(x)
	reversedNumber := ""

	for length := len(strNumber); length > 0; length-- {
		reversedNumber += string(strNumber[length-1])
	}
	reverseNum, error := strconv.Atoi(reversedNumber)
	if error != nil {
		fmt.Println("Failure to cast String to int")
	}
	if isMinus == true {
		reverseNum = reverseNum * -1
	}
	if reverseNum > 2147483647 || reverseNum < -2147483647 {
		return 0
	}
	return reverseNum
}
