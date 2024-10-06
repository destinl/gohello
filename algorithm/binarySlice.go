package algorithm

import (
	"sort"
	"strconv"
)

// BinarySlice 自定义类型，实现 sort.Interface 接口
type BinarySlice []string

func (bs BinarySlice) Len() int {
	return len(bs)
}

func (bs BinarySlice) Less(i, j int) bool {
	// strI := bs[i]
	// strJ := bs[j]
	// lenI := len(strI)
	// lenJ := len(strJ)
	// indexI := 0
	// indexJ := 0
	// countI := 0
	// countJ := 0

	// for indexI < lenI && indexJ < lenJ {
	// 	if strI[indexI] == '1' {
	// 		countI++
	// 	} else {
	// 		countI = 0
	// 	}
	// 	if strJ[indexJ] == '1' {
	// 		countJ++
	// 	} else {
	// 		countJ = 0
	// 	}
	// 	if countI > countJ {
	// 		return true
	// 	}
	// 	if countJ > countI {
	// 		return false
	// 	}
	// 	indexI++
	// 	indexJ++
	// }

	// // 如果其中一个遍历完了
	// if indexI == lenI && indexJ < lenJ && strJ[indexJ] == '0' {
	// 	return true
	// } else if indexI == lenI && indexJ < lenJ && strJ[indexJ] == '1'  {
	// 	return false
	// }
	// if indexJ == lenJ && indexI < lenI && strI[indexI] == '0' {
	// 	return false
	// } else if indexJ == lenJ && indexI < lenI && strI[indexI] == '1'  {
	// 	return true
	// }

	// return false
	a, _ := strconv.Atoi(bs[i] + bs[j])
	b, _ := strconv.Atoi(bs[j] + bs[i])
	return a > b
}

func (bs BinarySlice) Swap(i, j int) {
	bs[i], bs[j] = bs[j], bs[i]
}

func getBinaryFromFirstOne(num int64) string {
	binary := strconv.FormatInt(num, 2)
	firstOneIndex := 0
	for i, char := range binary {
		if char == '1' {
			firstOneIndex = i
			break
		}
	}
	return binary[firstOneIndex:]
}

func maxGoodNumber(nums []int) int {

	binary := make([]string, len(nums))
	for i, num := range nums {
		// binary[i] = strconv.Itoa(num)
		binary[i] = getBinaryFromFirstOne(int64(num))
	}
	sort.Sort(BinarySlice(binary))

	resStr := ""
	for _, num := range binary {
		resStr += num
	}
	resInt, _ := strconv.ParseInt(resStr, 2, 64)
	return int(resInt)
}
