package main

import "fmt"

func main() {
	testArr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(getSum(testArr))
	max := getMax(testArr)
	fmt.Printf("the max number of this arr is %d", max)
	fmt.Println()

	testArr2 := []int{1, 1, 2, 2, 3, 4, 5, 5}
	result := getSame(testArr2)
	fmt.Printf("the origin arr is %d", result)

}

//将数组作为参数传入
func getSum(arr [5]int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum

}

func getMax(arr [5]int) int {
	max := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

func hasEl(arr []int, el int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == el {
			return true
		}
	}
	return false
}

//数组去重
// func removeDuplicate(arr []int)[]int {
//     newArr :=[] int{}
//     for i:=0; i<len(arr); i++{
//         if(!hasEl(newArr, arr[i])){
//             length :=len(newArr)
//             newArr[length]=arr[i]
//             }
//     }
//     return newArr

// }

func getSame(arr []int) []int {
	return arr
}
