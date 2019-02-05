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

func boggle(board [][]string, word string) {
	fmt.Print("")
}

func isOneChangeAway(word string, word2 string) bool {
	r := len(word2) - len(word)

	if r == 0 {
		sum := 0
		for i := range word2 {
			if word2[i] != word[i] {
				sum++
			}
		}

		if sum == 1 {
			return true
		}
	} else if r == -1 {
		for i := range word {
			if i == len(word2) {
				word2 = word2 + string(word[i])
				if word2 == word {
					return true
				}
			} else {
				if word2[i] != word[i] {
					word2 = word2[:i] + string(word[i]) + word2[i:len(word2)]
					if word2 == word {
						return true
					}
				}
			}
		}
	} else if r == 1 {
		for i := range word2 {
			if i == len(word) {
				word2 = word2[:len(word2)-1]
				if word2 == word {
					return true
				}
			} else {
				if word2[i] != word[i] {
					word2 = word2[:i] + word2[i+1:len(word2)]
					if word2 == word {
						return true
					}
					return false
				}
			}
		}
	}

	return false
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
	fmt.Println(isOneChangeAway("abc", "bbc"))
	fmt.Println(isOneChangeAway("abc", "aba"))
	fmt.Println(isOneChangeAway("abc", "acc"))
	fmt.Println(isOneChangeAway("abc", "cbc"))
	fmt.Println(isOneChangeAway("abc", "bca"))
	fmt.Println(isOneChangeAway("abc", "c"))
	fmt.Println(isOneChangeAway("abc", "ab"))
	fmt.Println(isOneChangeAway("abc", "ac"))
	fmt.Println(isOneChangeAway("abc", "bc"))
	fmt.Println(isOneChangeAway("abc", "abca"))
	fmt.Println(isOneChangeAway("abc", "aabc"))
	fmt.Println(isOneChangeAway("abc", "abbc"))
	fmt.Println(isOneChangeAway("abc", "dabc"))
	fmt.Println(isOneChangeAway("abc", "cbassss"))

	//fibo(10)
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
