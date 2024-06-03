package main

func main() {
	println(maximumLength("aaaa"))
	//fmt.Println(BubbleSort([]int{1, 2, 3, 46, 4, 5, 6}, false))
}

func maximumLength(s string) int {
	// 初始化字母表map
	alphabetMap := make(map[rune][]int)
	for i := 'a'; i <= 'z'; i++ {
		alphabetMap[i] = []int{}
	}
	// 遍历s
	count := 0
	var char rune
	for i := 0; i < len(s); i++ {
		char = rune(s[i])
		if i+1 == len(s) {
			alphabetMap[char] = append(alphabetMap[char], count+1)
			break
		}
		if s[i+1] == s[i] {
			count++
		} else {
			alphabetMap[char] = append(alphabetMap[char], count+1)
			count = 0
		}
	}
	// 排序字母表map的值
	var sl []int
	var result int
	for i := 'a'; i <= 'z'; i++ {
		sl = alphabetMap[i]
		if len(sl) != 0 {
			sl = BubbleSort(alphabetMap[i], false)
			sl = append(sl, []int{0, 0}...)
			if max(sl[0]-2, min(sl[0]-1, sl[0]), sl[2]) > result {
				result = max(sl[0]-2, min(sl[0]-1, sl[1]), sl[2])
			}
		}

	}
	if result == 0 {
		return -1
	}

	return result
}

// 排序
func BubbleSort(src []int, mode bool) []int {
	// mode 为true从小到大；false 则从大到小
	length := len(src)
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			if mode && src[j] > src[j+1] {
				src[j], src[j+1] = src[j+1], src[j]
			} else if src[j] < src[j+1] {
				src[j], src[j+1] = src[j+1], src[j]
			}
		}
	}
	return src
}
