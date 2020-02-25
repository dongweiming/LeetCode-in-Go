package problem0017

var letterMap = []string{
	" ",    //0
	"",     //1
	"abc",  //2
	"def",  //3
	"ghi",  //4
	"jkl",  //5
	"mno",  //6
	"pqrs", //7
	"tuv",  //8
	"wxyz", //9
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}
	res := []string{}
	findCombination(&res, &digits, 0, "")
	return res
}

func findCombination(res *[]string, digits *string, index int, s string) {
	if index == len(*digits) {
		*res = append(*res, s)
		return
	}
	num := (*digits)[index]
	letter := letterMap[num-'0']
	for i := 0; i < len(letter); i++ {
		findCombination(res, digits, index+1, s+string(letter[i]))
	}
	return
}
