package main

// EvenOrOdd は偶数か奇数かを返す関数
func EvenOrOdd(number int) string {
	if number%2 == 0 {
		return "even"
	} else {
		return "odd"
	}
}
