// +build ignore

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	totalAmount := 100 * 100 // 总金额，分
	totalPeople := 8         // 总人数
	fmt.Println(totalAmount/100, totalPeople)
	ret := doubleMeanAlgo(totalAmount, totalPeople)
	var tmp float64
	for i := range ret {
		yuan := float64(ret[i]) / 100
		fmt.Println(yuan)
		tmp += yuan
	}
	fmt.Println(tmp)
}

// 红包算法，二倍均值法
// totalAmount，总金额，单位：分
// totalPeople，总人数
func doubleMeanAlgo(totalAmount, totalPeople int) []int {
	arr := make([]int, 0, totalPeople)
	for {
		money := rand.Intn(totalAmount/totalPeople*2-1) + 1
		arr = append(arr, money)
		totalAmount -= money
		totalPeople--
		if totalPeople == 1 {
			break
		}
	}
	arr = append(arr, totalAmount)
	return arr
}
