package day2

import "fmt"

func MiniQuizDay2() {
	// Quiz 1
	type Int int
	arrInt := []Int{1, 2, 3, 4, 1, 1, 1, 2, 3, 4, 5, 6, 7}
	mapInt := map[Int]int{}

	for _, val := range arrInt {
		mapInt[val]++
	}

	for key, val := range mapInt {
		fmt.Println(key, val)
	}

	// Quiz 2
	str := []string{"cal", "cal", "cal", "man", "man", "tar", "tar", "ra", "ra", "ra"}
	mapStr := map[string]int{}
	mapDuplicate := map[string]bool{}

	var newArr []string
	for _, v := range str {
		mapStr[v]++
		if (mapStr[v] >= 3) && !mapDuplicate[v] {
			newArr = append(newArr, v)
			mapDuplicate[v] = true
		}
	}
	fmt.Println("Jawaban Quiz 2:")
	fmt.Println(newArr)
}
