package problems

type number struct {
	num int
}

var numbers map[string]number

func romanToInt(s string) int {
	n := make(map[string]number)
	n["I"] = number{1}
	n["V"] = number{5}
	n["X"] = number{10}
	n["L"] = number{50}
	n["C"] = number{100}
	n["D"] = number{500}
	n["M"] = number{1000}
	returnValue := 0
	//IV

	for i, val := range s {
		if i > 0 {
			if n[string(s[i-1])].num < n[string(s[i])].num {
				returnValue -= n[string(s[i-1])].num * 2
			}
			returnValue += n[string(val)].num
		} else {
			returnValue += n[string(val)].num
		}
	}

	return returnValue

}
