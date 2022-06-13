package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	// data set [1, X, 8, 17, Y, Z, 78, 113]
	/* Pseudocode
	Write a function foo()
	For loop X between (not included) 1 and 8
		For loop Y between (not included) 17 and Z
			For loop Z betwwen (not included) Y and 78
				dataSet = [1, X, 8, 17, Y, Z, 78, 113]
				If dataSet is valid
					Print("value of X = %v", X)
					Print("value of Y = %v", Y)
					Print("value of Z = %v", Z)
	*/

	fmt.Println("find value of x, y and z...")
	count := 0
	var wg sync.WaitGroup
	countCh := make(chan int)
	totalCountCh := make(chan int)

	go func() {
		totalCount := 0
		for num := range countCh {
			totalCount += num
		}
		totalCountCh <- totalCount
	}()

	for x := 1; x < 8; x++ {
		for y := 18; y < 77; y++ {
			for z := y + 1; z < 78; z++ {
				count++
				wg.Add(1)
				go func(i, j, k int) {
					defer wg.Done()
					dataSet := []int{1, i, 8, 17, j, k, 78, 113}
					if checkDataSet(dataSet) {
						fmt.Printf("value of (x, y, z) = (%v, %v, %v) in data set %v\n", i, j, k, dataSet)
						countCh <- 1
					}
				}(x, y, z)
			}
		}
	}
	wg.Wait()
	close(countCh)

	fmt.Printf("possible values that make data set valid: %v out of %v\n", <-totalCountCh, count)
}

func checkDataSet(dataSet []int) bool {
	/*
		##############################################################
		### Specify how data set element relate to each other here ###
		##############################################################
	*/
	/* Ex. If mod 2 and result is in order of 1, 0, 0, 1,...
	for i, elem := range dataSet {
		switch i % 4 {
		case 0, 3:
			if elem%2 != 1 {
				return false
			}
		case 1, 2:
			if elem%2 != 0 {
				return false
			}
		}
	}
	*/

	// If check prime number and result is in order of no, yes, no, yes,...
	for i, elem := range dataSet {
		res := isPrime(elem)
		switch i % 2 {
		case 0:
			if res {
				return false
			}
		case 1:
			if !res {
				return false
			}
		}
	}
	return true
}

func isPrime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}
