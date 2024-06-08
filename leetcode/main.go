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
	reversed := 0
	for x != 0 {
		reversed = reversed*10 + x%10
		x /= 10
	}
	return true
}

func main() {
	//fmt.Println("Hello, 世界")
	//var s = "anagram"
	//var t = "nagaram"
	//fmt.Println(isAnagram(s, t))
	//nums := []int{2, 7, 11, 15}
	//num := twoSum(nums, 9)
	//fmt.Println(num)
	fmt.Println(isPalindrome(121))
}
