package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

func main() {
	question1()
	question2()
}
func question1() {

	x := os.Args[1]
	num, _ := strconv.Atoi(x)

	// var a [num + 1]int
	// for i := 0; i <= num; i++ {
	//	a[i] = i
	//}

	// make slice
	firstTimeSlice := time.Now()
	var s []int
	for i := 0; i <= num; i++ {
		s = append(s, i)
	}
	lastTimeSlice := time.Now()
	diffSlice := lastTimeSlice.Sub(firstTimeSlice)
	fmt.Println(diffSlice)

	// make map
	firstTimeMap := time.Now()
	m := make(map[int]int)
	for i := 0; i <= num; i++ {
		m[i] = i
	}
	lastTimeMap := time.Now()
	diffMap := lastTimeMap.Sub(firstTimeMap)
	fmt.Println(diffMap)

	// add 1 to each element in slice
	firstTimeSlice = time.Now()
	for i := 0; i <= num; i++ {
		s[i] += 1
	}
	lastTimeSlice = time.Now()
	diffSlice = lastTimeSlice.Sub(firstTimeSlice)
	fmt.Println(diffSlice)

	// add 1 to each value in map
	firstTimeMap = time.Now()
	for i := 0; i <= num; i++ {
		m[i] += 1
	}
	lastTimeMap = time.Now()
	diffMap = lastTimeMap.Sub(firstTimeMap)
	fmt.Println(diffMap)
}

func question2() {
	x := os.Args[1]
	num, _ := strconv.Atoi(x)
	rand.Int()

	var s []int
	for i := 0; i <= num; i++ {
		s = append(s, rand.Int())
	}

	firstTimeSlice := time.Now()
	sort.Slice(s, func(i int, j int) bool { return i < j })
	lastTimeSlice := time.Now()
	diffSlice := lastTimeSlice.Sub(firstTimeSlice)
	fmt.Println(diffSlice)

	for i := 0; i <= num; i++ {
		s[i] = rand.Int()
	}
	firstTimeSlice = time.Now()
	sort.SliceStable(s, func(i int, j int) bool { return i < j })
	lastTimeSlice = time.Now()
	diffSlice = lastTimeSlice.Sub(firstTimeSlice)
	fmt.Println(diffSlice)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
