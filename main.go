package main

import (
	"fmt"
)

func calculateProductOfAll(arr []int) (result []int) {

	//write code here
	for i := 0; i < len(arr); i++ {
		result = append(result, calculateArray(arr, i))
	}

	return result
}

func calculateArray(arr []int, pos int) int {
	result := 1

	for i, x := range arr {
		if i != pos {
			result *= x
		}
	}

	return result
}

func fibo(n int) {
	x, y := 0, 0
	f := make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			//fmt.Print(i, " ")
			x = i
			f[i] = x
		} else if i == 1 {
			//fmt.Print(i, " ")
			y = 1
			f[i] = y
		} else {

			tx := y
			ty := x + y

			f[i] = ty

			//fmt.Print(x+y, " ")

			x = tx
			y = ty
		}
	}

	fmt.Print(f)
}

// Bootstrap code {
// Code snippet in grey is not editable and is necessary
func main() {
	fibo(10)
	// scanner := bufio.NewScanner(os.Stdin)
	// scanner.Scan() // use `for scanner.Scan()` to keep reading
	// line := scanner.Text()
	// arr := strings.Split(line, " ")

	// var array = []int{}

	// for _, i := range arr {
	// 	j, err := strconv.Atoi(i)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	array = append(array, j)
	// }

	// var result = calculateProductOfAll(array)
	// for _, element := range result {
	// 	fmt.Print(element, " ")
	// }

}

//}
