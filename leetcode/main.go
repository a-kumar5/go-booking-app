package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	count_s := map[string]int{}
	count_t := map[string]int{}

	for _, c := range s {
		if _, count := count_s[string(c)]; count == false {
			count_s[string(c)] = 0
		}
		count_s[string(c)] += 1
	}
	for _, c := range t {
		if _, count := count_s[string(c)]; count == false {
			count_s[string(c)] = 0
		}
		count_t[string(c)] += 1
	}

	for key, _ := range count_s {
		_, ok := count_t[key]

		if ok == false || count_t[key] != count_s[key] {
			return false
		}
	}
	return true
}

/*
	func twoSum(nums []int, target int) []int {
		var targetSlice []int
		for i := 0; i < len(nums); i++ {
			//fmt.Printf("%d %d\n", i, nums[i])
			for j := i + 1; j < len(nums); j++ {
				//fmt.Printf("%d\n", nums[j])
				if nums[i]+nums[j] == target {
					targetSlice = append(targetSlice, i)
					targetSlice = append(targetSlice, j)
				}
			}
		}
		return targetSlice
	}
*/
func twoSum(nums []int, target int) []int {
	mp := make(map[int]int)
	for i, val := range nums {
		if val, ok := mp[target-val]; ok {
			return []int{val, i}
		}
		mp[val] = i
	}

	return []int{}
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	original := x
	reversed := 0
	for x != 0 {
		reversed = reversed*10 + x%10
		x /= 10
	}
	return reversed == original
}

/*
Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
*/
func romanToInt(s string) int {
	romanLiteral := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	//fmt.Println(romanLiteral)
	sum := 0
	for i := 0; i < len(s); i++ {
		val := romanLiteral[s[i]]
		if i < len(s)-1 && val < romanLiteral[s[i+1]] {
			sum -= val
		} else {
			sum += val
		}
	}
	return sum
}

func main() {
	//fmt.Println("Hello, 世界")
	//var s = "anagram"
	//var t = "nagaram"
	//fmt.Println(isAnagram(s, t))
	//nums := []int{2, 7, 11, 15}
	//num := twoSum(nums, 9)
	//fmt.Println(num)
	//fmt.Println(isPalindrome(121))
	s := "MCMXCIV"
	fmt.Println(romanToInt(s))
}
